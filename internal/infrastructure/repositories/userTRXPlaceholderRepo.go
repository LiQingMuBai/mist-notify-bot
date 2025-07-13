package repositories

import (
	"context"
	"gorm.io/gorm"
	"ushield_bot/internal/domain"
)

type UserTRXSubscriptionsRepository struct {
	db *gorm.DB
}

func NewUserTRXSubscriptionsRepository(db *gorm.DB) *UserTRXSubscriptionsRepository {
	return &UserTRXSubscriptionsRepository{
		db: db,
	}
}
func (r *UserTRXSubscriptionsRepository) ListAll(ctx context.Context) ([]domain.UserTRXSubscriptions, error) {
	var subscriptions []domain.UserTRXSubscriptions
	err := r.db.WithContext(ctx).
		Model(&domain.UserTRXSubscriptions{}).
		Select("id", "name", "amount").
		Where("status = ?", 0).
		Scan(&subscriptions).Error
	return subscriptions, err

}
