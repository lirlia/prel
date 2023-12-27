package model_test

import (
	"prel/internal/model"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {

	DescribeTable("CanJudge", func(role model.UserRole, isOwnRequest, isPending, expected bool) {
		user := model.NewUser("googleID", "email", role, time.Time{}, time.Time{})

		var requesterUserID model.UserID
		if isOwnRequest {
			requesterUserID = user.ID()
		} else {
			requesterUserID = model.UserID("requester-user-id")
		}
		req := model.NewRequest(requesterUserID, "project-id", []string{"iam-role"}, "reason", time.Time{}, time.Time{})
		if !isPending {
			req.Approve(user, time.Time{})
		}

		if expected {
			Expect(user.CanJudge(req)).NotTo(HaveOccurred())
		} else {
			Expect(user.CanJudge(req)).To(HaveOccurred())
		}
	},
		// pending
		Entry("when role is admin", model.UserRoleAdmin, false, true, true),
		Entry("when role is judger", model.UserRoleJudger, false, true, true),
		Entry("when role is requester", model.UserRoleRequester, false, true, false),
		Entry("when role is admin and is own request", model.UserRoleAdmin, true, true, false),
		Entry("when role is judger and is own request", model.UserRoleJudger, true, true, false),
		Entry("when role is requester and is own request", model.UserRoleRequester, true, true, false),
		// approved
		Entry("when role is admin and is own request", model.UserRoleAdmin, true, false, false),
	)

	DescribeTable("CanDelete", func(role model.UserRole, isOwnRequest, isPending, expected bool) {
		user := model.NewUser("googleID", "email", role, time.Time{}, time.Time{})

		var requesterUserID model.UserID
		if isOwnRequest {
			requesterUserID = user.ID()
		} else {
			requesterUserID = model.UserID("requester-user-id")
		}
		req := model.NewRequest(requesterUserID, "project-id", []string{"iam-role"}, "reason", time.Time{}, time.Time{})
		if !isPending {
			req.Approve(user, time.Time{})
		}

		if expected {
			Expect(user.CanDelete(req)).NotTo(HaveOccurred())
		} else {
			Expect(user.CanDelete(req)).To(HaveOccurred())
		}
	},
		// pending
		Entry("when role is admin", model.UserRoleAdmin, false, true, true),
		Entry("when role is judger", model.UserRoleJudger, false, true, false),
		Entry("when role is requester", model.UserRoleRequester, false, true, false),
		Entry("when role is admin and is own request", model.UserRoleAdmin, true, true, true),
		Entry("when role is judger and is own request", model.UserRoleJudger, true, true, true),
		Entry("when role is requester and is own request", model.UserRoleRequester, true, true, true),
		// approved
		Entry("when role is admin and is own request", model.UserRoleAdmin, true, false, false),
	)

	DescribeTable("Session Expired", func(sessionExpiredAt time.Time, expected bool) {
		user := model.NewUser("googleID", "email", model.UserRoleAdmin, sessionExpiredAt, time.Time{})
		Expect(user.IsSessionExpired(time.Now())).To(Equal(expected))
	},
		Entry("when session is expired", time.Now().Add(-10*time.Second), true),
		Entry("when session is not expired", time.Now().Add(10*time.Second), false),
	)
})
