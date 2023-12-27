package integration_test

import (
	"context"
	"io"
	api "prel/api/prel_api"
	"prel/internal/model"
	"prel/test/testutil"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {

	Describe("Get", func() {
		var (
			count  = 50
			helper *testutil.TestHelper
			ctx    context.Context
		)
		BeforeEach(func() {
			helper = testutil.NewTestHelper()
			helper.StartServer()
			ctx = context.Background()
			for i := 0; i < count; i++ {
				req := testutil.NewTestUser()
				Expect(helper.UserRepo.Save(ctx, req)).To(Succeed())
			}
		})

		Context("when user is valid", func() {
			It("should return expected users", func() {
				Expect(helper.ChangeRole(helper.User, model.UserRoleAdmin)).To(Succeed())
				res, err := helper.ApiClient.APIUsersGet(ctx, api.APIUsersGetParams{
					PageID: 1,
					Size:   api.PageSize25,
				})

				Expect(err).NotTo(HaveOccurred())
				Expect(res).NotTo(BeNil())
				Expect(res.(*api.APIUsersGetOK)).NotTo(BeNil())
				Expect(res.(*api.APIUsersGetOK).TotalPage).To(Equal(3)) // own user + 50 users
				Expect(res.(*api.APIUsersGetOK).CurrentPage).To(Equal(1))
				Expect(res.(*api.APIUsersGetOK).Users).To(HaveLen(25))

				By("+50 Users")
				for i := 0; i < count; i++ {
					req := testutil.NewTestUser()
					Expect(helper.UserRepo.Save(ctx, req)).To(Succeed())
				}
				res, err = helper.ApiClient.APIUsersGet(ctx, api.APIUsersGetParams{
					PageID: 2,
					Size:   api.PageSize50,
				})

				Expect(err).NotTo(HaveOccurred())
				Expect(res).NotTo(BeNil())
				Expect(res.(*api.APIUsersGetOK)).NotTo(BeNil())
				Expect(res.(*api.APIUsersGetOK).TotalPage).To(Equal(3)) // own user + 100 users
				Expect(res.(*api.APIUsersGetOK).CurrentPage).To(Equal(2))
				Expect(res.(*api.APIUsersGetOK).Users).To(HaveLen(50))
			})
		})

		Context("when user is not admin", func() {
			It("should return error", func() {
				res, err := helper.ApiClient.APIUsersGet(ctx, api.APIUsersGetParams{
					PageID: 1,
					Size:   api.PageSize25,
				})

				Expect(err).NotTo(HaveOccurred())
				Expect(res).NotTo(BeNil())
				Expect(res.(*api.Forbidden)).NotTo(BeNil())
			})
		})
	})

	Describe("Update", func() {
		var (
			targetUser  *model.User
			admin       *model.User
			adminClient api.Invoker
			ctx         context.Context
			helper      *testutil.TestHelper
		)

		BeforeEach(func() {
			helper = testutil.NewTestHelper()
			helper.StartServer()
			ctx = context.Background()
			var err error
			targetUser = helper.User
			admin, adminClient, err = helper.NewUserAndClient(testutil.WithRole(model.UserRoleAdmin))
			Expect(err).To(BeNil())
		})

		Context("when update", func() {
			It("should update user", func() {
				res, err := adminClient.APIUsersUserIDPatch(ctx, &api.APIUsersUserIDPatchReq{
					IsAvailable: false,
					Role:        api.UserRoleAdmin,
				}, api.APIUsersUserIDPatchParams{
					UserID: string(targetUser.ID()),
				})

				Expect(err).To(BeNil())
				Expect(res).ToNot(BeNil())
				Expect(res.(*api.APIUsersUserIDPatchNoContent)).ToNot(BeNil())

				user, err := helper.UserRepo.FindByID(ctx, targetUser.ID())
				Expect(err).To(BeNil())

				Expect(user.IsAvailable()).To(BeFalse())
				Expect(user.Role()).To(Equal(model.UserRoleAdmin))
			})
		})

		Context("when update by not admin user", func() {
			It("should return error", func() {
				Expect(helper.ChangeRole(admin, model.UserRoleJudger)).To(Succeed())

				res, err := adminClient.APIUsersUserIDPatch(ctx, &api.APIUsersUserIDPatchReq{
					IsAvailable: false,
					Role:        api.UserRoleAdmin,
				}, api.APIUsersUserIDPatchParams{
					UserID: string(targetUser.ID()),
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
