package model_test

import (
	"prel/internal/model"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	DescribeTable("Session Expired", func(sessionExpiredAt time.Time, expected bool) {
		user := model.NewUser("googleID", "email", model.UserRoleAdmin, sessionExpiredAt, time.Time{})
		Expect(user.IsSessionExpired(time.Now())).To(Equal(expected))
	},
		Entry("when session is expired", time.Now().Add(-10*time.Second), true),
		Entry("when session is not expired", time.Now().Add(10*time.Second), false),
	)
})
