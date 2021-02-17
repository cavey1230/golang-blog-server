package v1

import (
	"github.com/gin-gonic/gin"
	"goblog/model"
	"goblog/utils/errmsg"
	"net/http"
	"strconv"
)

// 查询用户是否存在
func UserExist(context *gin.Context) {

}

// 添加用户
func AddUser(context *gin.Context) {
	var data model.User
	_ = context.ShouldBindJSON(&data)
	code := model.CheckUser(data.Username)
	if code == errmsg.SUCCSE {
		model.CreateUser(&data)
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个用户

// 查询用户列表
func GetUsers(context *gin.Context) {
	var code int
	pageSize, _ := strconv.Atoi(context.Query("pageSize"))
	pageNum, _ := strconv.Atoi(context.Query("pageNum"))
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

}

// 删除用户
func DeleteUser(context *gin.Context) {

}
