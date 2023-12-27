package model

import (
	"context"
	"fmt"
	"prel/pkg/custom_error"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
)

type UserRole string

const (
	UserRoleAdmin     UserRole = "admin"
	UserRoleRequester UserRole = "requester"
	UserRoleJudger    UserRole = "judger"
)

type UserRoles []UserRole

// The user's role is rendered in this order.
// Sort by the smallest authority.
func SortedUserRoles() UserRoles {
	return UserRoles{
		UserRoleRequester,
		UserRoleJudger,
		UserRoleAdmin,
	}
}

type UserID string
type SessionID string

type User struct {
	id       UserID
	googleID string
	email    string
	// isAvailable is true when the user is available to login.
	isAvailable      bool
	role             UserRole
	sessionId        SessionID
	sessionExpiredAt time.Time
	lastSignInAt     time.Time
}

type Users []*User

func (u *User) ID() UserID {
	return u.id
}

func (u *User) GoogleID() string {
	return u.googleID
}

func (u *User) Email() string {
	return u.email
}

func (u *User) IsAvailable() bool {
	return u.isAvailable
}

func (u *User) SetAvailable(isAvailable bool) {
	u.isAvailable = isAvailable
}

func (u *User) SetRole(role UserRole) {
	u.role = role
}

func (u *User) Role() UserRole {
	return u.role
}

func (u *User) SessionID() SessionID {
	return u.sessionId
}

func generateSessionID() SessionID {
	return SessionID(uuid.New().String())
}

func (u *User) UpdateSession(sessionExpiredAt time.Time) {
	u.sessionId = generateSessionID()
	u.sessionExpiredAt = sessionExpiredAt
}

func (u *User) SessionExpiredAt() time.Time {
	return u.sessionExpiredAt
}

func (u *User) IsSessionExpired(t time.Time) bool {
	return u.sessionExpiredAt.Before(t)
}

func (u *User) Signout() {
	u.sessionExpiredAt = time.Time{}
}

func (u *User) LastSignInAt() time.Time {
	return u.lastSignInAt
}

func (u *User) UpdateLastSignInAt(now time.Time) {
	u.lastSignInAt = now
}

func (u *User) IsAdmin() bool {
	return u.role == UserRoleAdmin
}

// CanJudge returns true if the user can judge the request.
func (u *User) CanJudge(req *Request) error {

	if req.status != RequestStatusPending {
		return errors.WithDetail(errors.Newf("invalid status(%s), status must be pending", req.status), string(custom_error.InvalidArgument))
	}

	if u.Role() == UserRoleRequester {
		return errors.WithDetail(errors.Newf("user is requester", req.status), string(custom_error.InvalidArgument))
	}

	if req.requesterUserID == u.ID() {
		return errors.WithDetail(errors.Newf("can't judge own request", req.status), string(custom_error.InvalidArgument))
	}

	return nil
}

// CanDelete returns true if the user can delete the request.
func (u *User) CanDelete(req *Request) error {

	if req.status != RequestStatusPending {
		return errors.WithDetail(errors.Newf("invalid status(%s), status must be pending", req.status), string(custom_error.InvalidArgument))
	}

	if u.Role() == UserRoleAdmin {
		return nil
	}

	if u.ID() == req.requesterUserID {
		return nil
	}

	return errors.WithDetail(errors.Newf("user(%s) can't delete request(%s)", u.ID(), req.ID()), string(custom_error.OnlyAdmin))
}

// https://cloud.google.com/iam/docs/principal-identifiers
func (u *User) Principal() string {
	return fmt.Sprintf("%s:%s", u.PrincipalType(), u.Email())
}

func (u *User) PrincipalType() string {
	return "user"
}

func Reconstruct(id, googleID, email string, isAvailable bool, role UserRole, sessionID string, sessionExpiredAt, lastSignInAt time.Time) *User {
	s := SessionID(sessionID)
	return newUser(UserID(id), googleID, email, isAvailable, role, s, sessionExpiredAt, lastSignInAt)
}

func NewUser(googleID, email string, role UserRole, sessionExpiredAt, now time.Time) *User {
	id := UserID(uuid.New().String())
	sessionID := generateSessionID()
	return newUser(id, googleID, email, true, role, sessionID, sessionExpiredAt, now)
}

func newUser(id UserID, googleID, email string, isAvailable bool, role UserRole, sessionID SessionID, sessionExpiredAt, lastSignInAt time.Time) *User {
	return &User{
		id:               id,
		googleID:         googleID,
		email:            email,
		isAvailable:      isAvailable,
		role:             role,
		sessionId:        sessionID,
		sessionExpiredAt: sessionExpiredAt,
		lastSignInAt:     lastSignInAt,
	}
}

func (u Users) ByID(id UserID) *User {
	for _, user := range u {
		if user.ID() == id {
			return user
		}
	}

	return nil
}

type sessionKey struct{}

func SetUser(ctx context.Context, user *User) context.Context {
	return context.WithValue(ctx, sessionKey{}, user)
}

func GetUser(ctx context.Context) *User {
	v, ok := ctx.Value(sessionKey{}).(*User)
	if !ok {
		return nil
	}
	return v
}
