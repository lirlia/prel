package model

import (
	"context"
	"time"
)

type UserRepository interface {
	Count(ctx context.Context) (int, error)
	Save(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id UserID) (*User, error)
	FindByIDs(ctx context.Context, ids []UserID) (Users, error)
	FindByGoogleID(ctx context.Context, id string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindBySessionID(ctx context.Context, id SessionID) (*User, error)
}

type RequestRepository interface {
	Count(ctx context.Context) (int, error)
	Save(ctx context.Context, req *Request) error
	Delete(ctx context.Context, req *Request) error
	FindByID(ctx context.Context, requestId string) (*Request, error)
	FindRequestByStatusAndExpiredAt(ctx context.Context, status RequestStatus, t time.Time) (Requests, error)
	FindPaged(ctx context.Context, start int, limit int) (Requests, error)
}

type InvitationRepository interface {
	Save(ctx context.Context, invitation *Invitation) error
	Delete(ctx context.Context, invitation *Invitation) error
	FindByInviteeMail(ctx context.Context, mail string) (*Invitation, error)
	FindByInviteeMailsAndExpiredAt(ctx context.Context, mail []string, until time.Time) (Invitations, error)
}

type IamRoleFilteringRuleRepository interface {
	FindByID(ctx context.Context, id string) (*IamRoleFilteringRule, error)
	FindAll(ctx context.Context) (IamRoleFilteringRules, error)
	Save(ctx context.Context, rule *IamRoleFilteringRule) error
	Delete(ctx context.Context, rule *IamRoleFilteringRule) error
}

type TransactionManager interface {
	Transaction(ctx context.Context, f func(ctx context.Context) error) error
}

type UserAndInvitationRepository interface {
	Count(ctx context.Context) (int, error)
	FindUserAndInvitationPagedByExpiredAt(ctx context.Context, start int, limit int, until time.Time) (Users, error)
}

type SettingRepository interface {
	Find(ctx context.Context) (*Setting, error)
	Save(ctx context.Context, setting *Setting) error
}
