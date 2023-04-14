package internal

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"time"
)

// InitDB 初始化database
func InitDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(8.141.175.100:3306)/online_chat?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		zap.S().Errorf("connect mysql error: %v", err)
		return nil
	}
	db.Model(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(10)
	db.DB().SetConnMaxIdleTime(time.Second * 10)
	return db
}
