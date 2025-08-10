package domain

type UserTRXPlaceholders struct {
	Id          int64  `json:"id" form:"id" gorm:"primarykey;column:id;size:20;"`         //id字段
	Status      int64  `json:"status" form:"status" gorm:"column:status;"`                //   `db:"user_id"`
	Placeholder string `json:"placeholder" form:"placeholder" gorm:"column:placeholder;"` // `db:"times"`
}

// TableName ronUsers表 RonUsers自定义表名 ron_users
func (UserTRXPlaceholders) TableName() string {
	return "user_trx_placeholders"
}
