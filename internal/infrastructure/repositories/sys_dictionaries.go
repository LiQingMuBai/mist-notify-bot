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
