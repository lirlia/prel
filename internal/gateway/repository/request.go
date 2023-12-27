package repository

import (
	"context"
	"database/sql"
	"prel/internal/gateway/postgresql"
	"prel/internal/model"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/jackc/pgx/v5/pgtype"
)

type requestRepository struct{}

func NewRequestRepository() *requestRepository {
	return &requestRepository{}
}

func (r *requestRepository) Save(ctx context.Context, req *model.Request) error {

	judger := pgtype.Text{}
	if req.JudgerUserID() != "" {
		judger = pgtype.Text{String: string(req.JudgerUserID()), Valid: true}
	}

	judgedAt := sql.NullTime{}
	if !req.JudgedAt().IsZero() {
		judgedAt = sql.NullTime{Time: req.JudgedAt(), Valid: true}
	}

	err := postgresql.GetQuery(ctx).UpsertRequest(ctx, postgresql.UpsertRequestParams{
		ID:              req.ID(),
		RequesterUserID: string(req.RequesterUserID()),
		JudgerUserID:    judger,
		Status:          string(req.Status()),
		ProjectID:       req.ProjectID(),
		IamRoles:        req.ConcatIamRole(),
		Reason:          req.Reason(),
		RequestedAt:     Timestamptz(req.RequestedAt()),
		ExpiredAt:       Timestamptz(req.ExpiredAt()),
		JudgedAt:        TimestamptzNullTime(judgedAt),
	})
	if err != nil {
		return errors.Wrap(err, "failed to save user")
	}

	return nil
}

func (r *requestRepository) FindRequestByStatusAndExpiredAt(ctx context.Context, status model.RequestStatus, t time.Time) (model.Requests, error) {
	reqs, err := postgresql.GetQuery(ctx).FindRequestByStatusAndExpiredAt(ctx, postgresql.FindRequestByStatusAndExpiredAtParams{
		Status:    string(status),
		ExpiredAt: Timestamptz(t),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to find request by status")
	}

	requests := make(model.Requests, 0, len(reqs))
	for _, req := range reqs {
		r, err := model.ReconstructRequest(
			req.ID,
			req.RequesterUserID,
			req.RequesterEmail.String,
			req.JudgerUserID.String,
			req.JudgerEmail.String,
			req.Status,
			req.ProjectID,
			req.IamRoles,
			req.Reason,
			req.RequestedAt.Time,
			req.ExpiredAt.Time,
			req.JudgedAt.Time,
		)

		if err != nil {
			return nil, errors.Wrap(err, "failed to reconstruct request")
		}
		requests = append(requests, r)
	}

	return requests, nil
}

func (r *requestRepository) FindByID(ctx context.Context, id string) (*model.Request, error) {
	req, err := postgresql.GetQuery(ctx).FindRequestByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find request by id")
	}

	return model.ReconstructRequest(
		req.ID,
		req.RequesterUserID,
		req.RequesterEmail.String,
		req.JudgerUserID.String,
		req.JudgerEmail.String,
		req.Status,
		req.ProjectID,
		req.IamRoles,
		req.Reason,
		req.RequestedAt.Time,
		req.ExpiredAt.Time,
		req.JudgedAt.Time,
	)
}

func (r *requestRepository) FindPaged(ctx context.Context, start int, limit int) (model.Requests, error) {
	reqs, err := postgresql.GetQuery(ctx).FindRequestPaged(ctx, postgresql.FindRequestPagedParams{
		Offset: int32(start),
		Limit:  int32(limit),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to find request paged")
	}

	requests := make(model.Requests, 0, len(reqs))
	for _, req := range reqs {
		r, err := model.ReconstructRequest(
			req.ID,
			req.RequesterUserID,
			req.RequesterEmail.String,
			req.JudgerUserID.String,
			req.JudgerEmail.String,
			req.Status,
			req.ProjectID,
			req.IamRoles,
			req.Reason,
			req.RequestedAt.Time,
			req.ExpiredAt.Time,
			req.JudgedAt.Time,
		)

		if err != nil {
			return nil, errors.Wrap(err, "failed to reconstruct request")
		}
		requests = append(requests, r)
	}

	return requests, nil
}

func (r *requestRepository) Count(ctx context.Context) (int, error) {
	count, err := postgresql.GetQuery(ctx).CountRequest(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "failed to count request")
	}
	return int(count), nil
}

func (r *requestRepository) Delete(ctx context.Context, req *model.Request) error {
	err := postgresql.GetQuery(ctx).DeleteRequest(ctx, req.ID())
	if err != nil {
		return errors.Wrap(err, "failed to delete request")
	}
	return nil
}
