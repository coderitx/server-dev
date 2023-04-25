package responsex

import (
	"blog-server/common/errorx"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"msg"`
}

const (
	Success = 0
	Err     = 7
)

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Data:    data,
		Message: msg,
	})
}

func Ok(data any, msg string, c *gin.Context) {
	Result(Success, data, msg, c)
}

func OkWithData(data any, c *gin.Context) {
	Result(Success, data, "成功", c)
}

func OkWithMessage(msg string, c *gin.Context) {
	Result(Success, map[string]any{}, msg, c)
}

func OkWith(c *gin.Context) {
	Result(Success, map[string]any{}, "成功", c)
}

func Fail(data any, msg string, c *gin.Context) {
	Result(Err, data, msg, c)
}

func FailWithMessage(msg string, c *gin.Context) {
	Result(Err, map[string]any{}, msg, c)
}

func FailWithCode(code errorx.ErrorCode, msg string, c *gin.Context) {
	msg, ok := errorx.ErrorMap[code]
	if ok {
		Result(int(code), map[string]any{}, msg, c)
	}
	Result(Err, map[string]any{}, msg, c)
}
