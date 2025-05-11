package services

import (
	"ushield_bot/internal/application/interfaces"
	"ushield_bot/internal/infrastructure/repositories"
)

type Service struct {
	interfaces.IUserService
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		IUserService: NewUserService(repos.IUserRepository),
	}
}
