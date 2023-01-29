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
	var requestErr = ctx.ShouldBind(&request)

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
	//userAuthDao.Token = *request.Token
	userAuthDao.Token = request.GetToken()
	var userExitErr = userAuthDao.Get()
	if userExitErr != nil {
		statusMsg = userExitErr.Error()
		ctx.JSON(http.StatusOK, &response)
		log.Println("User Not Login")
		return
	}

	//2. 配置视频信息
	//videoDao.Title = *request.Title
	videoDao.Title = request.GetTitle()
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
	filePath.WriteString(strconv.FormatInt(videoDao.Id, 10))
	fileExt := strings.SplitAfter(video.Filename, ".")[1]
	filePath.WriteString("." + fileExt)

	fileSaveErr := ctx.SaveUploadedFile(video, filePath.String())
	if fileSaveErr != nil {
		statusMsg = fileSaveErr.Error()
		ctx.JSON(http.StatusOK, &response)
		log.Println("File Save Fail")
		videoDao.CancelUpload()
		return
	}
	uploadErr = videoDao.FinishUpload()
	if uploadErr != nil {
		statusMsg = uploadErr.Error()
		ctx.JSON(http.StatusOK, &response)
		log.Println("File Upload Fail")
		return
	}
	statusMsg = "上传成功,id" + strconv.FormatInt(videoDao.Id, 10)
	ctx.JSON(http.StatusOK, &response)

}

func GetPublishList(ctx *gin.Context) {
	request := proto.DouyinPublishListRequest{}
	response := proto.DouyinPublishListResponse{}
	//request绑定
	var requestErr = ctx.ShouldBindQuery(&request)
	if requestErr != nil {
		ctx.JSON(http.StatusBadRequest, requestErr)
		return
	}

	//respond相关信息
	statusMsg := "获取成功"
	var statusCode int32 = 0
	var videos = make([]*proto.Video, 0)
	response.StatusMsg = &statusMsg
	response.StatusCode = &statusCode
	//从数据库中获取视频
	videoDaos, queryErr := daos.GetUploadedVideos(request.GetUserId())
	if queryErr != nil {
		statusMsg = queryErr.Error()
		ctx.JSON(http.StatusOK, &response)
		return
	}
	//遍历查询到的数据,填写至video
	for _, dao := range videoDaos {
		var (
			video   = proto.Video{}
			playUrl = "http://localhost:8080/file/?id=" + strconv.FormatInt(*video.Id, 10)
		)
		videos = append(videos, &video)
		video.Title = &dao.Title
		video.Id = &dao.Id
		video.PlayUrl = &playUrl
	}
	//这个变量不是指针变量,所以需要重新赋值
	response.VideoList = videos
	ctx.JSON(http.StatusOK, &response)
}
