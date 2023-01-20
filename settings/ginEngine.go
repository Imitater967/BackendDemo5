package settings

import (
	"ByteTechTraining/apis/douyin"
	"ByteTechTraining/middlewares"
	"github.com/gin-gonic/gin"
)

func InitGinEngine() (*gin.Engine, error) {
	engine := gin.New()
	gin.SetMode(gin.DebugMode)

	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(middlewares.TokenVerify())    // 加载token校验中间件
	engine.Use(middlewares.CorsMiddleware()) // 加载跨域处理中间件

	// 初始化路由组
	douyin.InitApiRouter(engine)

	return engine, nil
}
