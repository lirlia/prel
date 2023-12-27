package middleware_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	api "prel/api/prel_api"
	"prel/internal/gateway/repository"
	"prel/internal/model"
	"prel/pkg/custom_error"
	"prel/pkg/middleware"
	"prel/test/testutil"
	"time"

	"github.com/cockroachdb/errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Session Middleware", func() {
	var (
		req       *http.Request
		err       error
		rr        *httptest.ResponseRecorder
		ctx       context.Context
		mw        middleware.Middleware
		handler   http.Handler
		validUser *model.User
		userRepo  = repository.NewUserRepository()
	)

	Describe("session middleware", func() {
		BeforeEach(func() {
			ctx = context.Background()
			testutil.SetupTestDB()

			validUser = testutil.NewTestUser()
			Expect(userRepo.Save(ctx, validUser)).To(Succeed())

			req, err = http.NewRequest("GET", "/test", nil)
			Expect(err).NotTo(HaveOccurred())

			rr = httptest.NewRecorder()
			mw = middleware.Session()
		})

		Context("when session is expired", func() {
			It("should user is nil", func() {
				testutil.NewTestHelper().Travel(validUser.SessionExpiredAt().Add(1*time.Hour), func() {
					req.AddCookie(&http.Cookie{Name: "token", Value: string(validUser.SessionID())})

					handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						user := model.GetUser(r.Context())
						Expect(user).To(BeNil())
					})
					mw(handler).ServeHTTP(rr, req)
					Expect(rr.Code).To(Equal(http.StatusOK))
				})
			})
		})

		Context("when user is not found", func() {
			It("should user is nil", func() {
				newUser := testutil.NewTestUser()
				req.AddCookie(&http.Cookie{Name: "token", Value: string(newUser.SessionID())})

				handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					user := model.GetUser(r.Context())
					Expect(user).To(BeNil())
				})
				mw(handler).ServeHTTP(rr, req)
				Expect(rr.Code).To(Equal(http.StatusOK))
			})
		})
	})

	Describe("HandleCookieAuth", func() {

		BeforeEach(func() {
			ctx = context.Background()
			ctx = model.SetClock(ctx, &model.RealClock{})
			testutil.SetupTestDB()

			validUser = testutil.NewTestUser()
			Expect(userRepo.Save(ctx, validUser)).To(Succeed())

			req, err = http.NewRequest("GET", "/test", nil)
			Expect(err).NotTo(HaveOccurred())

			rr = httptest.NewRecorder()
			mw = middleware.Session()
		})

		Context("when session is valid", func() {
			It("should add user to context for valid session", func() {
				cookie := api.CookieAuth{
					APIKey: string(validUser.SessionID()),
				}

				ctx := model.SetUser(ctx, validUser)
				sessionValidator := middleware.NewSessionValidator()
				_, err = sessionValidator.HandleCookieAuth(ctx, "User", cookie)
				Expect(err).To(BeNil())
			})
		})

		Context("when user is not found", func() {
			It("should return error", func() {
				cookie := api.CookieAuth{
					APIKey: "invalid-session",
				}

				sessionValidator := middleware.NewSessionValidator()
				_, err = sessionValidator.HandleCookieAuth(ctx, "User", cookie)
				Expect(err).To(HaveOccurred())
				Expect(errors.GetAllDetails(err)[0]).To(Equal(string(custom_error.InvalidArgument)))
			})
		})

		Context("when user is not available", func() {
			It("should return error", func() {
				cookie := api.CookieAuth{
					APIKey: string(validUser.SessionID()),
				}

				validUser.SetAvailable(false)
				Expect(userRepo.Save(ctx, validUser)).To(Succeed())

				ctx := model.SetUser(ctx, validUser)
				sessionValidator := middleware.NewSessionValidator()
				_, err = sessionValidator.HandleCookieAuth(ctx, "User", cookie)
				Expect(err).To(HaveOccurred())
				Expect(errors.GetAllDetails(err)[0]).To(Equal(string(custom_error.UnavailableUser)))
			})
		})

		Context("when user is not admin and access /admin", func() {
			It("should return error", func() {
				cookie := api.CookieAuth{
					APIKey: string(validUser.SessionID()),
				}

				ctx := model.SetUser(ctx, validUser)
				sessionValidator := middleware.NewSessionValidator()
				_, err = sessionValidator.HandleCookieAuth(ctx, "AdminUser", cookie)
				Expect(err).To(HaveOccurred())
				Expect(errors.GetAllDetails(err)[0]).To(Equal(string(custom_error.OnlyAdmin)))
			})
		})
	})
})
