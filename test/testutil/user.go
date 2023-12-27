package testutil

import (
	"fmt"
	"prel/internal/model"
	"prel/pkg/utils"
	"time"
)

func NewTestUser(opts ...UserOption) *model.User {
	e1, _ := utils.RandomString(10)
	e2, _ := utils.RandomString(10)
	googleID, _ := utils.RandomString(10)
	u := &TestUser{
		googleID:         googleID,
		email:            fmt.Sprintf("%s@%s", e1, e2),
		role:             model.UserRoleRequester,
		sessionExpiredAt: time.Now().Add(1 * time.Hour),
		lastSignInAt:     time.Now(),
	}

	for _, opt := range opts {
		opt(u)
	}

	return model.NewUser(u.googleID, u.email, u.role, u.sessionExpiredAt, u.lastSignInAt)
}

type TestUser struct {
	googleID         string
	email            string
	role             model.UserRole
	sessionExpiredAt time.Time
	lastSignInAt     time.Time
}

type UserOption func(*TestUser)

func WithRole(role model.UserRole) UserOption {
	return func(u *TestUser) {
		u.role = role
	}
}

func WithGoogleID(googleID string) UserOption {
	return func(u *TestUser) {
		u.googleID = googleID
	}
}

func WithEmail(email string) UserOption {
	return func(u *TestUser) {
		u.email = email
	}
}

func WithSessionExpiredAt(sessionExpiredAt time.Time) UserOption {
	return func(u *TestUser) {
		u.sessionExpiredAt = sessionExpiredAt
	}
}
