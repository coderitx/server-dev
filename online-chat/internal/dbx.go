package internal

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"online-chat/config"
	"online-chat/global"
	"os"
	"time"
)

// InitDB 初始化database
func InitDB(c config.DBConfig) {
	logx := logger.New(
		log.New(os.Stdout, "\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		})
	dbCfg := mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.Username, c.Password, c.Addr, c.Port, c.Database),
	}
	db, err := gorm.Open(mysql.New(dbCfg), &gorm.Config{Logger: logx})
	if err != nil {
		zap.S().Errorf("connect mysql error: %v", err)
		return
	}
	db.Model(true)
	global.DB = db
	return
}
