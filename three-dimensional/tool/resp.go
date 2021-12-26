package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(ctx *gin.Context)  {
	ctx.JSON(http.StatusOK,gin.H{
		"info":"success",
	})
}

func Failure(ctx *gin.Context)  {
	ctx.JSON(http.StatusOK,gin.H{
		"info":"failure",
	})
}
