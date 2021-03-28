package v1

import (
	"github.com/gin-gonic/gin"
	"goblog/model"
	"goblog/utils/errmsg"
	"net/http"
	"strconv"
)

func GetOneBlogroll(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	blogroll, code := model.GetOneBlogroll(id)
	if code == errmsg.ERROR {
		code = errmsg.ERROR_BLOGROLL_GET_ERROR
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    blogroll,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询前十的工具链接地址
func GetAllBlogroll(context *gin.Context) {
	var blogroll []model.Blogroll
	var code int
	blogroll, code = model.GetBlogroll()
	if code == errmsg.ERROR {
		code = errmsg.ERROR_BLOGROLL_GET_ERROR
	} else {
		code = errmsg.SUCCSE
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    blogroll,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑单条工具链接地址
func EditBlogroll(context *gin.Context) {
	var code int
	id, _ := strconv.Atoi(context.Param("id"))
	var blogroll model.Blogroll
	_ = context.ShouldBindJSON(&blogroll)
	code = model.EditBlogroll(id, &blogroll)
	if code != errmsg.SUCCSE {
		code = errmsg.ERROR_BLOGROLL_EDIT_ERROR
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 添加单条工具链接地址
func AddBlogroll(context *gin.Context) {
	var blogroll model.Blogroll
	var code int
	_ = context.ShouldBindJSON(&blogroll)
	code = model.CreateBlogroll(&blogroll)
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    blogroll,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除个人信息
func DeleteBlogroll(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	code := model.DeleteBlogroll(id)
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
