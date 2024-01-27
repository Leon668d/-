package main

import (
	"fmt"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

func initMiddleware_1(c *gin.Context) {
	fmt.Println("1.我是一个中间件")
	// 调用该请求的剩余处理程序（全部）
	c.Next()
	fmt.Println("2.我是一个中间件")
}

func initMiddleware_4(c *gin.Context) {
	fmt.Println("1.我是一个中间件")
	// 调用该请求的剩余处理程序（全部）
	c.Next()
	fmt.Println("2.我是一个中间件")
}

func initMiddleware_2(c *gin.Context) {
	fmt.Println("1.我是一个中间件")
	// 调用该请求的剩余处理程序（全部）
	c.Next()
	fmt.Println("2.我是一个中间件")
}

func initMiddleware_3(c *gin.Context) {
	fmt.Println("1.我是一个中间件")
	// 调用该请求的剩余处理程序（全部）
	c.Next()
	fmt.Println("2.我是一个中间件")
}

// 时间戳转化为日期
func UnixToTime(timestamp int) string {
	fmt.Println(timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

// gin模板渲染和模板语法

func main() {
	r := gin.Default()
	// 自定义模板函数 注意要吧函数放在加载模板前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
	})
	// 加载模板 放在配置路由上面
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/", initMiddleware_1, func(ctx *gin.Context) {
		fmt.Println("这是一个首页")
		ctx.String(200, "gin首页")
	})

	r.GET("/news", func(ctx *gin.Context) {
		ctx.String(200, "新闻页面")
	})

	r.Run()
}
