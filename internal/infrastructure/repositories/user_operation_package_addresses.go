package repositories

import (
	"context"
	_ "github.com/go-sql-driver/mysql"

	"gorm.io/gorm"
	"ushield_bot/internal/domain"
)

type UserOperationPackageAddressesRepo struct {
	db *gorm.DB
}

func NewUserOperationPackageAddressesRepo(db *gorm.DB) *UserOperationPackageAddressesRepo {
	return &UserOperationPackageAddressesRepo{
		db: db,
	}
}

func (r *UserOperationPackageAddressesRepo) Create(ctx context.Context, userAddress *domain.UserOperationPackageAddresses) error {
	return r.db.WithContext(ctx).Create(userAddress).Error
}
func (r *UserOperationPackageAddressesRepo) Update(ctx context.Context, chatID int64, address string) error {
	r.db.WithContext(ctx).Model(&domain.UserOperationPackageAddresses{}).
		Where("chat_id = ?", chatID).Where("address = ?", address).
		Update("status", 1)

	r.db.WithContext(ctx).Model(&domain.UserOperationPackageAddresses{}).
		Where("chat_id = ?", chatID).Where("address != ?", address).
		Update("status", 0)

	return nil

}
func (r *UserOperationPackageAddressesRepo) Remove(ctx context.Context, _chatID int64, _address string) error {
	return r.db.WithContext(ctx).Delete(&domain.UserOperationPackageAddresses{}, "chat_id = ? AND address = ?", _chatID, _address).Error
}

func (r *UserOperationPackageAddressesRepo) Query(ctx context.Context, _chatID int64) ([]domain.UserOperationPackageAddresses, error) {
	var subscriptions []domain.UserOperationPackageAddresses
	err := r.db.WithContext(ctx).
		Model(&domain.UserOperationPackageAddresses{}).
		Select("id", "address", "status", "remark").
		Where("chat_id = ?", _chatID).
		Scan(&subscriptions).Error
	return subscriptions, err

}
