package template

import (
	"embed"
	"html/template"
	"prel/internal/model"
)

//go:embed templates/*
var Files embed.FS

const (
	AdminRequestPageTpl          = "templates/admin_request.tpl"
	AdminUserPageTpl             = "templates/admin_user.tpl"
	AdminIamRoleFilteringPageTpl = "templates/admin_iam_role_filtering.tpl"
	AdminSettingPageTpl          = "templates/admin_setting.tpl"
	HeaderTpl                    = "templates/_header.tpl"
	ErrorPageTpl                 = "templates/error.tpl"
	IndexPageTpl                 = "templates/index.tpl"
	RequestPageTpl               = "templates/request.tpl"
	RequestFormPageTpl           = "templates/request_form.tpl"
)

type ErrorPageData struct {
	Name        string
	Description template.HTML
}

type PageData struct {
	HeaderData       HeaderData
	RequestPage      RequestPage
	RequestFormPage  RequestFormPage
	AdminListPage    AdminListPage
	AdminSettingPage AdminSettingPage
}

type HeaderData struct {
	IsAdmin bool
	AppName string
}

type RequestFormPage struct {
	Email    string
	Projects []*Project
	Periods  []*Period
}

type Project struct {
	Name      string
	ProjectID string
}

type IamRole struct {
	Name string
}

type Period struct {
	Key   model.PeriodKey
	Value string
}

type RequestPage struct {
	Requests []*Request
}

type Request struct {
	ID          string
	CanJudge    bool
	CanDelete   bool
	Requester   string
	Judger      string
	Status      string
	ProjectID   string
	IamRoles    []string
	Period      string
	Reason      string
	RequestedAt string
	JudgedAt    string
	ExpiredAt   string
}

type AdminListPage struct {
	Options   []int
	UserRoles []string
}

type AdminSettingPage struct {
	NotificationMessageForRequest string
	NotificationMessageForJudge   string
}

func NewHeaderData(isAdmin bool, appName string) HeaderData {
	return HeaderData{
		IsAdmin: isAdmin,
		AppName: appName,
	}
}
