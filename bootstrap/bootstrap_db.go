package bootstrap

import (
	"time"

	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"dbdoc/conf"
	"dbdoc/dao"
	"dbdoc/log"
)

/*
* 启动数据库
 */
func StartDb() {
	db_promotion, err := gorm.Open("mysql", conf.Config.GetString("database.mysqlConn"))
	if err != nil {
		log.Println("startDb error:%v", err)
		panic("StartDB Error")
	}
	db_promotion.DB().SetConnMaxLifetime(time.Minute * 5)
	db_promotion.DB().SetMaxIdleConns(conf.Config.GetInt("database.mysqlMaxActive"))
	db_promotion.DB().SetMaxOpenConns(conf.Config.GetInt("database.mysqlMaxIdle"))
	db_promotion.LogMode(true)
	db_promotion.SingularTable(true)
	db_promotion.SetLogger(logs.GetLogger())

	dao.DbEngin = db_promotion
}
