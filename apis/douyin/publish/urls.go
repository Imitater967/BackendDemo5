package publish

import "github.com/gin-gonic/gin"

var (
	Api *gin.RouterGroup
)

func InitPublishRouterGroup(engine *gin.RouterGroup) {
	Api = engine.Group("/publish")
	Api.POST("/action/", PostPublishAction)
	Api.GET("/list/", GetPublishList)
}
