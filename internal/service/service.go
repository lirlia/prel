package service

import (
	"prel/internal/gateway/repository"
	"prel/internal/model"
)

type Service struct {
	userRepo    model.UserRepository
	requestRepo model.RequestRepository
}

func NewService() *Service {
	return &Service{
		userRepo:    repository.NewUserRepository(),
		requestRepo: repository.NewRequestRepository(),
	}
}
