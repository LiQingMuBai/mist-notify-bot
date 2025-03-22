package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id          uuid.UUID `db:"id"`
	Username    string    `db:"username"`
	Amount      string    `db:"amount"`
	Associates  string    `db:"associates"`
	TronAmount  string    `db:"tron_amount"`
	TronAddress string    `db:"tron_address"`
	EthAddress  string    `db:"eth_address"`
	EthAmount   string    `db:"eth_amount"`

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

func NewUser(username, _amount, _Associates, _TronAmount, _TronAddress, _EthAddress, _EthAmount string) *User {
	return &User{
		Username:    username,
		Amount:      _amount,
		Associates:  _Associates,
		TronAmount:  _TronAmount,
		TronAddress: _TronAddress,
		EthAddress:  _EthAddress,
		EthAmount:   _EthAmount,
	}
}
