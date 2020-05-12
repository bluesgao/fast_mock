package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SUCCESS(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{"code": "0000", "msg": "成功", "data": data})
}

func FAIL(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(http.StatusNotFound, gin.H{"code": "9999", "msg": "失败:" + msg, "data": data})
}
