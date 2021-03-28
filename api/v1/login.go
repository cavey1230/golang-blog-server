package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goblog/middleware"
	"goblog/model"
	"goblog/utils/errmsg"
	"net/http"
)

//登录验证
func Login(context *gin.Context) {
	var user model.User
	var token string
	var code int
	_ = context.ShouldBindJSON(&user)
	fmt.Println(user.Username, user.Password)
	code = model.CheckLogin(user.Username, user.Password)
	if code == errmsg.SUCCSE {
		token, code = middleware.SetToken(user.Username)
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"token":   token,
		"message": errmsg.GetErrMsg(code),
	})
}

//后台登录验证
func AdminLogin(context *gin.Context) {
	var user model.User
	var token string
	var code int
	_ = context.ShouldBindJSON(&user)
	fmt.Println(user.Username, user.Password)
	code = model.CheckAdminLogin(user.Username, user.Password)
	if code == errmsg.SUCCSE {
		token, code = middleware.SetToken(user.Username)
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"token":   token,
		"message": errmsg.GetErrMsg(code),
	})
}
