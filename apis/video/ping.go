package video

import "github.com/gin-gonic/gin"

func Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status_code": 200,
		"message":     "pong",
	})
}
