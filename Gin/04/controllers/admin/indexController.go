package admin

import "github.com/gin-gonic/gin"

type IndexController struct{}

func (con IndexController) Index(ctx *gin.Context) {
	ctx.String(200, "用户列表")
}

func (con IndexController) Add(ctx *gin.Context) {
	ctx.String(200, "用户列表-add")
}

func (con IndexController) Edit(ctx *gin.Context) {
	ctx.String(200, "用户列表-edit")
}
