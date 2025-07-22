package repositories

import (
	"context"
	"gorm.io/gorm"
	"ushield_bot/internal/domain"
)

type SysUsersRepository struct {
	db *gorm.DB
}

func NewSysUsersRepository(db *gorm.DB) *SysUsersRepository {
	return &SysUsersRepository{
		db: db,
	}
}

func (r *SysUsersRepository) Find(ctx context.Context, _username string) (address, depositAddress string, err error) {
	var sysUser domain.SysUser
	result := r.db.WithContext(ctx).
		Model(&domain.UserTRXPlaceholders{}).
		Select("address, deposit_address").
		Where("username = ?", _username).
		First(&sysUser)
	if result.Error != nil {
		return "", "", result.Error
	}

	return sysUser.Address, sysUser.DepositAddress, nil
}
