package models

import (
	"fmt"
	"github.com/TtMyth123/DengG/Lottery61BS/models/mconst"
	"github.com/astaxie/beego/orm"
	"time"
)

type SysMenu struct {
	Id       int64
	Pid      int64
	Icon     string `orm:"size(256)"`
	Title    string `orm:"size(256)"`
	UrlFor   string `orm:"size(256)";json:"index"`
	MenuType int

	IconM   string `orm:"size(256)"`
	TitleM  string `orm:"size(256)"`
	UrlForM string `orm:"size(256)"`

	Seq       int
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
	Level     int       `orm:"-"` //第几级，从0开始
}

func (a *SysMenu) TableName() string {
	return mconst.TableName_SysMenu
}

func iniSysMenu(o orm.Ormer) error {
	c, _ := o.QueryTable(mconst.TableName_SysMenu).Count()
	if c == 0 {
		arrSysMenu := make([]SysMenu, 0)
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 100, Id: 100, Pid: 0, Icon: "home", IconM: "home", Title: "主页", UrlForM: "/", UrlFor: "/", MenuType: mconst.MenuType_Fun})

		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 1000, Id: 1000, Pid: 0, Icon: "manage_accounts", IconM: "", Title: "权限管理", UrlForM: "", UrlFor: "", MenuType: mconst.MenuType_Group})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 1001, Id: 1001, Pid: 1000, Icon: "people", IconM: "home", Title: "系统用户列表", UrlForM: "SysUserList", UrlFor: "SysUserList", MenuType: mconst.MenuType_Fun})

		_, e := o.InsertMulti(len(arrSysMenu), &arrSysMenu)
		if e != nil {
			panic(e)
		}
		iniSysUserMenuRel(o)

		return e
	}
	return nil
}

type GroupMenu struct {
	Id          int
	Title       string `json:"title"`
	Icon        string `json:"icon"`
	Seq         int
	UrlFor      string `json:"path";orm:"size(256)"`
	IconM       string
	UrlForM     string
	SubMenuList []*GroupMenu `json:"children"`
	MenuType    int
	Power       int `json:"power"`
}

func GetSysMenuList(userId int64, IsSuper bool) []*GroupMenu {
	arrGroupMenu := GetGroupMenu(0)
	if IsSuper {
		return arrGroupMenu
	}

	mpMenu := GetSysUserMenuRel(userId)
	newGroupMenu := GetUserMenuRelList(arrGroupMenu, mpMenu)
	return newGroupMenu
}

func GetUserMenuRelList(arrGroupMenu []*GroupMenu, mpMenu map[int]int) []*GroupMenu {
	arrGroupMenu1 := make([]*GroupMenu, 0)
	for i := 0; i < len(arrGroupMenu); i++ {
		if v, ok := mpMenu[arrGroupMenu[i].Id]; ok {
			if v > 0 {
				arrGroupMenu1 = append(arrGroupMenu1, arrGroupMenu[i])
			}
		}
	}
	for i := 0; i < len(arrGroupMenu1); i++ {
		arrGroupMenu1[i].SubMenuList = GetUserMenuRelList(arrGroupMenu1[i].SubMenuList, mpMenu)
		//arrGroupMenu1[i].MenuType = mconst.MenuType_Group
		if len(arrGroupMenu1[i].SubMenuList) == 0 {
			arrGroupMenu1[i].SubMenuList = nil
			//arrGroupMenu1[i].MenuType = mconst.MenuType_Fun
		}
	}

	arrGroupMenu2 := make([]*GroupMenu, 0)
	for i := 0; i < len(arrGroupMenu1); i++ {
		if arrGroupMenu1[i].MenuType == mconst.MenuType_Group {
			if len(arrGroupMenu1[i].SubMenuList) > 0 {
				arrGroupMenu2 = append(arrGroupMenu2, arrGroupMenu1[i])
			}
		} else {
			arrGroupMenu2 = append(arrGroupMenu2, arrGroupMenu1[i])
		}
	}

	return arrGroupMenu2
}

func GetGroupMenu(Pid int) []*GroupMenu {
	arrGroupMenu := make([]*GroupMenu, 0)
	sql := `SELECT * from sys_menu where pid = ? order by seq `
	o := orm.NewOrm()
	_, e := o.Raw(sql, Pid).QueryRows(&arrGroupMenu)

	if e != nil {
		panic(e)
	}
	for _, gMenu := range arrGroupMenu {
		gMenu.SubMenuList = GetGroupMenu(gMenu.Id)
		//gMenu.MenuType = mconst.MenuType_Group
		if len(gMenu.SubMenuList) == 0 {
			//gMenu.MenuType = mconst.MenuType_Fun
			gMenu.SubMenuList = nil
		}
		if gMenu.UrlFor == "" {
			gMenu.UrlFor = fmt.Sprintf("A%d", gMenu.Id)
		}
	}

	return arrGroupMenu
}
