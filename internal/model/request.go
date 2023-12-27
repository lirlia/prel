package model

import (
	"strings"
	"time"

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
		reason, RequestStatus(status), requestedAt, expiredAt, judgedAt), nil
}

func NewRequest(
	requesterUserID UserID,
	projectID string,
	roles []string,
	reason string,
	requestedAt time.Time,
	until time.Time) *Request {
	id := uuid.New().String()
	judgerUserID := UserID("")
	status := RequestStatusPending
	return newRequest(id, requesterUserID, "", judgerUserID, "", projectID, roles, reason, status, requestedAt, until, time.Time{})
}

func newRequest(
	id string,
	requesterUserID UserID,
	requesterEmail string,
	judgerUserID UserID,
	judgerEmail string,
	projectID string,
	iamRoles []string,
	reason string,
	status RequestStatus,
	requestedAt time.Time,
	until time.Time,
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
		reason:          reason,
		requestedAt:     requestedAt,
		judgedAt:        judgedAt,
		expiredAt:       until,
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

func (r *Request) Reason() string {
	return r.reason
}

func (r *Request) RequestedAt() time.Time {
	return r.requestedAt
}

func (r *Request) ExpiredAt() time.Time {
	return r.expiredAt
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
		r.reason,
		r.status,
		r.requestedAt,
		r.expiredAt,
		r.judgedAt)
}
