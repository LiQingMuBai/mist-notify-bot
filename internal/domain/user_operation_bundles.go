package domain

type UserOperationBundles struct {
	Id     int64  `json:"id" form:"id" gorm:"primarykey;column:id;size:20;"` //id字段
	Status int64  `json:"status" form:"status" gorm:"column:status;"`        //   `db:"user_id"`
	Name   string `json:"name" form:"name" gorm:"column:name;"`              // `db:"times"`
	Token  string `json:"token" form:"token" gorm:"column:token;"`           // `db:"times"`
	Amount string `json:"amount" form:"amount" gorm:"column:amount;"`        // `db:"times"`
	//CreatedAt time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"` //createdAt字段 `db:"create_at"`
	//UpdatedAt time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"` //updatedAt字段`db:"update_at"`
}

// TableName ronUsers表 RonUsers自定义表名 ron_users
func (UserOperationBundles) TableName() string {
	return "user_operation_bundles"
}
