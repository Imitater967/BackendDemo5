package favorite

import (
	"ByteTechTraining/proto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostFavoriteAction(ctx *gin.Context) {
	var request = proto.DouyinFavoriteActionRequest{}
	var response = proto.DouyinFavoriteActionResponse{}

	err := ctx.ShouldBindQuery(&request)
	if err != nil {
		var statusMsg string = err.Error()
		var statusCode int32 = http.StatusBadRequest
		response.StatusMsg = &statusMsg
		response.StatusCode = &statusCode
		ctx.JSON(http.StatusBadRequest, &response)
		return
	}

	// 实现逻辑
	if 1 == request.GetActionType() {

	} else if 2 == request.GetActionType() {

	}

	var statusMsg string = "successful"
	var statusCode int32 = 0
	response.StatusMsg = &statusMsg
	response.StatusCode = &statusCode
	ctx.JSON(http.StatusOK, &response)
	return
}

func GetFavoriteList(ctx *gin.Context) {
	var request = proto.DouyinFavoriteListRequest{}
	var response = proto.DouyinFavoriteListResponse{}

	err := ctx.ShouldBindQuery(&request)
	if err != nil {
		var statusMsg string = err.Error()
		var statusCode int32 = http.StatusBadRequest
		var videoList []*proto.Video // // 待更新

		response.StatusMsg = &statusMsg
		response.StatusCode = &statusCode
		response.VideoList = videoList
		ctx.JSON(http.StatusBadRequest, &response)
		return
	}

	// 实现逻辑

	var statusMsg string = "successful"
	var statusCode int32 = 0
	var videoList []*proto.Video // 待更新

	response.StatusMsg = &statusMsg
	response.StatusCode = &statusCode
	response.VideoList = videoList
	ctx.JSON(http.StatusOK, &response)
	return
}
