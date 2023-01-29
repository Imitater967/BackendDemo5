package comment

import (
	"ByteTechTraining/daos"
	"ByteTechTraining/proto"
	"github.com/gin-gonic/gin"
	"log"
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
	userAuthDao := daos.UserAuthDao{}
	//相关变量
	var statusMsg string = "填写成功"
	var statusCode int32 = 0
	var comment *proto.Comment
	response.StatusMsg = &statusMsg
	response.StatusCode = &statusCode
	response.Comment = comment
	var commentObj daos.CommentDao
	//1.从表单中获取Token,检索相关用户
	userAuthDao.Token = request.GetToken()
	var userExitErr = userAuthDao.Get()
	if userExitErr != nil {
		statusMsg = userExitErr.Error()
		ctx.JSON(http.StatusOK, &response)
		log.Println("User Not Login")
		return
	}

	// action_type = 1表示写评论， action_type = 2表示删评论
	//缺少用户校队
	if 1 == request.GetActionType() {
		commentObj.Content = request.GetCommentText()
		commentObj.VideoId = request.GetVideoId()
		commentObj.Date = time.Now()
		err := commentObj.Add()
		if err != nil {
			statusCode = 1
			statusMsg = err.Error()
			ctx.JSON(http.StatusOK, &response)
			return
		}
		statusMsg = "填写成功"
		ctx.JSON(http.StatusOK, &response)

	}
	//删除评论
	//缺少用户校队
	if 2 == request.GetActionType() {
		commentObj.Id = request.GetCommentId()
		err := commentObj.Delete()
		if err != nil {
			statusCode = 2
			statusMsg = err.Error()
			ctx.JSON(http.StatusOK, &response)
			return
		}
		statusMsg = "删除成功"
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
