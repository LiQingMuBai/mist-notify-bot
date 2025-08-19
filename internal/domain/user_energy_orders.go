// 自动生成模板UserEnergyOrders
package domain

import "time"

// userEnergyOrders表 结构体  UserEnergyOrders
type UserEnergyOrders struct {
	ID          uint      `gorm:"primarykey" json:"ID"`                                               // 主键ID
	OrderNo     string    `json:"orderNo" form:"orderNo" gorm:"column:order_no;size:50;"`             //orderNo字段
	TxId        string    `json:"txId" form:"txId" gorm:"column:tx_id;size:50;"`                      //txId字段
	FromAddress string    `json:"fromAddress" form:"fromAddress" gorm:"column:from_address;size:50;"` //fromAddress字段
	ToAddress   string    `json:"toAddress" form:"toAddress" gorm:"column:to_address;size:50;"`       //toAddress字段
	Token       string    `json:"token" form:"token" gorm:"column:token;size:50;"`                    //token字段
	Amount      float64   `json:"amount" form:"amount" gorm:"column:amount;size:22;"`                 //amount字段
	ChatId      string    `json:"chatId" form:"chatId" gorm:"column:chat_id;size:50;"`                //chatId字段
	Remark      string    `json:"remark" form:"remark" gorm:"column:remark;size:255;"`                //remark字段
	CreatedAt   time.Time // 创建时间
}

// TableName userEnergyOrders表 UserEnergyOrders自定义表名 user_energy_orders
func (UserEnergyOrders) TableName() string {
	return "user_energy_orders"
}
