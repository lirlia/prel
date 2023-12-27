package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"prel/pkg/logger"
	"prel/pkg/middleware"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Logger Middleware", func() {
	It("should log the request and response information", func() {
		req, err := http.NewRequest("GET", "/test", nil)
		Expect(err).NotTo(HaveOccurred())
		rr := httptest.NewRecorder()

		reqBuf := gbytes.NewBuffer()
		appBuf := gbytes.NewBuffer()
		requestBuilder := logger.NewBuilder().WithOutput(reqBuf)
		appBuilder := logger.NewBuilder().WithOutput(appBuf)

		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()
			logger.Get(ctx).Info("this is an app test")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		})

		middleware := middleware.Logger(requestBuilder, appBuilder)
		middleware(handler).ServeHTTP(rr, req)

		reqLog := string(reqBuf.Contents())
		Expect(reqLog).To(
			MatchRegexp("{\"time\":\".*\",\"level\":\"INFO\",\"msg\":\"/test\",\"httpRequest\":{\"requestMethod\":\"GET\",\"requestUrl\":\"/test\",\"remoteIp\":\"\",\"status\":200,\"latency\":\".*\"}}\n"))

		appLog := string(appBuf.Contents())
		Expect(appLog).To(
			MatchRegexp("{\"time\":\".*\",\"level\":\"INFO\",\"msg\":\"this is an app test\",\"httpRequest\":{\"requestMethod\":\"GET\",\"requestUrl\":\"/test\",\"remoteIp\":\"\"}}\n"))

	})
})
