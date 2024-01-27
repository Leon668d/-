package main

import "github.com/gin-gonic/gin"

func main() {
	//创建一个默认的路由引擎
	r := gin.Default()

	//配置路由
	r.GET("/", func(c *gin.Context) {
		c.String(200, "值：%v", "你好GIN")
	})
	r.GET("/news", func(c *gin.Context) {
		c.String(200, "我是山里灵活的狗")
	})

	r.POST("/add", func(ctx *gin.Context) {
		ctx.String(200, "这是一个POST请求,主要用于编辑数据")
	})

	r.PUT("/edit", func(c *gin.Context) {
		c.String(200, "这是一个PUT请求，主要用于编辑数据")
	})

	r.DELETE("/delete", func(c *gin.Context) {
		c.String(200, "这是一个delete请求，主要用于删除数据")
	})

	///r.Run() //启动一个web服务 默认在0.0.0.0:8080启动服务

	r.Run(":8000")

}
