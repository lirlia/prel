package usecase_test

import (
	"prel/internal/usecase"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Usecase", func() {
	Describe("FilterRoleID", func() {
		Context("when patterns are present in roleIDs", func() {
			It("should return a slice of roleIDs that match the patterns", func() {
				roleIDs := []string{"admin123", "user456", "guest789", "admin987"}
				patterns := []string{"admin", "guest"}

				filteredIDs := usecase.FilterRoleIDs(roleIDs, patterns)

				Expect(filteredIDs).To(ConsistOf("admin123", "guest789", "admin987"))
			})
		})

		Context("when no patterns are present in roleIDs", func() {
			It("should return an empty slice", func() {
				roleIDs := []string{"user456", "user789"}
				patterns := []string{"admin", "guest"}

				filteredIDs := usecase.FilterRoleIDs(roleIDs, patterns)

				Expect(filteredIDs).To(BeEmpty())
			})
		})
	})
})
