package user

import "github.com/gin-gonic/gin"

var (
	Api *gin.RouterGroup
)

func InitUserRouterGroup(engine *gin.RouterGroup) {
	Api = engine.Group("/user")
	Api.POST("/register/", PostUserRegister)
	Api.POST("/login/", PostUserLogin)
	Api.GET("/", GetUser)
}
