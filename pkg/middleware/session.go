package middleware

import (
	"context"
	"net/http"
	api "prel/api/prel_api"
	"prel/pkg/custom_error"
	"strings"

	"prel/internal/gateway/repository"
	"prel/internal/model"

	"github.com/cockroachdb/errors"
)

type sessionValidator struct{}

func (s *sessionValidator) HandleCookieAuth(ctx context.Context, operationName string, t api.CookieAuth) (context.Context, error) {

	user := model.GetUser(ctx)
	if user == nil {
		return nil, errors.WithDetail(errors.New("user is not set"), string(custom_error.InvalidArgument))
	}

	// session expired check
	now := model.GetClock(ctx).Now()
	if user.IsSessionExpired(now) {
		return nil, errors.WithDetail(errors.New("session is expired"), string(custom_error.SessionExpired))
	}

	// user availability check
	if !user.IsAvailable() {
		return nil, errors.WithDetail(errors.New("user is not available"), string(custom_error.UnavailableUser))
	}

	// role admin check
	if strings.Contains(operationName, "Admin") {
		if !user.IsAdmin() {
			return nil, errors.WithDetail(errors.New("user is not admin"), string(custom_error.OnlyAdmin))
		}
	}

	return ctx, nil
}

func NewSessionValidator() *sessionValidator {
	return &sessionValidator{}
}

func Session() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()

			// same as api.CookieAuth
			cookie, err := r.Cookie("token")
			if err != nil || cookie == nil || cookie.Value == "" {
				next.ServeHTTP(w, r)
				return
			}

			// through error if session is invalid
			// HandleCookieAuth will handle the error
			//
			// in ogen, HandleCookieAuth is called after middlewares
			// so we need to apply user id to context here for Logger middleware
			id := model.SessionID(cookie.Value)
			user, err := repository.NewUserRepository().FindBySessionID(ctx, id)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			ctx = model.SetUser(ctx, user)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}
