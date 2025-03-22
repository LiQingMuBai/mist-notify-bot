package services

import (
	"github.com/google/uuid"
	"homework_bot/internal/application/interfaces"
	"homework_bot/internal/domain"
)

type UserService struct {
	repos interfaces.IUserRepository
}

func NewUserService(repos interfaces.IUserRepository) *UserService {
	return &UserService{
		repos: repos,
	}
}

func (s *UserService) Create(user domain.User) error {
	user.Id = uuid.New()
	return s.repos.Create(user)
}

func (s *UserService) Update(user domain.User) error {
	return s.repos.Update(user)
}

func (s *UserService) GetByUsername(username string) (domain.User, error) {
	return s.repos.GetByUsername(username)
}
