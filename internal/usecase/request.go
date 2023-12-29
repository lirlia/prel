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

func (uc *Usecase) CreateRequest(
	ctx context.Context,
	url, projectID, reason string,
	roles []string,
	period model.PeriodKey,
	requestExpireSeconds int) (req *model.Request, err error) {

	now := model.GetClock(ctx).Now()

	var user *model.User
	err = repository.NewTransactionManager().Transaction(ctx, func(ctx context.Context) error {

		err = uc.validateRoles(ctx, projectID, roles)
		if err != nil {
			return err
		}

		user = model.GetUser(ctx)
		requestExpiredAt := now.Add(time.Duration(requestExpireSeconds) * time.Second)
		req, err = model.NewRequest(user.ID(), projectID, roles, period, reason, now, requestExpiredAt)
		if err != nil {
			return err
		}
		return uc.requestRepo.Save(ctx, req)
	})

	if err != nil {
		return nil, err
	}

	if uc.n.CanSend() {
		url = fmt.Sprintf("%s/request/%s", url, req.ID())
		_, err = uc.n.SendRequestMessage(ctx, user.Email(), url, projectID, req.PeriodViewValue(), reason, roles, req.ExpiredAt())
		if err != nil {
			// if failed to send notification, only log the error
			logger.Get(ctx).Error(fmt.Sprintf("failed to send notification: %s", err))
			return req, nil
		}
	}

	return req, nil
}

func (uc *Usecase) JudgeRequest(ctx context.Context, url, requestID string, status model.RequestStatus) error {

	var req, oldReq *model.Request
	var requester *model.User
	var err error

	judger := model.GetUser(ctx)
	now := model.GetClock(ctx).Now()

	err = repository.NewTransactionManager().Transaction(ctx, func(ctx context.Context) error {
		req, err = uc.requestRepo.FindByID(ctx, requestID)
		if err != nil {
			return err
		}
		oldReq = req.Clone()

		// check roles is permitted to request
		err = uc.validateRoles(ctx, req.ProjectID(), req.IamRoles())
		if err != nil {
			return err
		}

		requester, err = uc.userRepo.FindByID(ctx, req.RequesterUserID())
		if err != nil {
			return err
		}

		return uc.service.Judge(ctx, req, status, requester, judger, now)
	})

	if err != nil {
		return err
	}

	// update iam policy
	if status == model.RequestStatusApproved {
		err = uc.c.SetIamPolicy(ctx, req.ProjectID(), req.IamRoles(), requester, req.CalculateRoleBindingExpiry(now))
		if err != nil {
			// if failed to update iam policy, rollback request status
			dbErr := uc.requestRepo.Save(ctx, oldReq)
			if dbErr != nil {
				logger.Get(ctx).Error(fmt.Sprintf("failed to rollback request status: %s", dbErr))
			}
			return err
		}
	}

	// send notification
	if uc.n.CanSend() {
		url = fmt.Sprintf("%s/request/%s", url, req.ID())
		_, err = uc.n.SendJudgeMessage(
			ctx, req.Status(), requester.Email(), judger.Email(), url, req.ProjectID(), req.Reason(), req.IamRoles(), req.CalculateRoleBindingExpiry(now))

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

		err = uc.service.CanDeleteRequest(req, user)
		if err != nil {
			return err
		}

		return uc.requestRepo.Delete(ctx, req)
	})
}
