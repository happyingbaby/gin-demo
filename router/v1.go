package router

import (
	"jwtdemo/controllers"
	"jwtdemo/middleware"

	"github.com/gin-gonic/gin"
)

func V1CollectRouter(r *gin.Engine) *gin.Engine {

	authRoute := r.Group("/api/auth")
	{
		authRoute.POST("/register", controllers.Register)
		authRoute.POST("/login", controllers.Login)
		authRoute.POST("/info", middleware.AuthMiddleware(), controllers.Info)
	}

	//以下的接口，都使用Authorize()中间件身份验证
	v1 := r.Group("/api/v1")
	v1.Use(middleware.AuthMiddleware())

	//upload file

	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	r.MaxMultipartMemory = 100 << 20 // 100 MiB
	v1.POST("upload", controllers.UploadFile)

	//add edit del detail book
	bookRoute := v1.Group("book")
	{
		bookRoute.POST("list", controllers.AddBook)
	}

	return r
}
