package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null;comment:用户名"`
	Telephone string `gorm:"type:varchar(11);not null;comment:手机号"`
	Password  string `gorm:"type:varchar(255);not null;comment:密码"`
}
