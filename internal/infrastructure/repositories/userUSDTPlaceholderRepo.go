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
	var Placeholders []domain.UserUsdtPlaceholders
	err := r.db.WithContext(ctx).
		Model(&domain.UserUsdtPlaceholders{}).
		Select("id", "placeholder").
		Where("status = ?", 0).
		Scan(&Placeholders).Error
	return Placeholders, err

}
