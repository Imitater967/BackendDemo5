package middlewares

import "github.com/gin-gonic/gin"

// TokenVerify token校验中间件 负责校验user身份
func TokenVerify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 前置操作

		ctx.Next()

		// 后置操作
	}
}
