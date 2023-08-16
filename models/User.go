package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"unique_index; not null"`
	Tasks     []Task `gorm:"foreignKey:UserID"`
}
