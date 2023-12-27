package middleware

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"prel/config"
	"prel/internal/model"
	"prel/pkg/logger"
)

var ignorePaths = []string{
	"/health",
	"/favicon.ico",
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// get status code from response for logging
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func Logger(requestBuilder, appBuilder logger.Builder) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			wrapped := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

			ctx := r.Context()
			now := model.GetClock(ctx).Now()

			httpRequestArgs := []any{
				slog.String("requestMethod", r.Method),
				slog.String("requestUrl", r.URL.Path),
				slog.String("remoteIp", r.RemoteAddr),
			}

			args := []any{}

			user := model.GetUser(ctx)
			if user != nil {
				args = append(args, slog.Group("metadata",
					slog.String("user_id", string(user.ID())),
					slog.String("email", user.Email())),
				)
			}

			appArgs := append(args, slog.Group("httpRequest", httpRequestArgs...))
			appLogger := appBuilder.Build()
			appLogger = appLogger.With(appArgs...)
			ctx = logger.Set(ctx, appLogger)
			r = r.WithContext(ctx)

			for _, path := range ignorePaths {
				if path == r.URL.Path {
					next.ServeHTTP(w, r)
					return
				}
			}

			if config.IsDebug() {
				// output request payload
				if r.Method == http.MethodPost || r.Method == http.MethodPut || r.Method == http.MethodPatch {
					var payload []byte
					if r.Body != nil {
						payload, _ = io.ReadAll(r.Body)
					}
					r.Body = io.NopCloser(bytes.NewBuffer(payload))

					// show beautiful payload in log
					slog.Info("request payload", slog.String("payload", string(payload)))
				}
			}

			next.ServeHTTP(wrapped, r)
			latency := model.GetClock(ctx).Now().Sub(now).String()
			httpRequestArgs = append(httpRequestArgs, slog.Int("status", wrapped.statusCode))
			httpRequestArgs = append(httpRequestArgs, slog.String("latency", latency))
			reqArgs := append(args, slog.Group("httpRequest", httpRequestArgs...))

			requestLogger := requestBuilder.Build()
			requestLogger = requestLogger.With(reqArgs...)

			if wrapped.statusCode >= http.StatusInternalServerError {
				requestLogger.Error(r.URL.Path)
			} else {
				requestLogger.Info(r.URL.Path)
			}
		})
	}
}
