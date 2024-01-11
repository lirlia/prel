package middleware_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"prel/pkg/middleware"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sanitizer Middleware", func() {
	var (
		req       *http.Request
		rr        *httptest.ResponseRecorder
		sanitizer middleware.Middleware
		handler   http.Handler
	)

	Context("when request method is POST, PUT, or PATCH", func() {
		BeforeEach(func() {
			sanitizer = middleware.Sanitizer()
			handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				msg := "hello&lt;script&gt;alert('xss')&lt;/script&gt;"
				Expect(r.Header.Get("Content-Length")).To(Equal(strconv.Itoa(len(msg))))

				body := r.Body
				defer body.Close()
				buf := new(bytes.Buffer)
				buf.ReadFrom(body)
				bodyString := buf.String()
				Expect(bodyString).To(Equal(msg))
			})

			rr = httptest.NewRecorder()
		})

		It("should sanitize request body", func() {
			msg := "hello<script>alert('xss')</script>"

			By("POST")
			req = httptest.NewRequest("POST", "/test", bytes.NewBufferString(msg))
			sanitizer(handler).ServeHTTP(rr, req)

			By("PUT")
			req = httptest.NewRequest("PUT", "/test", bytes.NewBufferString(msg))
			sanitizer(handler).ServeHTTP(rr, req)

			By("PATCH")
			req = httptest.NewRequest("PATCH", "/test", bytes.NewBufferString(msg))
			sanitizer(handler).ServeHTTP(rr, req)
		})
	})

	Context("when request not have sanitizable message in body", func() {
		BeforeEach(func() {
			sanitizer = middleware.Sanitizer()
			handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				Expect(r.Header.Get("Content-Length")).To(Equal("5"))

				body := r.Body
				defer body.Close()
				buf := new(bytes.Buffer)
				buf.ReadFrom(body)
				bodyString := buf.String()
				Expect(bodyString).To(Equal("hello"))
			})

			rr = httptest.NewRecorder()
		})

		It("should not sanitize request body", func() {
			By("POST")
			req = httptest.NewRequest("POST", "/test", bytes.NewBufferString("hello"))
			sanitizer(handler).ServeHTTP(rr, req)

			By("PUT")
			req = httptest.NewRequest("PUT", "/test", bytes.NewBufferString("hello"))
			sanitizer(handler).ServeHTTP(rr, req)

			By("PATCH")
			req = httptest.NewRequest("PATCH", "/test", bytes.NewBufferString("hello"))
			sanitizer(handler).ServeHTTP(rr, req)
		})
	})

	Context("when request header has encoding gzip", func() {
		BeforeEach(func() {
			sanitizer = middleware.Sanitizer()
			handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				body := r.Body
				defer body.Close()
				buf := new(bytes.Buffer)
				buf.ReadFrom(body)
				bodyString := buf.String()
				Expect(bodyString).To(Equal("hello"))
			})

			rr = httptest.NewRecorder()
		})

		It("should return error", func() {
			req = httptest.NewRequest("POST", "/test", bytes.NewBufferString("hello"))
			req.Header.Set("Content-Encoding", "gzip")
			sanitizer(handler).ServeHTTP(rr, req)
			Expect(rr.Code).To(Equal(http.StatusBadRequest))
			Expect(rr.Body.String()).To(Equal("Bad Request: gzip is not supported\n"))
		})
	})

	Context("when request body is too large", func() {
		var msg = "hello"
		BeforeEach(func() {
			sanitizer = middleware.Sanitizer()
			middleware.SetMaxBodySize(int64(len(msg) - 1))

			handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
			rr = httptest.NewRecorder()
		})

		It("should return error", func() {
			req = httptest.NewRequest("POST", "/test", bytes.NewBufferString(msg))
			sanitizer(handler).ServeHTTP(rr, req)
			Expect(rr.Code).To(Equal(http.StatusBadRequest))
			Expect(rr.Body.String()).To(Equal("Bad Request: Body size is too large\n"))
		})
	})
})
