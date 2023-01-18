package main

import (
	"ByteTechTraining/orm"
	route "ByteTechTraining/route"
	"github.com/gin-gonic/gin"
	_ "gorm.io/driver/sqlite"
	_ "gorm.io/gorm"
	_ "net/http"
)

func main() {
	orm.ConnectToDatabase()
	server := gin.New()
	server.Use(gin.Logger())
	server.Use(gin.Recovery())
	server.Use(tokenVerify()) // 使用token校验中间件
	registerBasicHandles(server)
	registerInteractHandles(server)
	server.Run(":8080")
}

// token校验中间件 负责校验user身份
func tokenVerify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 前置操作

		ctx.Next()

		// 后置操作
	}
}

func registerInteractHandles(engine *gin.Engine) {
	var favoriteActionHandle = route.PostFavoriteAction
	var favoriteListHandle = route.GetFavoriteList
	var commentActionHandle = route.PostCommentAction
	var commentListHandle = route.GetCommentList

	engine.POST("/douyin/favorite/action/", favoriteActionHandle)
	engine.GET("/douyin/favorite/list/", favoriteListHandle)
	engine.POST("/douyin/comment/action", commentActionHandle)
	engine.GET("/douyin/comment/list", commentListHandle)

}
func registerBasicHandles(engine *gin.Engine) {
	var feedHandle gin.HandlerFunc = route.GetFeed
	var userLoginHandle gin.HandlerFunc = route.PostUserLogin
	var userRegisterHandle = route.PostUserRegister
	var userHandle = route.GetUser
	var publishActionHandle = route.PostPublishAction
	var publishListHandle = route.GetPublishList
	engine.GET("/douyin/feed/", feedHandle)
	engine.POST("/douyin/user/register", userRegisterHandle)
	engine.POST("/douyin/user/login", userLoginHandle)
	engine.GET("/douyin/user", userHandle)
	engine.POST("/douyin/publish/action", publishActionHandle)
	engine.GET("/douyin/publish/list", publishListHandle)
}
