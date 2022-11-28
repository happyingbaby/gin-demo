package router

import (
	"jwtdemo/controllers"
	"jwtdemo/middleware"

	"github.com/gin-gonic/gin"
)

func V1CollectRouter(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controllers.Register)
	r.POST("/api/auth/login", controllers.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controllers.Info)

	return r
}
