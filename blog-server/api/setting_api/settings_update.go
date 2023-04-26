package setting_api

import (
	"blog-server/common/errorx"
	"blog-server/common/responsex"
	"blog-server/config/internal_config"
	"blog-server/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"io/fs"
	"io/ioutil"
)

// SettingsInfoUpdate 修改配置文件
func (s *SettingApi) SettingsInfoUpdate(c *gin.Context) {
	var su SettingUri
	err := c.ShouldBindUri(&su)
	if err != nil {
		responsex.FailWithMessage(err.Error(), c)
		return
	}
	switch su.Name {
	case "site":
		var info internal_config.SiteInfo
		err = c.ShouldBindJSON(&info)
		if err != nil {
			responsex.FailWithCode(errorx.ArgumentError, "", c)
			return
		}
		global.GlobalC.SiteInfo = info
	case "email":
		var info internal_config.Email
		err = c.ShouldBindJSON(&info)
		if err != nil {
			responsex.FailWithCode(errorx.ArgumentError, "", c)
			return
		}
		global.GlobalC.Email = info
	case "qq":
		var info internal_config.QQ
		err = c.ShouldBindJSON(&info)
		if err != nil {
			responsex.FailWithCode(errorx.ArgumentError, "", c)
			return
		}
		global.GlobalC.QQ = info
	case "tencent":
		var info internal_config.COS
		err = c.ShouldBindJSON(&info)
		if err != nil {
			responsex.FailWithCode(errorx.ArgumentError, "", c)
			return
		}
		global.GlobalC.Tencent = info

	}
	err = UpdateYaml()
	if err != nil {
		responsex.FailWithMessage(err.Error(), c)
		return
	}
	responsex.OkWith(c)
}

// UpdateYaml 更新yaml文件
func UpdateYaml() error {
	byteData, err := yaml.Marshal(global.GlobalC)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(global.ConfigPath, byteData, fs.ModePerm)
	if err != nil {
		return err
	}
	zap.S().Infoln("配置文件更新成功")
	return nil
}
