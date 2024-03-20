package models

import (
	"github.com/TtMyth123/DengG/Lottery61BS/models/mconst"
	"github.com/astaxie/beego/orm"
	"time"
)

type LoAgentType struct {
	LoBaseInfo
	Title   string `orm:"size(200)"`
	SiteUrl string `orm:"size(400)"`
}

func (a *LoAgentType) TableName() string {
	return mconst.TableName_LoAgentType
}

func (this *LoAgentType) Read(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.Read(this)
	return e
}
func (this *LoAgentType) Delete(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	_, e := o.Delete(this)
	return e
}

func (this *LoAgentType) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}

	this.CreatedAt = time.Now()
	this.UpdatedAt = this.CreatedAt
	id, e := o.Insert(this)
	this.Id = id
	return e
}

func (this *LoAgentType) Update(o orm.Ormer, cols ...string) error {
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
func InitLoAgentType(o orm.Ormer) {
	c, _ := o.QueryTable(mconst.TableName_LoAgentType).Count()
	if c == 0 {
		arrData := make([]LoAgentType, 0)
		arrData = append(arrData, LoAgentType{LoBaseInfo: LoBaseInfo{Id: mconst.AgentSiteType_Id_CS2}, Title: mconst.AgentSiteType_CS2, SiteUrl: mconst.AgentSiteType_Url_CS2})
		_, e := o.InsertMulti(len(arrData), &arrData)
		if e != nil {
			panic(e)
		}
	}
}
