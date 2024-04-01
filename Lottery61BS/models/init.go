package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//_ "github.com/mattn/go-sqlite3"
)

func Init() {
	InitRegisterModel()
	initDatabase()
	InitData()
}

func InitRegisterModel() {
	initRegisterModelSys()
	initRegisterModelBo()
}
func InitData() {
	o := orm.NewOrm()
	initDataSys(o)
	initDataBo(o)
}

func initDatabase() {
	dbType := beego.AppConfig.String("db_type")
	switch dbType {
	case "sqlite3":
		initSQLiteDatabase()
	case "mysql":
		initMysqlDatabase()
	default:
		panic("有问题的数据库类型")
	}
}
