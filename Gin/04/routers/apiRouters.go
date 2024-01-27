package routers

import "github.com/gin-gonic/gin"

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", func(ctx *gin.Context) {
			ctx.String(200, "我是一个api接口")
		})
		apiRouters.GET("/userlist", func(ctx *gin.Context) {
			ctx.String(200, "我是一个api接口-userlist")
		})
	}
}
