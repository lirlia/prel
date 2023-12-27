package repository

import (
	"context"
	"prel/internal/gateway/postgresql"
	"prel/internal/model"
	"time"

	"github.com/cockroachdb/errors"
)

type userAndInvitationRepository struct{}

func NewUserAndInvitationRepository() *userAndInvitationRepository {
	return &userAndInvitationRepository{}
}

func (u *userAndInvitationRepository) FindUserAndInvitationPagedByExpiredAt(ctx context.Context, start int, limit int, until time.Time) (model.Users, error) {
	users, err := postgresql.GetQuery(ctx).FindUserAndInvitationPagedByExpiredAt(ctx, postgresql.FindUserAndInvitationPagedByExpiredAtParams{
		ExpiredAt: Timestamptz(until),
		Limit:     int32(limit),
		Offset:    int32(start),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to find user paged")
	}

	var result model.Users
	for _, user := range users {
		result = append(result, model.Reconstruct(
			user.ID,
			user.GoogleID,
			user.Email,
			user.IsAvailable,
			model.UserRole(user.Role),
			user.SessionID,
			user.SessionExpiredAt.Time,
			user.LastSigninAt.Time))
	}

	return result, nil
}

func (u *userAndInvitationRepository) Count(ctx context.Context) (int, error) {
	countUser, err := postgresql.GetQuery(ctx).CountUser(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "failed to count user")
	}

	countInvitation, err := postgresql.GetQuery(ctx).CountInvitation(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "failed to count invitation")
	}

	return int(countUser + countInvitation), nil
}
