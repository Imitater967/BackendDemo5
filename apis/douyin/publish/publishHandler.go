package publish

import (
	"ByteTechTraining/proto"
	"ByteTechTraining/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func PostPublishAction(ctx *gin.Context) {
	request := proto.DouyinPublishActionRequest{}
	response := proto.DouyinPublishActionResponse{}
	var requestErr = ctx.ShouldBindQuery(&request)
	if requestErr != nil {
		ctx.JSON(http.StatusBadRequest, requestErr)
		return
	}
	statusMsg := "上传成功"
	var statusCode int32 = 0
	response.StatusMsg = &statusMsg
	response.StatusCode = &statusCode
	video, fileErr := ctx.FormFile("data")
	if fileErr != nil {
		statusMsg = fileErr.Error()
		return
	}
	filePath := strings.Builder{}
	filePath.WriteString(utils.GetFilePath())
	filePath.WriteString(video.Filename)
	ctx.SaveUploadedFile(video, filePath.String())
	ctx.JSON(http.StatusOK, &response)

}

func GetPublishList(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status_code": 200,
		"message":     "list",
	})
}
