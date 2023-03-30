package models

import (
	"gorm.io/gorm"
	"time"
)

type Goods struct {
	GoodsID   string         `gorm:"primarykey" json:"id"`
	Name      string         `gorm:"column:name" json:"name"`
	Describe  string         `gorm:"column:describe" json:"describe"'`
	Price     string         `gorm:"column:price" json:"price"`
	Picture   string         `gorm:"column:picture" json:"picture"`
	CreatedAt time.Time      `gorm:"type:timestamp;" json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;" json:"deleted_at"`
}

func (table *Goods) TableName() string {
	return "goods"
}

func GetAllGoods() *gorm.DB {
	tx := DB.Model(&Goods{}).Limit(10).Find(&Goods{})
}
