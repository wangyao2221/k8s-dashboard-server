package configs

const (
	// MinGoVersion 最小 Go 版本
	MinGoVersion = 1.16

	// ProjectVersion 项目版本
	ProjectVersion = "v1.0.0"

	// ProjectName 项目名称
	ProjectName = "k8s_dashboard_server"

	// ProjectPort 项目端口
	ProjectPort = ":8080"

	// ProjectAccessLogFile 项目访问日志存放文件
	ProjectAccessLogFile = "./logs/" + ProjectName + "-access.log"
)
