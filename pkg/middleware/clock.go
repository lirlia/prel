package middleware

import (
	"net/http"
	"prel/internal/model"
)

func Clock(clock model.Clock) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()
			ctx = model.SetClock(ctx, clock)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
