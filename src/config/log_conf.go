package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

var Syslog *logrus.Logger

func ApiLogToFile() gin.HandlerFunc {

	logFilePath := Yml.Log.ApiLogPath
	logFileName := Yml.Log.ApiLogFile

	//日志文件
	fileName := path.Join(logFilePath, logFileName)

	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	//实例化
	logger := logrus.New()
	//设置输出
	logger.Out = src
	//设置日志级别
	logger.SetLevel(logrus.InfoLevel)
	//设置日志格式
	logger.SetFormatter(&logrus.JSONFormatter{})

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		// 日志格式
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}

func SysLogToFile() {

	logFilePath := Yml.Log.SysLogPath
	logFileName := Yml.Log.SysLogFile
	//日志文件
	fileName := path.Join(logFilePath, logFileName)

	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	//实例化
	Syslog = logrus.New()
	//设置输出
	Syslog.Out = src
	//设置日志级别
	Syslog.SetLevel(logrus.InfoLevel)
	//设置日志格式
	Syslog.SetFormatter(&logrus.JSONFormatter{})
}

func SysLog(message string) *logrus.Entry {
	return Syslog.WithFields(logrus.Fields{})
}

func CatchErr(err interface{}) {
	if err != nil {
		SysLog("").Errorf("%s", err)
		return
	}
}

func CatchInfo(err interface{}) {
	if err != nil {
		SysLog("").Infof("%s", err)
		return
	}
}

func CatchWarn(err interface{}) {
	if err != nil {
		SysLog("").Warnf("%s", err)
		return
	}
}

func CatchFatal(err interface{}) {
	if err != nil {
		SysLog("").Fatalf("%s", err)
		return
	}
}
