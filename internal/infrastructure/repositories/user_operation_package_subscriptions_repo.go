package repositories

import (
	"context"
	"gorm.io/gorm"
	"ushield_bot/internal/domain"
	"ushield_bot/internal/request"
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
func (r *UserPackageSubscriptionsRepository) Query(ctx context.Context, ID string) (domain.UserPackageSubscriptions, error) {
	var subscriptions []domain.UserPackageSubscriptions
	err := r.db.WithContext(ctx).
		Model(&domain.UserPackageSubscriptions{}).
		Select("id", "times", "bundle_name", "bundle_id", "amount", "address").
		Where("id = ?", ID).
		Scan(&subscriptions).Error
	return subscriptions[0], err

}

// Create 创建新套餐
func (r *UserPackageSubscriptionsRepository) Create(ctx context.Context, pkg *domain.UserPackageSubscriptions) error {
	return r.db.WithContext(ctx).Create(pkg).Error
}

// Update 更新套餐
func (r *UserPackageSubscriptionsRepository) Update(ctx context.Context, pkg *domain.UserPackageSubscriptions) error {
	return r.db.WithContext(ctx).Save(pkg).Error
}

// Update 更新套餐
func (r *UserPackageSubscriptionsRepository) UpdateStatus(ctx context.Context, ID int64, _status int64) error {
	return r.db.WithContext(ctx).Model(&domain.UserPackageSubscriptions{}).
		Where("id = ?", ID).
		Update("status", _status).Error
}

// Update 更新套餐
func (r *UserPackageSubscriptionsRepository) UpdateTimes(ctx context.Context, ID int64, _times int64) error {
	return r.db.WithContext(ctx).Model(&domain.UserPackageSubscriptions{}).
		Where("id = ?", ID).
		Update("times", _times).Error
}

// Delete 删除套餐
func (r *UserPackageSubscriptionsRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.UserPackageSubscriptions{}, id).Error
}
func (r *UserPackageSubscriptionsRepository) GetUserPackageSubscriptionsInfoList(ctx context.Context, info request.UserAddressDetectionSearch, _chatID int64) (list []domain.UserPackageSubscriptions, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := r.db.Model(&domain.UserPackageSubscriptions{}).Select("id,status,amount,times,bundle_name,bundle_id,address, DATE_FORMAT(created_at, '%m-%d') as created_date").Where("chat_id = ?", _chatID)
	var UserPackageSubscriptions []domain.UserPackageSubscriptions
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(int(limit)).Offset(int(offset)).Order("id DESC")
	}

	err = db.Find(&UserPackageSubscriptions).Error
	return UserPackageSubscriptions, total, err
}
