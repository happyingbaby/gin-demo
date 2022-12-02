package models

import "gorm.io/gorm"

type book struct {
	gorm.Model
	UserId   uint
	BookName string
	BookImg  string
	Desc     string
	Borrow   uint
}
