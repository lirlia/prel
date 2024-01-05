package google_cloud_test

import (
	"context"
	"prel/internal/gateway/google_cloud"
	google_cloud_mock "prel/internal/gateway/google_cloud/mock"
	"prel/internal/model"
	"prel/test/testutil"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
	"google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/iam/v1"
)

var _ = Describe("Google Cloud", func() {
	var (
		client                 *google_cloud.Client
		ctx                    context.Context
		projectID              string
		ctrl                   *gomock.Controller
		resourceManagerService *google_cloud_mock.MockResourceManagerService
		iamService             *google_cloud_mock.MockIamService
		user                   *model.User
	)

	BeforeEach(func() {
		ctx = context.Background()
		projectID = "test-project"
		ctrl = gomock.NewController(GinkgoT())
		resourceManagerService = google_cloud_mock.NewMockResourceManagerService(ctrl)
		iamService = google_cloud_mock.NewMockIamService(ctrl)
		client = google_cloud.NewClient(ctx, projectID, resourceManagerService, iamService)
		user = testutil.NewTestUser()
	})

	Describe("SettableBinding", func() {
		var bindings []*cloudresourcemanager.Binding
		BeforeEach(func() {
			bindings = []*cloudresourcemanager.Binding{
				// false
				{Role: "roles/spanner.admin"},
				{Role: "roles/spanner.admin", Condition: &cloudresourcemanager.Expr{Expression: "test"}},
				{Role: "roles/spanner.admin", Condition: &cloudresourcemanager.Expr{Title: "test"}},
				{Role: "roles/spanner.admin", Condition: &cloudresourcemanager.Expr{Location: "test"}},
				// true
				{Role: "roles/spanner.admin", Condition: &cloudresourcemanager.Expr{Title: "generated_by_prel_test"}},
			}
		})

		Context("when set binding", func() {
			It("should be set binding", func() {
				b := google_cloud.SettableBinding(bindings, "roles/spanner.admin")
				Expect(b.Role).To(Equal("roles/spanner.admin"))
				Expect(b.Condition.Title).To(Equal("generated_by_prel_test"))
			})
		})
	})

	Describe("ExcludeBasicRole", func() {
		Context("when get iam roles", func() {
			It("should be excluded basic roles(viewer/editor/owner)", func() {
				call := google_cloud_mock.NewMockRolesQueryGrantableRolesCall(ctrl)
				call.EXPECT().Do().Return(&iam.QueryGrantableRolesResponse{
					Roles: []*iam.Role{
						{Name: "roles/viewer"},
						{Name: "roles/editor"},
						{Name: "roles/owner"},
						{Name: "roles/test"},
					}}, nil)
				iamService.EXPECT().QueryGrantableRoles(gomock.Any()).Return(call)

				roles, err := client.GetIamRoles(time.Now(), projectID, user)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(roles).To(HaveLen(1))
				Expect(roles[0].Name).To(Equal("roles/test"))
			})
		})
	})

	Describe("ExcludeRoleByPrincipalType", func() {
		Context("when get iam roles", func() {
			It("should be excluded service account role", func() {
				call := google_cloud_mock.NewMockRolesQueryGrantableRolesCall(ctrl)
				call.EXPECT().Do().Return(&iam.QueryGrantableRolesResponse{
					Roles: []*iam.Role{
						{Name: "roles/aiplatform.customCodeServiceAgent"},
						{Name: "roles/anthospolicycontroller.serviceAgent"},
						{Name: "roles/test"},
					}}, nil)
				iamService.EXPECT().QueryGrantableRoles(gomock.Any()).Return(call)

				roles, err := client.GetIamRoles(time.Now(), projectID, user)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(roles).To(HaveLen(1))
				Expect(roles[0].Name).To(Equal("roles/test"))
			})
		})
	})
})
