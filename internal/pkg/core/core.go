package core

import (
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"go.uber.org/zap"

	"k8s-dashboard-server/pkg/color"
	"k8s-dashboard-server/pkg/env"
)

// see https://patorjk.com/software/taag/#p=testall&f=Graffiti&t=go-gin-api TODO
const _UI = `
██╗  ██╗ █████╗ ███████╗      ██████╗  █████╗ ███████╗██╗  ██╗██████╗  ██████╗  █████╗ ██████╗ ██████╗       ███████╗███████╗██████╗ ██╗   ██╗███████╗██████╗ 
██║ ██╔╝██╔══██╗██╔════╝      ██╔══██╗██╔══██╗██╔════╝██║  ██║██╔══██╗██╔═══██╗██╔══██╗██╔══██╗██╔══██╗      ██╔════╝██╔════╝██╔══██╗██║   ██║██╔════╝██╔══██╗
█████╔╝ ╚█████╔╝███████╗█████╗██║  ██║███████║███████╗███████║██████╔╝██║   ██║███████║██████╔╝██║  ██║█████╗███████╗█████╗  ██████╔╝██║   ██║█████╗  ██████╔╝
██╔═██╗ ██╔══██╗╚════██║╚════╝██║  ██║██╔══██║╚════██║██╔══██║██╔══██╗██║   ██║██╔══██║██╔══██╗██║  ██║╚════╝╚════██║██╔══╝  ██╔══██╗╚██╗ ██╔╝██╔══╝  ██╔══██╗
██║  ██╗╚█████╔╝███████║      ██████╔╝██║  ██║███████║██║  ██║██████╔╝╚██████╔╝██║  ██║██║  ██║██████╔╝      ███████║███████╗██║  ██║ ╚████╔╝ ███████╗██║  ██║
╚═╝  ╚═╝ ╚════╝ ╚══════╝      ╚═════╝ ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝╚═════╝  ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚═════╝       ╚══════╝╚══════╝╚═╝  ╚═╝  ╚═══╝  ╚══════╝╚═╝  ╚═╝
                                                                                                                                                              
`

type Option func(*option)

type option struct {
	enableCors bool
}

// WithEnableCors 设置支持跨域
func WithEnableCors() Option {
	return func(opt *option) {
		opt.enableCors = true
	}
}

// 因为以面向接口的思维写代码，所以只有Mux是暴露出去的
// 而Mux是接口，没有属性，所以即使下面Engine是public的，外部也无法访问(外部只知道Mux接口，不知道mux实现)
// 所以要想使用GROUP,GET,POST,PUT,DELETE...就只能让Mux去封装一遍gin的这些方法
// 对外暴露接口的方式好处是，采用新的web框架(gin之外的)时其他代码不用修改，只需要重现一个Mux就可以
type Mux interface {
	Run(addr ...string) (err error)
}

type mux struct {
	Engine *gin.Engine
}

func (m *mux) Run(addr ...string) (err error) {
	return m.Engine.Run(addr...)
}

func New(logger *zap.Logger, options ...Option) (Mux, error) {
	if logger == nil {
		return nil, errors.New("logger required")
	}

	// TODO 看一下gin.ReleaseMode的作用
	gin.SetMode(gin.ReleaseMode)
	mux := &mux{
		Engine: gin.New(),
	}

	// TODO 控制台打印带颜色文件，有趣，玩一玩
	fmt.Println(color.Blue(_UI))

	opt := new(option)
	for _, f := range options {
		f(opt)
	}

	// TODO 弄清楚跨域这些参数的作用他
	if opt.enableCors {
		mux.Engine.Use(cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{
				http.MethodHead,
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodPatch,
				http.MethodDelete,
			},
			AllowedHeaders:     []string{"*"},
			AllowCredentials:   true,
			OptionsPassthrough: true,
		}))
	}

	// TODO ??? recover两次，防止处理时发生panic，尤其是在OnPanicNotify中。
	mux.Engine.Use(func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("got panic", zap.String("panic", fmt.Sprintf("%+v", err)), zap.String("stack", string(debug.Stack())))
			}
		}()

		ctx.Next()
	})

	mux.Engine.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	mux.Engine.Use(ginzap.RecoveryWithZap(logger, true))

	system := mux.Engine.Group("/system")
	{
		// 健康检查
		system.GET("/health", func(ctx *gin.Context) {
			resp := &struct {
				Timestamp   time.Time `json:"timestamp"`
				Environment string    `json:"environment"`
				Host        string    `json:"host"`
				Status      string    `json:"status"`
			}{
				Timestamp:   time.Now(),
				Environment: env.Active().Value(),
				Host:        ctx.Request.Host,
				Status:      "ok",
			}
			ctx.JSON(http.StatusOK, resp)
		})
	}

	return mux, nil
}
