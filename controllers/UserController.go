package controllers

import (
	"jwtdemo/common"
	"jwtdemo/dto"
	"jwtdemo/models"
	"jwtdemo/response"
	"jwtdemo/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserForm struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
}

func Register(ctx *gin.Context) {

	var userForm UserForm
	if err := ctx.ShouldBind(&userForm); err != nil {
		response.Fail(ctx, gin.H{}, "获取参数错误！")
		return
	}

	//数据验证
	if len(userForm.Telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}

	if len(userForm.Password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		return
	}

	if len(userForm.Name) == 0 {
		userForm.Name = utils.RandomString(10)
	}

	//判断手机号是否存在
	if isTelephoneExist(common.DB, userForm.Telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号已经存在")
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(userForm.Password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
		return

	}
	//创建用户
	newUser := models.User{
		Name:      userForm.Name,
		Telephone: userForm.Telephone,
		Password:  string(hashPassword),
	}
	common.DB.Create(&newUser)

	//返回结果
	response.Success(ctx, nil, "注册成功")
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user models.User
	db.Where("telephone = ?", telephone).First(&user)

	return user.ID != 0
}

func Login(ctx *gin.Context) {
	//获取参数
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	//数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}

	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		return
	}

	//判断手机号是否存在
	var user models.User
	common.DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Fail(ctx, nil, "密码错误")
		return
	}

	//发送token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate err : %v", err)
		return
	}
	//返回结果
	response.Success(ctx, gin.H{"token": token}, "登录成功")
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	data := gin.H{
		"user": dto.ToUserDto(user.(models.User)),
	}
	response.Success(ctx, data, "获取成功")
}
