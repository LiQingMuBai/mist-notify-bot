package repositories

import "gorm.io/gorm"

type SysDictionariesRepo struct {
	db *gorm.DB
}

func NewSysDictionariesRepo(db *gorm.DB) *SysDictionariesRepo {
	return &SysDictionariesRepo{
		db: db,
	}
}

func (r *SysDictionariesRepo) GetDictionary(_key string) (string, error) {
	var dict string
	err := r.db.Raw("SELECT description FROM sys_dictionaries where name ='" + _key + "'").Scan(&dict).Error
	return dict, err
}

func (r *SysDictionariesRepo) GetReceiveAddress(_agent string) (string, error) {
	var dict string
	err := r.db.Raw("SELECT address FROM sys_users where username ='" + _agent + "'").Scan(&dict).Error
	return dict, err
}

func (r *SysDictionariesRepo) GetDepositAddress(_agent string) (string, error) {
	var dict string
	err := r.db.Raw("SELECT deposit_address FROM sys_users where username ='" + _agent + "'").Scan(&dict).Error
	return dict, err
}
