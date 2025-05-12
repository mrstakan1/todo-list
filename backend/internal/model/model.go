package model

import "gorm.io/gorm"

type Todo struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UserID    uint   `gorm:"index;type:bigint" json:"-"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	gorm.Model
}

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex"`
	Password string
}
