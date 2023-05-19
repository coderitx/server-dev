package image_svc

import (
	"blog-server/apps/models"
	"blog-server/apps/models/ctype"
	"blog-server/global"
	"blog-server/plugins/tencent"
	"blog-server/utils"
	"fmt"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"path"
	"strings"
)

type FileUploadResponse struct {
	Filename  string `json:"filename"`   // 上传的文件名称
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Msg       string `json:"msg"`        // 消息
}

func (ImageService) ImageUploadService(file *multipart.FileHeader) (res FileUploadResponse) {
	fs, _ := file.Open()
	byteData, err := io.ReadAll(fs)
	if err != nil {
		zap.S().Errorf("read file error: %v", err)
		return
	}
	filename := file.Filename
	savePath := path.Join(global.GlobalC.Uploads.Path, filename)
	res.Filename = savePath
	// 是否存在白名单中
	if !ImgWhitelistVerification(filename) {
		zap.S().Warnf("%v 上传图片不合法,上传失败", filename)
		res.Msg = fmt.Sprintf("%v 上传图片不合法,上传失败", filename)
		return
	}

	size := float64(file.Size) / float64(1024*1024)
	if size > float64(global.GlobalC.Uploads.Size) {
		zap.S().Warnf("%v 文件大小超出设定大小，设定大小为%dMB，未保存", file.Filename, global.GlobalC.Uploads.Size)
		res.Msg = fmt.Sprintf("%v 文件大小超出设定大小，当前文件大小为: %.2fMB，设定大小为%dMB，未保存", filename, size, global.GlobalC.Uploads.Size)
		return
	}
	// 文件hash存储数据库
	imgHash := utils.MD5(byteData)
	var imgObj models.BannerModel
	err = global.DB.Take(&imgObj, "hash = ?", imgHash).Error
	if err == nil {
		res.Filename = imgObj.Path
		res.Msg = fmt.Sprintf("%v 图片已存在", filename)
		return
	}

	imageType := ctype.Local
	res.Msg = "图片上传成功"
	res.IsSuccess = true

	if global.GlobalC.Tencent.Enable {
		savePath, err = tencent.UploadImageTencent(byteData, filename)
		if err != nil {
			res.Msg = fmt.Sprintf("%v 文件上传失败", filename)
			res.IsSuccess = false
			return
		}
		res.Filename = savePath
		imageType = ctype.Tencent
	}

	if res.IsSuccess {
		global.DB.Create(&models.BannerModel{
			Path:      savePath,
			Hash:      imgHash,
			Name:      filename,
			ImageType: imageType,
		})
	}
	return
}

// ImgWhitelistVerification 上传图片白名单验证
func ImgWhitelistVerification(filename string) bool {
	ext := path.Ext(filename)
	for _, whiteExt := range WhiteImageList {
		if strings.ToLower(ext) == whiteExt {
			return true
		}
	}
	return false
}
