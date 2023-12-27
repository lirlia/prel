package integration_test

import (
	"context"
	api "prel/api/prel_api"
	"prel/internal/model"
	"prel/test/testutil"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
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
		helper.User.SetRole(model.UserRoleAdmin)
		Expect(helper.UserRepo.Save(ctx, helper.User)).To(Succeed())
	})

	AfterEach(func() {
		helper.Close()
	})

	Describe("Create", func() {
		Context("when iam role filtering is valid", func() {
			It("should create expected rule", func() {
				res, err := helper.ApiClient.APIIamRoleFilteringRulesPost(ctx, &api.APIIamRoleFilteringRulesPostReq{
					Pattern: "spanner",
				})

				Expect(err).NotTo(HaveOccurred())
				Expect(res).NotTo(BeNil())
				Expect(res.(*api.APIIamRoleFilteringRulesPostOK)).NotTo(BeNil())
				rule := res.(*api.APIIamRoleFilteringRulesPostOK).IamRoleFilteringRule
				Expect(rule).NotTo(BeNil())
				Expect(rule.ID).NotTo(BeEmpty())
				Expect(rule.Pattern).To(Equal("spanner"))

				dbRule, err := helper.IamRoleFilteringRepo.FindByID(ctx, rule.ID)
				Expect(err).NotTo(HaveOccurred())
				Expect(dbRule).NotTo(BeNil())
				Expect(dbRule.Pattern()).To(Equal("spanner"))
			})
		})
	})

	Describe("Delete", func() {
		var ruleID string
		BeforeEach(func() {
			res, err := helper.ApiClient.APIIamRoleFilteringRulesPost(ctx, &api.APIIamRoleFilteringRulesPostReq{
				Pattern: "spanner",
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(res).NotTo(BeNil())
			Expect(res.(*api.APIIamRoleFilteringRulesPostOK)).NotTo(BeNil())
			ruleID = res.(*api.APIIamRoleFilteringRulesPostOK).IamRoleFilteringRule.ID

			dbRule, err := helper.IamRoleFilteringRepo.FindByID(ctx, ruleID)
			Expect(err).NotTo(HaveOccurred())
			Expect(dbRule).NotTo(BeNil())
			Expect(dbRule.Pattern()).To(Equal("spanner"))
		})

		Context("delete", func() {
			It("should delete expected rule", func() {
				res, err := helper.ApiClient.APIIamRoleFilteringRulesRuleIDDelete(ctx, api.APIIamRoleFilteringRulesRuleIDDeleteParams{
					RuleID: ruleID,
				})

				Expect(err).NotTo(HaveOccurred())
				Expect(res).NotTo(BeNil())
				Expect(res.(*api.APIIamRoleFilteringRulesRuleIDDeleteNoContent)).NotTo(BeNil())

				dbRule, err := helper.IamRoleFilteringRepo.FindByID(ctx, ruleID)
				Expect(err).To(HaveOccurred())
				Expect(dbRule).To(BeNil())
			})
		})
	})

	Describe("Get", func() {
		BeforeEach(func() {
			for _, pattern := range []string{"spanner", "storage", "bigquery"} {
				_, err := helper.ApiClient.APIIamRoleFilteringRulesPost(ctx, &api.APIIamRoleFilteringRulesPostReq{
					Pattern: pattern,
				})
				Expect(err).NotTo(HaveOccurred())
			}
		})

		Context("when request is valid", func() {
			It("should return expected rules", func() {
				res, err := helper.ApiClient.APIIamRoleFilteringRulesGet(ctx)

				Expect(err).NotTo(HaveOccurred())
				Expect(res).NotTo(BeNil())
				Expect(res.(*api.APIIamRoleFilteringRulesGetOK)).NotTo(BeNil())
				Expect(res.(*api.APIIamRoleFilteringRulesGetOK).IamRoleFilteringRules).To(HaveLen(3))
			})
		})
	})

})
