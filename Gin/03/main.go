package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string `json:"title" xml:"title"`
	Content string `json:"content" xml:"content"`
}

type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
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

	// // 配置静态web目录，第一个参数表示路由，第二个参数表示映射的目录
	// r.Static("/aaa", "templates/static")

	r.GET("/default/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "default/index.html", gin.H{
			"title": "首页222",
		})
	})

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "首页",
			"msg":   "我是msg",
			"score": 89,
			"date":  1706010308,
			"hobby": []string{"吃饭", "睡觉", "打豆豆"},
			"news": &Article{
				Title:   "新闻标题",
				Content: "新闻详情",
			},
		})
	})

	// GET 请求传值
	r.GET("/trans", func(ctx *gin.Context) {

		username := ctx.Query("username")
		age := ctx.Query("age")
		page := ctx.DefaultQuery("page", "1")
		ctx.JSON(http.StatusOK, gin.H{
			"username": username,
			"age":      age,
			"page":     page,
		})
	})

	//POST演示
	r.GET("/user", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "default/user.html", gin.H{})
	})

	// 传值绑定到结构体上
	r.GET("/getUser", func(ctx *gin.Context) {
		user := &UserInfo{}
		if err1 := ctx.ShouldBind(&user); err1 == nil {
			ctx.JSON(http.StatusOK, user)
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"err1": err1.Error(),
			})
		}
	})
	//获取表单post过来的数据
	r.POST("/doAddUser", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		//age := c.DefaultPostForm("age", "20")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			//	"age":      age,
		})
	})

	// 获取POST XML数据
	r.POST("/xml", func(c *gin.Context) {
		article := &Article{}
		xmlSlice, _ := c.GetRawData() // 获取 c.Request.Body 读取请求数据
		if err := xml.Unmarshal(xmlSlice, article); err == nil {
			c.JSON(http.StatusOK, article)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}
	})

	// 动态路由传值
	//list/123	list/456
	r.GET("/list/:cid", func(ctx *gin.Context) {
		cid := ctx.Param("cid")
		ctx.String(200, "%v", cid)
	})

	r.GET("/list/:cid")
	r.Run()
}
