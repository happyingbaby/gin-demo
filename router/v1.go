package router

import (
	"jwtdemo/controllers"
	"jwtdemo/middleware"

	"github.com/gin-gonic/gin"
)

func V1CollectRouter(r *gin.Engine) *gin.Engine {

	v1 := r.Group("/api/v1")

	authRoute := v1.Group("auth")
	{
		authRoute.POST("/register", controllers.Register)
		authRoute.POST("/login", controllers.Login)
		authRoute.POST("/info", middleware.AuthMiddleware(), controllers.Info)
	}

	return r
}
