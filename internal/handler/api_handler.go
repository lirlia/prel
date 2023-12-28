package handler

import (
	"context"
	api "prel/api/prel_api"
	"prel/internal/gateway/repository"
	"prel/internal/model"
	"prel/pkg/custom_error"
	"sort"

	"github.com/cockroachdb/errors"
)

func (h *Handler) APIRequestsGet(ctx context.Context, params api.APIRequestsGetParams) (api.APIRequestsGetRes, error) {

	reqs, paging, err := h.usecase.RequestsWithPaging(ctx, params.PageID, int(params.Size))
	if err != nil {
		return nil, err
	}

	return &api.APIRequestsGetOK{
		TotalPage:   paging.TotalPage,
		CurrentPage: paging.CurrentPage,
		Requests:    convertRequests(reqs),
	}, nil
}

func (h *Handler) APIRequestsPost(ctx context.Context, req *api.APIRequestsPostReq) (api.APIRequestsPostRes, error) {

	request, err := h.usecase.CreateRequest(
		ctx, h.config.URL, req.GetProjectID(), req.GetReason(), req.GetIamRoles(), model.PeriodKey(req.GetPeriod()))

	if err != nil {
		return nil, err
	}

	return &api.APIRequestsPostOK{
		RequestID: request.ID(),
	}, nil
}

func (h *Handler) APIRequestsRequestIDPatch(ctx context.Context, req *api.APIRequestsRequestIDPatchReq, params api.APIRequestsRequestIDPatchParams) (api.APIRequestsRequestIDPatchRes, error) {
	err := h.usecase.JudgeRequest(ctx, h.config.URL, params.RequestID, convertRequestStatus(req.GetStatus()))
	if err != nil {
		return nil, err
	}
	return &api.APIRequestsRequestIDPatchNoContent{}, nil
}

func (h *Handler) APIRequestsRequestIDDelete(ctx context.Context, params api.APIRequestsRequestIDDeleteParams) (api.APIRequestsRequestIDDeleteRes, error) {
	err := h.usecase.DeleteRequest(ctx, params.RequestID)
	if err != nil {
		return nil, err
	}
	return &api.APIRequestsRequestIDDeleteNoContent{}, nil
}

func (h *Handler) APIUsersGet(ctx context.Context, params api.APIUsersGetParams) (api.APIUsersGetRes, error) {

	users, invitations, paging, err := h.usecase.UserWithPaging(ctx, params.PageID, int(params.Size))
	if err != nil {
		return nil, err
	}

	return &api.APIUsersGetOK{
		TotalPage:   paging.TotalPage,
		CurrentPage: paging.CurrentPage,
		Users:       convertUsers(users, invitations),
	}, nil
}

func (h *Handler) APIUsersUserIDPatch(ctx context.Context, req *api.APIUsersUserIDPatchReq, params api.APIUsersUserIDPatchParams) (api.APIUsersUserIDPatchRes, error) {
	err := h.usecase.UpdateUser(ctx, params.UserID, req.IsAvailable, model.UserRole(req.Role))
	if err != nil {
		return nil, err
	}

	return &api.APIUsersUserIDPatchNoContent{}, nil
}

func (h *Handler) APIInvitationsPost(ctx context.Context, req *api.APIInvitationsPostReq) (api.APIInvitationsPostRes, error) {
	err := h.usecase.Invite(ctx, req.Email, model.UserRole(req.Role))
	if err != nil {
		return nil, err
	}

	return &api.APIInvitationsPostNoContent{}, nil
}

func (h *Handler) APIIamRolesGet(ctx context.Context, params api.APIIamRolesGetParams) (api.APIIamRolesGetRes, error) {
	roleIDs, err := h.usecase.GetIamRoles(ctx, params.ProjectID)
	if err != nil {
		return nil, err
	}

	return &api.APIIamRolesGetOK{
		IamRoles: roleIDs,
	}, nil
}

func (h *Handler) APIIamRoleFilteringRulesGet(ctx context.Context) (api.APIIamRoleFilteringRulesGetRes, error) {
	user := model.GetUser(ctx)
	if !user.IsAdmin() {
		return nil, errors.WithDetail(errors.New("user is not admin"), string(custom_error.OnlyAdmin))
	}

	repo := repository.NewIamRoleFilteringRuleRepository()
	rules, err := repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	sort.Slice(rules, func(i, j int) bool {
		return rules[i].Pattern() < rules[j].Pattern()
	})

	return &api.APIIamRoleFilteringRulesGetOK{
		IamRoleFilteringRules: convertIamRoleFilteringRules(rules),
	}, nil
}

func (h *Handler) APIIamRoleFilteringRulesPost(ctx context.Context, req *api.APIIamRoleFilteringRulesPostReq) (api.APIIamRoleFilteringRulesPostRes, error) {
	user := model.GetUser(ctx)
	if !user.IsAdmin() {
		return nil, errors.WithDetail(errors.New("user is not admin"), string(custom_error.OnlyAdmin))
	}

	repo := repository.NewIamRoleFilteringRuleRepository()
	rule, err := model.NewIamRoleFilteringRule(req.Pattern, user.ID())
	if err != nil {
		return nil, err
	}

	err = repo.Save(ctx, rule)
	if err != nil {
		return nil, err
	}

	return &api.APIIamRoleFilteringRulesPostOK{
		IamRoleFilteringRule: convertIamRoleFilteringRule(rule),
	}, nil
}

func (h *Handler) APIIamRoleFilteringRulesRuleIDDelete(ctx context.Context, params api.APIIamRoleFilteringRulesRuleIDDeleteParams) (api.APIIamRoleFilteringRulesRuleIDDeleteRes, error) {
	user := model.GetUser(ctx)
	if !user.IsAdmin() {
		return nil, errors.WithDetail(errors.New("user is not admin"), string(custom_error.OnlyAdmin))
	}

	repo := repository.NewIamRoleFilteringRuleRepository()
	rule, err := repo.FindByID(ctx, params.RuleID)

	if err != nil {
		return nil, err
	}

	if rule == nil {
		return nil, errors.Newf("iam role filtering rule not found")
	}

	err = repo.Delete(ctx, rule)
	if err != nil {
		return nil, err
	}

	return &api.APIIamRoleFilteringRulesRuleIDDeleteNoContent{}, nil
}
