package repositories

import (
	"context"
	_ "github.com/go-sql-driver/mysql"

	"gorm.io/gorm"
	"ushield_bot/internal/domain"
)

type UserEnergyOrdersRepo struct {
	db *gorm.DB
}

func NewUserEnergyOrdersRepo(db *gorm.DB) *UserEnergyOrdersRepo {
	return &UserEnergyOrdersRepo{
		db: db,
	}
}

func (r *UserEnergyOrdersRepo) Create(ctx context.Context, userAddress *domain.UserEnergyOrders) error {
	return r.db.WithContext(ctx).Create(userAddress).Error
}
