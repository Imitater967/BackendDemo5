package feed

import "github.com/gin-gonic/gin"

var (
	Api *gin.RouterGroup
)

func InitFeedRouterGroup(engine *gin.RouterGroup) {
	Api = engine.Group("/feed")
	Api.POST("/", GetFeed)
}
