package routers

import (
	"fast_gin/api"

	"github.com/gin-gonic/gin"
)

func ImagesRouter(g *gin.RouterGroup) {
	app := api.App.ImagesApi
	g.POST("images/upload", app.UploadView)
}
