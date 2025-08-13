package repositories

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"ushield_bot/internal/request"

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

func (r *UserAddressMonitorEventRepo) Find(ctx context.Context, _ID string) (domain.UserAddressMonitorEvent, error) {
	var event domain.UserAddressMonitorEvent
	err := r.db.WithContext(ctx).
		Find(&event, "id = ?", _ID).Error
	return event, err

}

func (r *UserAddressMonitorEventRepo) Create(ctx context.Context, userAddress *domain.UserAddressMonitorEvent) error {
	return r.db.WithContext(ctx).Create(userAddress).Error
}
func (r *UserAddressMonitorEventRepo) Remove(ctx context.Context, _chatID int64, _address string) error {
	//return r.db.WithContext(ctx).del(userAddress).Error

	return r.db.WithContext(ctx).Delete(&domain.UserAddressMonitor{}, "chat_id = ? AND address = ?", _chatID, _address).Error
}
func (r *UserAddressMonitorEventRepo) Close(ctx context.Context, _ID string) error {
	return r.db.WithContext(ctx).Delete(&domain.UserAddressMonitorEvent{}, "id = ? ", _ID).Error
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
func (r *UserAddressMonitorEventRepo) RemoveAll(ctx context.Context, _chatID int64) error {
	//return r.db.WithContext(ctx).del(userAddress).Error

	return r.db.WithContext(ctx).Delete(&domain.UserAddressMonitorEvent{}, "chat_id = ?", _chatID).Error
}

func (r *UserAddressMonitorEventRepo) GetAddressMonitorEventInfoList(ctx context.Context, info request.UserAddressDetectionSearch, _chatID int64) (list []domain.UserAddressMonitorEvent, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := r.db.Model(&domain.UserAddressMonitorEvent{}).Select("id,amount,address, DATE_FORMAT(created_at, '%m-%d') as created_date").Where("chat_id = ?", _chatID)
	var UserAddressMonitorEvent []domain.UserAddressMonitorEvent
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(int(limit)).Offset(int(offset)).Order("id DESC")
	}

	err = db.Find(&UserAddressMonitorEvent).Error
	return UserAddressMonitorEvent, total, err
}
