package cmd_test

import (
	"os"
	"prel/cmd/prel/cmd"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"
)

var _ = Describe("Cmd", func() {

	Describe("Execute", func() {
		AfterEach(func() {
			os.Unsetenv("CLIENT_ID")
			os.Unsetenv("CLIENT_SECRET")
			os.Unsetenv("PROJECT_ID")
			os.Unsetenv("DB_PASSWORD")
			os.Unsetenv("ADDRESS")
		})

		Context("when execute with required flag", func() {
			It("should return nil", func() {
				os.Setenv("CLIENT_ID", "valid")
				os.Setenv("CLIENT_SECRET", "valid")
				os.Setenv("PROJECT_ID", "valid")
				os.Setenv("DB_PASSWORD", "valid")
				Expect(cmd.Validate()).To(Succeed())
			})
		})

		Context("when cmdline flag and os environment set together", func() {
			It("should return cmdline flag value", func() {

				os.Setenv("CLIENT_ID", "valid")
				os.Setenv("CLIENT_SECRET", "valid")
				os.Setenv("PROJECT_ID", "valid")
				os.Setenv("DB_PASSWORD", "valid")
				os.Setenv("ADDRESS", "2.2.2.2")

				args := []string{
					"--address", "1.1.1.1",
				}
				cmd.RootCmd.SetArgs(args)

				cmd.ExecuteTest(func() {
					Expect(viper.GetString("address")).To(Equal("1.1.1.1"))
				})
			})
		})

		Context("when execute without required flag", func() {
			It("should return error", func() {
				os.Setenv("CLIENT_ID", "valid")
				os.Setenv("CLIENT_SECRET", "valid")
				os.Setenv("PROJECT_ID", "valid")
				err := cmd.Validate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("required flag is not set: db-password"))
			})
		})

		Context("when db-type is invalid", func() {
			It("should return error", func() {
				os.Setenv("CLIENT_ID", "valid")
				os.Setenv("CLIENT_SECRET", "valid")
				os.Setenv("PROJECT_ID", "valid")
				os.Setenv("DB_PASSWORD", "valid")
				os.Setenv("DB_TYPE", "invalid")
				err := cmd.Validate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("invalid db-type: invalid"))
			})
		})
	})
})
