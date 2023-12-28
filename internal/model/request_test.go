package model_test

import (
	"prel/internal/model"
	"prel/test/testutil"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Request", func() {
	Describe("Approve/Reject", func() {
		Context("approve request", func() {
			It("should return expected request", func() {
				now := time.Now()
				req, err := model.NewRequest("user-id", "project-id", []string{"iam-role"}, 5, "reason", now, now.Add(1*time.Hour))
				Expect(err).NotTo(HaveOccurred())
				judger := model.NewUser("judger-id", "judger-email", model.UserRoleJudger, now, now.Add(1*time.Hour))

				req.Approve(judger, now)
				Expect(req.Status()).To(Equal(model.RequestStatusApproved))
				Expect(req.JudgerUserID()).To(Equal(judger.ID()))
				Expect(req.JudgerEmail()).To(Equal(judger.Email()))
				Expect(req.JudgedAt()).To(Equal(now))
			})
		})

		Context("reject request", func() {
			It("should return expected request", func() {
				now := time.Now()
				req, err := model.NewRequest("user-id", "project-id", []string{"iam-role"}, 5, "reason", now, now.Add(1*time.Hour))
				Expect(err).NotTo(HaveOccurred())
				judger := model.NewUser("judger-id", "judger-email", model.UserRoleJudger, now, now.Add(1*time.Hour))

				req.Reject(judger, now)
				Expect(req.Status()).To(Equal(model.RequestStatusRejected))
				Expect(req.JudgerUserID()).To(Equal(judger.ID()))
				Expect(req.JudgerEmail()).To(Equal(judger.Email()))
				Expect(req.JudgedAt()).To(Equal(now))
			})
		})
	})

	Describe("CalculateRoleBindingExpiry", func() {
		clock := testutil.NewClock()
		// fixed time
		clock.Set(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))

		DescribeTable("table", func(period model.PeriodKey, expected time.Time, isError bool) {
			req, err := model.NewRequest("user-id", "project-id", []string{"iam-role"}, period, "reason", clock.Now(), clock.Now().Add(1*time.Hour))
			if isError {
				Expect(err).To(HaveOccurred())
				return
			}
			Expect(err).NotTo(HaveOccurred())
			Expect(req.CalculateRoleBindingExpiry(clock.Now())).To(Equal(expected))
		},
			Entry("5 min", model.PeriodKey5, clock.Now().Add(5*time.Minute), false),
			Entry("10 min", model.PeriodKey10, clock.Now().Add(10*time.Minute), false),
			Entry("30 min", model.PeriodKey30, clock.Now().Add(30*time.Minute), false),
			Entry("1 hour", model.PeriodKey60, clock.Now().Add(1*time.Hour), false),
			Entry("5 hour", model.PeriodKey(300), clock.Now().Add(5*time.Hour), true),
		)
	})
})
