package repositories

import (
	"context"
	_ "github.com/go-sql-driver/mysql"

	"gorm.io/gorm"
	"ushield_bot/internal/domain"
)

type UserUSDTDepositsRepo struct {
	db *gorm.DB
}

func NewUserUSDTDepositsRepository(db *gorm.DB) *UserUSDTDepositsRepo {
	return &UserUSDTDepositsRepo{
		db: db,
	}
}

func (r *UserUSDTDepositsRepo) Create(ctx context.Context, USDTDeposit *domain.UserUSDTDeposits) error {
	return r.db.WithContext(ctx).Create(USDTDeposit).Error
}
