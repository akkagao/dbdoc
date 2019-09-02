package route

import (
	"strings"

	"github.com/gin-gonic/gin"

	"dbdoc/tools"
)

func Middleware(engine *gin.Engine) {
	// 增加中间件
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(setRequestId)
}

// 设置接口的RequestId
func setRequestId(c *gin.Context) {
	requestId, _ := tools.GenUUID()
	c.Set(tools.RequestKey, strings.Replace(requestId, "-", "", -1))
}
