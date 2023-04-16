package requestx

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"online-chat/common/errorx"
	"online-chat/global"
	"online-chat/utils"
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

func (u *UserBasic) FindUserByID() (*UserBasic, int) {
	user := UserBasic{}
	err := global.DB.Model(&UserBasic{}).Where("id = ?", u.ID).Find(&user).Debug().Error
	if err != nil {
		return nil, errorx.ServerErrorCode
	}
	return &user, errorx.SuccessCode
}

func (u *UserBasic) FindUserByName() (*UserBasic, int) {
	user := UserBasic{}
	err := global.DB.Model(&UserBasic{}).Where("name = ?", u.Name).Find(&user).Debug().Error
	if err != nil {
		return nil, errorx.ServerErrorCode
	}
	return &user, errorx.SuccessCode
}

func (u *UserBasic) FindUserByPhone() (*UserBasic, int) {
	user := UserBasic{}
	err := global.DB.Model(&UserBasic{}).Where("phone = ?", u.Phone).Find(&user).Debug().Error
	if err != nil {
		return nil, errorx.ServerErrorCode
	}
	return &user, errorx.SuccessCode
}

func (u *UserBasic) FindUserByEmail() (*UserBasic, int) {
	user := UserBasic{}
	err := global.DB.Model(&UserBasic{}).Where("email = ?", u.Email).Find(&user).Debug().Error
	if err != nil {
		return nil, errorx.ServerErrorCode
	}
	return &user, errorx.SuccessCode
}

func (u *UserBasic) GetUserList() ([]UserBasic, int) {
	var userBasic []UserBasic
	err := global.DB.Model(&UserBasic{}).Scan(&userBasic).Debug().Error
	if err != nil {
		zap.S().Errorf("find user list error: %v", err)
		return nil, errorx.ServerErrorCode
	}
	return userBasic, errorx.SuccessCode
}

func (u *UserBasic) CreateUser() int {
	user := UserBasic{}
	err := global.DB.Model(&UserBasic{}).Where("name = ?", u.Name).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		zap.S().Errorf("check userinfo exists error: %v", err)
		return errorx.ServerErrorCode
	}
	u.Password = utils.MD5(u.Password)
	if user.ID != 0 {
		return errorx.AlreadyExistsCode
	}
	err = global.DB.Create(&u).Debug().Error
	if err != nil {
		zap.S().Errorf("create user error: %v", err)
		return errorx.ServerErrorCode
	}
	return errorx.SuccessCode
}

func (u *UserBasic) DeleteUser() int {
	err := global.DB.Model(&UserBasic{}).Delete(&u).Where("id = ?", u.ID).Error
	if err != nil {
		return errorx.ServerErrorCode
	}
	return errorx.SuccessCode
}

func (u *UserBasic) UpdateUser() int {
	err := global.DB.Model(&UserBasic{}).Where("id = ?", u.ID).Updates(&u).Debug().Error
	if err != nil {
		return errorx.ServerErrorCode
	}
	return errorx.SuccessCode
}

func (u *UserBasic) FindUserByNameOrEmailOrPhoneAndPwd() (*UserBasic, int) {
	name := u.Name
	phone := u.Phone
	email := u.Email
	if name != "" {
		user := UserBasic{}
		err := global.DB.Model(&UserBasic{}).Where("name = ? AND password = ?", u.Name, utils.MD5(u.Password)).Find(&user).Debug().Error
		if err != nil {
			return nil, errorx.ServerErrorCode
		}
		return &user, errorx.SuccessCode
	}

	if phone != "" {
		user := UserBasic{}
		err := global.DB.Model(&UserBasic{}).Where("phone = ? AND password = ?", u.Phone, utils.MD5(u.Password)).Find(&user).Debug().Error
		if err != nil {
			return nil, errorx.ServerErrorCode
		}
		return &user, errorx.SuccessCode
	}

	if email != "" {
		user := UserBasic{}
		err := global.DB.Model(&UserBasic{}).Where("email = ? AND password = ?", u.Email, utils.MD5(u.Password)).Find(&user).Debug().Error
		if err != nil {
			return nil, errorx.ServerErrorCode
		}
		return &user, errorx.SuccessCode
	}
	return nil, errorx.UnregisteredErrorCode
}
