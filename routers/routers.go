package routers

import "github.com/gin-gonic/gin"

type Option func(engine *gin.Engine)

var options []Option

func IncludeOpts(opts ...Option) {
	options = append(options, opts...)
}

func InitRouter() *gin.Engine {
	// 创建一个没有任何中间件的路由
	r := gin.New()
	// 添加自定义的中间件
	//r.Use(middleware.LoggerToFile(), middleware.Cors(), gin.Recovery())
	for _, opt := range options {
		opt(r)
	}
	return r
}
