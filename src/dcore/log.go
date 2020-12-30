/**
* @Author: Jason
* @Date: 2020/9/29 14:26
* @File : log
* @Software: GoLand
**/

package dcore

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

/**
初始化日志
*/
func init() {
	file, _ := os.OpenFile("run.log", os.O_CREATE|os.O_WRONLY, 0666)
	log.SetOutput(io.MultiWriter(file, os.Stdout))
	log.SetLevel(log.InfoLevel)
	// 设置日志格式
	// 设置输出样式，自带的只有两种模式 log.JSONFormatter{}和log.TextFormatter{}
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
}
