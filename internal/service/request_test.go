package service_test

import (
	"prel/internal/model"
	"prel/internal/service"
	"prel/test/testutil"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Request", func() {

	var (
		srv *service.Service
	)

	BeforeEach(func() {
		srv = service.NewService()
	})

	Describe("Can Test", func() {
		var (
			clock *testutil.Clock
			now   time.Time
		)

		BeforeEach(func() {
			clock = testutil.NewClock()
			clock.Set(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))
			now = clock.Now()
		})

		DescribeTable("CanJudgeRequest", func(role model.UserRole, isOwnRequest, isPending, isExpired, expected bool) {

			user := testutil.NewTestUser(testutil.WithRole(role))
			opts := []testutil.RequestOption{}
			if isOwnRequest {
				opts = append(opts, testutil.WithRequesterUserID(user.ID()))
			} else {
				opts = append(opts, testutil.WithRequesterUserID(model.UserID("requester-user-id")))
			}

			if !isPending {
				opts = append(opts, testutil.WithStatus(model.RequestStatusApproved))
			}

			if isExpired {
				opts = append(opts, testutil.WithExpiredAt(now.Add(-1*time.Hour)))
			} else {
				opts = append(opts, testutil.WithExpiredAt(now.Add(1*time.Hour)))
			}

			req := testutil.NewTestRequest(opts...)
			err := srv.CanJudgeRequest(req, user, now)
			if expected {
				Expect(err).NotTo(HaveOccurred())
			} else {
				Expect(err).To(HaveOccurred())
			}
		},
			// pending
			Entry("when role is admin", model.UserRoleAdmin, false, true, false, true),
			Entry("when role is judger", model.UserRoleJudger, false, true, false, true),
			Entry("when role is requester", model.UserRoleRequester, false, true, false, false),
			Entry("when role is admin and is own request", model.UserRoleAdmin, true, true, false, false),
			Entry("when role is judger and is own request", model.UserRoleJudger, true, true, false, false),
			Entry("when role is requester and is own request", model.UserRoleRequester, true, true, false, false),
			// approved
			Entry("when role is admin and is own request", model.UserRoleAdmin, true, false, false, false),
			// expired
			Entry("when role is judger and request is expired", model.UserRoleJudger, false, true, true, false),
		)

		DescribeTable("CanDeleteRequest", func(role model.UserRole, isOwnRequest, isPending, expected bool) {
			user := testutil.NewTestUser(testutil.WithRole(role))
			opts := []testutil.RequestOption{}
			if isOwnRequest {
				opts = append(opts, testutil.WithRequesterUserID(user.ID()))
			} else {
				opts = append(opts, testutil.WithRequesterUserID(model.UserID("requester-user-id")))
			}

			if !isPending {
				opts = append(opts, testutil.WithStatus(model.RequestStatusApproved))
			}

			req := testutil.NewTestRequest(opts...)

			if expected {
				Expect(srv.CanDeleteRequest(req, user)).NotTo(HaveOccurred())
			} else {
				Expect(srv.CanDeleteRequest(req, user)).To(HaveOccurred())
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
			Entry("when role is admin and isn't status pending", model.UserRoleAdmin, true, false, false),
		)
	})
})
