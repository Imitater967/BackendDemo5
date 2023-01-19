package feed

import (
	"ByteTechTraining/proto"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetFeed(context *gin.Context) {
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
