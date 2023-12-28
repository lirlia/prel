package model

import (
	"prel/pkg/custom_error"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
)

type Request struct {
	id              string
	requesterUserID UserID
	requesterEmail  string
	judgerUserID    UserID
	judgerEmail     string
	status          RequestStatus
	projectID       string
	iamRoles        []string
	period          PeriodKey
	reason          string
	requestedAt     time.Time
	// available until (iam condition expired at)
	expiredAt time.Time
	// last reviewed at (approved or rejected)
	judgedAt time.Time
}

type Requests []*Request
type RequestStatus string

const (
	RequestStatusUnknown  RequestStatus = "unknown"
	RequestStatusPending  RequestStatus = "pending"
	RequestStatusApproved RequestStatus = "approved"
	RequestStatusRejected RequestStatus = "rejected"
)

func ReconstructRequest(
	id string,
	requesterUserID string,
	requesterEmail string,
	judgerUserID string,
	judgerEmail string,
	status string,
	projectID string,
	roles string,
	period int32,
	reason string,
	requestedAt time.Time,
	expiredAt time.Time,
	judgedAt time.Time) (*Request, error) {

	return newRequest(
		id, UserID(requesterUserID),
		requesterEmail,
		UserID(judgerUserID),
		judgerEmail,
		projectID,
		SplitRole(roles),
		PeriodKey(period),
		reason, RequestStatus(status), requestedAt, expiredAt, judgedAt), nil
}

func NewRequest(
	requesterUserID UserID,
	projectID string,
	roles []string,
	period PeriodKey,
	reason string,
	now time.Time,
	requestExpiredAt time.Time,
) (*Request, error) {

	_, ok := periodMap[period]
	if !ok {
		return nil, errors.WithDetail(errors.Newf("invalid period %v", period), string(custom_error.InvalidArgument))
	}
	id := uuid.New().String()
	judgerUserID := UserID("")
	status := RequestStatusPending
	return newRequest(id, requesterUserID, "", judgerUserID, "", projectID, roles, period, reason, status, now, requestExpiredAt, time.Time{}), nil
}

func newRequest(
	id string,
	requesterUserID UserID,
	requesterEmail string,
	judgerUserID UserID,
	judgerEmail string,
	projectID string,
	iamRoles []string,
	period PeriodKey,
	reason string,
	status RequestStatus,
	requestedAt time.Time,
	expiredAt time.Time,
	judgedAt time.Time) *Request {
	return &Request{
		id:              id,
		requesterUserID: requesterUserID,
		requesterEmail:  requesterEmail,
		judgerUserID:    judgerUserID,
		judgerEmail:     judgerEmail,
		status:          status,
		projectID:       projectID,
		iamRoles:        iamRoles,
		period:          period,
		reason:          reason,
		requestedAt:     requestedAt,
		judgedAt:        judgedAt,
		expiredAt:       expiredAt,
	}
}

func (r *Request) ID() string {
	return r.id
}

func (r *Request) RequesterUserID() UserID {
	return r.requesterUserID
}

func (r *Request) RequesterEmail() string {
	return r.requesterEmail
}

func (r *Request) JudgerUserID() UserID {
	return r.judgerUserID
}

func (r *Request) JudgerEmail() string {
	return r.judgerEmail
}

func (r *Request) Status() RequestStatus {
	return r.status
}

func (r *Request) ProjectID() string {
	return r.projectID
}

func (r *Request) IamRoles() []string {
	return r.iamRoles
}

const separator = ","

func (r *Request) ConcatIamRole() string {
	return strings.Join(r.iamRoles, separator)
}

func SplitRole(concatRole string) []string {
	return strings.Split(concatRole, separator)
}

func (r *Request) Period() PeriodKey {
	return r.period
}

// returns the value
func (r *Request) PeriodViewValue() string {
	return periodMap[r.period]
}

func (r *Request) PeriodDuration() time.Duration {
	return periodTimeMap[r.period]
}

// returns the time when the IAM role binding time based condition is valid.
func (r *Request) CalculateRoleBindingExpiry(now time.Time) time.Time {
	return now.Add(periodTimeMap[r.period])
}

func (r *Request) Reason() string {
	return r.reason
}

func (r *Request) RequestedAt() time.Time {
	return r.requestedAt
}

func (r *Request) ExpiredAt() time.Time {
	return r.expiredAt
}

func (r *Request) IsExpired(now time.Time) bool {
	return r.expiredAt.Before(now)
}

func (r *Request) JudgedAt() time.Time {
	return r.judgedAt
}

func (r *Request) Approve(judger *User, t time.Time) {
	r.judge(judger, RequestStatusApproved, t)
}

func (r *Request) Reject(judger *User, t time.Time) {
	r.judge(judger, RequestStatusRejected, t)
}

func (r *Request) judge(judger *User, status RequestStatus, t time.Time) {
	r.status = status
	r.judgerUserID = judger.ID()
	r.judgerEmail = judger.Email()
	r.judgedAt = t
}

func (r *Request) IsApprove() bool {
	return r.status == RequestStatusApproved
}

func (r *Request) Clone() *Request {
	return newRequest(
		r.id,
		r.requesterUserID,
		r.requesterEmail,
		r.judgerUserID,
		r.judgerEmail,
		r.projectID,
		r.iamRoles,
		r.period,
		r.reason,
		r.status,
		r.requestedAt,
		r.expiredAt,
		r.judgedAt)
}
