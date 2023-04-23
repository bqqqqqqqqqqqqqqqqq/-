package models

import "gorm.io/gorm"

type Picture struct {
	gorm.Model
	Path      string `grom:"column:path;NOT NULL" json:"path"`
	UserId    string `gorm:"column:user_id;NULL" json:"user_id"`
	ProblemId string `gorm:"column:problem_id;NULL" json:"problem_id"`
}

func (table *Picture) TableName() string {
	return "picture"
}
