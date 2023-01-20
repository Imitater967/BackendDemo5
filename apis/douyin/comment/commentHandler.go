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
		var statusMsg string = err.Error()
		var statusCode int32 = http.StatusBadRequest
		var comment *proto.Comment // 待更新

		response.StatusMsg = &statusMsg
		response.StatusCode = &statusCode
		response.Comment = comment
		ctx.JSON(http.StatusBadRequest, &response)
		return
	}

	// 实现逻辑(！！！！后续需要更改)

	// action_type = 1表示写评论， action_type = 2表示删评论
	if 1 == request.GetActionType() {
		var commentObj daos.CommentDao
		commentObj.CommentID = 1
		commentObj.CommenterID = 2
		commentObj.Content = "this is content"
		commentObj.CreateTime = time.Now()
		commentObj.IsDeleted = 0

		err := commentObj.Add()
		if err != nil {
			var statusMsg string = err.Error()
			var statusCode int32 = http.StatusBadRequest
			var comment *proto.Comment // 待更新

			response.StatusMsg = &statusMsg
			response.StatusCode = &statusCode
			response.Comment = comment
			ctx.JSON(http.StatusBadRequest, &response)
			return
		}

	}

	var statusMsg string = "successful"
	var statusCode int32 = 0
	var comment *proto.Comment // 待更新

	response.StatusMsg = &statusMsg
	response.StatusCode = &statusCode
	response.Comment = comment
	ctx.JSON(http.StatusOK, &response)
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
