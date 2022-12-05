package controllers

import (
	"jwtdemo/response"

	"github.com/gin-gonic/gin"
)

type AddBookForm struct {
	BookName     string `json:"book_name"`
	TestamentKey string `json:"testament_key"`
	ShortName    string `json:"short_name"`
}

func AddBook(ctx *gin.Context) {
	var addBookForm AddBookForm
	//获取参数
	if err := ctx.ShouldBind(&addBookForm); err != nil {
		response.Fail(ctx, gin.H{}, "获取参数错误！")
		return
	}

	response.Success(ctx, gin.H{}, "获取成功")
}

func ListBook(ctx *gin.Context) {

}
