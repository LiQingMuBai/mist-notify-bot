package services

import (
	"homework_bot/internal/application/interfaces"
	"homework_bot/internal/infrastructure/repositories"
)

type Service struct {
	interfaces.IHomeworkService
	interfaces.IScheduleService
	interfaces.IUserService
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		IHomeworkService: NewHomeworkService(repos.IHomeworkRepository),
		IScheduleService: NewScheduleFefuService(),
		IUserService:     NewUserService(repos.IUserRepository),
	}
}
