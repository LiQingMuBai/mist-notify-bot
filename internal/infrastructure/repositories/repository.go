package repositories

import (
	"github.com/jmoiron/sqlx"
	"ushield_bot/internal/application/interfaces"
)

type Repository struct {
	interfaces.IUserRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		IUserRepository: NewUserRepository(db),
	}
}
