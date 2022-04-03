package core

import (
	"errors"
	"fmt"
	"go.uber.org/multierr"
	"k8s-dashboard-server/internal/code"
	"net/http"
	"net/url"
	"runtime/debug"
	"time"

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

// RouterGroup 包装gin的RouterGroup
type RouterGroup interface {
	Group(string, ...HandlerFunc) RouterGroup
	IRoutes
}

var _ IRoutes = (*router)(nil)

// IRoutes 包装gin的IRoutes
// TODO 研究一下为什么要包装IRoutes，用起来用什么更方便的地方，注：HandlerFunc也是自定义的
type IRoutes interface {
	Any(string, ...HandlerFunc)
	GET(string, ...HandlerFunc)
	POST(string, ...HandlerFunc)
	DELETE(string, ...HandlerFunc)
	PATCH(string, ...HandlerFunc)
	PUT(string, ...HandlerFunc)
	OPTIONS(string, ...HandlerFunc)
	HEAD(string, ...HandlerFunc)
}

type router struct {
	group *gin.RouterGroup
}

func (r *router) Group(relativePath string, handlers ...HandlerFunc) RouterGroup {
	group := r.group.Group(relativePath, wrapHandlers(handlers...)...)
	return &router{group: group}
}

// 貌似一个请求可以有一连串的handler
func (r *router) Any(relativePath string, handlers ...HandlerFunc) {
	r.group.Any(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) GET(relativePath string, handlers ...HandlerFunc) {
	r.group.GET(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) POST(relativePath string, handlers ...HandlerFunc) {
	r.group.POST(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) DELETE(relativePath string, handlers ...HandlerFunc) {
	r.group.DELETE(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) PATCH(relativePath string, handlers ...HandlerFunc) {
	r.group.PATCH(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) PUT(relativePath string, handlers ...HandlerFunc) {
	r.group.PUT(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) OPTIONS(relativePath string, handlers ...HandlerFunc) {
	r.group.OPTIONS(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) HEAD(relativePath string, handlers ...HandlerFunc) {
	r.group.HEAD(relativePath, wrapHandlers(handlers...)...)
}

// 把自定义的HandlerFunc数组转换成gin.HandlerFunc数组
func wrapHandlers(handlers ...HandlerFunc) []gin.HandlerFunc {
	funcs := make([]gin.HandlerFunc, len(handlers))
	for i, handler := range handlers {
		handler := handler
		funcs[i] = func(c *gin.Context) {
			ctx := newContext(c)
			defer releaseContext(ctx)

			handler(ctx)
		}
	}

	return funcs
}

// 因为以面向接口的思维写代码，所以只有Mux是暴露出去的
// 而Mux是接口，没有属性，所以即使下面Engine是public的，外部也无法访问(外部只知道Mux接口，不知道mux实现)
// 所以要想使用GROUP,GET,POST,PUT,DELETE...就只能让Mux去封装一遍gin的这些方法
// 对外暴露接口的方式好处是，采用新的web框架(gin之外的)时其他代码不用修改，只需要重现一个Mux就可以
type Mux interface {
	http.Handler
	Group(relativePath string, handlers ...HandlerFunc) RouterGroup
}

type mux struct {
	engine *gin.Engine
}

func (m *mux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	m.engine.ServeHTTP(w, req)
}

func (m *mux) Group(relativePath string, handlers ...HandlerFunc) RouterGroup {
	return &router{
		group: m.engine.Group(relativePath, wrapHandlers(handlers...)...),
	}
}

func New(logger *zap.Logger, options ...Option) (Mux, error) {
	if logger == nil {
		return nil, errors.New("logger required")
	}

	// TODO 看一下gin.ReleaseMode的作用
	gin.SetMode(gin.ReleaseMode)
	mux := &mux{
		engine: gin.New(),
	}

	// TODO 控制台打印带颜色文件，有趣，玩一玩
	fmt.Println(color.Blue(_UI))

	opt := new(option)
	for _, f := range options {
		f(opt)
	}

	// TODO 弄清楚跨域这些参数的作用他
	if opt.enableCors {
		mux.engine.Use(cors.New(cors.Options{
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
	mux.engine.Use(func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("got panic", zap.String("panic", fmt.Sprintf("%+v", err)), zap.String("stack", string(debug.Stack())))
			}
		}()

		ctx.Next()
	})

	mux.engine.Use(func(ctx *gin.Context) {
		if ctx.Writer.Status() == http.StatusNotFound {
			return
		}

		ts := time.Now()

		context := newContext(ctx)
		defer releaseContext(context)

		context.init()
		context.setLogger(logger)

		defer func() {
			var (
				response        interface{}
				businessCode    int
				businessCodeMsg string
				abortErr        error
			)

			// region 发生 Panic 异常发送告警提醒
			if err := recover(); err != nil {
				stackInfo := string(debug.Stack())
				logger.Error("got panic", zap.String("panic", fmt.Sprintf("%+v", err)), zap.String("stack", stackInfo))
				context.AbortWithError(Error(
					http.StatusInternalServerError,
					10101,
					"服务器异常",
				))
			}
			// endregion

			// region 发生错误，进行返回
			if ctx.IsAborted() {
				for i := range ctx.Errors {
					multierr.AppendInto(&abortErr, ctx.Errors[i])
				}

				if err := context.abortError(); err != nil { // customer err
					multierr.AppendInto(&abortErr, err.StackError())
					businessCode = err.BusinessCode()
					businessCodeMsg = err.Message()
					response = &code.Failure{
						Code:    businessCode,
						Message: businessCodeMsg,
					}
					ctx.JSON(err.HTTPCode(), response)
				}
			}
			// endregion

			// region 正确返回
			payload := context.getPayload()
			if payload != nil {
				ctx.JSON(http.StatusOK, payload)
			}

			decodedURL, _ := url.QueryUnescape(ctx.Request.URL.RequestURI())

			success := !ctx.IsAborted() && (ctx.Writer.Status() == http.StatusOK)
			costSeconds := time.Since(ts).Seconds()

			logger.Info("trace-log",
				zap.Any("method", ctx.Request.Method),
				zap.Any("path", decodedURL),
				zap.Any("http_code", ctx.Writer.Status()),
				zap.Any("business_code", businessCode),
				zap.Any("success", success),
				zap.Any("cost_seconds", costSeconds),
				zap.Error(abortErr),
			)
			// endregion
		}()

		ctx.Next()
	})

	system := mux.engine.Group("/system")
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
