package interceptor

import (
	"github.com/gin-gonic/gin"
)

func (i *interceptor) CheckRBAC() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 补充授权逻辑
	}
}
