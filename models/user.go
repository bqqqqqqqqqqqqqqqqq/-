package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           uint           `gorm:"primarykey;" json:"id"`
	Name         string         `gorm:"column:name" json:"name"`
	IsAdmination string         `gorm:"column:isadmination:type:varchat(2)" json:"isAdmination"`
	Password     string         `gorm:"column:password:type:varchat(30)" `
	CreatedAt    time.Time      `gorm:"type:timestamp;" json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index;" json:"deleted_at"`
}

func (table *User) TableName() string {
	return "user"
}

type userAPI struct {
	ID           uint   `gorm:"primarykey;" json:"id"`
	IsAdmination string `gorm:"column:isadmination:type:varchat(2)" json:"isAdmination"`
}

func GetAllUser() *gorm.DB {
	tx := DB.Model(&User{}).Limit(10).Find(&userAPI{}).Where("isadmination = 0")
	return tx.Order("User.id DESC")
}
