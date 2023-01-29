package favorite

import (
	"ByteTechTraining/daos"
	"ByteTechTraining/models"
	"ByteTechTraining/proto"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func PostFavoriteAction(ctx *gin.Context) {
	//定义基本的requsetg和response
	var request = proto.DouyinFavoriteActionRequest{}
	var response = proto.DouyinFavoriteActionResponse{}
	//定义返回用的变量
	var statusMsg string = "未知代码"
	var statusCode int32 = 0
	response.StatusMsg = &statusMsg
	response.StatusCode = &statusCode
	//绑定query
	bindErr := ctx.ShouldBindQuery(&request)
	if bindErr != nil {
		statusMsg = bindErr.Error()
		statusCode = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, &response)
		return
	}
	//用户登录
	userAuthDao := daos.UserAuthDao{}
	userAuthDao.Token = request.GetToken()
	var userExitErr = userAuthDao.Get()
	if userExitErr != nil {
		statusMsg = userExitErr.Error()
		ctx.JSON(http.StatusOK, &response)
		log.Println("User Not Login")
		return
	}
	//创建Dao用于相关操作
	favoriteDao := daos.VideoFavoriteDao{}
	favoriteDao.User = userAuthDao.Id
	favoriteDao.Video = request.GetVideoId()
	// 点赞操作
	if 1 == request.GetActionType() {
		markErr := favoriteDao.Mark()
		if markErr != nil {
			statusMsg = markErr.Error()
			ctx.JSON(http.StatusOK, &response)
			return
		}
		statusMsg = "喜欢成功"
	}
	//取消点赞操作
	if 2 == request.GetActionType() {
		markErr := favoriteDao.Unmark()
		if markErr != nil {
			statusMsg = markErr.Error()
			ctx.JSON(http.StatusOK, &response)
			return
		}
		statusMsg = "取消喜欢成功"
	}

	ctx.JSON(http.StatusOK, &response)
	return
}

func GetFavoriteList(ctx *gin.Context) {
	//请求和回复的变量
	var request = proto.DouyinFavoriteListRequest{}
	var response = proto.DouyinFavoriteListResponse{}

	var statusMsg string = "successful"
	var statusCode int32 = 0
	var videos = make([]*proto.Video, 0)
	response.StatusMsg = &statusMsg
	response.StatusCode = &statusCode
	//请求数据绑定
	bindErr := ctx.ShouldBindQuery(&request)
	if bindErr != nil {
		statusMsg = bindErr.Error()
		statusCode = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, &response)
		return
	}
	//用户登录
	userAuthDao := daos.UserAuthDao{}
	userAuthDao.Token = request.GetToken()
	var userExitErr = userAuthDao.Get()
	if userExitErr != nil {
		statusMsg = userExitErr.Error()
		ctx.JSON(http.StatusOK, &response)
		log.Println("User Not Login")
		return
	}
	favoriteDao, sqlErr := daos.GetFavoriteVideos(userAuthDao.Id)
	if sqlErr != nil {
		statusMsg = sqlErr.Error()
		ctx.JSON(http.StatusOK, &response)
		return
	}
	for _, dao := range favoriteDao {
		videoDao := daos.VideoDao{VideoModel: models.VideoModel{Id: dao.Video}}
		var sqlErr = videoDao.Get()
		if sqlErr != nil {
			continue
		}
		var video = proto.Video{}
		videos = append(videos, &video)
		video.Title = &videoDao.Title
		video.Id = &videoDao.Id
		playUrl := "http://localhost:8080/file/?id=" + strconv.FormatInt(*video.Id, 10)
		video.PlayUrl = &playUrl
	}
	//这个变量不是指针变量,所以需要重新赋值
	response.VideoList = videos
	ctx.JSON(http.StatusOK, &response)
	return
}
