package domain

import (
	"time"
)

type UserPackageSubscriptions struct {
	Id        int64     `json:"id" form:"id" gorm:"primarykey;column:id;size:20;"`    //id字段
	Status    int64     `json:"status" form:"status" gorm:"column:status;"`           //   `db:"user_id"`
	UserID    int64     `json:"user_id" form:"user_id" gorm:"column:user_id;"`        //   `db:"user_id"`
	BundleID  int64     `json:"bundle_id" form:"bundle_id" gorm:"column:bundle_id;"`  //   `db:"user_id"`
	CreatedAt time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"` //createdAt字段 `db:"create_at"`
	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"` //updatedAt字段`db:"update_at"`
}

// TableName ronUsers表 RonUsers自定义表名 ron_users
func (UserPackageSubscriptions) TableName() string {
	return "user_package_subscriptions"
}
