package log

import (
	"fmt"

	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"

	"dbdoc/tools"
)

func Error(c *gin.Context, format string, params ...interface{}) {
	requestId, _ := c.Get(tools.RequestKey)
	logs.Error("RequestId:%v %v", requestId, fmt.Sprintf(format, params...))
}

func Debug(c *gin.Context, format string, params ...interface{}) {
	requestId, _ := c.Get(tools.RequestKey)
	logs.Debug("RequestId:%v %v", requestId, fmt.Sprintf(format, params...))
}

func Info(c *gin.Context, format string, params ...interface{}) {
	requestId, _ := c.Get(tools.RequestKey)
	logs.Info("RequestId:%v %v", requestId, fmt.Sprintf(format, params...))
}

func PrintlnError(format string, params ...interface{}) {
	logs.Error(fmt.Sprintf(format, params...))
}

func Println(format string, params ...interface{}) {
	logs.Info(fmt.Sprintf(format, params...))
}
