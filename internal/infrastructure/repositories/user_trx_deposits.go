package repositories

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"ushield_bot/internal/request"

	"gorm.io/gorm"
	"ushield_bot/internal/domain"
)

type UserTRXDepositsRepo struct {
	db *gorm.DB
}

func NewUserTRXDepositsRepository(db *gorm.DB) *UserTRXDepositsRepo {
	return &UserTRXDepositsRepo{
		db: db,
	}
}

func (r *UserTRXDepositsRepo) Create(ctx context.Context, trxDeposit *domain.UserTRXDeposits) error {
	return r.db.WithContext(ctx).Create(trxDeposit).Error
}
func (r *UserTRXDepositsRepo) ListAll(ctx context.Context, _chatID int64, _status int64) ([]domain.UserTRXDeposits, error) {
	var subscriptions []domain.UserTRXDeposits

	err := r.db.Select("id,amount,order_no, DATE_FORMAT(created_at, '%m-%d') as created_date").
		Where("user_id = ?", _chatID).
		Where("status = ?", _status).
		Find(&subscriptions).Error
	return subscriptions, err

}
func (r *UserTRXDepositsRepo) GetUserTrxDepositsInfoList(ctx context.Context, info request.UserTrxDepositsSearch, _chatID int64) (list []domain.UserTRXDeposits, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := r.db.Model(&domain.UserTRXDeposits{}).Select("id,amount,order_no, DATE_FORMAT(created_at, '%m-%d') as created_date").Where("user_id = ?", _chatID).Where("status = ?", 1)
	var userTrxDepositss []domain.UserTRXDeposits
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(int(limit)).Offset(int(offset))
	}

	err = db.Find(&userTrxDepositss).Error
	return userTrxDepositss, total, err
}
func (r *UserTRXDepositsRepo) Find(ctx context.Context, orderNo string) (domain.UserTRXDeposits, error) {
	var depositRecords []domain.UserTRXDeposits
	err := r.db.WithContext(ctx).
		Model(&domain.UserTRXDeposits{}).
		Select("id", "placeholder").
		Where("order_no = ?", orderNo).
		Scan(&depositRecords).Error
	return depositRecords[0], err

}
