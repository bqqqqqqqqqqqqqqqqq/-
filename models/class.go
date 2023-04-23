package models

import (
	"gorm.io/gorm"
	"time"
)

type Class struct {
	ID        uint           `gorm:"primarykey;" json:"id"`
	ClassName string         `grom:"column:classname" json:"classname"`
	StuID     uint           `gorm:"column:stuid" json:"stuid"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;" json:"deleted_at"`
}

func (table *Class) TableName() string {
	return "class"
}
