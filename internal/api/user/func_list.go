package user

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	"time"

	"k8s-dashboard-server/internal/services/user"
)

type listRequest struct {
	Page     int    `json:"page"`      // 第几页
	PageSize int    `json:"page_size"` // 每页显示条数
	Username string `json:"username"`  // 用户名
}

type listData struct {
	Id        int32     `json:"id"`         // 主键
	Username  string    `json:"username"`   // 用户名
	Nickname  string    `json:"nickname"`   // 昵称
	IsDeleted int32     `json:"is_deleted"` // 是否删除 1:是  -1:否
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
}

type listResponse struct {
	List       []listData `json:"list"`
	Pagination struct {
		Total        int `json:"total"`
		CurrentPage  int `json:"current_page"`
		PerPageCount int `json:"per_page_count"`
	} `json:"pagination"`
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
// @Router /user/md5/{str} [get]
func (h *Handler) List() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := new(listRequest)
		res := new(listResponse)

		if err := ctx.ShouldBindJSON(req); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		page := req.Page
		if page == 0 {
			page = 1
		}

		pageSize := req.PageSize
		if pageSize == 0 {
			pageSize = 10
		}

		searchData := new(user.SearchData)
		searchData.Page = page
		searchData.PageSize = pageSize
		searchData.Username = req.Username

		resListData, err := h.userService.List(ctx, searchData)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		//resCountData, err := h.userService.PageListCount(ctx, searchData)
		//if err != nil {
		//	ctx.AbortWithError(http.StatusBadRequest, err)
		//	return
		//}
		resCountData := len(resListData)
		res.Pagination.Total = cast.ToInt(resCountData)
		res.Pagination.PerPageCount = pageSize
		res.Pagination.CurrentPage = page
		res.List = make([]listData, len(resListData))

		for k, v := range resListData {
			data := listData{
				Id:        v.Id,
				Username:  v.Username,
				Nickname:  v.Nickname,
				IsDeleted: v.IsDeleted,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			}

			res.List[k] = data
		}

		ctx.JSON(http.StatusOK, res)
	}
}
