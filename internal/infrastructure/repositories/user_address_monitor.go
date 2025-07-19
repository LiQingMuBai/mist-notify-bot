package repositories

import (
	"context"
	_ "github.com/go-sql-driver/mysql"

	"gorm.io/gorm"
	"ushield_bot/internal/domain"
)

type UserAddressMonitorRepo struct {
	db *gorm.DB
}

func NewUserAddressMonitorRepo(db *gorm.DB) *UserAddressMonitorRepo {
	return &UserAddressMonitorRepo{
		db: db,
	}
}

func (r *UserAddressMonitorRepo) Create(ctx context.Context, userAddress *domain.UserAddressMonitor) error {
	return r.db.WithContext(ctx).Create(userAddress).Error
}

func (r *UserAddressMonitorRepo) Remove(ctx context.Context, _chatID int64, _address string) error {
	//return r.db.WithContext(ctx).del(userAddress).Error

	return r.db.WithContext(ctx).Delete(&domain.UserAddressMonitor{}, "chat_id = ? AND address = ?", _chatID, _address).Error
}

func (r *UserAddressMonitorRepo) Query(ctx context.Context, _chatID int64) ([]domain.UserAddressMonitor, error) {
	var subscriptions []domain.UserAddressMonitor
	err := r.db.WithContext(ctx).
		Model(&domain.UserAddressMonitor{}).
		Select("id", "address", "network").
		Where("chat_id = ?", _chatID).
		Scan(&subscriptions).Error
	return subscriptions, err

}
