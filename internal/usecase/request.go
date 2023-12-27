package usecase

import (
	"context"
	"fmt"
	"prel/internal/gateway/repository"
	"prel/internal/model"
	"prel/pkg/custom_error"
	"prel/pkg/logger"
	"slices"
	"sort"
	"time"

	"github.com/cockroachdb/errors"
)

type Paging struct {
	TotalPage   int
	CurrentPage int
}

func (uc *Usecase) RequestsWithPaging(ctx context.Context, page int, pageSize int) (model.Requests, Paging, error) {
	user := model.GetUser(ctx)
	if !user.IsAdmin() {
		return nil, Paging{}, errors.WithDetail(errors.New("user is not admin"), string(custom_error.OnlyAdmin))
	}

	return uc.requestWithPaging(ctx, page, pageSize)
}

func (uc *Usecase) requestWithPaging(ctx context.Context, page, pageSize int) (model.Requests, Paging, error) {

	totalPage, currentPage, start, end, err := calcPager(ctx, uc.requestRepo, page, pageSize)
	if err != nil {
		return nil, Paging{}, err
	}

	reqs, err := uc.requestRepo.FindPaged(ctx, start, end-start)
	if err != nil {
		return nil, Paging{}, err
	}

	sort.Slice(reqs, func(i, j int) bool {
		return reqs[i].RequestedAt().After(reqs[j].RequestedAt())
	})

	return reqs, Paging{totalPage, currentPage}, nil
}

// check roles is permitted to request
func (uc *Usecase) validateRoles(ctx context.Context, projectID string, roles []string) error {

	roleIDs, err := uc.GetIamRoles(ctx, projectID)
	if err != nil {
		return err
	}

	for _, role := range roles {
		if !slices.Contains(roleIDs, role) {
			return errors.WithDetail(
				errors.Newf("invalid iam role: %s", role),
				string(custom_error.InvalidArgument))
		}
	}

	return nil
}

func (uc *Usecase) CreateRequest(ctx context.Context, url, projectID, reason string, roles []string, duration time.Duration) (req *model.Request, err error) {

	now := model.GetClock(ctx).Now()
	until := now.Add(duration)

	var user *model.User

	err = repository.NewTransactionManager().Transaction(ctx, func(ctx context.Context) error {

		err = uc.validateRoles(ctx, projectID, roles)
		if err != nil {
			return err
		}

		user = model.GetUser(ctx)
		req = model.NewRequest(user.ID(), projectID, roles, reason, now, until)
		return uc.requestRepo.Save(ctx, req)
	})

	if err != nil {
		return nil, err
	}

	if uc.n.CanSend() {
		url = fmt.Sprintf("%s/request/%s", url, req.ID())
		_, err = uc.n.SendRequestMessage(ctx, user.Email(), url, projectID, reason, roles, until)
		if err != nil {
			// if failed to send notification, only log the error
			logger.Get(ctx).Error(fmt.Sprintf("failed to send notification: %s", err))
			return req, nil
		}
	}

	return req, nil
}

func (uc *Usecase) JudgeRequest(ctx context.Context, url, requestID string, status model.RequestStatus) error {

	var req *model.Request
	var oldReq *model.Request
	var judger, requester *model.User
	var err error

	err = repository.NewTransactionManager().Transaction(ctx, func(ctx context.Context) error {
		req, err = uc.requestRepo.FindByID(ctx, requestID)
		if err != nil {
			return err
		}

		// check roles is permitted to request
		err = uc.validateRoles(ctx, req.ProjectID(), req.IamRoles())
		if err != nil {
			return err
		}

		oldReq = req.Clone()

		judgerUser := model.GetUser(ctx)
		err = judgerUser.CanJudge(req)
		if err != nil {
			return errors.Wrapf(err, "user(%s) can't judge request(%s)", judgerUser.ID(), req.ID())
		}

		ids := []model.UserID{req.RequesterUserID(), judgerUser.ID()}
		users, err := uc.userRepo.FindByIDs(ctx, ids)
		if err != nil {
			return err
		}

		judger = users.ByID(judgerUser.ID())
		if judger == nil {
			return errors.Newf("failed to find judger by id(%s)", judgerUser.ID())
		}

		requester = users.ByID(req.RequesterUserID())
		if requester == nil {
			return errors.Newf("failed to find request user by id(%s)", req.RequesterUserID())
		}

		now := model.GetClock(ctx).Now()
		switch status {
		case model.RequestStatusApproved:
			req.Approve(judger, now)
		case model.RequestStatusRejected:
			req.Reject(judger, now)
		default:
			return errors.Newf("invalid status(%s), status must be approved or rejected", status)
		}

		return uc.requestRepo.Save(ctx, req)
	})

	if err != nil {
		return err
	}

	pj := req.ProjectID()
	roles := req.IamRoles()
	until := req.ExpiredAt()

	// update iam policy
	if req.IsApprove() {
		err = uc.c.SetIamPolicy(ctx, pj, roles, requester, until)
		if err != nil {
			// if failed to update iam policy, rollback request status
			dbErr := uc.requestRepo.Save(ctx, oldReq)
			if dbErr != nil {
				logger.Get(ctx).Error(fmt.Sprintf("failed to rollback request status: %s", dbErr))
			}
			return err
		}
	}

	if uc.n.CanSend() {
		url = fmt.Sprintf("%s/request/%s", url, req.ID())
		_, err = uc.n.SendJudgeMessage(
			ctx, req.Status(), requester.Email(), judger.Email(), url, pj, req.Reason(), roles, until)

		if err != nil {
			// if failed to send notification, only log the error
			logger.Get(ctx).Error(fmt.Sprintf("failed to send notification: %s", err))
		}
	}
	return nil
}

func (uc *Usecase) DeleteRequest(ctx context.Context, requestID string) error {

	user := model.GetUser(ctx)
	return repository.NewTransactionManager().Transaction(ctx, func(ctx context.Context) error {
		req, err := uc.requestRepo.FindByID(ctx, requestID)
		if err != nil {
			return err
		}

		err = user.CanDelete(req)
		if err != nil {
			return err
		}

		return uc.requestRepo.Delete(ctx, req)
	})
}
