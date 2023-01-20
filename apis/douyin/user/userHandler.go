package user

import (
	"ByteTechTraining/daos"
	"ByteTechTraining/proto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostUserLogin(ctx *gin.Context) {

}

// 注册时候,调用登录,进行token的生成与管理
func PostUserRegister(ctx *gin.Context) {
	var request = proto.DouyinUserRegisterRequest{}
	var response = proto.DouyinUserRegisterResponse{}
	var err = ctx.ShouldBind(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	var userAuth daos.UserAuthDao
	userAuth.Password = *request.Password
	userAuth.Name = *request.Username
	var statusMsg string = "注册成功"
	var statusCode int32 = 0
	var userId int64 = 1
	var addErr = userAuth.Add()
	if addErr != nil {
		statusCode = 1
		statusMsg = addErr.Error()
	}
	Login(&userAuth)
	response.Token = &userAuth.Token
	response.StatusMsg = &statusMsg
	response.StatusCode = &statusCode
	response.UserId = &userId
	ctx.ProtoBuf(http.StatusOK, response)
}

func GetUser(ctx *gin.Context) {

}

// 用户登录, 生成新的Token并刷新token时间
func Login(dao *daos.UserAuthDao) {

}
