package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goblog/model"
	"goblog/utils/errmsg"
	"goblog/utils/myValidator"
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
	var code int
	var msg string
	_ = context.ShouldBindJSON(&user)
	msg, code = myValidator.Validator(&user)
	if code != errmsg.SUCCSE {
		context.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": msg,
		})
		return
	}
	code = model.CheckUser(user.Username)
	if code == errmsg.SUCCSE {
		code = model.CreateUser(&user)
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 模糊查询用户
func FindUser(context *gin.Context) {
	var code int
	type PageInFor struct {
		PageSize string `form:"pageSize"`
		PageNum  string `form:"pageNum"`
		Username string `form:"username"`
		Role     int    `form:"role"`
	}
	type DataObj struct {
		Total int64        `json:"total"`
		Data  []model.User `json:"data"`
	}
	var pageInFor PageInFor
	_ = context.ShouldBind(&pageInFor)
	pageSize, _ := strconv.Atoi(pageInFor.PageSize)
	pageNum, _ := strconv.Atoi(pageInFor.PageNum)
	fmt.Println(pageSize, pageNum)
	users, total := model.FindUsers(pageSize,
		pageNum, pageInFor.Username, pageInFor.Role)
	if len(users) == 0 {
		code = errmsg.ERROR_USER_NOT_EXIST
	} else {
		code = errmsg.SUCCSE
	}
	context.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": DataObj{
			Total: total,
			Data:  users,
		},
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询用户列表
func GetUsers(context *gin.Context) {
	var code int
	type PageInFor struct {
		PageSize string `form:"pageSize"`
		PageNum  string `form:"pageNum"`
	}
	type DataObj struct {
		Total int64        `json:"total"`
		Data  []model.User `json:"data"`
	}
	var pageInFor PageInFor
	_ = context.ShouldBind(&pageInFor)
	pageSize, _ := strconv.Atoi(pageInFor.PageSize)
	pageNum, _ := strconv.Atoi(pageInFor.PageNum)
	fmt.Println(pageSize, pageNum)
	users, total := model.GetUsers(pageSize, pageNum)
	if len(users) == 0 {
		code = errmsg.ERROR_USERS_PAGEINFO_ERROR
	} else {
		code = errmsg.SUCCSE
	}
	context.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": DataObj{
			Total: total,
			Data:  users,
		},
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
