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
func AddReply(context *gin.Context) {
	var reply model.Reply
	_ = context.ShouldBindJSON(&reply)
	code := model.CreateReply(&reply)
	if code == errmsg.ERROR {
		code = errmsg.ERROR_REPLY_CREATE_ERROR
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    reply,
		"message": errmsg.GetErrMsg(code),
	})
}

// 模糊查询所有文章
func FindAllReply(context *gin.Context) {
	var code int
	type PageInFor struct {
		PageSize         string `form:"pageSize"`
		PageNum          string `form:"pageNum"`
		Content          string `form:"content"`
		ArticleId        string `form:"articleId"`
		ReplierId        string `form:"replierId"`
		ReplyToCommentId string `form:"replyToCommentId"`
		Guestbook        string `form:"guestbook"`
	}
	type DataObj struct {
		Total int64                 `json:"total"`
		Data  []model.ReplyWithUser `json:"data"`
	}
	var pageInFor PageInFor
	_ = context.ShouldBind(&pageInFor)
	fmt.Println(pageInFor)
	pageSize, _ := strconv.Atoi(pageInFor.PageSize)
	pageNum, _ := strconv.Atoi(pageInFor.PageNum)
	reply, total := model.FindAllReply(pageSize, pageNum,
		pageInFor.Content, pageInFor.ArticleId,
		pageInFor.ReplierId, pageInFor.ReplyToCommentId,
		pageInFor.Guestbook,
	)
	if len(reply) == 0 {
		code = errmsg.ERROR_REPLY_FIND_ERROR
	} else {
		code = errmsg.SUCCSE
	}
	context.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": DataObj{
			Total: total,
			Data:  reply,
		},
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除文章
func DeleteReply(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	code := model.DeleteReply(id)
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
