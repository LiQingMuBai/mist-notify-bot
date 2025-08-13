package repositories

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"ushield_bot/internal/request"

	"gorm.io/gorm"
	"ushield_bot/internal/domain"
)

type UserAddressDetectionRepo struct {
	db *gorm.DB
}

func NewUserAddressDetectionRepository(db *gorm.DB) *UserAddressDetectionRepo {
	return &UserAddressDetectionRepo{
		db: db,
	}
}

func (r *UserAddressDetectionRepo) Create(ctx context.Context, trxDeposit *domain.UserAddressDetection) error {
	return r.db.WithContext(ctx).Create(trxDeposit).Error
}
func (r *UserAddressDetectionRepo) ListAll(ctx context.Context, _chatID int64, _status int64) ([]domain.UserAddressDetection, error) {
	var subscriptions []domain.UserAddressDetection

	err := r.db.Select("id,amount,order_no, DATE_FORMAT(created_at, '%m-%d') as created_date").
		Where("user_id = ?", _chatID).
		Where("status = ?", _status).
		Find(&subscriptions).Error
	return subscriptions, err

}
func (r *UserAddressDetectionRepo) GetUserAddressDetectionInfoList(ctx context.Context, info request.UserAddressDetectionSearch, _chatID int64) (list []domain.UserAddressDetection, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := r.db.Model(&domain.UserAddressDetection{}).Select("id,amount,address, DATE_FORMAT(created_at, '%m-%d') as created_date").Where("chat_id = ?", _chatID)
	var UserAddressDetections []domain.UserAddressDetection
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(int(limit)).Offset(int(offset)).Order("id DESC")
	}

	err = db.Find(&UserAddressDetections).Error
	return UserAddressDetections, total, err
}
