package integration_test

import (
	"context"
	api "prel/api/prel_api"
	"prel/internal/model"
	"prel/test/testutil"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Setting", func() {

	Describe("Setting", func() {
		var (
			helper *testutil.TestHelper
			client api.Invoker
			ctx    context.Context
		)
		BeforeEach(func() {
			helper = testutil.NewTestHelper()
			helper.StartServer()
			client = helper.ApiClient
			Expect(helper.ChangeRole(helper.User, model.UserRoleAdmin)).To(Succeed())
			ctx = context.Background()
		})

		Context("When the notificationMessageForRequest is set", func() {
			It("returns 204", func() {
				res, err := client.APISettingsPatch(ctx, &api.APISettingsPatchReq{
					NotificationMessageForRequest: api.OptString{
						Value: "this is test",
						Set:   true,
					},
				})
				Expect(err).To(BeNil())
				Expect(res).ToNot(BeNil())

				setting, err := helper.SettingRepo.Find(ctx)
				Expect(err).To(BeNil())
				Expect(setting.NotificationMessageForRequest()).To(Equal("this is test"))
				Expect(setting.NotificationMessageForJudge()).To(Equal(""))
			})
		})

		Context("When the notificationMessageForJudge is set", func() {
			It("returns 204", func() {
				res, err := client.APISettingsPatch(ctx, &api.APISettingsPatchReq{
					NotificationMessageForJudge: api.OptString{
						Value: "this is test",
						Set:   true,
					},
				})
				Expect(err).To(BeNil())
				Expect(res).ToNot(BeNil())

				setting, err := helper.SettingRepo.Find(ctx)
				Expect(err).To(BeNil())
				Expect(setting.NotificationMessageForRequest()).To(Equal(""))
				Expect(setting.NotificationMessageForJudge()).To(Equal("this is test"))
			})
		})
	})
})
