// Code generated by sqlc. DO NOT EDIT.

package postgresql

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type IamRoleFilteringRule struct {
	ID        string
	Pattern   string
	UserID    string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

type Invitation struct {
	ID            string
	InviterUserID string
	InviteeMail   string
	InviteeRole   string
	InvitedAt     pgtype.Timestamptz
	ExpiredAt     pgtype.Timestamptz
	CreatedAt     pgtype.Timestamptz
	UpdatedAt     pgtype.Timestamptz
}

type Request struct {
	ID              string
	RequesterUserID string
	JudgerUserID    pgtype.Text
	Status          string
	ProjectID       string
	IamRoles        string
	Period          int32
	Reason          string
	RequestedAt     pgtype.Timestamptz
	ExpiredAt       pgtype.Timestamptz
	JudgedAt        pgtype.Timestamptz
	CreatedAt       pgtype.Timestamptz
	UpdatedAt       pgtype.Timestamptz
}

type Setting struct {
	ID                            string
	NotificationMessageForRequest pgtype.Text
	NotificationMessageForJudge   pgtype.Text
	CreatedAt                     pgtype.Timestamptz
	UpdatedAt                     pgtype.Timestamptz
}

type User struct {
	ID               string
	GoogleID         string
	Email            string
	IsAvailable      bool
	Role             string
	SessionID        string
	SessionExpiredAt pgtype.Timestamptz
	LastSigninAt     pgtype.Timestamptz
	CreatedAt        pgtype.Timestamptz
	UpdatedAt        pgtype.Timestamptz
}
