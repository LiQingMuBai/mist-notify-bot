package domain

import (
	"time"
)

type UserUSDTDeposits struct {
	Id          int64     `json:"id" form:"id" gorm:"primarykey;column:id;size:20;"`         //id字段
	UserID      int64     `json:"user_id" form:"user_id" gorm:"column:user_id;"`             //   `db:"user_id"`
	Status      int64     `json:"status" form:"status" gorm:"column:status;"`                //   `db:"user_id"`
	Placeholder string    `json:"placeholder" form:"placeholder" gorm:"column:placeholder;"` // `db:"times"`
	OrderNO     string    `json:"order_no" form:"order_no" gorm:"column:order_no;"`          // `db:"times"`
	Address     string    `json:"address" form:"address" gorm:"column:address;"`             // `db:"times"`
	TxHash      string    `json:"tx_hash" form:"tx_hash" gorm:"column:tx_hash;"`             // `db:"times"`
	Block       string    `json:"block" form:"block" gorm:"column:block;"`                   // `db:"times"`
	Amount      string    `json:"amount" form:"amount" gorm:"column:amount;"`                //  `db:"amount"`
	CreatedAt   time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`      //createdAt字段 `db:"create_at"`
	UpdatedAt   time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`      //updatedAt字段`db:"update_at"`
}

// TableName ronUsers表 RonUsers自定义表名 ron_users
func (UserUSDTDeposits) TableName() string {
	return "user_usdt_deposits"
}
