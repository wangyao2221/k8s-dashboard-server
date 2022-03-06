package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"k8s-dashboard-server/pkg/shutdown"
	"net/http"
	"time"

	"k8s-dashboard-server/configs"
	"k8s-dashboard-server/internal/router"
	"k8s-dashboard-server/pkg/env"
	"k8s-dashboard-server/pkg/logger"
	"k8s-dashboard-server/pkg/timeutil"
)

func main() {
	// 初始化 access logger
	accessLogger, err := logger.NewJSONLogger(
		logger.WithDisableConsole(),
		logger.WithField("domain", fmt.Sprintf("%s[%s]", configs.ProjectName, env.Active().Value())),
		logger.WithTimeLayout(timeutil.CSTLayout),
		logger.WithFileP(configs.ProjectAccessLogFile),
	)
	if err != nil {
		panic(err)
	}

	// 进程结束前将日志缓存刷入文件
	defer func() {
		_ = accessLogger.Sync()
		//_ = cronLogger.Sync()
	}()

	// 初始化 HTTP 服务
	s, err := router.NewHTTPServer(accessLogger)
	if err != nil {
		panic(err)
	}

	// 上一步创建的Server在这里启动()
	server := &http.Server{
		Addr:    configs.ProjectPort,
		Handler: s.Mux,
	}
	go func() {
		// 没用gin.Engine的run方法，只是借用他里面的handler TODO 自己可以不用这个模式，直接用gin的框架
		// 为什么mux里全部都是使用的gin.Engine的东西却要包装自己包装一层
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			accessLogger.Fatal("http server startup err", zap.Error(err))
		}
	}()

	// 优雅关闭
	// 优雅关闭
	shutdown.NewHook().Close(
		// 关闭 http server
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				accessLogger.Error("server shutdown err", zap.Error(err))
			}
		},

		// 关闭 db
		func() {
			if s.Db != nil {
				if err := s.Db.DbWClose(); err != nil {
					accessLogger.Error("dbw close err", zap.Error(err))
				}

				if err := s.Db.DbRClose(); err != nil {
					accessLogger.Error("dbr close err", zap.Error(err))
				}
			}
		},

		// 关闭 cache
		func() {
			if s.Cache != nil {
				if err := s.Cache.Close(); err != nil {
					accessLogger.Error("cache close err", zap.Error(err))
				}
			}
		},

		// 关闭 cron Server
		//func() {
		//	if s.CronServer != nil {
		//		s.CronServer.Stop()
		//	}
		//},
	) // 优雅关闭
	shutdown.NewHook().Close(
		// 关闭 http server
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				accessLogger.Error("server shutdown err", zap.Error(err))
			}
		},

		// 关闭 db
		func() {
			if s.Db != nil {
				if err := s.Db.DbWClose(); err != nil {
					accessLogger.Error("dbw close err", zap.Error(err))
				}

				if err := s.Db.DbRClose(); err != nil {
					accessLogger.Error("dbr close err", zap.Error(err))
				}
			}
		},

		// 关闭 cache
		func() {
			if s.Cache != nil {
				if err := s.Cache.Close(); err != nil {
					accessLogger.Error("cache close err", zap.Error(err))
				}
			}
		},

		// 关闭 cron Server
		//func() {
		//	if s.CronServer != nil {
		//		s.CronServer.Stop()
		//	}
		//},
	)
}
