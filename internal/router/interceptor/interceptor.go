package interceptor

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"k8s-dashboard-server/internal/repository/mysql"
	"k8s-dashboard-server/internal/repository/redis"
)

var _ Interceptor = (*interceptor)(nil)

type Interceptor interface {
	// CheckLogin 验证是否登录
	CheckLogin() gin.HandlerFunc

	// CheckRBAC 验证 RBAC 权限是否合法
	CheckRBAC() gin.HandlerFunc

	// i 为了避免被其他包实现
	i()
}

type interceptor struct {
	logger *zap.Logger
	cache  redis.Repo
	db     mysql.Repo
	//authorizedService authorized.Service
	//adminService      admin.Service
}

func New(logger *zap.Logger, cache redis.Repo, db mysql.Repo) Interceptor {
	return &interceptor{
		logger: logger,
		cache:  cache,
		db:     db,
		//authorizedService: authorized.New(db, cache),
		//adminService:      admin.New(db, cache),
	}
}

func (i *interceptor) i() {}
