package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	TestamentKey  string `gorm:"type:varchar(20);not null;default:;comment:TOT/TNT"`
	TestamentName string `gorm:"type:varchar(20);not null;default:;comment:新约/旧约"`
	BookName      string `gorm:"type:varchar(20);not null;default:;comment:书卷"`
	ShortName     string `gorm:"type:varchar(20);not null;default:;comment:书卷缩写"`
	TimesCited    int64  `gorm:"type:bigint(20);not null;default:0;comment:引用次数"`
}
