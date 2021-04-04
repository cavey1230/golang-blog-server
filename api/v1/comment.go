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
func AddComment(context *gin.Context) {
	var comment model.Comment
	_ = context.ShouldBindJSON(&comment)
	code := model.CreateComment(&comment)
	if code == errmsg.ERROR {
		code = errmsg.ERROR_COMMENT_CREATE_ERROR
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    comment,
		"message": errmsg.GetErrMsg(code),
	})
}

// 模糊查询所有文章
func FindAllComment(context *gin.Context) {
	var code int
	type PageInFor struct {
		PageSize    string `form:"pageSize"`
		PageNum     string `form:"pageNum"`
		Content     string `form:"content"`
		ArticleId   string `form:"articleId"`
		CommenterId string `form:"commenterId"`
		Guestbook   string `form:"guestbook"`
	}
	type DataObj struct {
		Total int64           `json:"total"`
		Data  []model.CommentWithUser `json:"data"`
	}
	var pageInFor PageInFor
	_ = context.ShouldBind(&pageInFor)
	fmt.Println(pageInFor)
	pageSize, _ := strconv.Atoi(pageInFor.PageSize)
	pageNum, _ := strconv.Atoi(pageInFor.PageNum)
	comment, total := model.FindAllComment(pageSize, pageNum,
		pageInFor.Content, pageInFor.ArticleId,
		pageInFor.CommenterId, pageInFor.Guestbook,
	)
	if len(comment) == 0 {
		code = errmsg.ERROR_COMMENT_FIND_ERROR
	} else {
		code = errmsg.SUCCSE
	}
	context.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": DataObj{
			Total: total,
			Data:  comment,
		},
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除文章
func DeleteComment(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	code := model.DeleteComment(id)
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
