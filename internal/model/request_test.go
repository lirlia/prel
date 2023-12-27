package model_test

import (
	"prel/internal/model"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Request", func() {
	Describe("Approve/Reject", func() {
		Context("approve request", func() {
			It("should return expected request", func() {
				now := time.Now()
				req := model.NewRequest("user-id", "project-id", []string{"iam-role"}, "reason", now, now.Add(1*time.Hour))
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
				req := model.NewRequest("user-id", "project-id", []string{"iam-role"}, "reason", now, now.Add(1*time.Hour))
				judger := model.NewUser("judger-id", "judger-email", model.UserRoleJudger, now, now.Add(1*time.Hour))

				req.Reject(judger, now)
				Expect(req.Status()).To(Equal(model.RequestStatusRejected))
				Expect(req.JudgerUserID()).To(Equal(judger.ID()))
				Expect(req.JudgerEmail()).To(Equal(judger.Email()))
				Expect(req.JudgedAt()).To(Equal(now))
			})
		})
	})
})
