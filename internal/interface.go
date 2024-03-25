package internal

import (
	"context"
	"net/http"
	"prel/internal/gateway/google_cloud"
	"prel/internal/model"
	"time"

	"google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/iam/v1"
)

type GoogleCloudClient interface {
	GetProjects() ([]*cloudresourcemanager.Project, error)
	GetIamRoles(now time.Time, projectID string, member google_cloud.BindingMember) ([]*iam.Role, error)
	SetIamPolicy(ctx context.Context, projectID string, roles []string, member google_cloud.BindingMember, until time.Time) error
}

type NotificationClient interface {
	CanSend() bool
	SendRequestMessage(ctx context.Context, message, requesterEmail, requestUrl, projectID, period, reason string, roles []string, requestExpiredAt time.Time) (*http.Response, error)
	SendJudgeMessage(ctx context.Context, judge model.RequestStatus, message, requesterEmail, judgerEmail, requestUrl, projectID, reason string, roles []string, until time.Time) (*http.Response, error)
}
