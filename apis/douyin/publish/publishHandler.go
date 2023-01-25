package publish

import (
	"ByteTechTraining/daos"
	"ByteTechTraining/proto"
	"ByteTechTraining/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func PostPublishAction(ctx *gin.Context) {
	request := proto.DouyinPublishActionRequest{}
	response := proto.DouyinPublishActionResponse{}
	//request绑定
	var requestErr = ctx.ShouldBindQuery(&request)
	if requestErr != nil {
		ctx.JSON(http.StatusBadRequest, requestErr)
		return
	}

	//respond相关信息
	statusMsg := "上传成功"
	var statusCode int32 = 0
	response.StatusMsg = &statusMsg
	response.StatusCode = &statusCode
	videoDao := daos.VideoDao{}
	userAuthDao := daos.UserAuthDao{}
	//1.从表单中获取Token,检索相关用户
	userAuthDao.Token = *request.Token
	var userExitErr = userAuthDao.Get()
	if userExitErr != nil {
		statusMsg = userExitErr.Error()
		ctx.JSON(http.StatusOK, &response)
		log.Println("User Not Login")
		return
	}
	//2. 配置视频信息
	videoDao.Title = *request.Title
	videoDao.Uploader = userAuthDao.Id
	//从body中名为data的表单中提取文件
	video, fileErr := ctx.FormFile("data")
	if fileErr != nil {
		statusMsg = fileErr.Error()
		ctx.JSON(http.StatusOK, &response)
		log.Println("File PreUpload Fail")
		return
	}
	//文件上传成功,在数据库同样上传
	uploadErr := videoDao.PreUpload()
	if uploadErr != nil {
		statusMsg = uploadErr.Error()
		ctx.JSON(http.StatusOK, &response)
		log.Println("Sql PreUpload Fail")
		return
	}
	filePath := strings.Builder{}
	//uplaoded-video/
	filePath.WriteString(utils.GetFilePath())
	//生成的文件id这句是错误的
	//filePath.WriteString(string(videoDao.Id))
	filePath.WriteString(strconv.FormatInt(videoDao.Id, 10))
	fileSaveErr := ctx.SaveUploadedFile(video, filePath.String())
	if fileSaveErr != nil {
		statusMsg = fileSaveErr.Error()
		ctx.JSON(http.StatusOK, &response)
		log.Println("File Save Fail")
		videoDao.CancelUpload()
		return
	}
	videoDao.FinishUpload()
	statusMsg = "上传成功,id" + strconv.FormatInt(videoDao.Id, 10)
	ctx.JSON(http.StatusOK, &response)

}

func GetPublishList(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status_code": 200,
		"message":     "list",
	})
}
