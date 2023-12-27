package integration_test

import (
	"context"
	api "prel/api/prel_api"
	"prel/test/testutil"

	google_cloud_mock "prel/internal/gateway/google_cloud/mock"
	"prel/internal/model"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
	"google.golang.org/api/iam/v1"
)

var _ = Describe("Iam Role Filtering", func() {

	var (
		helper *testutil.TestHelper
		ctx    context.Context
	)
	BeforeEach(func() {
		helper = testutil.NewTestHelper()
		helper.StartServer()
		ctx = context.Background()
	})

	AfterEach(func() {
		helper.Close()
	})

	Describe("GetIamRole", func() {
		BeforeEach(func() {
			call := google_cloud_mock.NewMockRolesQueryGrantableRolesCall(helper.Ctrl)
			call.EXPECT().Do().Return(&iam.QueryGrantableRolesResponse{
				Roles: []*iam.Role{
					{Name: "roles/spanner.viewer"},
					{Name: "roles/spanner.admin"},
					{Name: "roles/bigquery.admin"},
					{Name: "roles/logging.viewer"},
					{Name: "roles/owner"},
					{Name: "roles/editor"},
					{Name: "roles/viewer"},
				}}, nil)

			helper.GoogleCloudIamSrv.EXPECT().QueryGrantableRoles(gomock.Any()).Return(call)
		})

		Context("when call get roles", func() {
			It("should return roles without basic role", func() {
				res, err := helper.ApiClient.APIIamRolesGet(ctx, api.APIIamRolesGetParams{
					ProjectID: "test-project",
				})

				Expect(err).NotTo(HaveOccurred())
				Expect(res).NotTo(BeNil())
				Expect(res.(*api.APIIamRolesGetOK)).NotTo(BeNil())
				Expect(res.(*api.APIIamRolesGetOK).IamRoles).To(HaveLen(4)) // 4 roles without basic role
			})
		})

		Context("when request api call", func() {
			BeforeEach(func() {
				helper.User.SetRole(model.UserRoleAdmin)
				Expect(helper.UserRepo.Save(ctx, helper.User)).To(Succeed())
				for _, pattern := range []string{"spanner", "bigquery"} {
					_, err := helper.ApiClient.APIIamRoleFilteringRulesPost(ctx, &api.APIIamRoleFilteringRulesPostReq{
						Pattern: pattern,
					})
					Expect(err).NotTo(HaveOccurred())
				}
			})
			It("should return filtered roles", func() {
				res, err := helper.ApiClient.APIIamRolesGet(ctx, api.APIIamRolesGetParams{
					ProjectID: "test-project",
				})

				Expect(err).NotTo(HaveOccurred())
				Expect(res).NotTo(BeNil())
				Expect(res.(*api.APIIamRolesGetOK)).NotTo(BeNil())
				Expect(res.(*api.APIIamRolesGetOK).IamRoles).To(HaveLen(3)) // 3 roles without basic role and filtered roles
			})
		})
	})
})
