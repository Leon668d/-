package admin

import "github.com/gin-gonic/gin"

type ArticleController struct{}

func (con ArticleController) Index(ctx *gin.Context) {
	ctx.String(200, "文章列表")
}

func (con ArticleController) Add(ctx *gin.Context) {
	ctx.String(200, "文章列表-add")
}

func (con ArticleController) Edit(ctx *gin.Context) {
	ctx.String(200, "文章列表-edit")
}
