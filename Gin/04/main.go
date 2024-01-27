package main

import (
	"04/routers"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 加载模板 放在配置路由上面
	r.LoadHTMLGlob("templates/**/*")

	// 配置session中间件

	// 创建基于 cookie 的存储引擎， secret11111 参数是用于加密的密钥
	store := cookie.NewStore([]byte("secret11111"))
	// store是前面创建的存储引擎，可以替换为其他存储引擎
	r.Use(sessions.Sessions("mysession", store))

	// 路由分组
	routers.AdminRoutersInit(r)
	routers.ApiRoutersInit(r)
	routers.DefaultRoutersInit(r)

	r.Run()
	fmt.Println("你好")

}
