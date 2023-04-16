package wsx

import "gorm.io/gorm"

// 群组
type GroupBasic struct {
	gorm.Model
	Name    string
	OwnerId uint
	Icon    string
	Desc    string
	Type    int
}

func (*GroupBasic) TableName() string {
	return "group_basic"
}
