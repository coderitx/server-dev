package models

import "time"

type RepositoryPool struct {
	ID        uint       `gorm:"primary_key"`
	Identity  string     `json:"identity"`
	Hash      string     `json:"hash"`
	Name      string     `json:"name"`
	Ext       string     `json:"ext"`
	Size      int64      `json:"size"`
	Path      string     `json:"path"`
	CreatedAt time.Time  `gorm:"column:created_at;default:null" `
	UpdatedAt time.Time  `gorm:"column:updated_at;default:null"`
	DeletedAt *time.Time `gorm:"column:deleted_at;default:null"`
}

func (table RepositoryPool) TableName() string {
	return "repository_pool"
}
