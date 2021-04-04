package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goblog/model"
	"goblog/utils/errmsg"
	"net/http"
	"strconv"
)

// 添加时间轴
func AddTimeline(context *gin.Context) {
	var timeline model.Timeline
	_ = context.ShouldBindJSON(&timeline)
	code := model.CreateTimeline(&timeline)
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    timeline,
		"message": errmsg.GetErrMsg(code),
	})
}

// 模糊查询时间轴
func FindAllTimeline(context *gin.Context) {
	var code int
	type PageInFor struct {
		PageSize string `form:"pageSize"`
		PageNum  string `form:"pageNum"`
		Title    string `form:"title"`
		Color    string `form:"color"`
		Com      string `form:"com"`
	}
	type DataObj struct {
		Total int64            `json:"total"`
		Data  []model.Timeline `json:"data"`
	}
	var pageInFor PageInFor
	_ = context.ShouldBind(&pageInFor)
	fmt.Println(pageInFor)
	pageSize, _ := strconv.Atoi(pageInFor.PageSize)
	pageNum, _ := strconv.Atoi(pageInFor.PageNum)
	timeline, total := model.FindAllTimeline(pageSize, pageNum,
		pageInFor.Title, pageInFor.Color, pageInFor.Com,
	)
	if len(timeline) == 0 {
		code = errmsg.ERROR_TIMELINE_FIND_ERROR
	} else {
		code = errmsg.SUCCSE
	}
	context.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": DataObj{
			Total: total,
			Data:  timeline,
		},
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑时间轴
func EditTimeline(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var timeline model.Timeline
	_ = context.ShouldBindJSON(&timeline)
	fmt.Printf("%v", timeline)
	code := model.EditTimeline(id, &timeline)
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除时间轴
func DeleteTimeline(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	code := model.DeleteTimeline(id)
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
