package repositories

import (
	"context"
	_ "github.com/go-sql-driver/mysql"

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
