package menu_api

import (
	"blog-server/apps/models"
	"blog-server/common/responsex"
	"blog-server/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// MenuDeleteList 删除菜单
// @Tags 菜单管理
// @Summary 批量删除菜单
// @Description 批量删除菜单
// @Param data body models.RemoveRequest true  "图片id列表"
// @Router /api/menus [delete]
// @Produce json
// @Success 200 {object} responsex.Response{data=string}
func (i *MenuApi) MenuDeleteList(c *gin.Context) {
	var req models.RemoveRequest
	c.ShouldBindJSON(&req)
	var menusList []models.MenuModel
	count := global.DB.Debug().Where("id in (?)", req.IDList).Find(&menusList).RowsAffected
	if count == 0 {
		responsex.FailWithMessage("菜单不存在", c)
		return
	}
	tx := global.DB.Begin()
	err := tx.Model(&menusList).Association("Banners").Clear()
	if err != nil {
		zap.S().Errorf("删除菜单关联图片失败 [ERROR]: %v", err)

	}
	err = tx.Delete(&menusList).Error
	if err != nil {
		responsex.FailWithMessage("删除菜单关联图片失败", c)
		return
	}
	tx.Commit()
	deleteData := []string{}
	for _, m := range menusList {
		deleteData = append(deleteData, m.Title)
	}
	responsex.Ok(deleteData, fmt.Sprintf("删除 %d 条菜单", count), c)
	return
}
