/**
* @Author: Jason
* @Date: 2020/9/29 14:26
* @File : log
* @Software: GoLand
**/

package dcore

import (
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

func init() {
	file, _ := os.OpenFile("run.log", os.O_CREATE|os.O_WRONLY, 0666)
	log.SetOutput(io.MultiWriter(file, os.Stdout))
	log.SetLevel(log.InfoLevel)
	// 设置日志格式
	// 设置输出样式，自带的只有两种模式 logrus.JSONFormatter{}和logrus.TextFormatter{}
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
}
