package repositories

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"ushield_bot/internal/request"

	"gorm.io/gorm"
	"ushield_bot/internal/domain"
)

type UserUSDTDepositsRepo struct {
	db *gorm.DB
}

func NewUserUSDTDepositsRepository(db *gorm.DB) *UserUSDTDepositsRepo {
	return &UserUSDTDepositsRepo{
		db: db,
	}
}

func (r *UserUSDTDepositsRepo) Create(ctx context.Context, USDTDeposit *domain.UserUSDTDeposits) error {
	return r.db.WithContext(ctx).Create(USDTDeposit).Error
}

func (r *UserUSDTDepositsRepo) ListAll(ctx context.Context, _chatID int64, _status int64) ([]domain.UserUSDTDeposits, error) {
	var subscriptions []domain.UserUSDTDeposits
	err := r.db.Select("id,amount,order_no, DATE_FORMAT(created_at, '%m-%d') as created_date").
		Where("user_id = ?", _chatID).
		Where("status = ?", _status).
		Find(&subscriptions).Error
	return subscriptions, err

}
func (r *UserUSDTDepositsRepo) GetUserUsdtDepositsInfoList(ctx context.Context, info request.UserUsdtDepositsSearch, _chatID int64) (list []domain.UserUSDTDeposits, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := r.db.Model(&domain.UserUSDTDeposits{}).Select("id,amount,order_no, DATE_FORMAT(created_at, '%m-%d') as created_date").Where("user_id = ?", _chatID).Where("status = ?", 1)
	var userUsdtDepositss []domain.UserUSDTDeposits
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(int(limit)).Offset(int(offset))
	}

	err = db.Find(&userUsdtDepositss).Error
	return userUsdtDepositss, total, err
}
