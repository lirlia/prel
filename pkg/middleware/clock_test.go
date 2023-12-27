package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"prel/internal/model"
	"prel/pkg/middleware"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Clock Middleware", func() {
	It("should set the clock in the request context", func() {
		req, err := http.NewRequest("GET", "/", nil)
		Expect(err).NotTo(HaveOccurred())
		rr := httptest.NewRecorder()

		now := time.Now()
		testClock := &model.FrozenClock{Time: now}

		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			clock := model.GetClock(ctx)
			Expect(clock.Now()).To(Equal(testClock.Now()))
		})

		middleware := middleware.Clock(testClock)
		middleware(handler).ServeHTTP(rr, req)
	})
})
