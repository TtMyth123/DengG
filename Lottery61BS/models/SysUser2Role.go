package models

import (
	"github.com/TtMyth123/DengG/Lottery61BS/models/mconst"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type SysUser2Role struct {
	BaseInfo
	SysUserId int64 `orm:"default(0);description(用户ID) "`
	SysRoleId int64 `orm:"default(0);description(角色ID) "`
}

func (a *SysUser2Role) TableName() string {
	return mconst.TableName_SysUser2Role
}

func InitSysUserRole(o orm.Ormer) {
	if o == nil {
		o = orm.NewOrm()
	}
	c, _ := o.QueryTable(mconst.TableName_SysUser2Role).Count()
	if c == 0 {
		arrData := make([]SysUser2Role, 0)
		arrData = append(arrData, SysUser2Role{BaseInfo: BaseInfo{Id: 4}, SysUserId: 1, SysRoleId: 1})
		arrData = append(arrData, SysUser2Role{BaseInfo: BaseInfo{Id: 5}, SysUserId: 1, SysRoleId: 2})
		_, e := o.InsertMulti(len(arrData), &arrData)
		if e != nil {
			logs.Error(e)
		}
	}
}
