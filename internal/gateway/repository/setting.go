package repository

import (
	"context"
	"prel/internal/gateway/postgresql"
	"prel/internal/model"

	"github.com/cockroachdb/errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type settingRepository struct{}

func NewSettingRepository() *settingRepository {
	return &settingRepository{}
}

func (r *settingRepository) Save(ctx context.Context, req *model.Setting) error {

	err := postgresql.GetQuery(ctx).UpsertSetting(ctx, postgresql.UpsertSettingParams{
		ID:                            req.ID(),
		NotificationMessageForRequest: pgtype.Text{String: req.NotificationMessageForRequest(), Valid: true},
		NotificationMessageForJudge:   pgtype.Text{String: req.NotificationMessageForJudge(), Valid: true},
	})
	if err != nil {
		return errors.Wrap(err, "failed to save user")
	}

	return nil
}

func (r *settingRepository) Find(ctx context.Context) (*model.Setting, error) {
	setting, err := postgresql.GetQuery(ctx).FindSetting(ctx)

	// return default setting if not found
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return model.NewSetting("", ""), nil
	}

	if err != nil {
		return nil, errors.Wrap(err, "failed to find setting")
	}

	return model.ReconstructSetting(
		setting.ID,
		setting.NotificationMessageForRequest.String,
		setting.NotificationMessageForJudge.String,
	), nil
}
