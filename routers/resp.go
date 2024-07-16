package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	200 - 成功
	500 - 服务器错误
	400 - 访问错误
	409 - 重复访问错误
	4001 - Token 过期
	4002 - Token 在黑名单
*/

// ResponseMsg
// @desc 响应消息体
type ResponseMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func MakeResp(ctx *gin.Context, statusCode int, msg string, data any) {
	ctx.JSON(
		statusCode,
		&ResponseMsg{
			Code: statusCode,
			Msg:  msg,
			Data: data,
		},
	)

	ctx.Abort()
}

func Ok(ctx *gin.Context, msg string, data any) {
	MakeResp(ctx, http.StatusOK, msg, data)
}

func Err(ctx *gin.Context, msg string, data any) {
	MakeResp(ctx, http.StatusInternalServerError, msg, data)
}

func BadRequest(ctx *gin.Context, msg string, data any) {
	MakeResp(ctx, http.StatusBadRequest, msg, data)
}

func ConflictRequest(ctx *gin.Context, msg string, data any) {
	MakeResp(ctx, http.StatusConflict, msg, data)
}

// TokenIsExpired Token 过期
func TokenIsExpired(ctx *gin.Context, msg string, data any) {
	MakeResp(ctx, http.StatusUnauthorized, msg, data)
}

// TokenInBlackList Token 在黑名单中
func TokenInBlackList(ctx *gin.Context, msg string, data any) {
	MakeResp(ctx, http.StatusUnauthorized, msg, data)
}
