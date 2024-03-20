package models

import (
	"github.com/TtMyth123/DengG/Lottery61BS/models/mconst"
	"github.com/astaxie/beego/orm"
	"time"
)

type LoBetType struct {
	LoBaseInfo
	Title   string `orm:"size(200)"`
	SiteUrl string `orm:"size(400)"`
}

func (a *LoBetType) TableName() string {
	return mconst.TableName_LoBetType
}

func (this *LoBetType) Read(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.Read(this)
	return e
}
func (this *LoBetType) Delete(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	_, e := o.Delete(this)
	return e
}

func (this *LoBetType) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}

	this.CreatedAt = time.Now()
	this.UpdatedAt = this.CreatedAt
	id, e := o.Insert(this)
	this.Id = id
	return e
}

func (this *LoBetType) Update(o orm.Ormer, cols ...string) error {
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
func InitLoBetType(o orm.Ormer) {
	c, _ := o.QueryTable(mconst.TableName_LoBetType).Count()
	if c == 0 {
		arrData := make([]LoBetType, 0)
		arrData = append(arrData, LoBetType{LoBaseInfo: LoBaseInfo{Id: mconst.BetSiteType_Id_CS2}, Title: mconst.BetSiteType_CS2, SiteUrl: mconst.BetSiteType_Url_CS2})
		_, e := o.InsertMulti(len(arrData), &arrData)
		if e != nil {
			panic(e)
		}
	}
}
