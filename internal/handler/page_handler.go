package handler

import (
	"bytes"
	"context"
	"html/template"
	"net/http"
	"net/url"
	api "prel/api/prel_api"
	"prel/config"
	"prel/internal/gateway/repository"
	"prel/internal/model"
	"prel/pkg/custom_error"
	"prel/pkg/utils"
	tpl "prel/web/template"
	"sort"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/jackc/pgx/v5"
)

func (h *Handler) Get(ctx context.Context, params api.GetParams) (api.GetRes, error) {

	// if token is set and valid, redirect to request form page
	if params.Token.Set && params.Token.Value != "" {
		user, err := repository.NewUserRepository().FindBySessionID(ctx, model.SessionID(params.Token.Value))
		if err != nil || user == nil {
			goto index
		}

		now := model.GetClock(ctx).Now()
		if user.IsSessionExpired(now) {
			goto index
		}

		if !user.IsAvailable() {
			goto index
		}

		url, err := url.Parse("request-form")
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse url")
		}
		return &api.GetSeeOtherHeaders{
			Location: api.NewOptURI(*url),
		}, nil
	}

index:

	b := &bytes.Buffer{}
	tmpl, err := template.ParseFS(tpl.Files, tpl.IndexPageTpl)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse template")
	}

	err = tmpl.Execute(b, &tpl.PageData{
		HeaderData: tpl.NewHeaderData(false, config.AppName),
	})

	if err != nil {
		return nil, errors.Wrap(err, "failed to execute template")
	}

	return &api.GetOK{
		Data: b,
	}, nil
}

func (h *Handler) SigninPost(ctx context.Context) (api.SigninPostRes, error) {
	state, err := utils.RandomString(32)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate state")
	}

	url, err := url.Parse(h.oauthConfig.AuthCodeURL(state))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse url")
	}

	return &api.SigninPostTemporaryRedirect{
		Location:  api.NewOptURI(*url),
		SetCookie: api.NewOptString(generateCookie(h.config.URL, "state", state, time.Now().Add(5*time.Minute))),
	}, nil
}

func (h *Handler) AuthGoogleCallbackGet(ctx context.Context, params api.AuthGoogleCallbackGetParams) (api.AuthGoogleCallbackGetRes, error) {
	if params.QueryState != params.CookieState.Value {
		return nil, errors.WithDetail(errors.Errorf("invalid oauth state given %s", params.QueryState), string(custom_error.InvalidArgument))
	}

	user, err := h.usecase.Signin(ctx, h.oauthConfig, params.Code)
	if err != nil {
		return nil, err
	}

	token := generateCookie(h.config.URL, "token", string(user.SessionID()), user.SessionExpiredAt())
	return &api.AuthGoogleCallbackGetTemporaryRedirect{
		Location:  api.NewOptString("/request-form"),
		SetCookie: api.NewOptString(token),
	}, nil
}

func (h *Handler) SignoutPost(ctx context.Context) (api.SignoutPostRes, error) {
	url, err := url.Parse("/")
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse url")
	}

	user := model.GetUser(ctx)
	err = repository.NewTransactionManager().Transaction(ctx, func(ctx context.Context) error {
		user.Signout()
		return repository.NewUserRepository().Save(ctx, user)
	})

	if err != nil {
		return nil, err
	}

	return &api.SignoutPostSeeOtherHeaders{
		SetCookie: api.NewOptString(generateCookie(h.config.URL, "token", "", time.Time{})),
		Location:  api.NewOptURI(*url),
	}, nil
}

func (h *Handler) RequestFormGet(ctx context.Context) (api.RequestFormGetRes, error) {
	user := model.GetUser(ctx)
	pjs, err := h.client.GetProjects()
	if err != nil {
		return nil, err
	}

	tmpl, err := template.ParseFS(tpl.Files, tpl.RequestFormPageTpl, tpl.HeaderTpl)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse template")
	}

	sort.Slice(pjs, func(i, j int) bool {
		return pjs[i].Name < pjs[j].Name
	})

	projectsTpl := convertTplProjects(pjs)
	periodsTpl := convertTplPeriods()

	b := &bytes.Buffer{}
	err = tmpl.Execute(b, &tpl.PageData{
		HeaderData: tpl.NewHeaderData(user.IsAdmin(), config.AppName),
		RequestFormPage: tpl.RequestFormPage{
			Email:    user.Email(),
			Projects: projectsTpl,
			Periods:  periodsTpl,
		}})

	if err != nil {
		return nil, errors.Wrap(err, "failed to execute template")
	}

	return &api.RequestFormGetOK{Data: b}, nil
}

func (h *Handler) RequestGet(ctx context.Context) (api.RequestGetRes, error) {

	now := model.GetClock(ctx).Now()
	repo := repository.NewRequestRepository()
	reqs, err := repo.FindRequestByStatusAndExpiredAt(ctx, model.RequestStatusPending, now)
	if err != nil {
		return nil, err
	}

	sort.Slice(reqs, func(i, j int) bool {
		return reqs[i].ExpiredAt().Before(reqs[j].ExpiredAt())
	})

	tmpl, err := template.ParseFS(tpl.Files, tpl.RequestPageTpl, tpl.HeaderTpl)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse template")
	}

	user := model.GetUser(ctx)
	b := &bytes.Buffer{}
	err = tmpl.Execute(b, &tpl.PageData{
		HeaderData: tpl.NewHeaderData(user.IsAdmin(), config.AppName),
		RequestPage: tpl.RequestPage{
			Requests: convertTplRequests(user, reqs),
		}})

	if err != nil {
		return nil, err
	}

	return &api.RequestGetOK{Data: b}, nil
}

func (h *Handler) RequestRequestIDGet(ctx context.Context, params api.RequestRequestIDGetParams) (api.RequestRequestIDGetRes, error) {

	repo := repository.NewRequestRepository()
	req, err := repo.FindByID(ctx, params.RequestID)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.WithDetail(err, string(custom_error.ResourceNotFound))
	}

	if err != nil {
		return nil, err
	}

	user := model.GetUser(ctx)
	reqsTpl := convertTplRequest(user, req)
	tmpl, err := template.ParseFS(tpl.Files, tpl.RequestPageTpl, tpl.HeaderTpl)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse template")
	}

	b := &bytes.Buffer{}
	err = tmpl.Execute(b, &tpl.PageData{
		HeaderData: tpl.NewHeaderData(user.IsAdmin(), config.AppName),
		RequestPage: tpl.RequestPage{
			Requests: []*tpl.Request{reqsTpl},
		}})

	if err != nil {
		return nil, errors.Wrap(err, "failed to execute template")
	}

	return &api.RequestRequestIDGetOK{Data: b}, nil
}

func (h *Handler) AdminIamRoleFilteringGet(ctx context.Context) (api.AdminIamRoleFilteringGetRes, error) {
	user := model.GetUser(ctx)

	b := &bytes.Buffer{}
	tmpl, err := template.ParseFS(tpl.Files, tpl.AdminIamRoleFilteringPageTpl, tpl.HeaderTpl)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse template")
	}

	err = tmpl.Execute(b, &tpl.PageData{
		HeaderData: tpl.NewHeaderData(user.IsAdmin(), config.AppName),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute template")
	}

	return &api.AdminIamRoleFilteringGetOK{
		Data: b,
	}, nil
}

func (h *Handler) HealthGet(ctx context.Context) (api.HealthGetRes, error) {
	return &api.HealthGetNoContent{}, nil
}

func generateCookie(url, key, value string, expiredAt time.Time) string {
	cookie := &http.Cookie{
		Name:     key,
		Value:    value,
		Path:     "/",
		Expires:  expiredAt,
		HttpOnly: true,
	}
	if strings.HasPrefix(url, "https") {
		cookie.Secure = true
	}
	return cookie.String()
}
