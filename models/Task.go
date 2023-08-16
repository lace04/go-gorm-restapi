package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `gorm:"type:varchar(125); not null" json:"title"`
	Description string `gorm:"not null" json:"description"`
	Completed   bool   `gorm:"default:false" json:"completed"`
	UserID      uint   `gorm:"not null" json:"user_id"`
}
