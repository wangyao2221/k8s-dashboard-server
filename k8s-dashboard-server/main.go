package main

import (
	"fmt"

	"k8s-dashboard-server/conifgs"
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

}
