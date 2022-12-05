package models

import "gorm.io/gorm"

type Chapter struct {
	gorm.Model
	BookId     uint   `gorm:"type:int(20);not null;default:0;comment:卷id"`
	BookName   string `gorm:"type:varchar(200);not null;default:;comment:卷名"`
	ShortName  string `gorm:"type:varchar(200);not null;default:;comment:卷名缩写"`
	ChapterId  uint   `gorm:"type:int(20);not null;default:0;comment:章"`
	SectionId  uint   `gorm:"type:int(20);not null;default:0;comment:节"`
	Content    string `gorm:"type:text;comment:内容"`
	TimesCited uint8  `gorm:"type:int(20);not null;default:0;comment:引用次数"`
}
