package v1

import (
	"github.com/gin-gonic/gin"
	"goblog/model"
	"goblog/utils/errmsg"
	"net/http"
)

func UpLoad(context *gin.Context) {
	file, fileHeader, _ := context.Request.FormFile("file")
	fileSize := fileHeader.Size
	url, code := model.UploadFile(file, fileSize)
	context.JSON(http.StatusOK, gin.H{
		"status":  code,
		"url":     url,
		"message": errmsg.GetErrMsg(code),
	})
}
