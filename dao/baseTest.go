package dao

import (
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"dbdoc/conf"
	"dbdoc/log"
	"dbdoc/tools"
)

func InitTest() *gin.Context {
	conf.InitConfig("../conf/conf-dev.yaml")
	initMysql()

	c := &gin.Context{}
	requestKey, _ := tools.GenUUID()
	c.Set(tools.RequestKey, requestKey)
	return c
}

func initMysql() {
	logs.Debug("Init Mysql Engine Now...")
	// 初始化Mysql
	DbEngin = initOneMysqlEngine()
	logs.Debug("Init Mysql Engine Done!!!")
}

func initOneMysqlEngine() *gorm.DB {
	db_promotion, err := gorm.Open("mysql", conf.Config.GetString("database.mysqlConn"))
	if err != nil {
		log.PrintlnError("startDb error:%v", err)
		panic("StartDB Error")
	}
	db_promotion.DB().SetConnMaxLifetime(time.Minute * 5)
	db_promotion.DB().SetMaxIdleConns(conf.Config.GetInt("database.mysqlMaxActive"))
	db_promotion.DB().SetMaxOpenConns(conf.Config.GetInt("database.mysqlMaxIdle"))
	db_promotion.LogMode(true)
	db_promotion.SingularTable(true)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "dr_" + defaultTableName
	}
	return db_promotion
}
