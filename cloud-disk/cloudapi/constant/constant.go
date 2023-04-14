package constant

import (
	"github.com/dgrijalva/jwt-go"
	"os"
)

type UserClaim struct {
	Id       uint64
	Identity string
	Name     string
	jwt.StandardClaims
}

const (
	SuccessCode = "1000"
	ErrorCode   = "2000"
)

// 生成token的key
var JwtKey = "cloud-disk-key"

// 邮箱的授权密码
var EmailPassword = os.Getenv("EmailPwd")

// 随机验证码的长度
var CodeSize = 6

// 验证码过期时间(s)
var CodeExpire = 300

var CosSecretKey = os.Getenv("CosSecretKey")
var CosSecretID = os.Getenv("CosSecretID")
var CosBucket = "https://cloud-disk-1317712916.cos.ap-beijing.myqcloud.com"

// 分页参数
var Page = 1
var Size = 20
