package repositories

import (
	"context"
	_ "github.com/go-sql-driver/mysql"

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
