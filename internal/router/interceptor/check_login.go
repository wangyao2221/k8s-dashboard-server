package interceptor

import (
	"github.com/gin-gonic/gin"
)

func (i *interceptor) CheckLogin() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 补充登录逻辑
	}
}
