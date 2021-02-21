package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goblog/model"
	"goblog/utils/errmsg"
	"net/http"
	"strconv"
)

// 添加文章
func AddArticle(context *gin.Context) {
	var article model.Article
	_ = context.ShouldBindJSON(&article)
	code := model.CheckArticle(article.Title)
	if code == errmsg.SUCCSE {
		code = model.CreateArticle(&article)
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    article,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个文章
func GetOneArticle(context *gin.Context) {
	var article model.Article
	var code int
	id, _ := strconv.Atoi(context.Param("id"))
	if model.GetOneArticle(id) == errmsg.ERROR_ARTICLE_NOT_DEFINE {
		code = model.GetOneArticle(id).(int)
	} else {
		article = model.GetOneArticle(id).(model.Article)
		code = errmsg.SUCCSE
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    article,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询所有文章
func GetAllArticles(context *gin.Context) {
	var code int
	type InFor struct {
		PageSize string `form:"pageSize"`
		PageNum  string `form:"pageNum"`
		Cid      string `form:"cid"`
	}
	type DataObj struct {
		Total    int64           `json:"total"`
		Articles []model.Article `json:"articles"`
	}
	var inFor InFor
	_ = context.ShouldBind(&inFor)
	pageSize, _ := strconv.Atoi(inFor.PageSize)
	pageNum, _ := strconv.Atoi(inFor.PageNum)
	cid, _ := strconv.Atoi(inFor.Cid)
	//fmt.Println(pageSize,pageNum,cid)
	article, total := model.GetAllArticle(pageSize, pageNum, cid)
	if len(article) == 0 {
		code = errmsg.ERROR_CATEGORY_PAGEINFO_ERROR
	} else {
		code = errmsg.SUCCSE
	}
	context.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": DataObj{
			Total:    total,
			Articles: article,
		},
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑文章
func EditArticle(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var article model.Article
	_ = context.ShouldBindJSON(&article)
	fmt.Println(article.Title)
	code := model.CheckArticle(article.Title)
	if code == errmsg.SUCCSE {
		code = model.EditArticle(id, &article)
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除文章
func DeleteArticle(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	code := model.DeleteArticle(id)
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
