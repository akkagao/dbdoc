package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"dbdoc/log"
	"dbdoc/tools"

	"dbdoc/errs"
)

/**
公共返回错误处理
*/
func responseError(c *gin.Context, err error) {
	log.Error(c, "response error:%v", err)
	if v, ok := err.(*errs.Error); ok {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": v.Code,
			"msg":  v.Msg,
		})
	} else {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": errs.ErrorService.Code,
			"msg":  errs.ErrorService.Msg,
		})
	}
}

/**
返回默认成功参数
*/
func responseSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": errs.Success.Code,
		"msg":  errs.Success.Msg,
	})
}

func responseSuccessArgs(c *gin.Context, m map[string]interface{}) {
	log.Info(c, "response value:%v", tools.Data2Str(m))
	m["code"] = errs.Success.Code
	m["msg"] = errs.Success.Msg
	c.JSON(http.StatusOK, m)
}
