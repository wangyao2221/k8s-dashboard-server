module k8s-dashboard-server

go 1.16

require (
	github.com/fsnotify/fsnotify v1.5.1
	github.com/gin-gonic/gin v1.7.7
	github.com/go-redis/redis/v7 v7.4.1
	github.com/pkg/errors v0.9.1
	github.com/spf13/viper v1.10.1
	go.uber.org/zap v1.21.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gorm.io/driver/mysql v1.2.3
	gorm.io/gorm v1.22.5
)
