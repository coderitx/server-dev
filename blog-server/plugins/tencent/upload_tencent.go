package tencent

import (
	"bytes"
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

// UploadImageTencent 上传图片到腾讯云
func UploadImageTencent(imageByte []byte, filename string) (cosPath string, err error) {
	tencentUrl := os.Getenv("TencentUrl")
	u, _ := url.Parse(tencentUrl)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("CosSecretID"),
			SecretKey: os.Getenv("CosSecretKey"),
		},
	})

	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType: getImageContentType(filename),
			//XCosMetaXXX: &http.Header{},
		},
		ACLHeaderOptions: &cos.ACLHeaderOptions{
			// 如果不是必要操作，建议上传文件时不要给单个文件设置权限，避免达到限制。若不设置默认继承桶的权限。
			XCosACL: "public-read",
		},
	}
	cosSubPath := generateCosPath(filepath.Base(filename))
	_, err = client.Object.Put(context.Background(), cosSubPath, bytes.NewReader(imageByte), opt)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v/%v", tencentUrl, cosSubPath), nil
}

// generateCosPath 按照日期构建腾讯云存储路径
func generateCosPath(fs string) string {
	dateFormat := time.Now().Format("20060102")
	return fmt.Sprintf("images/%v/%v", dateFormat, fs)
}

// getImageContentType 获取图片的content-type
func getImageContentType(img string) string {
	fmt.Println(fmt.Sprintf("image/%v", strings.TrimLeft(path.Ext(img), ".")))
	return fmt.Sprintf("image/%v", strings.TrimLeft(path.Ext(img), "."))
}
