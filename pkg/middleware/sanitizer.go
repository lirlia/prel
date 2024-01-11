package middleware

import (
	"bytes"
	"io"
	"net/http"
	"strconv"

	"github.com/cockroachdb/errors"
)

var maxBodySize = int64(1024 * 1024) // Max body size set to 1MiB

func SetMaxBodySize(size int64) {
	maxBodySize = size
}

func Sanitizer() Middleware {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.Method == http.MethodPost ||
				r.Method == http.MethodPut ||
				r.Method == http.MethodPatch {

				// if request body is gzip
				if r.Header.Get("Content-Encoding") == "gzip" {
					http.Error(w, "Bad Request: gzip is not supported", http.StatusBadRequest)
					return
				}

				// limit request Body size
				r.Body = http.MaxBytesReader(w, r.Body, maxBodySize)

				// read request Body
				buf, err := io.ReadAll(r.Body)
				if err != nil {
					if errors.Is(err, &http.MaxBytesError{}) {
						http.Error(w, "Bad Request: Body size is too large", http.StatusBadRequest)
						return
					}
					http.Error(w, "Bad Request: Failed to read body", http.StatusBadRequest)
					return
				}
				buf = bytes.ReplaceAll(buf, []byte("<"), []byte("&lt;"))
				buf = bytes.ReplaceAll(buf, []byte(">"), []byte("&gt;"))

				// set request Body
				r.Body = io.NopCloser(bytes.NewBuffer(buf))

				r.Header.Set("Content-Length", strconv.Itoa(len(buf)))
			}

			next.ServeHTTP(w, r)
		})
	}
}
