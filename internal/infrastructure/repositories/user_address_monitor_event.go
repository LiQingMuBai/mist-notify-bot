package repositories

import (
	"context"
	_ "github.com/go-sql-driver/mysql"

	"gorm.io/gorm"
	"ushield_bot/internal/domain"
)

type UserAddressMonitorEventRepo struct {
	db *gorm.DB
}

func NewUserAddressMonitorEventRepo(db *gorm.DB) *UserAddressMonitorEventRepo {
	return &UserAddressMonitorEventRepo{
		db: db,
	}
}

func (r *UserAddressMonitorEventRepo) Create(ctx context.Context, userAddress *domain.UserAddressMonitorEvent) error {
	return r.db.WithContext(ctx).Create(userAddress).Error
}

func (r *UserAddressMonitorEventRepo) Remove(ctx context.Context, _chatID int64, _address string) error {
	//return r.db.WithContext(ctx).del(userAddress).Error

	return r.db.WithContext(ctx).Delete(&domain.UserAddressMonitor{}, "chat_id = ? AND address = ?", _chatID, _address).Error
}

func (r *UserAddressMonitorEventRepo) Query(ctx context.Context, _chatID int64) ([]domain.UserAddressMonitorEvent, error) {
	var subscriptions []domain.UserAddressMonitorEvent
	err := r.db.WithContext(ctx).
		Model(&domain.UserAddressMonitorEvent{}).
		Select("id", "days", "address", "network").
		Where("chat_id = ?", _chatID).
		Scan(&subscriptions).Error
	return subscriptions, err

}
