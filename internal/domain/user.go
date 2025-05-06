package domain

import (
	"time"
)

type User struct {
	Id          int64  `db:"id"`
	UserID      string `db:"user_id"`
	Times       int64  `db:"times"`
	Username    string `db:"username"`
	Amount      string `db:"amount"`
	Address     string `db:"address"`
	Key         string `db:"private_key"`
	Associates  string `db:"associates"`
	TronAmount  string `db:"tron_amount"`
	TronAddress string `db:"tron_address"`
	EthAddress  string `db:"eth_address"`
	EthAmount   string `db:"eth_amount"`

	CreatedAt time.Time `db:"create_at"`
	//Deadline  time.Time `db:"deadline"`
	UpdatedAt time.Time `db:"update_at"`
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
