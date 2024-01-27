package admin

import (
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController // 继承BaseController
}

func (con UserController) Index(ctx *gin.Context) {
	ctx.String(200, "用户列表")
	//	con.success(ctx)
}

func (con UserController) Add(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/useradd.html", gin.H{})
}

func (con UserController) Edit(ctx *gin.Context) {
	username := ctx.PostForm("username")
	file, err := ctx.FormFile("face")
	// file.Filename 获取文件名称
	dst := path.Join("./static/upload", file.Filename)
	if err == nil {
		// 获取文件后缀名
		extName := path.Ext(file.Filename)

		allowExtMap := map[string]bool{
			".jpg":  true,
			".png":  true,
			".gif":  true,
			".jpeg": true,
		}

		_, ok := allowExtMap[extName]
		if ok == true {
			ctx.String(200, "上传的文件不合法")
			return
		}
		// 创建图片保存目录

		ctx.SaveUploadedFile(file, dst)
	}
	ctx.JSON(200, gin.H{
		"success":  true,
		"username": username,
		"dst":      dst,
	})
	ctx.String(200, "执行上传")
}
