package main

import (
	"jwtdemo/common"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	//读取配置文件
	InitConfig()

	//初始化数据库
	common.InitDB()

	//启动gin
	r := gin.Default()

	//进入路由
	r = CollectRouter(r)

	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}

	panic(r.Run())
}

func InitConfig() {
	path, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path + "/config")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

}
