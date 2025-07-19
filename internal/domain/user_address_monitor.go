package domain

import "time"

type UserAddressMonitor struct {
	Id        int64     `json:"id" form:"id" gorm:"primarykey;column:id;size:20;"`    //id字段
	ChatID    int64     `json:"chat_id" form:"chat_id" gorm:"column:chat_id;"`        //   `db:"user_id"`
	Status    int64     `json:"status" form:"status" gorm:"column:status;"`           //   `db:"user_id"`
	Network   string    `json:"network" form:"network" gorm:"column:network;"`        // `db:"times"`
	Address   string    `json:"address" form:"address" gorm:"column:address;"`        // `db:"times"`
	CreatedAt time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"` //createdAt字段 `db:"create_at"`
	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"` //updatedAt字段`db:"update_at"`
}

// TableName ronUsers表 RonUsers自定义表名 ron_users
func (UserAddressMonitor) TableName() string {
	return "user_address_monitor"
}
