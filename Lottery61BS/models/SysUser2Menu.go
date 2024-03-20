package models

import "github.com/TtMyth123/DengG/Lottery61BS/models/mconst"

type SysUser2Menu struct {
	BaseInfo
	SysUserId int64 `orm:"default(0);description(角色主键) "`
	SysMenuId int64 `orm:"default(0);description(菜单主键) "`
	Modify    int   `orm:"default(0);description(是否可能修改) "`
}

func (a *SysUser2Menu) TableName() string {
	return mconst.TableName_SysUser2Menu
}
