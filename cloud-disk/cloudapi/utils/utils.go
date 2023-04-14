package utils

import (
	"cloud-disk/cloudapi/constant"
	"context"
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/jordan-wright/email"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/zeromicro/go-zero/core/logx"
	"math/rand"
	"net/http"
	"net/smtp"
	"net/url"
	"os"
	"path"
	"strings"
	"time"
)

func AnalyzeToken(token string) (*constant.UserClaim, error) {
	uc := new(constant.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(constant.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("token is invalid")
	}
	return uc, err
}

// MD5
// md5加密
func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))

	return strings.ToLower(hex.EncodeToString(h.Sum(nil)))
}

// GenerateToken
// 根据用户信息生成token
func GenerateToken(id uint64, identity, name string, expiredAt int) (string, error) {
	uc := constant.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(expiredAt)).Unix(),
		},
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	token, err := claims.SignedString([]byte(constant.JwtKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

// SendCodeToEmail
// 发送验证码到邮箱
func SendCodeToEmail(mail, code string) error {
	e := email.NewEmail()
	e.From = "golang-developer <super_me0208@163.com>"
	e.To = []string{mail}
	e.Subject = "验证码请求"
	e.HTML = []byte("<h3>你的验证码为: " + code + "</h3>")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "super_me0208@163.com", constant.EmailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		return err
	}
	return err
}

// RandCode
// 生成随机验证码
func RandCode() string {
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < constant.CodeSize; i++ {
		code += fmt.Sprintf("%d", rand.Intn(10))
	}
	return code
}

func UUID() string {
	return uuid.New().String()
}

// UploadFileToCos
// 上传文件至cos
func UploadFileToCos(r *http.Request) (string, error) {
	u, _ := url.Parse(constant.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: os.Getenv(constant.CosSecretID), // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: os.Getenv(constant.CosSecretKey), // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
		},
	})

	// 文件夹名称
	file, fileHeader, err := r.FormFile("file")
	key := "home/private/" + UUID() + path.Ext(fileHeader.Filename)
	// 传递大小为0的输入流
	_, err = client.Object.Put(context.Background(), key, file, nil)
	if err != nil {
		// ERROR
		logx.Errorf("upload file to cos error: %v", err)
		return "", err
	}
	return constant.CosBucket + "/" + key, nil
}
