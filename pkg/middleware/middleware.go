package middleware

import "net/http"

type Middleware func(next http.Handler) http.Handler

func Wrap(h http.Handler, mw Middleware) http.Handler {
	return mw(h)
}
