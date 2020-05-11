package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseByOk(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{"code": "0000", "msg": "错误:" + msg, "data": data})
}

func ResponseByErr(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(http.StatusNotFound, gin.H{"code": "9999", "msg": "错误:" + msg, "data": data})
}
