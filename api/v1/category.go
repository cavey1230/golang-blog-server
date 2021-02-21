package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goblog/model"
	"goblog/utils/errmsg"
	"net/http"
	"strconv"
)

// 添加分类
func AddCategory(context *gin.Context) {
	var category model.Category
	_ = context.ShouldBindJSON(&category)
	fmt.Println(category.Name)
	code := model.CheckCategory(category.Name)
	if code == errmsg.SUCCSE {
		code = model.CreateCategory(&category)
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    category,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个分类

// 查询用户分类
func GetCategory(context *gin.Context) {
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
	category := model.GetCategory(pageSize, pageNum)
	if len(category) == 0 {
		code = errmsg.ERROR_CATEGORY_PAGEINFO_ERROR
	} else {
		code = errmsg.SUCCSE
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    category,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑分类
func EditCategory(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	fmt.Println(id)
	var category model.Category
	_ = context.ShouldBindJSON(&category)
	fmt.Println(category.Name)
	code := model.CheckCategory(category.Name)
	if code == errmsg.SUCCSE {
		code = model.EditCategory(id, &category)
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除分类
func DeleteCategory(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	code := model.DeleteCategory(id)
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
