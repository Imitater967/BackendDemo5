package favorite

import "github.com/gin-gonic/gin"

var (
	Api *gin.RouterGroup
)

func InitFavoriteRouterGroup(engine *gin.RouterGroup) {
	Api = engine.Group("/favorite")
	Api.POST("/action/", PostFavoriteAction)
	Api.GET("/list/", GetFavoriteList)
}
