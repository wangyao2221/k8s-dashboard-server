package helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type md5Request struct {
	Str string `uri:"str" binding:"required"` // 需要加密的字符串
}

type md5Response struct {
	Md5Str string `json:"md5_str"` // MD5后的字符串
}

// Md5 加密
// @Summary 加密
// @Description 加密
// @Tags Helper
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param str path string true "需要加密的字符串"
// @Success 200 {object} md5Response
// @Failure 400 {object} code.Failure
// @Router /helper/md5/{str} [get]
func (h *Handler) Md5() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := new(md5Request)
		res := new(md5Response)

		if err := ctx.ShouldBindUri(req); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		res.Md5Str = h.helperService.Md5(req.Str)
		ctx.JSON(http.StatusOK, res)
	}
}
