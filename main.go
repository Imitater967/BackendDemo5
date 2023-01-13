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
	var loginHandle gin.HandlerFunc = postLogin
	var registerHandle = postRegister
	engine.GET("/douyin/feed/", feedHandle)
	engine.POST("/douyin/user/register", registerHandle)
	engine.POST("/douyin/user/login", loginHandle)
}
