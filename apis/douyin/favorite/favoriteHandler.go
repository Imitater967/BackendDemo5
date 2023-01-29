package favorite

import (
	"ByteTechTraining/daos"
	"ByteTechTraining/proto"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
	var request = proto.DouyinFavoriteListRequest{}
	var response = proto.DouyinFavoriteListResponse{}

	var statusMsg string = "successful"
	var statusCode int32 = 0
	var videoList []*proto.Video
	response.StatusMsg = &statusMsg
	response.StatusCode = &statusCode
	response.VideoList = videoList
	err := ctx.ShouldBindQuery(&request)
	if err != nil {
		statusMsg = err.Error()
		statusCode = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, &response)
		return
	}

	ctx.JSON(http.StatusOK, &response)
	return
}
