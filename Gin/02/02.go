package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Gin返回值c.String() c.JSON() c.JSONP() c.XML() c.HTML()

func main() {
	r := gin.Default()
	// 配置模板的文件
	r.LoadHTMLGlob("../templates/*")

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "值:%v", "首页")

	})
	r.GET("/json1", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]interface{}{
			"success": true,
			"msg":     "你好",
		})

	})

	// gin.H <=> map[string] interface{}

	r.GET("/json2", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"success": true,
			"msg":     "你好",
		})

	})

	r.GET("/json3", func(ctx *gin.Context) {
		a := &Article{
			Title:   "我是标题",
			Desc:    "描述",
			Content: "我是山里灵活的狗",
		}
		ctx.JSON(200, a)

	})

	//响应Jsonp请求
	r.GET("/jsonp", func(ctx *gin.Context) {
		a := &Article{
			Title:   "我是标题-jsonp",
			Desc:    "描述",
			Content: "我是山里灵活的狗",
		}
		ctx.JSONP(200, a)

	})

	r.GET("/xml", func(ctx *gin.Context) {
		ctx.XML(http.StatusOK, gin.H{
			"success": true,
			"msg":     "你好gin 我是一个XML",
		})
	})

	r.GET("/news", func(ctx *gin.Context) {
		//注意r.LoadHTMLGlob("templates/*")-

		ctx.HTML(http.StatusOK, "news.html", gin.H{
			"title": "我是后台数据",
		})
	})

	r.Run()
}
