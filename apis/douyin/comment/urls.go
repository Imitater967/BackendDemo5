package comment

import "github.com/gin-gonic/gin"

var (
	Api *gin.RouterGroup
)

func InitCommentRouterGroup(engine *gin.RouterGroup) {
	Api = engine.Group("/comment")
	Api.POST("/action/", PostCommentAction)
	Api.GET("/list/", GetCommentList)
}
