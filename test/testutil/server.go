package testutil

import (
	"context"
	"errors"
	"net"
	"net/http"
	"net/http/httptest"
	"prel/internal"
	google_cloud_mock "prel/internal/gateway/google_cloud/mock"
	"prel/internal/gateway/repository"
	"time"

	api "prel/api/prel_api"
	"prel/internal/gateway/google_cloud"
	"prel/internal/gateway/notification"
	"prel/internal/handler"
	"prel/internal/model"

	"prel/config"
	"prel/pkg/logger"
	"prel/pkg/middleware"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega/gbytes"
	"go.uber.org/mock/gomock"
)

type TestHelper struct {
	Ctrl                          *gomock.Controller
	Clock                         *Clock
	Listener                      net.Listener
	NotificationServerListener    net.Listener
	Config                        *config.Config
	Ctx                           context.Context
	UserRepo                      model.UserRepository
	RequestRepo                   model.RequestRepository
	InvitationRepo                model.InvitationRepository
	IamRoleFilteringRepo          model.IamRoleFilteringRuleRepository
	SettingRepo                   model.SettingRepository
	ApiClient                     *ApiClient
	HttpClient                    *http.Client
	NotificationClient            internal.NotificationClient
	NotificationServer            *httptest.Server
	GoogleCloudClient             internal.GoogleCloudClient
	User                          *model.User
	GoogleCloudResourceManagerSrv *google_cloud_mock.MockResourceManagerService
	GoogleCloudIamSrv             *google_cloud_mock.MockIamService
	RequestLogBuffer              *gbytes.Buffer
	AppLogBuffer                  *gbytes.Buffer
}

type ApiClient struct {
	*api.Client
	securitySource securitySource
}

func NewTestHelper() *TestHelper {

	ctx := context.Background()

	SetupTestDB()

	c := &config.Config{
		ProjectID:            "test-project-id",
		Address:              "localhost",
		Port:                 "12345",
		URL:                  "http://localhost:12345",
		ClientID:             "",
		ClientSecret:         "",
		SessionExpireSeconds: 43200,
		RequestExpireSeconds: 86400,
		IsDebug:              false,
	}

	httpClient := &http.Client{Timeout: 5 * time.Second}

	return &TestHelper{
		Clock:                NewClock(),
		Ctx:                  ctx,
		Config:               c,
		UserRepo:             repository.NewUserRepository(),
		RequestRepo:          repository.NewRequestRepository(),
		InvitationRepo:       repository.NewInvitationRepository(),
		IamRoleFilteringRepo: repository.NewIamRoleFilteringRuleRepository(),
		SettingRepo:          repository.NewSettingRepository(),
		HttpClient:           httpClient,
	}
}

func (h *TestHelper) StartServer() {

	var err error
	h.Ctrl = gomock.NewController(ginkgo.GinkgoT())

	h.GoogleCloudResourceManagerSrv = google_cloud_mock.NewMockResourceManagerService(h.Ctrl)
	h.GoogleCloudIamSrv = google_cloud_mock.NewMockIamService(h.Ctrl)
	h.GoogleCloudClient = google_cloud.NewClient(h.Ctx, h.Config.ProjectID, h.GoogleCloudResourceManagerSrv, h.GoogleCloudIamSrv)

	h.NotificationServerListener, err = net.Listen("tcp", "127.0.0.1:0")
	h.NotificationClient = notification.NewSlackClient(h.HttpClient, "http://"+h.NotificationServerListener.Addr().String())
	if err != nil {
		panic(err)
	}
	hlr := handler.NewHandler(h.Config, nil, h.GoogleCloudClient, h.HttpClient, h.NotificationClient)

	apiServer, err := api.NewServer(
		hlr,
		middleware.NewSessionValidator(),
		api.WithErrorHandler(handler.ErrorHandler),
		api.WithNotFound(handler.NotFoundHandler),
	)
	if err != nil {
		panic(err)
	}

	h.RequestLogBuffer = gbytes.NewBuffer()
	h.AppLogBuffer = gbytes.NewBuffer()
	requestLoggerBuilder := logger.NewBuilder().WithOutput(h.RequestLogBuffer)
	appLoggerBuilder := logger.NewBuilder().WithOutput(h.AppLogBuffer)

	middlewares := []middleware.Middleware{
		middleware.Session(),
		middleware.Logger(requestLoggerBuilder, appLoggerBuilder),
		middleware.Clock(h.Clock),
		// middleware.Recover(appLogger),
	}

	var handler http.Handler = apiServer
	for _, mw := range middlewares {
		handler = middleware.Wrap(handler, mw)
	}

	h.Listener, err = net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	server := &http.Server{Handler: handler}

	go func() {
		defer ginkgo.GinkgoRecover()
		err = server.Serve(h.Listener)
		if !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	h.User, err = h.CreateUser()
	if err != nil {
		panic(err)
	}

	h.ApiClient, err = h.NewApiClient(string(h.User.SessionID()))
	if err != nil {
		panic(err)
	}
}

func (h *TestHelper) NewUserAndClient(opts ...UserOption) (*model.User, *ApiClient, error) {
	user, err := h.CreateUser(opts...)
	if err != nil {
		return nil, nil, err
	}
	client, err := h.NewApiClient(string(user.SessionID()))
	if err != nil {
		return nil, nil, err
	}
	return user, client, nil
}

func (h *TestHelper) CreateUser(opts ...UserOption) (*model.User, error) {
	user := NewTestUser(opts...)
	err := h.UserRepo.Save(h.Ctx, user)
	return user, err
}

func (h *TestHelper) NewApiClient(session string) (*ApiClient, error) {

	ss := securitySource{session: session}
	c, err := api.NewClient("http://"+h.Listener.Addr().String(), ss)
	if err != nil {
		return nil, err
	}
	return &ApiClient{c, ss}, nil
}

func (h *TestHelper) UpdateSession(c *ApiClient, expiredAt time.Time) (*ApiClient, error) {

	user, err := h.UserRepo.FindBySessionID(h.Ctx, model.SessionID(c.securitySource.session))
	if err != nil {
		return nil, err
	}

	user.UpdateSession(expiredAt)
	err = h.UserRepo.Save(h.Ctx, user)
	if err != nil {
		return nil, err
	}

	return h.NewApiClient(string(user.SessionID()))
}

type securitySource struct {
	session string
}

func (s securitySource) CookieAuth(ctx context.Context, operationName string) (api.CookieAuth, error) {
	return api.CookieAuth{
		APIKey: s.session,
	}, nil
}

func (h *TestHelper) StartNotificationServer(opts ...ServerOption) {
	server := h.startServer(opts...)
	h.NotificationServer = server
}

func (h *TestHelper) startServer(opts ...ServerOption) *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	server := &httptest.Server{
		Listener: h.NotificationServerListener,
		Config:   &http.Server{Handler: handler},
	}

	for _, opt := range opts {
		opt(server)
	}

	server.Start()
	return server
}

type ServerOption func(*httptest.Server)

func WithServerHandler(handler http.HandlerFunc) ServerOption {
	return func(s *httptest.Server) {
		s.Config.Handler = handler
	}
}

func CheckHandler(checkFn func(http.ResponseWriter, *http.Request), statusCode int) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		checkFn(w, r)
		w.WriteHeader(statusCode)
	}
}

func (h *TestHelper) Close() {
	if h.NotificationServer != nil {
		h.NotificationServer.Close()
	}
}

func (h *TestHelper) ChangeRole(user *model.User, role model.UserRole) error {
	user.SetRole(role)
	return h.UserRepo.Save(h.Ctx, user)
}
