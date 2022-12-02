package router

import (
	"jwtdemo/controllers"
	"jwtdemo/middleware"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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
	uploadMemory := viper.GetString("upload.maxMemory")
	mermory, _ := strconv.ParseInt(uploadMemory, 10, 64)
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	r.MaxMultipartMemory = mermory << 20 // 100 MiB
	v1.POST("upload", controllers.UploadFile)

	//add edit del detail book
	bookRoute := v1.Group("book")
	{
		bookRoute.POST("list", controllers.AddBook)
	}

	return r
}
