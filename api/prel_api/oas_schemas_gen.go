// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"net/url"
	"time"

	"github.com/go-faster/errors"
)

type APIIamRoleFilteringRulesGetOK struct {
	IamRoleFilteringRules []IamRoleFilteringRule `json:"iamRoleFilteringRules"`
}

// GetIamRoleFilteringRules returns the value of IamRoleFilteringRules.
func (s *APIIamRoleFilteringRulesGetOK) GetIamRoleFilteringRules() []IamRoleFilteringRule {
	return s.IamRoleFilteringRules
}

// SetIamRoleFilteringRules sets the value of IamRoleFilteringRules.
func (s *APIIamRoleFilteringRulesGetOK) SetIamRoleFilteringRules(val []IamRoleFilteringRule) {
	s.IamRoleFilteringRules = val
}

func (*APIIamRoleFilteringRulesGetOK) aPIIamRoleFilteringRulesGetRes() {}

type APIIamRoleFilteringRulesPostOK struct {
	IamRoleFilteringRule IamRoleFilteringRule `json:"iamRoleFilteringRule"`
}

// GetIamRoleFilteringRule returns the value of IamRoleFilteringRule.
func (s *APIIamRoleFilteringRulesPostOK) GetIamRoleFilteringRule() IamRoleFilteringRule {
	return s.IamRoleFilteringRule
}

// SetIamRoleFilteringRule sets the value of IamRoleFilteringRule.
func (s *APIIamRoleFilteringRulesPostOK) SetIamRoleFilteringRule(val IamRoleFilteringRule) {
	s.IamRoleFilteringRule = val
}

func (*APIIamRoleFilteringRulesPostOK) aPIIamRoleFilteringRulesPostRes() {}

type APIIamRoleFilteringRulesPostReq struct {
	Pattern string `json:"pattern"`
}

// GetPattern returns the value of Pattern.
func (s *APIIamRoleFilteringRulesPostReq) GetPattern() string {
	return s.Pattern
}

// SetPattern sets the value of Pattern.
func (s *APIIamRoleFilteringRulesPostReq) SetPattern(val string) {
	s.Pattern = val
}

// APIIamRoleFilteringRulesRuleIDDeleteNoContent is response for APIIamRoleFilteringRulesRuleIDDelete operation.
type APIIamRoleFilteringRulesRuleIDDeleteNoContent struct{}

func (*APIIamRoleFilteringRulesRuleIDDeleteNoContent) aPIIamRoleFilteringRulesRuleIDDeleteRes() {}

type APIIamRolesGetOK struct {
	IamRoles []string `json:"iamRoles"`
}

// GetIamRoles returns the value of IamRoles.
func (s *APIIamRolesGetOK) GetIamRoles() []string {
	return s.IamRoles
}

// SetIamRoles sets the value of IamRoles.
func (s *APIIamRolesGetOK) SetIamRoles(val []string) {
	s.IamRoles = val
}

func (*APIIamRolesGetOK) aPIIamRolesGetRes() {}

// APIInvitationsPostNoContent is response for APIInvitationsPost operation.
type APIInvitationsPostNoContent struct{}

func (*APIInvitationsPostNoContent) aPIInvitationsPostRes() {}

type APIInvitationsPostReq struct {
	Email string   `json:"email"`
	Role  UserRole `json:"role"`
}

// GetEmail returns the value of Email.
func (s *APIInvitationsPostReq) GetEmail() string {
	return s.Email
}

// GetRole returns the value of Role.
func (s *APIInvitationsPostReq) GetRole() UserRole {
	return s.Role
}

// SetEmail sets the value of Email.
func (s *APIInvitationsPostReq) SetEmail(val string) {
	s.Email = val
}

// SetRole sets the value of Role.
func (s *APIInvitationsPostReq) SetRole(val UserRole) {
	s.Role = val
}

type APIRequestsGetOK struct {
	// Total number of pages.
	TotalPage int `json:"totalPage"`
	// Current page number.
	CurrentPage int       `json:"currentPage"`
	Requests    []Request `json:"requests"`
}

// GetTotalPage returns the value of TotalPage.
func (s *APIRequestsGetOK) GetTotalPage() int {
	return s.TotalPage
}

// GetCurrentPage returns the value of CurrentPage.
func (s *APIRequestsGetOK) GetCurrentPage() int {
	return s.CurrentPage
}

// GetRequests returns the value of Requests.
func (s *APIRequestsGetOK) GetRequests() []Request {
	return s.Requests
}

// SetTotalPage sets the value of TotalPage.
func (s *APIRequestsGetOK) SetTotalPage(val int) {
	s.TotalPage = val
}

// SetCurrentPage sets the value of CurrentPage.
func (s *APIRequestsGetOK) SetCurrentPage(val int) {
	s.CurrentPage = val
}

// SetRequests sets the value of Requests.
func (s *APIRequestsGetOK) SetRequests(val []Request) {
	s.Requests = val
}

func (*APIRequestsGetOK) aPIRequestsGetRes() {}

type APIRequestsPostOK struct {
	RequestID string `json:"requestID"`
}

// GetRequestID returns the value of RequestID.
func (s *APIRequestsPostOK) GetRequestID() string {
	return s.RequestID
}

// SetRequestID sets the value of RequestID.
func (s *APIRequestsPostOK) SetRequestID(val string) {
	s.RequestID = val
}

func (*APIRequestsPostOK) aPIRequestsPostRes() {}

type APIRequestsPostReq struct {
	ProjectID string   `json:"projectID"`
	IamRoles  []string `json:"iamRoles"`
	// Available duration(minutes).
	Period APIRequestsPostReqPeriod `json:"period"`
	Reason string                   `json:"reason"`
}

// GetProjectID returns the value of ProjectID.
func (s *APIRequestsPostReq) GetProjectID() string {
	return s.ProjectID
}

// GetIamRoles returns the value of IamRoles.
func (s *APIRequestsPostReq) GetIamRoles() []string {
	return s.IamRoles
}

// GetPeriod returns the value of Period.
func (s *APIRequestsPostReq) GetPeriod() APIRequestsPostReqPeriod {
	return s.Period
}

// GetReason returns the value of Reason.
func (s *APIRequestsPostReq) GetReason() string {
	return s.Reason
}

// SetProjectID sets the value of ProjectID.
func (s *APIRequestsPostReq) SetProjectID(val string) {
	s.ProjectID = val
}

// SetIamRoles sets the value of IamRoles.
func (s *APIRequestsPostReq) SetIamRoles(val []string) {
	s.IamRoles = val
}

// SetPeriod sets the value of Period.
func (s *APIRequestsPostReq) SetPeriod(val APIRequestsPostReqPeriod) {
	s.Period = val
}

// SetReason sets the value of Reason.
func (s *APIRequestsPostReq) SetReason(val string) {
	s.Reason = val
}

// Available duration(minutes).
type APIRequestsPostReqPeriod int

const (
	APIRequestsPostReqPeriod5     APIRequestsPostReqPeriod = 5
	APIRequestsPostReqPeriod10    APIRequestsPostReqPeriod = 10
	APIRequestsPostReqPeriod30    APIRequestsPostReqPeriod = 30
	APIRequestsPostReqPeriod60    APIRequestsPostReqPeriod = 60
	APIRequestsPostReqPeriod720   APIRequestsPostReqPeriod = 720
	APIRequestsPostReqPeriod1440  APIRequestsPostReqPeriod = 1440
	APIRequestsPostReqPeriod2880  APIRequestsPostReqPeriod = 2880
	APIRequestsPostReqPeriod4320  APIRequestsPostReqPeriod = 4320
	APIRequestsPostReqPeriod10080 APIRequestsPostReqPeriod = 10080
	APIRequestsPostReqPeriod20160 APIRequestsPostReqPeriod = 20160
)

// AllValues returns all APIRequestsPostReqPeriod values.
func (APIRequestsPostReqPeriod) AllValues() []APIRequestsPostReqPeriod {
	return []APIRequestsPostReqPeriod{
		APIRequestsPostReqPeriod5,
		APIRequestsPostReqPeriod10,
		APIRequestsPostReqPeriod30,
		APIRequestsPostReqPeriod60,
		APIRequestsPostReqPeriod720,
		APIRequestsPostReqPeriod1440,
		APIRequestsPostReqPeriod2880,
		APIRequestsPostReqPeriod4320,
		APIRequestsPostReqPeriod10080,
		APIRequestsPostReqPeriod20160,
	}
}

// APIRequestsRequestIDDeleteNoContent is response for APIRequestsRequestIDDelete operation.
type APIRequestsRequestIDDeleteNoContent struct{}

func (*APIRequestsRequestIDDeleteNoContent) aPIRequestsRequestIDDeleteRes() {}

// APIRequestsRequestIDPatchNoContent is response for APIRequestsRequestIDPatch operation.
type APIRequestsRequestIDPatchNoContent struct{}

func (*APIRequestsRequestIDPatchNoContent) aPIRequestsRequestIDPatchRes() {}

type APIRequestsRequestIDPatchReq struct {
	// Request status.
	Status JudgeStatus `json:"status"`
}

// GetStatus returns the value of Status.
func (s *APIRequestsRequestIDPatchReq) GetStatus() JudgeStatus {
	return s.Status
}

// SetStatus sets the value of Status.
func (s *APIRequestsRequestIDPatchReq) SetStatus(val JudgeStatus) {
	s.Status = val
}

// APISettingsPatchNoContent is response for APISettingsPatch operation.
type APISettingsPatchNoContent struct{}

func (*APISettingsPatchNoContent) aPISettingsPatchRes() {}

type APISettingsPatchReq struct {
	NotificationMessageForRequest OptString `json:"notificationMessageForRequest"`
	NotificationMessageForJudge   OptString `json:"notificationMessageForJudge"`
}

// GetNotificationMessageForRequest returns the value of NotificationMessageForRequest.
func (s *APISettingsPatchReq) GetNotificationMessageForRequest() OptString {
	return s.NotificationMessageForRequest
}

// GetNotificationMessageForJudge returns the value of NotificationMessageForJudge.
func (s *APISettingsPatchReq) GetNotificationMessageForJudge() OptString {
	return s.NotificationMessageForJudge
}

// SetNotificationMessageForRequest sets the value of NotificationMessageForRequest.
func (s *APISettingsPatchReq) SetNotificationMessageForRequest(val OptString) {
	s.NotificationMessageForRequest = val
}

// SetNotificationMessageForJudge sets the value of NotificationMessageForJudge.
func (s *APISettingsPatchReq) SetNotificationMessageForJudge(val OptString) {
	s.NotificationMessageForJudge = val
}

type APIUsersGetOK struct {
	// Total number of pages.
	TotalPage int `json:"totalPage"`
	// Current page number.
	CurrentPage int    `json:"currentPage"`
	Users       []User `json:"users"`
}

// GetTotalPage returns the value of TotalPage.
func (s *APIUsersGetOK) GetTotalPage() int {
	return s.TotalPage
}

// GetCurrentPage returns the value of CurrentPage.
func (s *APIUsersGetOK) GetCurrentPage() int {
	return s.CurrentPage
}

// GetUsers returns the value of Users.
func (s *APIUsersGetOK) GetUsers() []User {
	return s.Users
}

// SetTotalPage sets the value of TotalPage.
func (s *APIUsersGetOK) SetTotalPage(val int) {
	s.TotalPage = val
}

// SetCurrentPage sets the value of CurrentPage.
func (s *APIUsersGetOK) SetCurrentPage(val int) {
	s.CurrentPage = val
}

// SetUsers sets the value of Users.
func (s *APIUsersGetOK) SetUsers(val []User) {
	s.Users = val
}

func (*APIUsersGetOK) aPIUsersGetRes() {}

// APIUsersUserIDPatchNoContent is response for APIUsersUserIDPatch operation.
type APIUsersUserIDPatchNoContent struct{}

func (*APIUsersUserIDPatchNoContent) aPIUsersUserIDPatchRes() {}

type APIUsersUserIDPatchReq struct {
	// User account available or not.
	IsAvailable bool     `json:"isAvailable"`
	Role        UserRole `json:"role"`
}

// GetIsAvailable returns the value of IsAvailable.
func (s *APIUsersUserIDPatchReq) GetIsAvailable() bool {
	return s.IsAvailable
}

// GetRole returns the value of Role.
func (s *APIUsersUserIDPatchReq) GetRole() UserRole {
	return s.Role
}

// SetIsAvailable sets the value of IsAvailable.
func (s *APIUsersUserIDPatchReq) SetIsAvailable(val bool) {
	s.IsAvailable = val
}

// SetRole sets the value of Role.
func (s *APIUsersUserIDPatchReq) SetRole(val UserRole) {
	s.Role = val
}

type AdminIamRoleFilteringGetOK struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s AdminIamRoleFilteringGetOK) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

func (*AdminIamRoleFilteringGetOK) adminIamRoleFilteringGetRes() {}

type AdminRequestGetOK struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s AdminRequestGetOK) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

func (*AdminRequestGetOK) adminRequestGetRes() {}

type AdminSettingGetOK struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s AdminSettingGetOK) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

func (*AdminSettingGetOK) adminSettingGetRes() {}

type AdminUserGetOK struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s AdminUserGetOK) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

func (*AdminUserGetOK) adminUserGetRes() {}

// AuthGoogleCallbackGetTemporaryRedirect is response for AuthGoogleCallbackGet operation.
type AuthGoogleCallbackGetTemporaryRedirect struct {
	Location  OptString
	SetCookie OptString
}

// GetLocation returns the value of Location.
func (s *AuthGoogleCallbackGetTemporaryRedirect) GetLocation() OptString {
	return s.Location
}

// GetSetCookie returns the value of SetCookie.
func (s *AuthGoogleCallbackGetTemporaryRedirect) GetSetCookie() OptString {
	return s.SetCookie
}

// SetLocation sets the value of Location.
func (s *AuthGoogleCallbackGetTemporaryRedirect) SetLocation(val OptString) {
	s.Location = val
}

// SetSetCookie sets the value of SetCookie.
func (s *AuthGoogleCallbackGetTemporaryRedirect) SetSetCookie(val OptString) {
	s.SetCookie = val
}

func (*AuthGoogleCallbackGetTemporaryRedirect) authGoogleCallbackGetRes() {}

type BadRequest struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s BadRequest) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

func (*BadRequest) aPIIamRoleFilteringRulesPostRes()         {}
func (*BadRequest) aPIIamRoleFilteringRulesRuleIDDeleteRes() {}
func (*BadRequest) aPIIamRolesGetRes()                       {}
func (*BadRequest) aPIInvitationsPostRes()                   {}
func (*BadRequest) aPIRequestsPostRes()                      {}
func (*BadRequest) aPIRequestsRequestIDDeleteRes()           {}
func (*BadRequest) aPIRequestsRequestIDPatchRes()            {}
func (*BadRequest) aPISettingsPatchRes()                     {}
func (*BadRequest) authGoogleCallbackGetRes()                {}
func (*BadRequest) signinPostRes()                           {}

type CookieAuth struct {
	APIKey string
}

// GetAPIKey returns the value of APIKey.
func (s *CookieAuth) GetAPIKey() string {
	return s.APIKey
}

// SetAPIKey sets the value of APIKey.
func (s *CookieAuth) SetAPIKey(val string) {
	s.APIKey = val
}

type Forbidden struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s Forbidden) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

func (*Forbidden) aPIIamRoleFilteringRulesGetRes()          {}
func (*Forbidden) aPIIamRoleFilteringRulesRuleIDDeleteRes() {}
func (*Forbidden) aPIInvitationsPostRes()                   {}
func (*Forbidden) aPIRequestsGetRes()                       {}
func (*Forbidden) aPIRequestsRequestIDDeleteRes()           {}
func (*Forbidden) aPIRequestsRequestIDPatchRes()            {}
func (*Forbidden) aPISettingsPatchRes()                     {}
func (*Forbidden) aPIUsersGetRes()                          {}
func (*Forbidden) aPIUsersUserIDPatchRes()                  {}
func (*Forbidden) adminIamRoleFilteringGetRes()             {}
func (*Forbidden) adminRequestGetRes()                      {}
func (*Forbidden) adminSettingGetRes()                      {}
func (*Forbidden) adminUserGetRes()                         {}
func (*Forbidden) authGoogleCallbackGetRes()                {}

type GetOK struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s GetOK) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

func (*GetOK) getRes() {}

type GetSeeOther struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s GetSeeOther) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

// GetSeeOtherHeaders wraps GetSeeOther with response headers.
type GetSeeOtherHeaders struct {
	Location OptURI
	Response GetSeeOther
}

// GetLocation returns the value of Location.
func (s *GetSeeOtherHeaders) GetLocation() OptURI {
	return s.Location
}

// GetResponse returns the value of Response.
func (s *GetSeeOtherHeaders) GetResponse() GetSeeOther {
	return s.Response
}

// SetLocation sets the value of Location.
func (s *GetSeeOtherHeaders) SetLocation(val OptURI) {
	s.Location = val
}

// SetResponse sets the value of Response.
func (s *GetSeeOtherHeaders) SetResponse(val GetSeeOther) {
	s.Response = val
}

func (*GetSeeOtherHeaders) getRes() {}

type HealthGetNoContent struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s HealthGetNoContent) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

func (*HealthGetNoContent) healthGetRes() {}

// Ref: #/components/schemas/iamRoleFilteringRule
type IamRoleFilteringRule struct {
	ID      string `json:"id"`
	Pattern string `json:"pattern"`
}

// GetID returns the value of ID.
func (s *IamRoleFilteringRule) GetID() string {
	return s.ID
}

// GetPattern returns the value of Pattern.
func (s *IamRoleFilteringRule) GetPattern() string {
	return s.Pattern
}

// SetID sets the value of ID.
func (s *IamRoleFilteringRule) SetID(val string) {
	s.ID = val
}

// SetPattern sets the value of Pattern.
func (s *IamRoleFilteringRule) SetPattern(val string) {
	s.Pattern = val
}

type InternalServerError struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s InternalServerError) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

// InternalServerErrorStatusCode wraps InternalServerError with StatusCode.
type InternalServerErrorStatusCode struct {
	StatusCode int
	Response   InternalServerError
}

// GetStatusCode returns the value of StatusCode.
func (s *InternalServerErrorStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *InternalServerErrorStatusCode) GetResponse() InternalServerError {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *InternalServerErrorStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *InternalServerErrorStatusCode) SetResponse(val InternalServerError) {
	s.Response = val
}

func (*InternalServerErrorStatusCode) aPIIamRoleFilteringRulesGetRes()          {}
func (*InternalServerErrorStatusCode) aPIIamRoleFilteringRulesPostRes()         {}
func (*InternalServerErrorStatusCode) aPIIamRoleFilteringRulesRuleIDDeleteRes() {}
func (*InternalServerErrorStatusCode) aPIIamRolesGetRes()                       {}
func (*InternalServerErrorStatusCode) aPIInvitationsPostRes()                   {}
func (*InternalServerErrorStatusCode) aPIRequestsGetRes()                       {}
func (*InternalServerErrorStatusCode) aPIRequestsPostRes()                      {}
func (*InternalServerErrorStatusCode) aPIRequestsRequestIDDeleteRes()           {}
func (*InternalServerErrorStatusCode) aPIRequestsRequestIDPatchRes()            {}
func (*InternalServerErrorStatusCode) aPISettingsPatchRes()                     {}
func (*InternalServerErrorStatusCode) aPIUsersGetRes()                          {}
func (*InternalServerErrorStatusCode) aPIUsersUserIDPatchRes()                  {}
func (*InternalServerErrorStatusCode) adminIamRoleFilteringGetRes()             {}
func (*InternalServerErrorStatusCode) adminRequestGetRes()                      {}
func (*InternalServerErrorStatusCode) adminSettingGetRes()                      {}
func (*InternalServerErrorStatusCode) adminUserGetRes()                         {}
func (*InternalServerErrorStatusCode) authGoogleCallbackGetRes()                {}
func (*InternalServerErrorStatusCode) getRes()                                  {}
func (*InternalServerErrorStatusCode) healthGetRes()                            {}
func (*InternalServerErrorStatusCode) requestFormGetRes()                       {}
func (*InternalServerErrorStatusCode) requestGetRes()                           {}
func (*InternalServerErrorStatusCode) requestRequestIDGetRes()                  {}
func (*InternalServerErrorStatusCode) signinPostRes()                           {}
func (*InternalServerErrorStatusCode) signoutPostRes()                          {}

// Ref: #/components/schemas/judgeStatus
type JudgeStatus string

const (
	JudgeStatusApprove JudgeStatus = "approve"
	JudgeStatusReject  JudgeStatus = "reject"
)

// AllValues returns all JudgeStatus values.
func (JudgeStatus) AllValues() []JudgeStatus {
	return []JudgeStatus{
		JudgeStatusApprove,
		JudgeStatusReject,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s JudgeStatus) MarshalText() ([]byte, error) {
	switch s {
	case JudgeStatusApprove:
		return []byte(s), nil
	case JudgeStatusReject:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *JudgeStatus) UnmarshalText(data []byte) error {
	switch JudgeStatus(data) {
	case JudgeStatusApprove:
		*s = JudgeStatusApprove
		return nil
	case JudgeStatusReject:
		*s = JudgeStatusReject
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type NotFound struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s NotFound) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

func (*NotFound) requestRequestIDGetRes() {}

// NewOptDateTime returns new OptDateTime with value set to v.
func NewOptDateTime(v time.Time) OptDateTime {
	return OptDateTime{
		Value: v,
		Set:   true,
	}
}

// OptDateTime is optional time.Time.
type OptDateTime struct {
	Value time.Time
	Set   bool
}

// IsSet returns true if OptDateTime was set.
func (o OptDateTime) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptDateTime) Reset() {
	var v time.Time
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptDateTime) SetTo(v time.Time) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptDateTime) Get() (v time.Time, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptDateTime) Or(d time.Time) time.Time {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptURI returns new OptURI with value set to v.
func NewOptURI(v url.URL) OptURI {
	return OptURI{
		Value: v,
		Set:   true,
	}
}

// OptURI is optional url.URL.
type OptURI struct {
	Value url.URL
	Set   bool
}

// IsSet returns true if OptURI was set.
func (o OptURI) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptURI) Reset() {
	var v url.URL
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptURI) SetTo(v url.URL) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptURI) Get() (v url.URL, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptURI) Or(d url.URL) url.URL {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/pageSize
type PageSize int

const (
	PageSize25  PageSize = 25
	PageSize50  PageSize = 50
	PageSize100 PageSize = 100
	PageSize200 PageSize = 200
)

// AllValues returns all PageSize values.
func (PageSize) AllValues() []PageSize {
	return []PageSize{
		PageSize25,
		PageSize50,
		PageSize100,
		PageSize200,
	}
}

// Request.
// Ref: #/components/schemas/request
type Request struct {
	Requester   string        `json:"requester"`
	Judger      string        `json:"judger"`
	ProjectID   string        `json:"projectID"`
	IamRoles    []string      `json:"iamRoles"`
	Period      string        `json:"period"`
	Reason      string        `json:"reason"`
	Status      RequestStatus `json:"status"`
	RequestTime time.Time     `json:"requestTime"`
	JudgeTime   time.Time     `json:"judgeTime"`
	ExpireTime  time.Time     `json:"expireTime"`
}

// GetRequester returns the value of Requester.
func (s *Request) GetRequester() string {
	return s.Requester
}

// GetJudger returns the value of Judger.
func (s *Request) GetJudger() string {
	return s.Judger
}

// GetProjectID returns the value of ProjectID.
func (s *Request) GetProjectID() string {
	return s.ProjectID
}

// GetIamRoles returns the value of IamRoles.
func (s *Request) GetIamRoles() []string {
	return s.IamRoles
}

// GetPeriod returns the value of Period.
func (s *Request) GetPeriod() string {
	return s.Period
}

// GetReason returns the value of Reason.
func (s *Request) GetReason() string {
	return s.Reason
}

// GetStatus returns the value of Status.
func (s *Request) GetStatus() RequestStatus {
	return s.Status
}

// GetRequestTime returns the value of RequestTime.
func (s *Request) GetRequestTime() time.Time {
	return s.RequestTime
}

// GetJudgeTime returns the value of JudgeTime.
func (s *Request) GetJudgeTime() time.Time {
	return s.JudgeTime
}

// GetExpireTime returns the value of ExpireTime.
func (s *Request) GetExpireTime() time.Time {
	return s.ExpireTime
}

// SetRequester sets the value of Requester.
func (s *Request) SetRequester(val string) {
	s.Requester = val
}

// SetJudger sets the value of Judger.
func (s *Request) SetJudger(val string) {
	s.Judger = val
}

// SetProjectID sets the value of ProjectID.
func (s *Request) SetProjectID(val string) {
	s.ProjectID = val
}

// SetIamRoles sets the value of IamRoles.
func (s *Request) SetIamRoles(val []string) {
	s.IamRoles = val
}

// SetPeriod sets the value of Period.
func (s *Request) SetPeriod(val string) {
	s.Period = val
}

// SetReason sets the value of Reason.
func (s *Request) SetReason(val string) {
	s.Reason = val
}

// SetStatus sets the value of Status.
func (s *Request) SetStatus(val RequestStatus) {
	s.Status = val
}

// SetRequestTime sets the value of RequestTime.
func (s *Request) SetRequestTime(val time.Time) {
	s.RequestTime = val
}

// SetJudgeTime sets the value of JudgeTime.
func (s *Request) SetJudgeTime(val time.Time) {
	s.JudgeTime = val
}

// SetExpireTime sets the value of ExpireTime.
func (s *Request) SetExpireTime(val time.Time) {
	s.ExpireTime = val
}

type RequestFormGetOK struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s RequestFormGetOK) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

func (*RequestFormGetOK) requestFormGetRes() {}

type RequestGetOK struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s RequestGetOK) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

func (*RequestGetOK) requestGetRes() {}

type RequestRequestIDGetOK struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s RequestRequestIDGetOK) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

func (*RequestRequestIDGetOK) requestRequestIDGetRes() {}

type RequestStatus string

const (
	RequestStatusApproved RequestStatus = "approved"
	RequestStatusRejected RequestStatus = "rejected"
	RequestStatusPending  RequestStatus = "pending"
)

// AllValues returns all RequestStatus values.
func (RequestStatus) AllValues() []RequestStatus {
	return []RequestStatus{
		RequestStatusApproved,
		RequestStatusRejected,
		RequestStatusPending,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s RequestStatus) MarshalText() ([]byte, error) {
	switch s {
	case RequestStatusApproved:
		return []byte(s), nil
	case RequestStatusRejected:
		return []byte(s), nil
	case RequestStatusPending:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *RequestStatus) UnmarshalText(data []byte) error {
	switch RequestStatus(data) {
	case RequestStatusApproved:
		*s = RequestStatusApproved
		return nil
	case RequestStatusRejected:
		*s = RequestStatusRejected
		return nil
	case RequestStatusPending:
		*s = RequestStatusPending
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// SigninPostSeeOther is response for SigninPost operation.
type SigninPostSeeOther struct {
	Location  OptURI
	SetCookie OptString
}

// GetLocation returns the value of Location.
func (s *SigninPostSeeOther) GetLocation() OptURI {
	return s.Location
}

// GetSetCookie returns the value of SetCookie.
func (s *SigninPostSeeOther) GetSetCookie() OptString {
	return s.SetCookie
}

// SetLocation sets the value of Location.
func (s *SigninPostSeeOther) SetLocation(val OptURI) {
	s.Location = val
}

// SetSetCookie sets the value of SetCookie.
func (s *SigninPostSeeOther) SetSetCookie(val OptString) {
	s.SetCookie = val
}

func (*SigninPostSeeOther) signinPostRes() {}

type SignoutPostSeeOther struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s SignoutPostSeeOther) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

// SignoutPostSeeOtherHeaders wraps SignoutPostSeeOther with response headers.
type SignoutPostSeeOtherHeaders struct {
	Location  OptURI
	SetCookie OptString
	Response  SignoutPostSeeOther
}

// GetLocation returns the value of Location.
func (s *SignoutPostSeeOtherHeaders) GetLocation() OptURI {
	return s.Location
}

// GetSetCookie returns the value of SetCookie.
func (s *SignoutPostSeeOtherHeaders) GetSetCookie() OptString {
	return s.SetCookie
}

// GetResponse returns the value of Response.
func (s *SignoutPostSeeOtherHeaders) GetResponse() SignoutPostSeeOther {
	return s.Response
}

// SetLocation sets the value of Location.
func (s *SignoutPostSeeOtherHeaders) SetLocation(val OptURI) {
	s.Location = val
}

// SetSetCookie sets the value of SetCookie.
func (s *SignoutPostSeeOtherHeaders) SetSetCookie(val OptString) {
	s.SetCookie = val
}

// SetResponse sets the value of Response.
func (s *SignoutPostSeeOtherHeaders) SetResponse(val SignoutPostSeeOther) {
	s.Response = val
}

func (*SignoutPostSeeOtherHeaders) signoutPostRes() {}

type Unauthorized struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s Unauthorized) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

func (*Unauthorized) aPIIamRoleFilteringRulesGetRes()          {}
func (*Unauthorized) aPIIamRoleFilteringRulesPostRes()         {}
func (*Unauthorized) aPIIamRoleFilteringRulesRuleIDDeleteRes() {}
func (*Unauthorized) aPIIamRolesGetRes()                       {}
func (*Unauthorized) aPIInvitationsPostRes()                   {}
func (*Unauthorized) aPIRequestsGetRes()                       {}
func (*Unauthorized) aPIRequestsPostRes()                      {}
func (*Unauthorized) aPIRequestsRequestIDDeleteRes()           {}
func (*Unauthorized) aPIRequestsRequestIDPatchRes()            {}
func (*Unauthorized) aPISettingsPatchRes()                     {}
func (*Unauthorized) aPIUsersGetRes()                          {}
func (*Unauthorized) aPIUsersUserIDPatchRes()                  {}
func (*Unauthorized) adminIamRoleFilteringGetRes()             {}
func (*Unauthorized) adminRequestGetRes()                      {}
func (*Unauthorized) adminSettingGetRes()                      {}
func (*Unauthorized) adminUserGetRes()                         {}
func (*Unauthorized) requestFormGetRes()                       {}
func (*Unauthorized) requestGetRes()                           {}
func (*Unauthorized) requestRequestIDGetRes()                  {}

// User.
// Ref: #/components/schemas/user
type User struct {
	ID             string      `json:"id"`
	Email          string      `json:"email"`
	IsAvailable    bool        `json:"isAvailable"`
	Role           UserRole    `json:"role"`
	LastSigninTime OptDateTime `json:"lastSigninTime"`
	IsInvited      bool        `json:"isInvited"`
}

// GetID returns the value of ID.
func (s *User) GetID() string {
	return s.ID
}

// GetEmail returns the value of Email.
func (s *User) GetEmail() string {
	return s.Email
}

// GetIsAvailable returns the value of IsAvailable.
func (s *User) GetIsAvailable() bool {
	return s.IsAvailable
}

// GetRole returns the value of Role.
func (s *User) GetRole() UserRole {
	return s.Role
}

// GetLastSigninTime returns the value of LastSigninTime.
func (s *User) GetLastSigninTime() OptDateTime {
	return s.LastSigninTime
}

// GetIsInvited returns the value of IsInvited.
func (s *User) GetIsInvited() bool {
	return s.IsInvited
}

// SetID sets the value of ID.
func (s *User) SetID(val string) {
	s.ID = val
}

// SetEmail sets the value of Email.
func (s *User) SetEmail(val string) {
	s.Email = val
}

// SetIsAvailable sets the value of IsAvailable.
func (s *User) SetIsAvailable(val bool) {
	s.IsAvailable = val
}

// SetRole sets the value of Role.
func (s *User) SetRole(val UserRole) {
	s.Role = val
}

// SetLastSigninTime sets the value of LastSigninTime.
func (s *User) SetLastSigninTime(val OptDateTime) {
	s.LastSigninTime = val
}

// SetIsInvited sets the value of IsInvited.
func (s *User) SetIsInvited(val bool) {
	s.IsInvited = val
}

// Ref: #/components/schemas/userRole
type UserRole string

const (
	UserRoleRequester UserRole = "requester"
	UserRoleJudger    UserRole = "judger"
	UserRoleAdmin     UserRole = "admin"
)

// AllValues returns all UserRole values.
func (UserRole) AllValues() []UserRole {
	return []UserRole{
		UserRoleRequester,
		UserRoleJudger,
		UserRoleAdmin,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s UserRole) MarshalText() ([]byte, error) {
	switch s {
	case UserRoleRequester:
		return []byte(s), nil
	case UserRoleJudger:
		return []byte(s), nil
	case UserRoleAdmin:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *UserRole) UnmarshalText(data []byte) error {
	switch UserRole(data) {
	case UserRoleRequester:
		*s = UserRoleRequester
		return nil
	case UserRoleJudger:
		*s = UserRoleJudger
		return nil
	case UserRoleAdmin:
		*s = UserRoleAdmin
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}
