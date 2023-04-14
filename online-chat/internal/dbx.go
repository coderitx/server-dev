package internal

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"online-chat/config"
	"online-chat/global"
	"time"
)

// InitDB 初始化database
func InitDB(c config.DBConfig) {
	stmt := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.Username, c.Password, c.Addr, c.Port, c.Database)
	fmt.Println("----dbx---", stmt)
	db, err := gorm.Open("mysql", stmt)
	if err != nil {
		zap.S().Errorf("connect mysql error: %v", err)
		return
	}
	db.Model(true)
	db.DB().SetMaxIdleConns(c.MaxConnect)
	db.DB().SetMaxOpenConns(c.OpenConnect)
	db.DB().SetConnMaxIdleTime(time.Second * 10)
	global.DB = db
	return
}
