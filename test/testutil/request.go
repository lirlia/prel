package testutil

import (
	"prel/internal/model"
	"time"
)

func NewRequest(opts ...RequestOption) *model.Request {
	r := &TestRequest{
		RequesterUserID: "requester-user-id",
		RequesterEmail:  "requester-email",
		JudgerUserID:    "judger-user-id",
		JudgerEmail:     "judger-email",
		Status:          model.RequestStatusPending,
		ProjectID:       "project-id",
		IamRoles:        []string{"iam-role"},
		Reason:          "reason",
		RequestedAt:     time.Now(),
		JudgedAt:        time.Now(),
		ExpiredAt:       time.Now().Add(1 * time.Hour),
	}

	for _, opt := range opts {
		opt(r)
	}

	return model.NewRequest(r.RequesterUserID, r.ProjectID, r.IamRoles, r.Reason, r.RequestedAt, r.ExpiredAt)
}

type TestRequest struct {
	RequesterUserID model.UserID
	RequesterEmail  string
	JudgerUserID    model.UserID
	JudgerEmail     string
	Status          model.RequestStatus
	ProjectID       string
	IamRoles        []string
	Reason          string
	RequestedAt     time.Time
	JudgedAt        time.Time
	ExpiredAt       time.Time
}

type RequestOption func(*TestRequest)

func WithRequesterUserID(requesterUserID model.UserID) RequestOption {
	return func(r *TestRequest) {
		r.RequesterUserID = requesterUserID
	}
}

func WithRequesterEmail(requesterEmail string) RequestOption {
	return func(r *TestRequest) {
		r.RequesterEmail = requesterEmail
	}
}

func WithJudgerUserID(judgerUserID model.UserID) RequestOption {
	return func(r *TestRequest) {
		r.JudgerUserID = judgerUserID
	}
}

func WithJudgerEmail(judgerEmail string) RequestOption {
	return func(r *TestRequest) {
		r.JudgerEmail = judgerEmail
	}
}

func WithStatus(status model.RequestStatus) RequestOption {
	return func(r *TestRequest) {
		r.Status = status
	}
}

func WithProjectID(projectID string) RequestOption {
	return func(r *TestRequest) {
		r.ProjectID = projectID
	}
}

func WithIamRoles(iamRoles []string) RequestOption {
	return func(r *TestRequest) {
		r.IamRoles = iamRoles
	}
}

func WithReason(reason string) RequestOption {
	return func(r *TestRequest) {
		r.Reason = reason
	}
}

func WithRequestedAt(requestedAt time.Time) RequestOption {
	return func(r *TestRequest) {
		r.RequestedAt = requestedAt
	}
}

func WithJudgedAt(judgedAt time.Time) RequestOption {
	return func(r *TestRequest) {
		r.JudgedAt = judgedAt
	}
}

func WithExpiredAt(expiredAt time.Time) RequestOption {
	return func(r *TestRequest) {
		r.ExpiredAt = expiredAt
	}
}
