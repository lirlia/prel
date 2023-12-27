package integration_test

import (
	"context"
	"io"
	api "prel/api/prel_api"
	"prel/internal/model"
	"prel/test/testutil"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Invitation", func() {

	Describe("Invite", func() {
		var (
			helper        *testutil.TestHelper
			inviterClient api.Invoker
			ctx           context.Context
		)
		BeforeEach(func() {
			helper = testutil.NewTestHelper()
			helper.StartServer()
			inviterClient = helper.ApiClient
			Expect(helper.ChangeRole(helper.User, model.UserRoleAdmin)).To(Succeed())
			ctx = context.Background()
		})

		Context("When the request is valid", func() {
			It("returns 200", func() {
				res, err := inviterClient.APIInvitationsPost(ctx, &api.APIInvitationsPostReq{
					Email: "test@test",
					Role:  api.UserRoleAdmin,
				})
				Expect(err).To(BeNil())
				Expect(res).ToNot(BeNil())
				Expect(res.(*api.APIInvitationsPostNoContent)).ToNot(BeNil())

				helper.InvitationRepo.FindByInviteeMailsAndExpiredAt(ctx, []string{"test@test"}, helper.Clock.Now())
			})
		})

		Context("When email is already invited(before expires)", func() {
			It("returns 400", func() {
				By("first invitation")
				res, err := inviterClient.APIInvitationsPost(ctx, &api.APIInvitationsPostReq{
					Email: "test@test",
					Role:  api.UserRoleAdmin,
				})
				Expect(err).To(BeNil())
				Expect(res).ToNot(BeNil())
				Expect(res.(*api.APIInvitationsPostNoContent)).ToNot(BeNil())

				By("second invitation")
				res, err = inviterClient.APIInvitationsPost(ctx, &api.APIInvitationsPostReq{
					Email: "test@test",
					Role:  api.UserRoleAdmin,
				})
				Expect(err).To(BeNil())
				Expect(res.(*api.BadRequest)).ToNot(BeNil())
				b, err := io.ReadAll(res.(*api.BadRequest).Data)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(b)).To(ContainSubstring("Already Invited"))
			})
		})

		Context("When email is already invited(after expires)", func() {
			It("returns 204", func() {
				By("first invitation")
				res, err := inviterClient.APIInvitationsPost(ctx, &api.APIInvitationsPostReq{
					Email: "test@test",
					Role:  api.UserRoleAdmin,
				})
				Expect(err).To(BeNil())
				Expect(res).ToNot(BeNil())
				Expect(res.(*api.APIInvitationsPostNoContent)).ToNot(BeNil())

				By("second invitation")
				helper.Travel(helper.Clock.Now().Add(model.InvitationExpiredAt+time.Hour), func() {

					inviterClient, err = helper.UpdateSession(inviterClient.(*testutil.ApiClient), helper.Clock.Now().Add(time.Hour))
					Expect(err).To(BeNil())

					res, err = inviterClient.APIInvitationsPost(ctx, &api.APIInvitationsPostReq{
						Email: "test@test",
						Role:  api.UserRoleAdmin,
					})
					Expect(err).To(BeNil())
					Expect(res).ToNot(BeNil())
					Expect(res.(*api.APIInvitationsPostNoContent)).ToNot(BeNil())

					helper.InvitationRepo.FindByInviteeMailsAndExpiredAt(ctx, []string{"test@test"}, helper.Clock.Now())
				})
			})
		})

		Context("When email is already registered in user table", func() {
			It("returns 400", func() {
				res, err := inviterClient.APIInvitationsPost(ctx, &api.APIInvitationsPostReq{
					Email: helper.User.Email(),
					Role:  api.UserRoleAdmin,
				})
				Expect(err).To(BeNil())
				Expect(res.(*api.BadRequest)).ToNot(BeNil())

				b, err := io.ReadAll(res.(*api.BadRequest).Data)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(b)).To(ContainSubstring("Already Registered"))
			})
		})

		Context("When user is not admin", func() {
			It("returns 403", func() {
				Expect(helper.ChangeRole(helper.User, model.UserRoleJudger)).To(Succeed())

				res, err := inviterClient.APIInvitationsPost(ctx, &api.APIInvitationsPostReq{
					Email: "test@test",
					Role:  api.UserRoleAdmin,
				})
				Expect(err).To(BeNil())
				Expect(res).ToNot(BeNil())

				b, err := io.ReadAll(res.(*api.Forbidden).Data)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(b)).To(ContainSubstring("Only Admin"))
			})
		})
	})
})
