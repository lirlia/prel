package model

import (
	"time"

	"github.com/google/uuid"
)

const (
	// InvitationExpiredAt is the default value of user invitation expiration
	InvitationExpiredAt = 7 * 24 * time.Hour
)

type Invitation struct {
	id            string
	inviterUserID UserID
	inviteeMail   string
	inviteeRole   UserRole
	invitedAt     time.Time
	expiredAt     time.Time
}

type Invitations []*Invitation

func (is Invitations) ToMap() map[string]*Invitation {
	m := make(map[string]*Invitation)
	for _, i := range is {
		m[i.ID()] = i
	}
	return m
}

func NewInvitation(inviterUserID UserID, inviteeMail string, inviteeRole UserRole, invitedAt time.Time) *Invitation {
	expiredAt := invitedAt.Add(InvitationExpiredAt)
	id := uuid.New().String()
	return newInvitation(id, inviterUserID, inviteeMail, inviteeRole, invitedAt, expiredAt)
}

func ReconstructInvitation(id, inviterUserID, inviteeMail, inviteeRole string, invitedAt time.Time, expiredAt time.Time) (*Invitation, error) {
	return newInvitation(
		id,
		UserID(inviterUserID),
		inviteeMail,
		UserRole(inviteeRole),
		invitedAt,
		expiredAt,
	), nil
}

func newInvitation(id string, inviterUserID UserID, inviteeMail string, inviteeRole UserRole, invitedAt time.Time, expiredAt time.Time) *Invitation {
	return &Invitation{
		id:            id,
		inviterUserID: inviterUserID,
		inviteeMail:   inviteeMail,
		inviteeRole:   inviteeRole,
		invitedAt:     invitedAt,
		expiredAt:     expiredAt,
	}
}

func (i *Invitation) ID() string {
	return i.id
}

func (i *Invitation) InviterUserID() UserID {
	return i.inviterUserID
}

func (i *Invitation) InviteeMail() string {
	return i.inviteeMail
}

func (i *Invitation) InviteeRole() UserRole {
	return i.inviteeRole
}

func (i *Invitation) InvitedAt() time.Time {
	return i.invitedAt
}

func (i *Invitation) ExpiredAt() time.Time {
	return i.expiredAt
}

func (i *Invitation) IsExpired(t time.Time) bool {
	return i.expiredAt.Before(t)
}

func (i *Invitation) UpdateExpiredAt(t time.Time) {
	i.expiredAt = t
}
