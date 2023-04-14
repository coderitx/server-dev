package models

import "time"

type UserRepository struct {
	ID                 uint       `gorm:"primary_key"`
	Identity           string     `json:"identity"`
	UserIdentity       string     `json:"user_identity"`
	ParentId           int64      `json:"parent_id"`
	RepositoryIdentity string     `json:"repository_identity"`
	Ext                string     `json:"ext"`
	Name               string     `json:"name"`
	CreatedAt          time.Time  `gorm:"column:created_at;default:null" `
	UpdatedAt          time.Time  `gorm:"column:updated_at;default:null"`
	DeletedAt          *time.Time `gorm:"column:deleted_at;default:null"`
}

func (table UserRepository) TableName() string {
	return "user_repository"
}
