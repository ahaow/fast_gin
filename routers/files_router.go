package routers

import (
	"fast_gin/api"

	"github.com/gin-gonic/gin"
)

func FilesRouter(g *gin.RouterGroup) {
	app := api.App.FilesApi
	g.POST("files/upload", app.UploadView)
}
