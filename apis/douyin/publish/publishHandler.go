package publish

import "github.com/gin-gonic/gin"

func PostPublishAction(ctx *gin.Context) {

}

func GetPublishList(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status_code": 200,
		"message":     "list",
	})
}
