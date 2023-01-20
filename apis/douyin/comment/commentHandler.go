package comment

import (
	"ByteTechTraining/daos"
	"ByteTechTraining/proto"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func PostCommentAction(ctx *gin.Context) {
	var request = proto.DouyinCommentActionRequest{}
	var response = proto.DouyinCommentActionResponse{}

	err := ctx.ShouldBindQuery(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	// action_type = 1表示写评论， action_type = 2表示删评论
	//缺少用户校队
	if 1 == request.GetActionType() {
		var commentObj daos.CommentDao
		//commentObj.Id = *request.CommentId
		commentObj.Content = request.GetCommentText()
		commentObj.VideoId = request.GetVideoId()
		commentObj.Date = time.Now()
		var statusMsg string = "填写成功"
		var statusCode int32 = 0
		var comment *proto.Comment
		err := commentObj.Add()
		if err != nil {
			statusCode = 1
			statusMsg = err.Error()
		}
		response.StatusMsg = &statusMsg
		response.StatusCode = &statusCode
		response.Comment = comment
		ctx.JSON(http.StatusOK, &response)

	}
	//删除评论
	//缺少用户校队
	if 2 == request.GetActionType() {
		var commentObj daos.CommentDao
		commentObj.Id = request.GetCommentId()
		err := commentObj.Delete()
		var statusMsg string = "删除成功"
		var statusCode int32 = 0
		if err != nil {
			statusCode = 2
			statusMsg = err.Error()
		}

		var comment *proto.Comment
		response.StatusMsg = &statusMsg
		response.StatusCode = &statusCode
		response.Comment = comment
		ctx.JSON(http.StatusOK, &response)
	}
	return
}
func GetCommentList(ctx *gin.Context) {
	var request = proto.DouyinCommentListRequest{}
	var response = proto.DouyinCommentListResponse{}

	err := ctx.ShouldBindQuery(&request)
	if err != nil {
		var statusMsg string = err.Error()
		var statusCode int32 = http.StatusBadRequest
		var commentList []*proto.Comment // 待更新

		response.StatusMsg = &statusMsg
		response.StatusCode = &statusCode
		response.CommentList = commentList
		ctx.JSON(http.StatusBadRequest, &response)
		return
	}

	// 实现逻辑

	var statusMsg string = "successful"
	var statusCode int32 = 0
	var commentList []*proto.Comment // 待更新

	response.StatusMsg = &statusMsg
	response.StatusCode = &statusCode
	response.CommentList = commentList
	ctx.JSON(http.StatusOK, &response)
	return
}
