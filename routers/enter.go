package routers

import (
	"fast_gin/errors"
	"fast_gin/global"

	"github.com/gin-gonic/gin"
)

func Run() {
	port := global.Config.App.Port
	gin.SetMode("debug") // debug
	r := gin.Default()
	r.Static("/uploads", "uploads")

	r.Use(errors.ErrorHandler())

	g := r.Group("api")
	UserRouter(g)
	ImagesRouter(g)
	FilesRouter(g)

	if port == "" {
		port = ":3000"
	}
	r.Run(port)
}
