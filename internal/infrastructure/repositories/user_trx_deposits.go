package repositories

import (
	"context"
	_ "github.com/go-sql-driver/mysql"

	"gorm.io/gorm"
	"ushield_bot/internal/domain"
)

type UserTRXDepositsRepo struct {
	db *gorm.DB
}

func NewUserTRXDepositsRepository(db *gorm.DB) *UserTRXDepositsRepo {
	return &UserTRXDepositsRepo{
		db: db,
	}
}

func (r *UserTRXDepositsRepo) Create(ctx context.Context, trxDeposit *domain.UserTRXDeposits) error {
	return r.db.WithContext(ctx).Create(trxDeposit).Error
}
