package handler

import (
	api "prel/api/prel_api"
	"prel/internal/model"
	tpl "prel/web/template"
	"sort"
	"time"

	"google.golang.org/api/cloudresourcemanager/v1"
)

func convertTplRequests(user *model.User, reqs []*model.Request) []*tpl.Request {
	requests := make([]*tpl.Request, 0, len(reqs))
	for _, r := range reqs {
		requests = append(requests, convertTplRequest(user, r))
	}
	return requests
}

func convertTplRequest(user *model.User, req *model.Request) *tpl.Request {

	j := req.JudgedAt()
	judgedAt := j.Format(time.RFC3339)
	if j.IsZero() {
		judgedAt = ""
	}

	return &tpl.Request{
		ID:          req.ID(),
		CanJudge:    user.CanJudge(req) == nil,
		CanDelete:   user.CanDelete(req) == nil,
		Requester:   req.RequesterEmail(),
		Judger:      req.JudgerEmail(),
		Status:      string(req.Status()),
		ProjectID:   req.ProjectID(),
		IamRoles:    req.IamRoles(),
		Period:      req.PeriodViewValue(),
		Reason:      req.Reason(),
		RequestedAt: req.RequestedAt().Format(time.RFC3339),
		JudgedAt:    judgedAt,
		ExpiredAt:   req.ExpiredAt().Format(time.RFC3339),
	}
}

func convertTplProjects(pjs []*cloudresourcemanager.Project) []*tpl.Project {
	projects := make([]*tpl.Project, 0, len(pjs))
	for _, p := range pjs {
		projects = append(projects, convertTplProject(p))
	}
	return projects
}

func convertTplProject(pj *cloudresourcemanager.Project) *tpl.Project {
	return &tpl.Project{
		Name:      pj.Name,
		ProjectID: pj.ProjectId,
	}
}

func convertUserRoles(roles model.UserRoles) []string {
	rolesTpl := make([]string, 0, len(roles))
	for _, r := range roles {
		rolesTpl = append(rolesTpl, convertUserRole(r))
	}
	return rolesTpl
}

func convertUserRole(role model.UserRole) string {
	return string(role)
}

func convertTplPeriods() []*tpl.Period {
	periods := make([]*tpl.Period, 0, len(model.PeriodMap()))
	for k, v := range model.PeriodMap() {
		periods = append(periods, &tpl.Period{
			Key:   k,
			Value: v,
		})
	}

	sort.Slice(periods, func(i, j int) bool {
		return periods[i].Key < periods[j].Key
	})

	return periods
}

func convertRequestStatus(status api.JudgeStatus) model.RequestStatus {
	switch status {
	case api.JudgeStatusApprove:
		return model.RequestStatusApproved
	case api.JudgeStatusReject:
		return model.RequestStatusRejected
	default:
		return model.RequestStatusUnknown
	}
}

func convertRequests(requests model.Requests) []api.Request {
	res := make([]api.Request, 0, len(requests))
	for _, request := range requests {
		res = append(res, convertRequest(request))
	}
	return res
}

func convertRequest(request *model.Request) api.Request {
	return api.Request{
		Requester:   request.RequesterEmail(),
		Judger:      request.JudgerEmail(),
		ProjectID:   request.ProjectID(),
		IamRoles:    request.IamRoles(),
		Period:      request.PeriodViewValue(),
		Reason:      request.Reason(),
		Status:      api.RequestStatus(request.Status()),
		RequestTime: request.RequestedAt(),
		JudgeTime:   request.JudgedAt(),
		ExpireTime:  request.ExpiredAt(),
	}
}

func convertUser(user *model.User, isInvited bool) api.User {
	// for invitee (not signin yet)
	var lastSigninTime api.OptDateTime
	if user.LastSignInAt().IsZero() {
		lastSigninTime = api.OptDateTime{
			Set: false,
		}
	} else {
		lastSigninTime = api.OptDateTime{
			Value: user.LastSignInAt(),
			Set:   true,
		}
	}
	return api.User{
		ID:             string(user.ID()),
		Email:          user.Email(),
		IsAvailable:    user.IsAvailable(),
		Role:           api.UserRole(user.Role()),
		LastSigninTime: lastSigninTime,
		IsInvited:      isInvited,
	}
}

func convertUsers(users model.Users, invitations model.Invitations) []api.User {
	m := invitations.ToMap()
	res := make([]api.User, 0, len(users))
	for _, user := range users {
		_, ok := m[string(user.ID())]
		res = append(res, convertUser(user, ok))
	}
	return res
}

func convertIamRoleFilteringRule(rule *model.IamRoleFilteringRule) api.IamRoleFilteringRule {
	return api.IamRoleFilteringRule{
		ID:      rule.ID(),
		Pattern: rule.Pattern(),
	}
}

func convertIamRoleFilteringRules(rules model.IamRoleFilteringRules) []api.IamRoleFilteringRule {
	res := make([]api.IamRoleFilteringRule, 0, len(rules))
	for _, rule := range rules {
		res = append(res, convertIamRoleFilteringRule(rule))
	}
	return res
}
