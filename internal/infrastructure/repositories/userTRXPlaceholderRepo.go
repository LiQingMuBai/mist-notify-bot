package repositories

import (
	"context"
	"gorm.io/gorm"
	"ushield_bot/internal/domain"
)

type UserTRXPlaceholdersRepository struct {
	db *gorm.DB
}

func NewUserTRXPlaceholdersRepository(db *gorm.DB) *UserTRXPlaceholdersRepository {
	return &UserTRXPlaceholdersRepository{
		db: db,
	}
}
func (r *UserTRXPlaceholdersRepository) ListAll(ctx context.Context) ([]domain.UserTRXPlaceholders, error) {
	var placeholders []domain.UserTRXPlaceholders
	err := r.db.WithContext(ctx).
		Model(&domain.UserTRXPlaceholders{}).
		Select("id", "placeholder").
		Where("status = ?", 0).
		Scan(&placeholders).Error
	return placeholders, err

}

func (r *UserTRXPlaceholdersRepository) Update(ctx context.Context, ID int64, _status int64) error {
	return r.db.WithContext(ctx).Model(&domain.UserTRXPlaceholders{}).
		Where("id = ?", ID).
		Update("status", _status).Error
}

func (r *UserTRXPlaceholdersRepository) UpdateByPlaceholder(ctx context.Context, _placeholder string, _status int64) error {
	return r.db.WithContext(ctx).Model(&domain.UserTRXPlaceholders{}).
		Where("placeholder = ?", _placeholder).
		Update("status", _status).Error
}

func (r *UserTRXPlaceholdersRepository) Find(ctx context.Context) (domain.UserTRXPlaceholders, error) {
	var placeholders []domain.UserTRXPlaceholders
	err := r.db.WithContext(ctx).
		Model(&domain.UserTRXPlaceholders{}).
		Select("id", "placeholder").
		Where("status = ?", 0).
		Scan(&placeholders).Error
	return placeholders[0], err

}
