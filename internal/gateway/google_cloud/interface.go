package google_cloud

//go:generate ../../../bin/mockgen -source=./interface.go -destination=./mock/google_cloud_mock.gen.go -package=google_cloud_mock

import (
	"google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/iam/v1"
)

type ListProjectsGetter interface {
	PageToken(token string) ListProjectsGetter               // *ProjectsListCall
	Do() (*cloudresourcemanager.ListProjectsResponse, error) // (*ListProjectsResponse, error)
}
type ResourceManagerService interface {
	GetIamPolicy(projectID string, request *cloudresourcemanager.GetIamPolicyRequest) IamPolicyGetter
	SetIamPolicy(projectID string, request *cloudresourcemanager.SetIamPolicyRequest) IamPolicySetter
	List() ListProjectsGetter //*ProjectsListCall
}

type IamPolicyGetter interface {
	Do() (*cloudresourcemanager.Policy, error)
}

type IamPolicySetter interface {
	Do() (*cloudresourcemanager.Policy, error)
}

type IamService interface {
	QueryGrantableRoles(req *iam.QueryGrantableRolesRequest) RolesQueryGrantableRolesCall
}

type RolesQueryGrantableRolesCall interface {
	Do() (*iam.QueryGrantableRolesResponse, error)
}

type ProjectsGetIamPolicyCallWrapper struct {
	Call *cloudresourcemanager.ProjectsGetIamPolicyCall
}

func (w *ProjectsGetIamPolicyCallWrapper) Do() (*cloudresourcemanager.Policy, error) {
	return w.Call.Do()
}

type ProjectsSetIamPolicyCallWrapper struct {
	Call *cloudresourcemanager.ProjectsSetIamPolicyCall
}

func (w *ProjectsSetIamPolicyCallWrapper) Do() (*cloudresourcemanager.Policy, error) {
	return w.Call.Do()
}

type ProjectsListCallWrapper struct {
	Call *cloudresourcemanager.ProjectsListCall
}

func (w *ProjectsListCallWrapper) PageToken(token string) ListProjectsGetter {
	return &ProjectsListCallWrapper{Call: w.Call.PageToken(token)}
}

func (w *ProjectsListCallWrapper) Do() (*cloudresourcemanager.ListProjectsResponse, error) {
	return w.Call.Do()
}

type RolesQueryGrantableRolesCallWrapper struct {
	Call *iam.RolesQueryGrantableRolesCall
}

func (w *RolesQueryGrantableRolesCallWrapper) Do() (*iam.QueryGrantableRolesResponse, error) {
	return w.Call.Do()
}

type ProjectsServiceWrapper struct {
	Service *cloudresourcemanager.ProjectsService
}

func (w *ProjectsServiceWrapper) GetIamPolicy(projectID string, request *cloudresourcemanager.GetIamPolicyRequest) IamPolicyGetter {
	return &ProjectsGetIamPolicyCallWrapper{Call: w.Service.GetIamPolicy(projectID, request)}
}

func (w *ProjectsServiceWrapper) SetIamPolicy(projectID string, request *cloudresourcemanager.SetIamPolicyRequest) IamPolicySetter {
	return &ProjectsSetIamPolicyCallWrapper{Call: w.Service.SetIamPolicy(projectID, request)}
}

func (w *ProjectsServiceWrapper) List() ListProjectsGetter {
	return &ProjectsListCallWrapper{Call: w.Service.List()}
}

type RolesServiceWrapper struct {
	Service *iam.RolesService
}

func (w *RolesServiceWrapper) QueryGrantableRoles(req *iam.QueryGrantableRolesRequest) RolesQueryGrantableRolesCall {
	return &RolesQueryGrantableRolesCallWrapper{Call: w.Service.QueryGrantableRoles(req)}
}
