package domain

type SysUser struct {
	Username       string `json:"userName" gorm:"column:username;comment:用户登录名"`                // 用户登录名
	Address        string `json:"address"  gorm:"comment:用户能量兑换地址"`                             // 用户登录密码
	DepositAddress string `json:"depositAddress"  gorm:"column:deposit_address;comment:用户存款地址"` // 用户登录密码
	Enable         int    `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`              //用户是否被冻结 1正常 2冻结
}

func (SysUser) TableName() string {
	return "sys_users"
}
