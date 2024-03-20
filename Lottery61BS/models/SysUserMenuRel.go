package models

import (
	"fmt"
	"github.com/TtMyth123/DengG/Lottery61BS/models/mconst"
	"github.com/TtMyth123/kit/ttLog"
	"github.com/astaxie/beego/orm"
	"time"
)

type SysUserMenuRel struct {
	Id        int64
	SysUserId int64
	SysMenuId int64
	Power     int       //权限:1:可查看，2：可修改
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

func (a *SysUserMenuRel) TableName() string {
	return mconst.TableName_SysUserMenuRel
}

func iniSysUserMenuRel(o orm.Ormer) error {
	c1, _ := o.QueryTable(mconst.TableName_SysUserMenuRel).Filter("SysUserId", mconst.SysUserAdminId).Count()
	c2, _ := o.QueryTable(mconst.TableName_SysMenu).Count()
	if c1 != c2 {
		sql := fmt.Sprintf(`delete from %s where sys_user_id=?`, mconst.TableName_SysUserMenuRel)
		o.Raw(sql, mconst.SysUserAdminId).Exec()

		arrMenu := make([]SysMenu, 0)
		o.QueryTable(mconst.TableName_SysMenu).All(&arrMenu)
		curTime := time.Now()
		arrData := make([]SysUserMenuRel, 0)
		for _, menu := range arrMenu {
			arrData = append(arrData, SysUserMenuRel{SysUserId: mconst.SysUserAdminId, SysMenuId: menu.Id, Power: mconst.Power_W, CreatedAt: curTime})
		}
		_, e := o.InsertMulti(len(arrData), &arrData)

		ttLog.LogDebug("添加菜单 完成 e:", e)
	}

	return nil
}

func GetSysUserMenuRel(userId int64) map[int]int {
	type Tmp struct {
		Id    int
		Pid   int
		Power int
	}

	arr := make([]Tmp, 0)
	o := orm.NewOrm()

	sql := fmt.Sprintf(`select b.id, b.pid, a.power from %s a, %s b where a.sys_menu_id =b.id and a.sys_user_id=?`,
		mconst.TableName_SysUserMenuRel, mconst.TableName_SysMenu)
	_, e := o.Raw(sql, userId).QueryRows(&arr)
	ttLog.LogWarning("添加菜单 完成 e:", e)
	mp := make(map[int]int)
	for _, v := range arr {
		//mp[v.Pid] = mconst.Power_R
		mp[v.Id] = v.Power
	}

	return mp
}
