package testutil

import "github.com/google/uuid"

func IsValidUUID(s string) bool {
	_, err := uuid.Parse(s)
	return err == nil
}
