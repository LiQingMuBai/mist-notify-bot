package repositories

import (
	"context"
	"gorm.io/gorm"
	"ushield_bot/internal/domain"
)

type UserPackageSubscriptionsRepository struct {
	db *gorm.DB
}

func NewUserPackageSubscriptionsRepository(db *gorm.DB) *UserPackageSubscriptionsRepository {
	return &UserPackageSubscriptionsRepository{
		db: db,
	}
}
func (r *UserPackageSubscriptionsRepository) ListAll(ctx context.Context) ([]domain.UserPackageSubscriptions, error) {
	var pkgs []domain.UserPackageSubscriptions
	err := r.db.WithContext(ctx).
		Model(&domain.UserPackageSubscriptions{}).
		Select("id", "name", "amount").
		Where("status = ?", 0).
		Scan(&pkgs).Error
	return pkgs, err

}

// Create 创建新套餐
func (r *UserPackageSubscriptionsRepository) Create(ctx context.Context, pkg *domain.UserPackageSubscriptions) error {
	return r.db.WithContext(ctx).Create(pkg).Error
}

// Update 更新套餐
func (r *UserPackageSubscriptionsRepository) Update(ctx context.Context, pkg *domain.UserPackageSubscriptions) error {
	return r.db.WithContext(ctx).Save(pkg).Error
}

// Delete 删除套餐
func (r *UserPackageSubscriptionsRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.UserPackageSubscriptions{}, id).Error
}
