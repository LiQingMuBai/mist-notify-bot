package services

import (
	"github.com/google/uuid"
	"ushield_bot/internal/application/interfaces"
	"ushield_bot/internal/domain"
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
	user.UserID = uuid.New().String()
	return s.repos.Create(user)
}

func (s *UserService) Update(user domain.User) error {
	return s.repos.Update(user)
}
func (s *UserService) UpdateAddress(user domain.User) error {
	return s.repos.UpdateAddress(user)
}

func (s *UserService) GetByUsername(username string) (domain.User, error) {
	return s.repos.GetByUsername(username)
}

func (s *UserService) GetByUserID(_userID string) (domain.User, error) {
	return s.repos.GetByUserID(_userID)
}

func (s *UserService) FetchNewestAddress() ([]domain.User, error) {
	return s.repos.FetchNewestAddress()
}
func (s *UserService) NotifyTronAddress() ([]domain.User, error) {
	return s.repos.NotifyTronAddress()
}
func (s *UserService) NotifyEthereumAddress() ([]domain.User, error) {
	return s.repos.NotifyEthereumAddress()
}

func (s *UserService) UpdateTimes(_times uint64, _username string) error {
	return s.repos.UpdateTimes(_times, _username)
}
func (s *UserService) BindTronAddress(_address string, _username string) error {
	return s.repos.BindTronAddress(_address, _username)
}

func (s *UserService) BindEthereumAddress(_address string, _username string) error {
	return s.repos.BindEthereumAddress(_address, _username)
}
func (s *UserService) DisableTronAddress(_address string) error {

	return s.repos.DisableTronAddress(_address)
}
func (s *UserService) BindChat(_associates string, _username string) error {
	return s.repos.BindChat(_associates, _username)
}
