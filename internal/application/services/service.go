package services

import (
	"homework_bot/internal/application/interfaces"
	"homework_bot/internal/infrastructure/repositories"
)

type Service struct {
	interfaces.IUserService
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		IUserService: NewUserService(repos.IUserRepository),
	}
}
