package handler

import (
	"context"
	api "prel/api/prel_api"
	"prel/internal/model"
	"sort"
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

	duration := model.PeriodTimeMap[model.PeriodKey(req.GetPeriod())]

	request, err := h.usecase.CreateRequest(
		ctx, h.config.URL, req.GetProjectID(), req.GetReason(), req.GetIamRoles(), duration)

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
	user := model.GetUser(ctx)

	// Currently, only roles that can be applied to user principal are returned.
	// so we set user as principal.
	roles, err := h.client.GetIamRoles(model.GetClock(ctx).Now(), params.ProjectID, user)
	if err != nil {
		return nil, err
	}

	roleIDs := make([]string, 0, len(roles))
	for _, role := range roles {
		roleIDs = append(roleIDs, role.Name)
	}

	sort.Slice(roleIDs, func(i, j int) bool {
		return roleIDs[i] < roleIDs[j]
	})

	return &api.APIIamRolesGetOK{
		IamRoles: roleIDs,
	}, nil
}
