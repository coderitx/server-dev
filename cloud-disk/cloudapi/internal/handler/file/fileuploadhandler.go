package file

import (
	"cloud-disk/cloudapi/internal/models"
	"cloud-disk/cloudapi/utils"
	"crypto/md5"
	"encoding/hex"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"path"

	"cloud-disk/cloudapi/internal/logic/file"
	"cloud-disk/cloudapi/internal/svc"
	"cloud-disk/cloudapi/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		f, fileHeder, err := r.FormFile("file")
		if err != nil {
			return
		}
		// 判断文件是否存在
		b := make([]byte, fileHeder.Size)
		f.Read(b)
		h := md5.New()
		h.Write(b)
		hash := hex.EncodeToString(h.Sum(nil))
		var repoInfo []models.RepositoryPool
		err = svcCtx.DB.Model(models.RepositoryPool{}).Where("hash = ?", hash).Scan(&repoInfo).Debug().Error
		if err != nil {
			logx.Errorf("query file exists by hash error: %v", err)
			return
		}
		if len(repoInfo) > 0 {
			fileInfo := repoInfo[0]
			httpx.OkJson(w, &types.FileUploadReply{
				Name:     fileInfo.Name,
				Ext:      fileInfo.Ext,
				Identity: fileInfo.Identity,
			})
			return
		}
		cosPath, err := utils.UploadFileToCos(r)
		if err != nil {
			logx.Errorf("upload file to cos error: %v", err)
			return
		}
		// 封装传递的的信息
		ext := path.Ext(fileHeder.Filename)
		req.Hash = hash
		req.Size = fileHeder.Size
		req.Ext = ext
		req.Name = fileHeder.Filename
		req.Path = cosPath
		l := file.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
