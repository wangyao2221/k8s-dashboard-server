package user

import (
	"github.com/gin-gonic/gin"

	"k8s-dashboard-server/internal/repository/mysql"
	"k8s-dashboard-server/internal/repository/mysql/user"
)

type SearchData struct {
	Page     int `json:"page"`      // 第几页
	PageSize int `json:"page_size"` // 每页显示条数
}

func (s *Service) List(c *gin.Context, searchData *SearchData) (listData []*user.User, err error) {
	qb := user.NewQueryBuilder()
	qb = qb.WhereIsDeleted(mysql.EqualPredicate, -1)

	listData, err = qb.
		OrderById(false).
		QueryAll(s.db.GetDbR().WithContext(c))
	//QueryAll(s.db.GetDbR().WithContext(c))

	if err != nil {
		return nil, err
	}

	return
}
