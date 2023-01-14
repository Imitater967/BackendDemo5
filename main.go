package main

import (
	route "ByteTechTraining/route"
	"github.com/gin-gonic/gin"
	_ "gorm.io/driver/sqlite"
	_ "gorm.io/gorm"
	_ "net/http"
)

func main() {
	server := gin.Default()
	registerBasicHandles(server)
	server.Run()
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
