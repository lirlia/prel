package middleware_test

import (
	"prel/test/testutil"

	. "github.com/onsi/ginkgo/v2"
)

var _ = AfterSuite(func() {
	testutil.StopTestDB()
})
