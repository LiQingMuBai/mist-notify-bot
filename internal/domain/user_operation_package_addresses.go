package domain

import "time"

type UserOperationPackageAddresses struct {
	Id        int64     `json:"id" form:"id" gorm:"primarykey;column:id;size:20;"`    //id字段
	ChatID    int64     `json:"chat_id" form:"chat_id" gorm:"column:chat_id;"`        //   `db:"user_id"`
	Status    int64     `json:"status" form:"status" gorm:"column:status;"`           //   `db:"user_id"`
	Address   string    `json:"address" form:"address" gorm:"column:address;"`        // `db:"times"`
	Remark    string    `json:"remark" form:"remark" gorm:"column:remark;"`           // `db:"times"`
	CreatedAt time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"` //createdAt字段 `db:"create_at"`
	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"` //updatedAt字段`db:"update_at"`
}

// TableName ronUsers表 RonUsers自定义表名 ron_users
func (UserOperationPackageAddresses) TableName() string {
	return "user_operation_package_addresses"
}
