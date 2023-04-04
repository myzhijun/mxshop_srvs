package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        int32     `gorm:"primarykey"`
	CreateAt  time.Time `gorm:"column:add_time"`
	UpdateAt  time.Time `gorm:"column:update_time"`
	DeleteAt  gorm.DeletedAt
	IsDeleted bool
}

/**
1.密文 2.密文不可反解
	1.对称加密
	2.非对称加密
	3.md5信息摘要算法
	密码如果不可以反解，用于找回密码
*/
type User struct {
	BaseModel
	Mobile   string     `gorm:"index:idx_mobile;unique;type:varchar(11);not null"` //加index的目的是后续利用手机号搜索的时候速度快。
	Password string     `gorm:"type:varchar(100);not null"`
	NickName string     `gorm:"type:varchar(200)"`
	Birthday *time.Time `gorm:"type:datetime"`
	Gender   string     `gorm:"column:gender;default:male;type:varchar(6) comment 'female表示女,male表示男'"`
	Role     int        `gorm:"column:role;defaule:1;type:int comment '1表示普通用户,2表示管理员' "`
}
