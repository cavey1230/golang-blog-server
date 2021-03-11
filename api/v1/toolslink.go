package v1

import (
	"github.com/gin-gonic/gin"
	"goblog/model"
	"goblog/utils/errmsg"
	"net/http"
	"strconv"
)

func GetOneToolsLink(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	toolsLink, code := model.GetOneToolsLink(id)
	if code == errmsg.ERROR {
		code = errmsg.ERROR_TOOLSLINK_GET_ERROR
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    toolsLink,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询前十的工具链接地址
func GetAllToolsLink(context *gin.Context) {
	var toolsLinkList []model.ToolsLink
	var code int
	toolsLinkList, code = model.GetToolsLinkList()
	if code == errmsg.ERROR {
		code = errmsg.ERROR_TOOLSLINK_GET_ERROR
	} else {
		code = errmsg.SUCCSE
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    toolsLinkList,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑单条工具链接地址
func EditToolsLink(context *gin.Context) {
	var code int
	id, _ := strconv.Atoi(context.Param("id"))
	var toolsLinkList model.ToolsLink
	_ = context.ShouldBindJSON(&toolsLinkList)
	code = model.EditToolsLink(id, &toolsLinkList)
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 添加单条工具链接地址
func AddToolsLink(context *gin.Context) {
	var toolsLinkList model.ToolsLink
	var code int
	_ = context.ShouldBindJSON(&toolsLinkList)
	code = model.CreateToolsLink(&toolsLinkList)
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    toolsLinkList,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除个人信息
func DeleteToolsLink(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	code := model.DeleteToolsLink(id)
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
