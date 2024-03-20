package models

import (
	"github.com/TtMyth123/DengG/Lottery61BS/models/mconst"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type SysRole2Menu struct {
	BaseInfo
	SysRoleId int64 `orm:"default(0);description(角色主键) "`
	SysMenuId int64 `orm:"default(0);description(菜单主键) "`
	Modify    int   `orm:"default(0);description(是否可能修改) "`
}

func (a *SysRole2Menu) TableName() string {
	return mconst.TableName_SysRole2Menu
}

func InitSysRoleMenu(o orm.Ormer) {
	if o == nil {
		o = orm.NewOrm()
	}
	c, _ := o.QueryTable(mconst.TableName_SysRole2Menu).Count()
	if c == 0 {
		arrData := make([]SysRole2Menu, 0)
		arrData = append(arrData, SysRole2Menu{BaseInfo: BaseInfo{Id: 0}, SysRoleId: 1, SysMenuId: 1})
		arrData = append(arrData, SysRole2Menu{BaseInfo: BaseInfo{Id: 0}, SysRoleId: 1, SysMenuId: 2})
		arrData = append(arrData, SysRole2Menu{BaseInfo: BaseInfo{Id: 0}, SysRoleId: 1, SysMenuId: 3})
		arrData = append(arrData, SysRole2Menu{BaseInfo: BaseInfo{Id: 0}, SysRoleId: 1, SysMenuId: 4})
		arrData = append(arrData, SysRole2Menu{BaseInfo: BaseInfo{Id: 0}, SysRoleId: 1, SysMenuId: 5})
		arrData = append(arrData, SysRole2Menu{BaseInfo: BaseInfo{Id: 0}, SysRoleId: 1, SysMenuId: 6})
		arrData = append(arrData, SysRole2Menu{BaseInfo: BaseInfo{Id: 0}, SysRoleId: 1, SysMenuId: 7})

		_, e := o.InsertMulti(len(arrData), &arrData)
		if e != nil {
			logs.Error(e)
		}
	}
}
