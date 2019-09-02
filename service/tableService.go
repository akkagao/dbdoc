package service

import (
	"github.com/gin-gonic/gin"

	"dbdoc/dao"
	"dbdoc/models"
)

type tableService struct {
}

/*
获取table列表
*/
func (s *tableService) Tables(c *gin.Context) ([]string, error) {
	return dao.TableDao.Tables(c)
}

func (s *tableService) TableDetail(c *gin.Context, tableName string) ([]models.TableField, error) {
	return dao.TableDao.TableDetail(c, tableName)
}
