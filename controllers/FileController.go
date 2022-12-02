package controllers

import (
	"fmt"
	"jwtdemo/response"
	"jwtdemo/utils"
	"log"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func UploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		response.Fail(ctx, gin.H{}, "上传错误")
		return
	}
	log.Println(file.Filename)
	uploadPath := viper.GetString("uploadpath")

	fileExt := strings.ToLower(path.Ext(file.Filename))
	fileName := utils.MD5(fmt.Sprintf("%s%s", file.Filename, time.Now().String()))

	filePath := filepath.Join(utils.Mkdir(uploadPath), "/", fileName+fileExt)

	// 上传文件至指定的完整文件路径
	ctx.SaveUploadedFile(file, filePath)

	response.Success(ctx, gin.H{}, "上传成功")
}
