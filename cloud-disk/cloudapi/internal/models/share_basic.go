package models

import "time"

type ShareBasic struct {
	ID                     uint       `gorm:"primary_key"`
	Identity               string     `json:"identity"`
	UserIdentity           string     `json:"user_identity"`
	UserRepositoryIdentity string     `json:"user_repository_identity"`
	RepositoryIdentity     string     `json:"repository_identity"`
	ExpiredTime            int        `json:"expired_time"`
	ClickNum               int        `json:"click_num"`
	CreatedAt              time.Time  `gorm:"column:created_at;default:null" `
	UpdatedAt              time.Time  `gorm:"column:updated_at;default:null"`
	DeletedAt              *time.Time `gorm:"column:deleted_at;default:null"`
}

func (table ShareBasic) TableName() string {
	return "share_basic"
}
