package douyin

import (
	"ByteTechTraining/apis/douyin/comment"
	"ByteTechTraining/apis/douyin/favorite"
	"ByteTechTraining/apis/douyin/feed"
	"ByteTechTraining/apis/douyin/publish"
	"ByteTechTraining/apis/douyin/user"
	"github.com/gin-gonic/gin"
)

var (
	Api *gin.RouterGroup
)

func InitApiRouter(engine *gin.Engine) {
	Api = engine.Group("/douyin")
	Api.Any("/ping", Ping)

	// 注册路由
	user.InitUserRouterGroup(Api)
	publish.InitPublishRouterGroup(Api)
	comment.InitCommentRouterGroup(Api)
	feed.InitFeedRouterGroup(Api)
	favorite.InitFavoriteRouterGroup(Api)
}
