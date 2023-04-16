package requestx

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"online-chat/global"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string
	Password      string
	Phone         string
	Email         string
	Identity      string
	ClientIp      string
	ClientPort    string
	LoginTime     *time.Time
	HeartbeatTime *time.Time
	LoginOutTime  *time.Time
	IsLogout      bool
	DeviceInfo    string
}

func (u *UserBasic) TableName() string {
	return "user_basic"
}

func (*UserBasic) GetUserList() ([]UserBasic, error) {
	var userBasic []UserBasic
	err := global.DB.Model(&UserBasic{}).Scan(&userBasic).Debug().Error
	if err != nil {
		zap.S().Errorf("find user list error: %v", err)
		return nil, err
	}
	return userBasic, nil
}
