package service

import (
	"context"
	"prel/internal/model"
	"prel/pkg/custom_error"
	"time"

	"github.com/cockroachdb/errors"
)

func (s *Service) Judge(ctx context.Context, req *model.Request, judgeStatus model.RequestStatus, requester, judger *model.User, now time.Time) error {

	err := s.CanJudgeRequest(req, judger, now)
	if err != nil {
		return err
	}

	switch judgeStatus {
	case model.RequestStatusApproved:
		req.Approve(judger, now)
	case model.RequestStatusRejected:
		req.Reject(judger, now)
	default:
		return errors.Newf("invalid status(%s), status must be approved or rejected", judgeStatus)
	}

	return s.requestRepo.Save(ctx, req)
}

// CanJudge checks whether the user can judge the request.
func (s *Service) CanJudgeRequest(req *model.Request, judger *model.User, now time.Time) error {

	// user unavailable check is already done in middleware
	// so we don't need to check it here

	if judger.Role() == model.UserRoleRequester {
		return errors.WithDetail(
			errors.Newf("user is requester", req.Status()),
			string(custom_error.NotAllowed))
	}

	if req.RequesterUserID() == judger.ID() {
		return errors.WithDetail(
			errors.Newf("can't judge own request", req.Status()),
			string(custom_error.InvalidArgument))
	}

	if !req.IsPending() {
		return errors.WithDetail(
			errors.Newf("invalid status(%s), status must be pending", req.Status()),
			string(custom_error.InvalidArgument))
	}

	if req.IsExpired(now) {
		return errors.WithDetail(
			errors.Newf("request is expired", req.Status()),
			string(custom_error.InvalidArgument))
	}

	return nil
}

func (s *Service) CanDeleteRequest(req *model.Request, judger *model.User) error {

	// user unavailable check is already done in middleware
	// so we don't need to check it here

	if !req.IsPending() {
		return errors.WithDetail(
			errors.Newf("invalid status(%s), status must be pending", req.Status()),
			string(custom_error.InvalidArgument))
	}

	if req.RequesterUserID() == judger.ID() {
		return nil
	}

	if judger.IsAdmin() {
		return nil
	}

	return errors.WithDetail(
		errors.Newf("can't delete request", req.Status()),
		string(custom_error.NotAllowed))
}
