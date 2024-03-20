package models

import (
	"github.com/astaxie/beego/orm"

	_ "github.com/mattn/go-sqlite3"
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
	initSQLiteDatabase()
}
