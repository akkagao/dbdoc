package route

import (
	"github.com/gin-gonic/gin"

	"dbdoc/controller"
)

/*
RouterInit 初始化路由
*/
func RouterInit(engine *gin.Engine) {
	// 静态文件目录
	engine.Static("/web", "./web")

	apiV1 := engine.Group("/api/v1")
	table := apiV1.Group("/table")
	{
		table.GET("/list", controller.Tables)
		table.GET("/detail", controller.TableDetail)
	}

}
