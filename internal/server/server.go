package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	api "prel/api/prel_api"
	"prel/internal"
	"prel/internal/gateway/google_cloud"
	"prel/internal/gateway/notification"
	"prel/internal/gateway/postgresql"
	"prel/internal/handler"
	"prel/internal/model"

	"prel/config"
	"prel/pkg/logger"
	"prel/pkg/middleware"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/iam/v1"
)

func Run(ctx context.Context) {

	requestLoggerBuilder := logger.NewBuilder().
		WithOutput(os.Stdout).
		WithOptions(slog.HandlerOptions{
			AddSource:   true,
			ReplaceAttr: logger.ReplaceAttr,
		})
	appLoggerBuilder := logger.NewBuilder().
		WithOutput(os.Stderr).
		WithOptions(slog.HandlerOptions{
			AddSource:   true,
			ReplaceAttr: logger.ReplaceAttr,
		})

	appLogger := appLoggerBuilder.Build()
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	defer func() {
		if err := recover(); err != nil {
			pc, file, line, ok := runtime.Caller(2)
			if ok {
				f := runtime.FuncForPC(pc)
				appLogger.Error(fmt.Sprintf("panic: %v, function: %s, file: %s, line: %d", err, f.Name(), file, line))
			} else {
				appLogger.Error(fmt.Sprintf("panic: %v", err))
			}
		}
	}()

	var c config.Config
	err := viper.Unmarshal(&c)

	if err != nil {
		panic(err)
	}

	config.SetDebug(c.IsDebug)

	switch c.DBType {
	case "fixed":
		err = postgresql.Initialize(c.DBUsername, c.DBPassword, c.DBName, postgresql.WithFixedDB(c.DBHost, c.DBPort, c.DBSslMode))
	case "cloud-sql-connector":
		err = postgresql.Initialize(c.DBUsername, c.DBPassword, c.DBName, postgresql.WithCloudSQLConnector(c.DBInstanceConnection))
	default:
		panic(fmt.Sprintf("invalid db type: %s", c.DBType))
	}

	if err != nil {
		panic(err)
	}

	crmService, err := cloudresourcemanager.NewService(ctx)
	if err != nil {
		panic(err)
	}

	iamService, err := iam.NewService(ctx)
	if err != nil {
		panic(err)
	}

	rSvc := google_cloud.ProjectsServiceWrapper{Service: crmService.Projects}
	iamSvc := google_cloud.RolesServiceWrapper{Service: iamService.Roles}
	client := google_cloud.NewClient(ctx, c.ProjectID, &rSvc, &iamSvc)

	httpClient := &http.Client{Timeout: 5 * time.Second}

	var notificationClient internal.NotificationClient
	switch c.NotificationType {
	case "slack":
		notificationClient = notification.NewSlackClient(httpClient, c.NotificationUrl)
	default:
		panic(fmt.Sprintf("invalid notification type: %s", c.NotificationType))
	}

	oauth := &oauth2.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		Endpoint:     google.Endpoint,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		RedirectURL: fmt.Sprintf("%s%s", c.URL, config.RedirectPath),
	}

	h := handler.NewHandler(&c, oauth, client, httpClient, notificationClient)
	apiServer, err := api.NewServer(
		h,
		middleware.NewSessionValidator(),
		api.WithErrorHandler(handler.ErrorHandler),
		api.WithNotFound(handler.NotFoundHandler),
	)
	if err != nil {
		panic(err)
	}

	addr := fmt.Sprintf("%s:%s", c.Address, c.Port)
	appLogger.Info("Server started on " + addr)

	clock := &model.RealClock{}

	// note: middleware order is important
	// note: middleware is executed in reverse order
	middlewares := []middleware.Middleware{
		middleware.Logger(requestLoggerBuilder, appLoggerBuilder),
		middleware.Session(),
		middleware.Clock(clock),
		middleware.Recover(appLogger),
	}

	var handler http.Handler = apiServer
	for _, mw := range middlewares {
		handler = middleware.Wrap(handler, mw)
	}

	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		err = server.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			appLogger.Info("Server closed")
		} else {
			appLogger.Error(fmt.Sprintf("failed to listen and serve: %v", err))
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = server.Shutdown(ctx)
	if err != nil {
		appLogger.Error("failed to shutdown: %v", err)
	}

	appLogger.Info("Server stopped")
}
