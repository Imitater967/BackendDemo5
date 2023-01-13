package main

import (
	"ByteTechTraining/proto"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// http://127.0.0.1:8080/douyin/feed/?Token=1&LatestTime=2
// 按理来说上面的应该会被成功执行,但结果并不是
func getFeed(context *gin.Context) {
	var request = proto.DouyinFeedRequest{}
	if context.ShouldBind(&request) == nil {
		log.Printf("Latest Time %d\n", *request.LatestTime)
		log.Printf("Token %s\n", *request.Token)
	}
	var response = &proto.DouyinFeedResponse{}
	//var videoList []*proto.Video = make([]*proto.Video, 30)
	var videoList []*proto.Video = make([]*proto.Video, 0)
	var msg = "welcome"
	var statusCode int32 = 1
	response.StatusMsg = &msg
	response.StatusCode = &statusCode
	response.VideoList = videoList
	context.ProtoBuf(http.StatusOK, response)
}

// http://127.0.0.1:8080/douyin/user/login
func postLogin(ctx *gin.Context) {

}

// curl -v --form Username=user --form Password=password http://127.0.0.1:8080/douyin/user/register
func postRegister(ctx *gin.Context) {
	var request = &proto.DouyinUserRegisterRequest{}
	var response = &proto.DouyinUserRegisterResponse{}
	if ctx.ShouldBind(request) == nil {
		log.Printf("Username %s\n", *request.Username)
		log.Printf("Password %s\n", *request.Password)
	}

	var token string = "123"
	var statusMsg string = "123"
	var statusCode int32 = 1
	var userId int64 = 1
	response.Token = &token
	response.StatusMsg = &statusMsg
	response.StatusCode = &statusCode
	response.UserId = &userId
	ctx.ProtoBuf(http.StatusOK, response)
}
