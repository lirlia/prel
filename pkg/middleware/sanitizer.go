package middleware

import (
	"bytes"
	"io"
	"net/http"
)

func Sanitizer() Middleware {
	maxBodySize := int64(1024 * 1024) // Max body size set to 1MiB

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.Method == http.MethodPost ||
				r.Method == http.MethodPut ||
				r.Method == http.MethodPatch {

				// limit request Body size
				r.Body = http.MaxBytesReader(w, r.Body, maxBodySize)

				// read request Body
				buf, err := io.ReadAll(r.Body)
				if err != nil {
					http.Error(w, "Bad Request: Failed to read body", http.StatusBadRequest)
					return
				}
				buf = bytes.ReplaceAll(buf, []byte("<"), []byte("&lt;"))
				buf = bytes.ReplaceAll(buf, []byte(">"), []byte("&gt;"))

				// set request Body
				r.Body = io.NopCloser(bytes.NewBuffer(buf))
			}

			next.ServeHTTP(w, r)
		})
	}
}
