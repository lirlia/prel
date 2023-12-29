package usecase

import (
	"net/http"
	"prel/config"
	"prel/internal"
	"prel/internal/gateway/repository"
	"prel/internal/model"
	"prel/internal/service"
)

type Usecase struct {
	config                   *config.Config
	c                        internal.GoogleCloudClient
	n                        internal.NotificationClient
	h                        *http.Client
	invitationRepo           model.InvitationRepository
	requestRepo              model.RequestRepository
	userRepo                 model.UserRepository
	userAndInvitationRepo    model.UserAndInvitationRepository
	iamRoleFilteringRuleRepo model.IamRoleFilteringRuleRepository
	service                  *service.Service
}

func NewUsecase(
	config *config.Config,
	c internal.GoogleCloudClient,
	h *http.Client,
	n internal.NotificationClient) *Usecase {
	return &Usecase{
		config:                   config,
		c:                        c,
		h:                        h,
		n:                        n,
		invitationRepo:           repository.NewInvitationRepository(),
		requestRepo:              repository.NewRequestRepository(),
		userRepo:                 repository.NewUserRepository(),
		userAndInvitationRepo:    repository.NewUserAndInvitationRepository(),
		iamRoleFilteringRuleRepo: repository.NewIamRoleFilteringRuleRepository(),
		service:                  service.NewService(),
	}
}
