package wsx

import "gorm.io/gorm"

// 人员关系
type Contact struct {
	gorm.Model
	OwnerId  uint // 谁的关系信息
	TargetId uint // 对应谁
	Type     uint // 对应类型： 0，1，2
	Desc     string
}

func (*Contact) TableName() string {
	return "contact"
}
