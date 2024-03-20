package models

import (
	"github.com/TtMyth123/DengG/Lottery61BS/models/mconst"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

type SysRole struct {
	BaseInfo
	Name        string `orm:"size(100);description(角色名称)" `
	AliasName   string `orm:"size(50);description(别名)" `
	Description string `orm:"size(500);description(描述)" `
	Status      int    `orm:"column(status);default(1);description(角色状态) "`
	Type        int    `orm:"column(type);default(1);description(属于哪个应用) "`
}

func (a *SysRole) TableName() string {
	return mconst.TableName_SysRole
}

func InitSysRole(o orm.Ormer) {
	if o == nil {
		o = orm.NewOrm()
	}
	c, _ := o.QueryTable(mconst.TableName_SysRole).Count()
	if c == 0 {
		arrData := make([]SysRole, 0)
		arrData = append(arrData, SysRole{BaseInfo: BaseInfo{Id: 1}, Status: 1, Type: 1, Name: `admin`, AliasName: `admin`, Description: `超级管理员具有所有权限`})
		arrData = append(arrData, SysRole{BaseInfo: BaseInfo{Id: 2}, Status: 1, Type: 1, Name: `editor`, AliasName: `editor`, Description: `运营者`})
		arrData = append(arrData, SysRole{BaseInfo: BaseInfo{Id: 3}, Status: 1, Type: 1, Name: `normal`, AliasName: `normal`, Description: `普通管理员`})

		_, e := o.InsertMulti(len(arrData), &arrData)
		if e != nil {
			logs.Error(e)
		}
	}
}

func (this *SysRole) Read(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.Read(this)
	return e
}

func (this *SysRole) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}

	this.CreatedAt = time.Now()
	this.UpdatedAt = this.CreatedAt
	id, e := o.Insert(this)
	this.Id = id
	return e

}

func (this *SysRole) Update(o orm.Ormer, cols ...string) error {
	if o == nil {
		o = orm.NewOrm()
	}
	this.UpdatedAt = time.Now()
	if cols != nil {
		cols = append(cols, "UpdatedAt")
	}

	_, e := o.Update(this, cols...)
	return e
}

func (this *SysRole) AddUpdate(o orm.Ormer, cols ...string) error {
	if o == nil {
		o = orm.NewOrm()
	}
	//_, e := o.InsertOrUpdate(this, "id")
	data := &SysRole{}
	data.Id = this.Id
	e := o.Read(data)
	if e == nil {
		e = this.Update(o, cols...)
	} else {
		e = this.Add(o)
	}

	return e
}
