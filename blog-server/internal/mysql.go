package internal

import (
	"blog-server/config/internal_config"
	"blog-server/global"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// InitDB 初始化database
func InitDB(c internal_config.MysqlConfig) error {
	var mysqlLogLeve logger.LogLevel
	if c.LogLevel == "info" {
		mysqlLogLeve = logger.Info
	} else {
		mysqlLogLeve = logger.Error
	}
	logx := logger.New(
		log.New(os.Stdout, "\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      mysqlLogLeve,
			Colorful:      true,
		})
	dbCfg := mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.DB),
	}

	db, err := gorm.Open(mysql.New(dbCfg), &gorm.Config{Logger: logx})
	if err != nil {
		zap.S().Errorf("connect mysql error: %v", err)
		return err
	}
	db.Model(true)
	global.DB = db
	return nil
}
