package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	retatelogs "github.com/lestrrat-go/file-rotatelogs" //日志分割
	"github.com/rifflock/lfshook"                       //自定义钩子
	"github.com/sirupsen/logrus"                        //自定义日志
	"math"
	"os"
	"time"
)

func WriteLog() gin.HandlerFunc {
	filePath := "log/allLog.log"
	linkName := "latest_log.log"
	scr, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("err", err)
	}
	logger := logrus.New()
	logger.Out = scr
	logger.SetLevel(logrus.DebugLevel)

	logWriter, _ := retatelogs.New(
		"log/%Y%m%d.log", //文件名字
		retatelogs.WithMaxAge(7*24*time.Hour),     //最大生存时间
		retatelogs.WithRotationTime(24*time.Hour), //分割时间
		retatelogs.WithLinkName(linkName),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.AddHook(Hook)

	return func(context *gin.Context) {
		startTime := time.Now()
		context.Next()
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds()/1000000.0))))
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		statusCode := context.Writer.Status()
		clientIp := context.ClientIP()
		userAgent := context.Request.UserAgent()
		dataSize := context.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := context.Request.Method
		path := context.Request.RequestURI
		entry := logger.WithFields(logrus.Fields{
			"Hostname":  hostName,
			"Status":    statusCode,
			"SpendTime": spendTime,
			"Ip":        clientIp,
			"Method":    method,
			"Path":      path,
			"DatsSize":  dataSize,
			"UserAgent": userAgent,
		})
		if len(context.Errors) > 0 {
			entry.Error(context.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
