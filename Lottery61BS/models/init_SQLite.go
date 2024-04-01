package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/mattn/go-sqlite3"
)

// 初始化数据连接
func initSQLiteDatabase() {
	//dbType := beego.AppConfig.String("db_type")
	db_name := beego.AppConfig.String("sqlite3::db_name")
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	//orm.RegisterDataBase("default", "sqlite3", "./conf/test.db")
	orm.RegisterDataBase("default", "sqlite3", db_name)
	orm.RunSyncdb("default", false, true)
}
