package models

import (
	"gorm.io/gorm"
)

type Evangelize struct {
	gorm.Model
	Ymd            int64  `gorm:"type:bigint(20);not null;default:0;comment:年月日"`
	Cleric         string `gorm:"type:varchar(255);not null;default:;comment:传教士"`
	Title          string `gorm:"type:varchar(500);not null;default:;comment:标题"`
	ChapterStr     string `gorm:"type:varchar(500);not null;default:;comment:引用经文"`
	Content        string `gorm:"type:text;not null;default:'';comment:内容"`
	CitedChapterId []uint `gorm:"type:json;comment:引用经文ID"`
}
