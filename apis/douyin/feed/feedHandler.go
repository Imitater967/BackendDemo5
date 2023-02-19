package feed

import (
	"ByteTechTraining/daos"
	"ByteTechTraining/models"
	"ByteTechTraining/proto"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func GetFeed(context *gin.Context) {
	var request = proto.DouyinFeedRequest{}
	if context.ShouldBindQuery(&request) == nil {
	}
	var response = &proto.DouyinFeedResponse{}
	//var videoList []*proto.Video = make([]*proto.Video, 30)
	var videoList []*proto.Video = make([]*proto.Video, 0)
	var msg = "获取成功"
	var statusCode int32 = 1
	var nextTime int64 = 0
	response.StatusMsg = &msg
	response.StatusCode = &statusCode
	response.NextTime = &nextTime
	var timeStamp = time.Unix(0, 0)
	if request.LatestTime != nil {
		timeStamp = time.Unix(*request.LatestTime, 0)
	}

	var videoDaos, time, err = daos.QueryFeed(timeStamp)
	if err != nil {
		msg = err.Error()
		context.JSON(http.StatusOK, &response)
		return
	}
	nextTime = time.Unix()
	for _, dao := range videoDaos {
		var video = proto.Video{}
		videoList = append(videoList, &video)
		video.Title = &dao.Title
		video.Id = &dao.Id
		playUrl := "http://192.168.31.195:8080/file/?id=" + strconv.FormatInt(*video.Id, 10)
		video.PlayUrl = &playUrl
		authDao := daos.UserAuthDao{models.UserAuthModel{Id: dao.Uploader}}
		err = authDao.QueryById()
		if err != nil {
			log.Println(err.Error())
		}
		var user = proto.User{}
		user.Name = &authDao.Name
		user.Id = &authDao.Id
		isFollow := false
		var follow int64 = 0
		user.IsFollow = &isFollow
		user.FollowerCount = &follow
		user.FollowCount = &follow
		video.Author = &user
	}
	response.VideoList = videoList
	context.JSON(http.StatusOK, &response)
}
