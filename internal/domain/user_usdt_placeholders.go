package domain

import (
	"time"
)

type UserUsdtPlaceholders struct {
	Id          int64     `json:"id" form:"id" gorm:"primarykey;column:id;size:20;"`         //id字段
	Status      int64     `json:"status" form:"status" gorm:"column:status;"`                //   `db:"user_id"`
	Placeholder string    `json:"placeholder" form:"placeholder" gorm:"column:placeholder;"` // `db:"times"`
	CreatedAt   time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`      //createdAt字段 `db:"create_at"`
	UpdatedAt   time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`      //updatedAt字段`db:"update_at"`
}

// TableName ronUsers表 RonUsers自定义表名 ron_users
func (UserUsdtPlaceholders) TableName() string {
	return "user_usdt_placeholders"
}
