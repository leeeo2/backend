package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Resp(ctx *gin.Context, httpStatus int, code int, data interface{}, msg string) {
	ctx.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

func Success(ctx *gin.Context, data interface{}, msg string) {
	Resp(ctx, http.StatusOK, 200, data, msg)
}

func Fail(ctx *gin.Context, data interface{}, msg string) {
	Resp(ctx, http.StatusOK, 400, data, msg)
}
