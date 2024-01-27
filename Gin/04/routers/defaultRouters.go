package routers

import "github.com/gin-gonic/gin"

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", func(ctx *gin.Context) {
			ctx.String(200, "我是一个前台接口")
		})
		defaultRouters.GET("/aaa", func(ctx *gin.Context) {
			ctx.String(200, "我是一个前台接口-aaa")
		})
	}
}
