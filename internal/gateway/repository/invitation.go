package repository

import (
	"context"
	"prel/internal/gateway/postgresql"
	"prel/internal/model"
	"time"

	"github.com/cockroachdb/errors"
)

type invitationRepository struct{}

func NewInvitationRepository() *invitationRepository {
	return &invitationRepository{}
}

func (i *invitationRepository) FindByInviteeMail(ctx context.Context, email string) (*model.Invitation, error) {
	req, err := postgresql.GetQuery(ctx).FindInvitationByInviteeMail(ctx, email)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find invitation by invitee mail")
	}

	return model.ReconstructInvitation(
		req.ID,
		req.InviterUserID,
		req.InviteeMail,
		req.InviteeRole,
		req.InvitedAt.Time,
		req.ExpiredAt.Time)
}

func (i *invitationRepository) FindByInviteeMailsAndExpiredAt(ctx context.Context, emails []string, until time.Time) (model.Invitations, error) {
	reqs, err := postgresql.GetQuery(ctx).FindInvitationByInviteeMailsAndExpiredAt(ctx, postgresql.FindInvitationByInviteeMailsAndExpiredAtParams{
		Column1:   emails,
		ExpiredAt: Timestamptz(until),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to find invitation by invitee mail")
	}

	ret := make(model.Invitations, 0, len(reqs))
	for _, req := range reqs {
		i, err := model.ReconstructInvitation(
			req.ID,
			req.InviterUserID,
			req.InviteeMail,
			req.InviteeRole,
			req.InvitedAt.Time,
			req.ExpiredAt.Time)

		if err != nil {
			return nil, errors.Wrap(err, "failed to reconstruct invitation")
		}

		ret = append(ret, i)
	}

	return ret, nil
}

func (i *invitationRepository) Save(ctx context.Context, invitation *model.Invitation) error {
	err := postgresql.GetQuery(ctx).UpsertInvitation(ctx, postgresql.UpsertInvitationParams{
		ID:            invitation.ID(),
		InviterUserID: string(invitation.InviterUserID()),
		InviteeMail:   invitation.InviteeMail(),
		InviteeRole:   string(invitation.InviteeRole()),
		InvitedAt:     Timestamptz(invitation.InvitedAt()),
		ExpiredAt:     Timestamptz(invitation.ExpiredAt()),
	})
	if err != nil {
		return errors.Wrap(err, "failed to save invitation")
	}

	return nil
}

func (i *invitationRepository) Delete(ctx context.Context, invitation *model.Invitation) error {
	err := postgresql.GetQuery(ctx).DeleteInvitation(ctx, invitation.ID())
	if err != nil {
		return errors.Wrap(err, "failed to delete invitation")
	}

	return nil
}

func (i *invitationRepository) Count(ctx context.Context) (int, error) {
	count, err := postgresql.GetQuery(ctx).CountInvitation(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "failed to count invitation")
	}

	return int(count), nil
}
