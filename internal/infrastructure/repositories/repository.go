package repositories

import (
	"gorm.io/gorm"
	"ushield_bot/internal/application/interfaces"
)

type Repository struct {
	interfaces.IUserRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		IUserRepository: NewUserRepository(db),
	}
}
