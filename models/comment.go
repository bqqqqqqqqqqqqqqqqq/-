package models

type Comment struct {
	Content string `gorm:"column:content" json:"content"`
	Goods   string `gorm:"column:goods" json:"goods"`
	Reply   string `gorm:"column:Reply" json:"reply"`
}

func (table *Comment) TableName() string {
	return "comment"
}
