package v1

import (
	"github.com/gin-gonic/gin"
	"goblog/model"
	"goblog/utils/errmsg"
	"net/http"
	"strconv"
)

// 查询单个个人信息
func GetOneInfo(context *gin.Context) {
	var info model.Info
	var code int
	info, code = model.GetInfo()
	if code == errmsg.ERROR {
		code = errmsg.ERROR_INFO_GET_ERROR
	} else {
		code = errmsg.SUCCSE
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    info,
		"message": errmsg.GetErrMsg(code),
	})
}

// 模糊查询个人信息
func FindInfo(context *gin.Context) {
	var code int
	type PageInFor struct {
		PageSize   string `form:"pageSize"`
		PageNum    string `form:"pageNum"`
		FillString string `form:"fill_string"`
		Image      string `form:"image"`
		Name       string `form:"name"`
		Wechat     string `form:"wechat"`
		Address    string `form:"address"`
		Checked    string `form:"checked"`
	}
	type DataObj struct {
		Total int64        `json:"total"`
		Data  []model.Info `json:"data"`
	}
	var pageInFor PageInFor
	_ = context.ShouldBind(&pageInFor)
	pageSize, _ := strconv.Atoi(pageInFor.PageSize)
	pageNum, _ := strconv.Atoi(pageInFor.PageNum)
	infos, total := model.FindInfo(
		pageSize, pageNum,
		pageInFor.FillString, pageInFor.Image,
		pageInFor.Name, pageInFor.Wechat,
		pageInFor.Address, pageInFor.Checked,
	)
	if len(infos) == 0 {
		code = errmsg.ERROR_INFO_NO_INFO
	} else {
		code = errmsg.SUCCSE
	}
	context.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": DataObj{
			Total: total,
			Data:  infos,
		},
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询所有个人信息
func GetAllInfo(context *gin.Context) {
	var code int
	type InFor struct {
		PageSize string `form:"pageSize"`
		PageNum  string `form:"pageNum"`
	}
	type DataObj struct {
		Total int64        `json:"total"`
		Data  []model.Info `json:"data"`
	}
	var inFor InFor
	_ = context.ShouldBind(&inFor)
	pageSize, _ := strconv.Atoi(inFor.PageSize)
	pageNum, _ := strconv.Atoi(inFor.PageNum)
	//fmt.Println(pageSize,pageNum)
	infos, total := model.GetAllInfo(pageSize, pageNum)
	if len(infos) == 0 {
		code = errmsg.ERROR_INFO_NO_INFO
	} else {
		code = errmsg.SUCCSE
	}
	context.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": DataObj{
			Total: total,
			Data:  infos,
		},
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑个人信息
func EditInfo(context *gin.Context) {
	var code int
	id, _ := strconv.Atoi(context.Param("id"))
	var info model.Info
	_ = context.ShouldBindJSON(&info)
	code = model.EditInfo(id, &info)
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 添加个人信息
func AddInfo(context *gin.Context) {
	var info model.Info
	var code int
	_ = context.ShouldBindJSON(&info)
	code = model.CreateInfo(&info)
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    info,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除个人信息
func DeleteInfo(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	code := model.DeleteInfo(id)
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
