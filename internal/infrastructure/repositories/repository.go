package repositories

import (
	"github.com/jmoiron/sqlx"
	"homework_bot/internal/application/interfaces"
)

type Repository struct {
	interfaces.IHomeworkRepository
	interfaces.IUserRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		IHomeworkRepository: NewHomeworkRepository(db),
		IUserRepository:     NewUserRepository(db),
	}
}
