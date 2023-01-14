package main

import (
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
	var feedHandle gin.HandlerFunc = getFeed
	var userLoginHandle gin.HandlerFunc = postUserLogin
	var userRegisterHandle = postUserRegister
	var userHandle = getUser
	var publishActionHandle = postPublishAction
	var publishListHandle = getPublishList
	engine.GET("/douyin/feed/", feedHandle)
	engine.POST("/douyin/user/register", userRegisterHandle)
	engine.POST("/douyin/user/login", userLoginHandle)
	engine.GET("/douyin/user", userHandle)
	engine.POST("/douyin/publish/action", publishActionHandle)
	engine.GET("/douyin/publish/list", publishListHandle)
}
