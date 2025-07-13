package repositories

import (
	"context"
	"gorm.io/gorm"
	"ushield_bot/internal/domain"
)

type UserUsdtSubscriptionsRepository struct {
	db *gorm.DB
}

func NewUserUsdtSubscriptionsRepository(db *gorm.DB) *UserUsdtSubscriptionsRepository {
	return &UserUsdtSubscriptionsRepository{
		db: db,
	}
}
func (r *UserUsdtSubscriptionsRepository) ListAll(ctx context.Context) ([]domain.UserUsdtSubscriptions, error) {
	var subscriptions []domain.UserUsdtSubscriptions
	err := r.db.WithContext(ctx).
		Model(&domain.UserUsdtSubscriptions{}).
		Select("id", "name", "amount").
		Where("status = ?", 0).
		Scan(&subscriptions).Error
	return subscriptions, err

}
