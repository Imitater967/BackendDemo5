package user

import (
	"ByteTechTraining/daos"
	"ByteTechTraining/proto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostUserLogin(ctx *gin.Context) {
	var request = proto.DouyinUserLoginRequest{}
	var response = proto.DouyinUserLoginResponse{}

	var err = ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	var code int32 = 1
	var msg = "登录成功"
	response.StatusCode = &code
	response.StatusMsg = &msg

	var authDao = daos.UserAuthDao{}
	authDao.Name = request.GetUsername()
	authDao.Password = request.GetPassword()
	authDao.Login()
}

// 注册时候,调用登录,进行token的生成与管理
func PostUserRegister(ctx *gin.Context) {
	var request = proto.DouyinUserRegisterRequest{}
	var response = proto.DouyinUserRegisterResponse{}
	var err = ctx.ShouldBindQuery(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	var userAuth daos.UserAuthDao
	userAuth.Name = request.GetUsername()
	userAuth.Password = request.GetPassword()
	var statusMsg string = "注册成功"
	var statusCode int32 = 0
	var userId int64 = 1
	var addErr = userAuth.Add()
	if addErr != nil {
		statusCode = 1
		statusMsg = addErr.Error()
	} else {
		userAuth.Login()
	}
	response.Token = &userAuth.Token
	response.StatusMsg = &statusMsg
	response.StatusCode = &statusCode
	response.UserId = &userId
	ctx.JSON(http.StatusOK, &response)
}

func GetUser(ctx *gin.Context) {

}
