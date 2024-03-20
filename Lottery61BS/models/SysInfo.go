package models

import (
	"github.com/TtMyth123/DengG/Lottery61BS/models/mconst"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

const (
	SysInfo_S_Title  = "澳门财运亨通"
	SysInfo_Id_Title = 1
)

type SysInfo struct {
	BaseInfo
	StrData string `orm:"size(200);description(数据)"`
	StrDes  string `orm:"size(200);description(描述)"`
}

func (a *SysInfo) TableName() string {
	return mconst.TableName_SysInfo
}
func InitSysInfo(o orm.Ormer) {
	if o == nil {
		o = orm.NewOrm()
	}
	c, _ := o.QueryTable(mconst.TableName_SysInfo).Count()
	if c == 0 {
		arrData := make([]SysInfo, 0)
		arrData = append(arrData, SysInfo{BaseInfo: BaseInfo{Id: SysInfo_Id_Title}, StrData: SysInfo_S_Title})
		_, e := o.InsertMulti(len(arrData), &arrData)
		if e != nil {
			logs.Error(e)
		}
	}
}

func (this *SysInfo) Read(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.Read(this)
	return e
}

func (this *SysInfo) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}

	this.CreatedAt = time.Now()
	this.UpdatedAt = this.CreatedAt
	id, e := o.Insert(this)
	this.Id = id
	return e

}

func (this *SysInfo) Update(o orm.Ormer, cols ...string) error {
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

func (this *SysInfo) AddUpdate(o orm.Ormer, cols ...string) error {
	if o == nil {
		o = orm.NewOrm()
	}
	//_, e := o.InsertOrUpdate(this, "id")
	data := &SysInfo{}
	data.Id = this.Id
	e := o.Read(data)
	if e == nil {
		e = this.Update(o, cols...)
	} else {
		e = this.Add(o)
	}

	return e
}
