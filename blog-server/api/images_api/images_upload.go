package images_api

import (
	"blog-server/common/responsex"
	"blog-server/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
	"path"
)

type FileUploadResponse struct {
	Filename  string `json:"filename"`   // 上传的文件名称
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Msg       string `json:"msg"`        // 消息
}

// ImagesUploadView 上传单个文件
func (i *ImagesApi) ImagesUploadView(c *gin.Context) {
	// 上传多个文件
	form, err := c.MultipartForm()
	if err != nil {
		responsex.FailWithMessage(err.Error(), c)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		responsex.FailWithMessage("文件不存在", c)
		return
	}
	// 判断保存文件的路径是否存在，不存在则创建
	err = IsExists(global.GlobalC.Uploads.Path)
	if err != nil {
		responsex.FailWithMessage(err.Error(), c)
		return
	}
	// 判断上传结果
	var resList []FileUploadResponse
	for _, file := range fileList {
		size := float64(file.Size) / float64(1024*1024)
		if size > float64(global.GlobalC.Uploads.Size) {
			zap.S().Warnf("%v 文件大小超出设定大小，设定大小为%dMB，未保存", file.Filename, global.GlobalC.Uploads.Size)
			resList = append(resList, FileUploadResponse{
				Filename:  file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("%v 文件大小超出设定大小，当前文件大小为: %.2fMB，设定大小为%dMB，未保存", file.Filename, size, global.GlobalC.Uploads.Size),
			})
			continue
		}
		savePath := path.Join(global.GlobalC.Uploads.Path, file.Filename)
		err := c.SaveUploadedFile(file, savePath)
		if err != nil {
			zap.S().Error("filename: %v filesize: %d 保存失败", file.Filename, file.Size)
			resList = append(resList, FileUploadResponse{
				Filename:  file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("%v 文件保存失败", file.Filename),
			})
			continue
		}
		resList = append(resList, FileUploadResponse{
			Filename:  file.Filename,
			IsSuccess: true,
			Msg:       "上传成功",
		})
	}
	responsex.OkWithData(resList, c)
	return
}

func IsExists(dir string) error {
	_, err := os.ReadDir(dir)
	if err != nil {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			zap.S().Errorf("创建文件夹: %v 失败，error: %v", dir, err.Error())
			return err
		}
	}
	return nil
}
