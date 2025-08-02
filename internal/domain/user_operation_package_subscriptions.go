package domain

import (
	"time"
)

type UserPackageSubscriptions struct {
	Id          int64     `json:"id" form:"id" gorm:"primarykey;column:id;size:20;"`         //id字段
	Status      int64     `json:"status" form:"status" gorm:"column:status;"`                //   `db:"user_id"`
	ChatID      int64     `json:"chat_id" form:"chat_id" gorm:"column:chat_id;"`             //   `db:"user_id"`
	BundleID    int64     `json:"bundle_id" form:"bundle_id" gorm:"column:bundle_id;"`       //   `db:"user_id"`
	Times       int64     `json:"times" form:"times" gorm:"column:times;"`                   //   `db:"user_id"`
	BundleName  string    `json:"bundle_name" form:"bundle_name" gorm:"column:bundle_name;"` // `db:"times"`
	Address     string    `json:"address" form:"address" gorm:"column:address;"`             // `db:"times"`
	CreatedAt   time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`      //createdAt字段 `db:"create_at"`
	UpdatedAt   time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`      //updatedAt字段`db:"update_at"`
	CreatedDate string    `json:"created_date"`
	Amount      string    `json:"amount" form:"amount" gorm:"column:amount;"` // `db:"times"`

}

// TableName ronUsers表 RonUsers自定义表名 ron_users
func (UserPackageSubscriptions) TableName() string {
	return "user_operation_package_subscriptions"
}
