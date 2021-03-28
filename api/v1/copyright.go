package v1

import (
	"github.com/gin-gonic/gin"
	"goblog/model"
	"goblog/utils/errmsg"
	"net/http"
	"strconv"
)

func GetOneCopyright(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	copyright, code := model.GetOneCopyright(id)
	if code == errmsg.ERROR {
		code = errmsg.ERROR_COPYRIGHT_GET_ERROR
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    copyright,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询前十的版权信息
func GetAllCopyright(context *gin.Context) {
	var copyright []model.Copyright
	var code int
	copyright, code = model.GetCopyright()
	if code == errmsg.ERROR {
		code = errmsg.ERROR_COPYRIGHT_GET_ERROR
	} else {
		code = errmsg.SUCCSE
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    copyright,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑单条版权信息
func EditCopyright(context *gin.Context) {
	var code int
	id, _ := strconv.Atoi(context.Param("id"))
	var copyright model.Copyright
	_ = context.ShouldBindJSON(&copyright)
	code = model.EditCopyright(id, &copyright)
	if code != errmsg.SUCCSE {
		code = errmsg.ERROR_COPYRIGHT_EDIT_ERROR
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 添加单条版权信息
func AddCopyright(context *gin.Context) {
	var copyright model.Copyright
	var code int
	_ = context.ShouldBindJSON(&copyright)
	code = model.CreateCopyright(&copyright)
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    copyright,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除版权信息
func DeleteCopyright(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	code := model.DeleteCopyright(id)
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
