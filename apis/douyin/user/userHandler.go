package user

import (
	"ByteTechTraining/proto"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func PostUserLogin(ctx *gin.Context) {

}

func PostUserRegister(ctx *gin.Context) {
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

func GetUser(ctx *gin.Context) {

}
