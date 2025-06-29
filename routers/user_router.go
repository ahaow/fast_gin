package routers

import (
	"fast_gin/api"
	"fast_gin/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouter(g *gin.RouterGroup) {
	app := api.App.UserApi
	g.POST("users/register", app.RegisterView)
	g.POST("users/login", middleware.LimitMiddleware(10), app.LoginView)
	g.GET("users/list", middleware.LimitMiddleware(10), middleware.AuthMiddleware, app.UserListView)
}
