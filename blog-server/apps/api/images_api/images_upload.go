package images_api

import (
	"blog-server/apps/service"
	"blog-server/apps/service/image_svc"
	"blog-server/common/errorx"
	"blog-server/common/responsex"
	"blog-server/global"
	"blog-server/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ImageUploadView 上传单个文件
func (i *ImagesApi) ImageUploadView(c *gin.Context) {
	// 上传多个文件
	form, err := c.MultipartForm()
	if err != nil {
		responsex.FailWithMessage(err.Error(), c)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		responsex.FailWithCode(errorx.ArgumentError, c)
		return
	}
	// 判断保存文件的路径是否存在，不存在则创建
	err = utils.IsExists(global.GlobalC.Uploads.Path)
	if err != nil {
		responsex.FailWithMessage(err.Error(), c)
		return
	}

	// 判断上传结果
	var resList []image_svc.FileUploadResponse
	for _, file := range fileList {
		// 	文件上传
		serviceRes := service.ServiceApp.ImageService.ImageUploadService(file)
		// 上传失败
		if !serviceRes.IsSuccess {
			resList = append(resList, serviceRes)
			continue
		}
		// 上传成功
		if !global.GlobalC.Tencent.Enable {
			err = c.SaveUploadedFile(file, serviceRes.Filename)
			if err != nil {
				zap.S().Error("filename: %v filesize: %d 保存失败", file.Filename, file.Size)
				serviceRes.Msg = fmt.Sprintf("%v 文件保存失败", file.Filename)
				continue
			}
			serviceRes.Msg = "上传成功"
			serviceRes.IsSuccess = true
		}
		resList = append(resList, serviceRes)
	}
	responsex.OkWithData(resList, c)
	return
}
