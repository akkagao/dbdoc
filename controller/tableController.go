package controller

import (
	"github.com/gin-gonic/gin"

	"dbdoc/errs"
	"dbdoc/service"
)

/*
获取table列表
*/
func Tables(c *gin.Context) {
	tables, err := service.TableService.Tables(c)
	if err != nil {
		responseError(c, errs.ErrorService)
		return
	}
	responseSuccessArgs(c, gin.H{"tables": tables})
}

/**
获取表定义详情
*/
func TableDetail(c *gin.Context) {
	tableName := c.DefaultQuery("table_name", "")
	if tableName == "" {
		responseError(c, errs.ErrorArgs)
		return
	}
	tableFields, err := service.TableService.TableDetail(c, tableName)
	if err != nil {
		responseError(c, errs.ErrorService)
		return
	}
	responseSuccessArgs(c, gin.H{"tableFields": tableFields})
}
