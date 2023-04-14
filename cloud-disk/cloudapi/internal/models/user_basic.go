package models

import (
	"time"
)

type UserBasic struct {
	ID        uint       `gorm:"primary_key"`
	Identity  string     `json:"identity"`
	Name      string     `json:"name"`
	Password  string     `json:"password"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `gorm:"column:created_at;default:null" `
	UpdatedAt time.Time  `gorm:"column:updated_at;default:null"`
	DeletedAt *time.Time `gorm:"column:deleted_at;default:null"`
}

func (table UserBasic) TableName() string {
	return "user_basic"
}
