package routers

import (
	"fast_gin/api"
	"fast_gin/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(g *gin.RouterGroup) {
	app := api.App.UserApi
	g.POST("users/login", middleware.LimitMiddleWare(1), app.LoginView)
	g.GET("users", middleware.LimitMiddleWare(10), middleware.AuthMiddleWare, app.UserListView)
}
