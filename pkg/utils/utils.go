package utils

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/cockroachdb/errors"
)

func RandomString(num int) (string, error) {
	b := make([]byte, num)
	if _, err := rand.Read(b); err != nil {
		return "", errors.Wrap(err, "failed to read random")
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// calcPageSize calculates start and end index of page for pagination.
func CalcPageSize(page int, pageSize int, count int) (start int, end int) {
	start = (page - 1) * pageSize
	end = page * pageSize
	if end > count {
		end = count
	}
	return
}
