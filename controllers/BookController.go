package controllers

import (
	"jwtdemo/response"

	"github.com/gin-gonic/gin"
)

func AddBook(ctx *gin.Context) {

	//获取参数
	// bookName := ctx.PostForm("book_name")
	// bookDesc := ctx.PostForm("book_desc")

	response.Success(ctx, gin.H{}, "获取成功")
}
