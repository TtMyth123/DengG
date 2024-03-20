package models

import "github.com/astaxie/beego/orm"

func initDataSys(o orm.Ormer) {
	InitSysInfo(o)
	iniSysUser(o)
	iniSysMenu(o)
}
func initRegisterModelSys() {
	orm.RegisterModel(new(SysInfo))
	orm.RegisterModel(new(SysLog))
	orm.RegisterModel(new(SysMenu))
	orm.RegisterModel(new(SysUser))
	orm.RegisterModel(new(SysUserMenuRel))
}
