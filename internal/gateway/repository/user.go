package repository

import (
	"context"
	"prel/internal/gateway/postgresql"
	"prel/internal/model"

	"github.com/cockroachdb/errors"
)

type userRepository struct{}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (u *userRepository) FindByID(ctx context.Context, id model.UserID) (*model.User, error) {
	user, err := postgresql.GetQuery(ctx).FindUserByID(ctx, string(id))
	if err != nil {
		return nil, errors.Wrap(err, "failed to find user by id")
	}

	return model.Reconstruct(user.ID, user.GoogleID, user.Email, user.IsAvailable, model.UserRole(user.Role), user.SessionID, user.SessionExpiredAt.Time, user.LastSigninAt.Time), nil
}

func (u *userRepository) FindByIDs(ctx context.Context, userIDs []model.UserID) (model.Users, error) {
	ids := make([]string, 0, len(userIDs))
	for _, id := range userIDs {
		ids = append(ids, string(id))
	}

	users, err := postgresql.GetQuery(ctx).FindUserByIDs(ctx, ids)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find users by ids")
	}

	var result model.Users
	for _, user := range users {
		result = append(result, model.Reconstruct(user.ID, user.GoogleID, user.Email, user.IsAvailable, model.UserRole(user.Role), user.SessionID, user.SessionExpiredAt.Time, user.LastSigninAt.Time))
	}

	return result, nil
}

func (u *userRepository) FindByGoogleID(ctx context.Context, id string) (*model.User, error) {
	user, err := postgresql.GetQuery(ctx).FindUserByGoogleID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find user by google id")
	}

	return model.Reconstruct(user.ID, user.GoogleID, user.Email, user.IsAvailable, model.UserRole(user.Role), user.SessionID, user.SessionExpiredAt.Time, user.LastSigninAt.Time), nil
}

func (u *userRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := postgresql.GetQuery(ctx).FindUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find user by email")
	}

	return model.Reconstruct(user.ID, user.GoogleID, user.Email, user.IsAvailable, model.UserRole(user.Role), user.SessionID, user.SessionExpiredAt.Time, user.LastSigninAt.Time), nil
}

func (u *userRepository) FindBySessionID(ctx context.Context, id model.SessionID) (*model.User, error) {
	user, err := postgresql.GetQuery(ctx).FindUserBySessionID(ctx, string(id))
	if err != nil {
		return nil, errors.Wrap(err, "failed to find user by session id")
	}

	return model.Reconstruct(user.ID, user.GoogleID, user.Email, user.IsAvailable, model.UserRole(user.Role), user.SessionID, user.SessionExpiredAt.Time, user.LastSigninAt.Time), nil
}

func (u *userRepository) Count(ctx context.Context) (int, error) {
	count, err := postgresql.GetQuery(ctx).CountUser(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "failed to count user")
	}

	return int(count), nil
}

func (u *userRepository) Save(ctx context.Context, user *model.User) error {
	err := postgresql.GetQuery(ctx).UpsertUser(ctx, postgresql.UpsertUserParams{
		ID:               string(user.ID()),
		GoogleID:         user.GoogleID(),
		Email:            user.Email(),
		IsAvailable:      user.IsAvailable(),
		Role:             string(user.Role()),
		SessionID:        string(user.SessionID()),
		SessionExpiredAt: Timestamptz(user.SessionExpiredAt()),
		LastSigninAt:     Timestamptz(user.LastSignInAt()),
	})
	if err != nil {
		return errors.Wrap(err, "failed to save user")
	}

	return nil
}
