package repositories

import (
	"context"
	"gorm.io/gorm"
	"ushield_bot/internal/domain"
)

type UserUsdtPlaceholdersRepository struct {
	db *gorm.DB
}

func NewUserUsdtPlaceholdersRepository(db *gorm.DB) *UserUsdtPlaceholdersRepository {
	return &UserUsdtPlaceholdersRepository{
		db: db,
	}
}
func (r *UserUsdtPlaceholdersRepository) ListAll(ctx context.Context) ([]domain.UserUsdtPlaceholders, error) {
	var placeholders []domain.UserUsdtPlaceholders
	err := r.db.WithContext(ctx).
		Model(&domain.UserUsdtPlaceholders{}).
		Select("id", "placeholder").
		Where("status = ?", 0).
		Scan(&placeholders).Error
	return placeholders, err

}

func (r *UserUsdtPlaceholdersRepository) UpdateByPlaceholder(ctx context.Context, _placeholder string, _status int64) error {
	return r.db.WithContext(ctx).Model(&domain.UserUsdtPlaceholders{}).
		Where("placeholder = ?", _placeholder).
		Update("status", _status).Error
}

func (r *UserUsdtPlaceholdersRepository) Update(ctx context.Context, ID int64, _status int64) error {
	return r.db.WithContext(ctx).Model(&domain.UserUsdtPlaceholders{}).
		Where("id = ?", ID).
		Update("status", _status).Error
}
func (r *UserUsdtPlaceholdersRepository) Find(ctx context.Context) (domain.UserUsdtPlaceholders, error) {
	var placeholders []domain.UserUsdtPlaceholders
	err := r.db.WithContext(ctx).
		Model(&domain.UserUsdtPlaceholders{}).
		Select("id", "placeholder").
		Where("status = ?", 0).
		Scan(&placeholders).Error
	return placeholders[0], err

}
