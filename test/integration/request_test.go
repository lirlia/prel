package integration_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	api "prel/api/prel_api"
	"prel/internal/gateway/google_cloud"
	google_cloud_mock "prel/internal/gateway/google_cloud/mock"
	"prel/internal/model"
	"prel/test/testutil"
	"strings"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
	"google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/iam/v1"
)

var _ = Describe("Request", func() {

	var (
		helper *testutil.TestHelper
		ctx    context.Context
	)
	BeforeEach(func() {
		helper = testutil.NewTestHelper()
		helper.StartServer()
		ctx = context.Background()

		call := google_cloud_mock.NewMockRolesQueryGrantableRolesCall(helper.Ctrl)
		call.EXPECT().Do().Return(&iam.QueryGrantableRolesResponse{
			Roles: []*iam.Role{
				{Name: "iam-role-a"},
				{Name: "iam-role-b"},
			}}, nil).AnyTimes()

		helper.GoogleCloudIamSrv.EXPECT().QueryGrantableRoles(gomock.Any()).Return(call).AnyTimes()
	})

	AfterEach(func() {
		helper.Close()
	})

	Describe("Create", func() {
		Context("when request is valid", func() {
			It("should create expected request", func() {
				// specific time
				now := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
				helper.Clock.Set(now)

				helper.StartNotificationServer(testutil.WithServerHandler(testutil.CheckHandler(func(w http.ResponseWriter, r *http.Request) {
					msg, err := io.ReadAll(r.Body)
					Expect(err).NotTo(HaveOccurred())
					Expect(string(msg)).To(ContainSubstring(helper.User.Email()))
					Expect(string(msg)).To(ContainSubstring(fmt.Sprintf("http://%s:%s", helper.Config.Address, helper.Config.Port)))
					Expect(string(msg)).To(ContainSubstring("project-id"))
					Expect(string(msg)).To(ContainSubstring("iam-role-a"))
					Expect(string(msg)).To(ContainSubstring("iam-role-b"))
					Expect(string(msg)).To(ContainSubstring("10 minutes"))
					Expect(string(msg)).To(ContainSubstring("this is a test"))
					Expect(string(msg)).To(ContainSubstring(now.Add(24 * time.Hour).Format("2006/01/02 15:04:05 MST")))
				}, http.StatusOK)))

				res, err := helper.ApiClient.APIRequestsPost(ctx, &api.APIRequestsPostReq{
					ProjectID: "project-id",
					IamRoles: []string{
						"iam-role-a",
						"iam-role-b",
					},
					Period: api.APIRequestsPostReqPeriod10,
					Reason: "this is a test",
				})

				Expect(err).NotTo(HaveOccurred())
				Expect(res).NotTo(BeNil())
				Expect(res.(*api.APIRequestsPostOK)).NotTo(BeNil())
				requestID := res.(*api.APIRequestsPostOK).GetRequestID()
				req, err := helper.RequestRepo.FindByID(ctx, requestID)
				Expect(err).NotTo(HaveOccurred())
				Expect(req).NotTo(BeNil())
				Expect(req.ID()).To(Equal(requestID))
			})
		})

		Context("when notification server is down", func() {
			It("should create expected request", func() {
				helper.StartNotificationServer(testutil.WithServerHandler(testutil.CheckHandler(func(w http.ResponseWriter, r *http.Request) {
				}, http.StatusInternalServerError)))

				res, err := helper.ApiClient.APIRequestsPost(ctx, &api.APIRequestsPostReq{
					ProjectID: "project-id",
					IamRoles: []string{
						"iam-role-a",
						"iam-role-b",
					},
					Period: api.APIRequestsPostReqPeriod10,
					Reason: "this is a test",
				})

				Expect(err).NotTo(HaveOccurred())
				Expect(res).NotTo(BeNil())
				Expect(res.(*api.APIRequestsPostOK)).NotTo(BeNil())
				requestID := res.(*api.APIRequestsPostOK).GetRequestID()

				req, err := helper.RequestRepo.FindByID(ctx, requestID)
				Expect(err).NotTo(HaveOccurred())
				Expect(req).NotTo(BeNil())
				Expect(req.ID()).To(Equal(requestID))
			})
		})

		Context("when request' role is not permitted", func() {
			It("should return error", func() {
				res, err := helper.ApiClient.APIRequestsPost(ctx, &api.APIRequestsPostReq{
					ProjectID: "project-id",
					IamRoles: []string{
						"roles/bigquery.admin",
					},
					Period: api.APIRequestsPostReqPeriod10,
					Reason: "this is a test",
				})

				Expect(err).NotTo(HaveOccurred())
				Expect(res).NotTo(BeNil())
				Expect(res.(*api.BadRequest)).NotTo(BeNil())

				b, err := io.ReadAll(res.(*api.BadRequest).Data)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(b)).To(ContainSubstring("Invalid Argument"))
			})
		})
	})

	Describe("Judge", func() {
		var (
			requestID       string
			judger          *model.User
			judgerClient    *testutil.ApiClient
			iamPolicyGetter *google_cloud_mock.MockIamPolicyGetter
			iamPolicySetter *google_cloud_mock.MockIamPolicySetter
		)
		BeforeEach(func() {
			res, err := helper.ApiClient.APIRequestsPost(ctx, &api.APIRequestsPostReq{
				ProjectID: "project-id",
				IamRoles: []string{
					"iam-role-a",
				},
				Period: api.APIRequestsPostReqPeriod10,
				Reason: "this is a test",
			})

			Expect(err).NotTo(HaveOccurred())
			Expect(res).NotTo(BeNil())
			Expect(res.(*api.APIRequestsPostOK)).NotTo(BeNil())
			requestID = res.(*api.APIRequestsPostOK).GetRequestID()

			judger, judgerClient, err = helper.NewUserAndClient(testutil.WithRole(model.UserRoleJudger))
			Expect(err).NotTo(HaveOccurred())

			iamPolicyGetter = google_cloud_mock.NewMockIamPolicyGetter(helper.Ctrl)
			iamPolicySetter = google_cloud_mock.NewMockIamPolicySetter(helper.Ctrl)
		})

		Context("approve", func() {
			Context("when request is valid", func() {
				BeforeEach(func() {
					// specific time
					now := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
					helper.Clock.Set(now)
					helper.StartNotificationServer(testutil.WithServerHandler(testutil.CheckHandler(func(w http.ResponseWriter, r *http.Request) {
						msg, err := io.ReadAll(r.Body)
						Expect(err).NotTo(HaveOccurred())
						if strings.Contains(string(msg), "New request") {
							return
						}
						Expect(string(msg)).To(ContainSubstring(helper.User.Email()))
						Expect(string(msg)).To(ContainSubstring(fmt.Sprintf("http://%s:%s", helper.Config.Address, helper.Config.Port)))
						Expect(string(msg)).To(ContainSubstring("project-id"))
						Expect(string(msg)).To(ContainSubstring("iam-role-a"))
						Expect(string(msg)).To(ContainSubstring(now.Add(10 * time.Minute).Format("2006/01/02 15:04:05 MST")))
						Expect(string(msg)).To(ContainSubstring("this is a test"))
					}, http.StatusOK)))
				})
				It("should update expected request", func() {
					iamPolicyGetter.EXPECT().Do().Return(&cloudresourcemanager.Policy{
						// already has no condition binding
						Bindings: []*cloudresourcemanager.Binding{
							{
								Members: []string{
									"dummy-user",
								},
								Role: "iam-role-a",
							},
							{
								Members: []string{
									"dummy-user",
								},
								Role: "iam-role-a",
								Condition: &cloudresourcemanager.Expr{
									Title: "dummy-title",
								},
							},
						},
					}, nil)
					helper.GoogleCloudResourceManagerSrv.EXPECT().
						GetIamPolicy("project-id", gomock.Any()).
						Return(iamPolicyGetter)

					iamPolicySetter.EXPECT().Do().Return(&cloudresourcemanager.Policy{}, nil)
					helper.GoogleCloudResourceManagerSrv.EXPECT().
						SetIamPolicy("project-id", gomock.Any()).
						DoAndReturn(func(projectID string, req *cloudresourcemanager.SetIamPolicyRequest) google_cloud.IamPolicySetter {
							// 3 bindings
							//   1: dummy-user and no condition binding
							//   2: dummy-user and condition binding (not prel)
							//   3: requester and condition binding
							// if Google Cloud Iam Policy already have no condition binding, prel don't change it.
							// if Google Cloud Iam Policy already have condition binding which didn't changed by prel, prel don't change it.
							// prel only change condition binding which generated by prel.
							Expect(req.Policy.Bindings).To(HaveLen(3))
							Expect(req.Policy.Bindings[2].Role).To(Equal("iam-role-a"))
							Expect(req.Policy.Bindings[2].Members).To(HaveLen(1))
							Expect(req.Policy.Bindings[2].Members[0]).To(Equal(helper.User.Principal()))
							Expect(req.Policy.Bindings[2].Condition).NotTo(BeNil())
							Expect(req.Policy.Bindings[2].Condition.Title).To(Equal("generated_by_prel_expires_after_2021_01_01_00_10_UTC"))
							Expect(req.Policy.Bindings[2].Condition.Description).To(Equal("Expiring at 2021/01/01 00:10:00 UTC generated by prel"))
							Expect(req.Policy.Bindings[2].Condition.Expression).To(Equal("request.time < timestamp(\"2021-01-01T00:10:00Z\")"))
							return iamPolicySetter
						})

					res, err := judgerClient.APIRequestsRequestIDPatch(ctx, &api.APIRequestsRequestIDPatchReq{
						Status: api.JudgeStatusApprove,
					}, api.APIRequestsRequestIDPatchParams{
						RequestID: requestID,
					})

					Expect(err).NotTo(HaveOccurred())
					Expect(res).NotTo(BeNil())
					Expect(res.(*api.APIRequestsRequestIDPatchNoContent)).NotTo(BeNil())

					req, err := helper.RequestRepo.FindByID(ctx, requestID)
					Expect(err).NotTo(HaveOccurred())
					Expect(req).NotTo(BeNil())
					Expect(req.Status()).To(Equal(model.RequestStatusApproved))
				})
			})

			Context("when judger do not have permission to judge(requester)", func() {
				It("should return error", func() {
					judger.SetRole(model.UserRoleRequester)
					Expect(helper.UserRepo.Save(ctx, judger)).NotTo(HaveOccurred())

					res, err := judgerClient.APIRequestsRequestIDPatch(ctx, &api.APIRequestsRequestIDPatchReq{
						Status: api.JudgeStatusApprove,
					}, api.APIRequestsRequestIDPatchParams{
						RequestID: requestID,
					})

					Expect(err).NotTo(HaveOccurred())
					Expect(res).NotTo(BeNil())

					b, err := io.ReadAll(res.(*api.Forbidden).Data)
					Expect(err).NotTo(HaveOccurred())
					Expect(string(b)).To(ContainSubstring("Not Allowed"))
				})
			})

			Context("when request is already approved", func() {
				It("should return error", func() {
					By("first approve")
					iamPolicyGetter.EXPECT().Do().Return(&cloudresourcemanager.Policy{}, nil)
					helper.GoogleCloudResourceManagerSrv.EXPECT().GetIamPolicy("project-id", gomock.Any()).Return(iamPolicyGetter)

					iamPolicySetter.EXPECT().Do().Return(&cloudresourcemanager.Policy{}, nil)
					helper.GoogleCloudResourceManagerSrv.EXPECT().
						SetIamPolicy("project-id", gomock.Any()).Return(iamPolicySetter)

					res, err := judgerClient.APIRequestsRequestIDPatch(ctx, &api.APIRequestsRequestIDPatchReq{
						Status: api.JudgeStatusApprove,
					}, api.APIRequestsRequestIDPatchParams{
						RequestID: requestID,
					})

					Expect(err).NotTo(HaveOccurred())
					Expect(res).NotTo(BeNil())
					Expect(res.(*api.APIRequestsRequestIDPatchNoContent)).NotTo(BeNil())

					By("second approve")
					res, err = judgerClient.APIRequestsRequestIDPatch(ctx, &api.APIRequestsRequestIDPatchReq{
						Status: api.JudgeStatusApprove,
					}, api.APIRequestsRequestIDPatchParams{
						RequestID: requestID,
					})

					Expect(err).NotTo(HaveOccurred())
					Expect(res).NotTo(BeNil())

					b, err := io.ReadAll(res.(*api.BadRequest).Data)
					Expect(err).NotTo(HaveOccurred())
					Expect(string(b)).To(ContainSubstring("Invalid Argument"))
				})
			})

			Context("when request has invalid role", func() {
				It("should return error", func() {
					newReq := testutil.NewTestRequest(
						testutil.WithRequesterUserID(helper.User.ID()),
						testutil.WithJudgerUserID(judger.ID()),
						testutil.WithIamRoles([]string{"invalid-role"}))
					Expect(helper.RequestRepo.Save(ctx, newReq)).To(Succeed())

					res, err := judgerClient.APIRequestsRequestIDPatch(ctx, &api.APIRequestsRequestIDPatchReq{
						Status: api.JudgeStatusApprove,
					}, api.APIRequestsRequestIDPatchParams{
						RequestID: newReq.ID(),
					})

					Expect(err).NotTo(HaveOccurred())
					Expect(res).NotTo(BeNil())

					b, err := io.ReadAll(res.(*api.BadRequest).Data)
					Expect(err).NotTo(HaveOccurred())
					Expect(string(b)).To(ContainSubstring("Invalid Argument"))
				})
			})
		})

		Context("reject", func() {
			Context("when request is valid", func() {
				It("should update expected request", func() {
					res, err := judgerClient.APIRequestsRequestIDPatch(ctx, &api.APIRequestsRequestIDPatchReq{
						Status: api.JudgeStatusReject,
					}, api.APIRequestsRequestIDPatchParams{
						RequestID: requestID,
					})

					Expect(err).NotTo(HaveOccurred())
					Expect(res).NotTo(BeNil())
					Expect(res.(*api.APIRequestsRequestIDPatchNoContent)).NotTo(BeNil())

					req, err := helper.RequestRepo.FindByID(ctx, requestID)
					Expect(err).NotTo(HaveOccurred())
					Expect(req).NotTo(BeNil())
					Expect(req.Status()).To(Equal(model.RequestStatusRejected))
				})
			})
		})
	})

	Describe("Delete", func() {
		var (
			requestID    string
			judgerClient *testutil.ApiClient
		)
		BeforeEach(func() {
			res, err := helper.ApiClient.APIRequestsPost(ctx, &api.APIRequestsPostReq{
				ProjectID: "project-id",
				IamRoles: []string{
					"iam-role-a",
				},
				Period: api.APIRequestsPostReqPeriod10,
				Reason: "this is a test",
			})

			Expect(err).NotTo(HaveOccurred())
			Expect(res).NotTo(BeNil())
			Expect(res.(*api.APIRequestsPostOK)).NotTo(BeNil())
			requestID = res.(*api.APIRequestsPostOK).GetRequestID()

			_, judgerClient, err = helper.NewUserAndClient(testutil.WithRole(model.UserRoleAdmin))
			Expect(err).NotTo(HaveOccurred())
		})

		Context("when request is valid", func() {
			It("should delete expected request", func() {
				res, err := judgerClient.APIRequestsRequestIDDelete(ctx, api.APIRequestsRequestIDDeleteParams{
					RequestID: requestID,
				})

				Expect(err).NotTo(HaveOccurred())
				Expect(res).NotTo(BeNil())
				Expect(res.(*api.APIRequestsRequestIDDeleteNoContent)).NotTo(BeNil())

				req, err := helper.RequestRepo.FindByID(ctx, requestID)
				Expect(err).To(HaveOccurred())
				Expect(req).To(BeNil())
			})
		})

		Context("when request is already approved", func() {
			It("should return error", func() {
				By("approve")
				iamPolicyGetter := google_cloud_mock.NewMockIamPolicyGetter(helper.Ctrl)
				iamPolicyGetter.EXPECT().Do().Return(&cloudresourcemanager.Policy{}, nil)
				helper.GoogleCloudResourceManagerSrv.EXPECT().GetIamPolicy("project-id", gomock.Any()).Return(iamPolicyGetter)

				iamPolicySetter := google_cloud_mock.NewMockIamPolicySetter(helper.Ctrl)
				iamPolicySetter.EXPECT().Do().Return(&cloudresourcemanager.Policy{}, nil)
				helper.GoogleCloudResourceManagerSrv.EXPECT().
					SetIamPolicy("project-id", gomock.Any()).Return(iamPolicySetter)

				res, err := judgerClient.APIRequestsRequestIDPatch(ctx, &api.APIRequestsRequestIDPatchReq{
					Status: api.JudgeStatusApprove,
				}, api.APIRequestsRequestIDPatchParams{
					RequestID: requestID,
				})

				Expect(err).NotTo(HaveOccurred())
				Expect(res).NotTo(BeNil())
				Expect(res.(*api.APIRequestsRequestIDPatchNoContent)).NotTo(BeNil())

				By("delete")
				deleteRes, err := judgerClient.APIRequestsRequestIDDelete(ctx, api.APIRequestsRequestIDDeleteParams{
					RequestID: requestID,
				})

				Expect(err).NotTo(HaveOccurred())
				Expect(deleteRes).NotTo(BeNil())

				b, err := io.ReadAll(deleteRes.(*api.BadRequest).Data)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(b)).To(ContainSubstring("Invalid Argument"))
			})
		})

		Context("when user is not admin", func() {
			It("should return error", func() {
				_, requesterClient, err := helper.NewUserAndClient(testutil.WithRole(model.UserRoleRequester))
				Expect(err).NotTo(HaveOccurred())

				res, err := requesterClient.APIRequestsRequestIDDelete(ctx, api.APIRequestsRequestIDDeleteParams{
					RequestID: requestID,
				})

				Expect(err).NotTo(HaveOccurred())
				Expect(res).NotTo(BeNil())

				b, err := io.ReadAll(res.(*api.Forbidden).Data)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(b)).To(ContainSubstring("Not Allowed"))
			})
		})
	})

	Describe("GetWithPaging", func() {
		var (
			requestCount = 50
			judger       *model.User
		)

		BeforeEach(func() {
			judger = testutil.NewTestUser(testutil.WithRole(model.UserRoleJudger))
			Expect(helper.UserRepo.Save(ctx, judger)).To(Succeed())

			for i := 0; i < requestCount; i++ {
				user := testutil.NewTestUser()
				req := testutil.NewTestRequest(testutil.WithRequesterUserID(user.ID()), testutil.WithJudgerUserID(judger.ID()))
				Expect(helper.UserRepo.Save(ctx, user)).To(Succeed())
				Expect(helper.RequestRepo.Save(ctx, req)).To(Succeed())
			}
		})

		Context("when request is valid", func() {
			It("should return expected requests", func() {
				Expect(helper.ChangeRole(helper.User, model.UserRoleAdmin)).To(Succeed())
				res, err := helper.ApiClient.APIRequestsGet(ctx, api.APIRequestsGetParams{
					PageID: 1,
					Size:   api.PageSize25,
				})

				Expect(err).NotTo(HaveOccurred())
				Expect(res).NotTo(BeNil())
				Expect(res.(*api.APIRequestsGetOK)).NotTo(BeNil())
				Expect(res.(*api.APIRequestsGetOK).TotalPage).To(Equal(2))
				Expect(res.(*api.APIRequestsGetOK).CurrentPage).To(Equal(1))
				Expect(res.(*api.APIRequestsGetOK).Requests).To(HaveLen(25))

				By("+50 requests")
				for i := 0; i < requestCount; i++ {
					user := testutil.NewTestUser()
					req := testutil.NewTestRequest(testutil.WithRequesterUserID(user.ID()), testutil.WithJudgerUserID(judger.ID()))
					Expect(helper.UserRepo.Save(ctx, user)).To(Succeed())
					Expect(helper.RequestRepo.Save(ctx, req)).To(Succeed())
				}
				res, err = helper.ApiClient.APIRequestsGet(ctx, api.APIRequestsGetParams{
					PageID: 2,
					Size:   api.PageSize50,
				})

				Expect(err).NotTo(HaveOccurred())
				Expect(res).NotTo(BeNil())
				Expect(res.(*api.APIRequestsGetOK)).NotTo(BeNil())
				Expect(res.(*api.APIRequestsGetOK).TotalPage).To(Equal(2))
				Expect(res.(*api.APIRequestsGetOK).CurrentPage).To(Equal(2))
				Expect(res.(*api.APIRequestsGetOK).Requests).To(HaveLen(50))
			})
		})

		Context("when user is not admin", func() {
			It("should return error", func() {
				res, err := helper.ApiClient.APIRequestsGet(ctx, api.APIRequestsGetParams{
					PageID: 1,
					Size:   api.PageSize25,
				})

				Expect(err).NotTo(HaveOccurred())
				Expect(res).NotTo(BeNil())
				Expect(res.(*api.Forbidden)).NotTo(BeNil())
			})
		})
	})
})
