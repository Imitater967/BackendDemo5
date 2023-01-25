package video

import (
	"ByteTechTraining/utils"
	"github.com/gin-gonic/gin"
)

var (
	Api *gin.RouterGroup
)

func InitApiRouter(engine *gin.Engine) {
	Api = engine.Group("/file")
	Api.Any("/ping", Ping)
	Api.GET("/", fileServer)

}
func fileServer(c *gin.Context) {
	path := utils.GetFilePath()
	fileName := path + c.Query("id")
	c.File(fileName)
}
