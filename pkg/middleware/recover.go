package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"
)

func Recover(l *slog.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				err := recover()
				if err != nil {
					stack := debug.Stack()
					slog.Error(fmt.Sprintf("panic: %+v, %s", err, stack))
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
