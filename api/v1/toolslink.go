package v1

import (
	"fmt"
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

// 模糊查询工具链接地址
func FindToolsLink(context *gin.Context) {
	var code int
	type PageInFor struct {
		PageSize  string `form:"pageSize"`
		PageNum   string `form:"pageNum"`
		IconImg   string `form:"icon_img"`
		Link      string `form:"link"`
		Title     string `form:"title"`
		Introduce string `form:"introduce"`
	}
	type DataObj struct {
		Total int64             `json:"total"`
		Data  []model.ToolsLink `json:"data"`
	}
	var pageInFor PageInFor
	_ = context.ShouldBind(&pageInFor)
	pageSize, _ := strconv.Atoi(pageInFor.PageSize)
	pageNum, _ := strconv.Atoi(pageInFor.PageNum)
	fmt.Println(pageSize, pageNum)
	toolsLink, total := model.FindToolsLink(
		pageSize, pageNum,
		pageInFor.IconImg, pageInFor.Link,
		pageInFor.Title, pageInFor.Introduce,
	)
	if len(toolsLink) == 0 {
		code = errmsg.ERROR_TOOLSLINK_GET_ERROR
	} else {
		code = errmsg.SUCCSE
	}
	context.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": DataObj{
			Total: total,
			Data:  toolsLink,
		},
		"message": errmsg.GetErrMsg(code),
	})
}

// 分页查询工具链接地址
func GetAllToolsLink(context *gin.Context) {
	var code int
	type PageInFor struct {
		PageSize string `form:"pageSize"`
		PageNum  string `form:"pageNum"`
	}
	type DataObj struct {
		Total int64             `json:"total"`
		Data  []model.ToolsLink `json:"data"`
	}
	var pageInFor PageInFor
	_ = context.ShouldBind(&pageInFor)
	pageSize, _ := strconv.Atoi(pageInFor.PageSize)
	pageNum, _ := strconv.Atoi(pageInFor.PageNum)
	toolsLinkList, total := model.GetToolsLinkList(pageSize, pageNum)
	if len(toolsLinkList) == 0 {
		code = errmsg.ERROR_TOOLSLINK_GET_ERROR
	} else {
		code = errmsg.SUCCSE
	}
	context.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": DataObj{
			Total: total,
			Data:  toolsLinkList,
		},
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
