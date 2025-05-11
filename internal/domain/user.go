package domain

import (
	"time"
)

type User struct {
	Id          int64     `json:"id" form:"id" gorm:"primarykey;column:id;size:20;"`            //id字段
	UserID      string    `json:"user_id" form:"user_id" gorm:"column:user_id;"`                //   `db:"user_id"`
	Times       int64     `json:"times" form:"times" gorm:"column:times;"`                      // `db:"times"`
	Username    string    `json:"username" form:"username" gorm:"column:username;"`             // `db:"times"`   `db:"username"`
	Amount      string    `json:"amount" form:"amount" gorm:"column:amount;"`                   //  `db:"amount"`
	Address     string    `json:"address" form:"address" gorm:"column:address;"`                //  `db:"amount"``db:"address"`
	Key         string    `json:"private_key" form:"private_key" gorm:"column:private_key;"`    //  db:"private_key"`
	Associates  string    `json:"associates" form:"associates" gorm:"column:associates;"`       //  db:"private_key"` ` db:"associates"`
	TronAmount  string    `json:"tron_amount" form:"tron_amount" gorm:"column:tron_amount;"`    //  db:"private_key"` ` db:"associates"` `db:"tron_amount"`
	TronAddress string    `json:"tron_address" form:"tron_address" gorm:"column:tron_address;"` //  db:"private_key"` ` db:"associates"` `db:"tron_amount"` `db:"tron_address"`
	EthAddress  string    `json:"eth_address" form:"eth_address" gorm:"column:eth_address;"`    //  db:"private_key"` ` db:"associates"` `db:"tron_amount"` `db:"eth_address"`
	EthAmount   string    `json:"eth_amount" form:"eth_amount" gorm:"column:eth_amount;"`       //  db:"private_key"` ` db:"associates"` `db:"tron_amount"` `db:"eth_amount"`
	CreatedAt   time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`         //createdAt字段 `db:"create_at"`
	UpdatedAt   time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`         //updatedAt字段`db:"update_at"`
}

// TableName ronUsers表 RonUsers自定义表名 ron_users
func (User) TableName() string {
	return "tg_users"
}

//associates VARCHAR(255),
//amount VARCHAR(255) ,
//tron_amount VARCHAR(255),
//tron_address VARCHAR(50),
//eth_address VARCHAR(50),
//eth_amount VARCHAR(255),

func NewUser(username, _amount, _Associates, _TronAmount, _TronAddress, _EthAddress, _EthAmount, _Address string) *User {
	return &User{
		//UserID:      _userId,
		Username:    username,
		Amount:      _amount,
		Associates:  _Associates,
		TronAmount:  _TronAmount,
		TronAddress: _TronAddress,
		EthAddress:  _EthAddress,
		EthAmount:   _EthAmount,
		Address:     _Address,
	}
}
