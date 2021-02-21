package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goblog/middleware"
	"goblog/model"
	"goblog/utils/errmsg"
	"net/http"
	"strconv"
)

// 查询用户是否存在
//func UserExist(context *gin.Context) {
//	username:=context.Query("username")
//	code := model.CheckUser(username)
//	context.JSON(http.StatusOK, gin.H{
//		"status":  code,
//		"message": errmsg.GetErrMsg(code),
//	})
//}

// 添加用户
func AddUser(context *gin.Context) {
	var user model.User
	_ = context.ShouldBindJSON(&user)
	code := model.CheckUser(user.Username)
	if code == errmsg.SUCCSE {
		code = model.CreateUser(&user)
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    user,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个用户

// 查询用户列表
func GetUsers(context *gin.Context) {
	var code int
	type PageInFor struct {
		PageSize string `form:"pageSize"`
		PageNum  string `form:"pageNum"`
	}
	var pageInFor PageInFor
	_ = context.ShouldBind(&pageInFor)
	pageSize, _ := strconv.Atoi(pageInFor.PageSize)
	pageNum, _ := strconv.Atoi(pageInFor.PageNum)
	fmt.Println(pageSize, pageNum)
	users := model.GetUsers(pageSize, pageNum)
	if len(users) == 0 {
		code = errmsg.ERROR_USERS_PAGEINFO_ERROR
	} else {
		code = errmsg.SUCCSE
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    users,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑用户
func EditUser(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var user model.User
	_ = context.ShouldBindJSON(&user)
	code := model.CheckUser(user.Username)
	if code == errmsg.SUCCSE {
		code = model.EditUser(id, &user)
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除用户
func DeleteUser(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	code := model.DeleteUser(id)
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

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
