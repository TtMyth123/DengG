package models

import (
	"github.com/TtMyth123/DengG/Lottery61BS/models/mconst"
	"github.com/astaxie/beego/orm"
	"time"
)

type LoInputBetM struct {
	LoBaseInfo
	Num  int
	BetM int `orm:"size(20)"`
}

func (a *LoInputBetM) TableName() string {
	return mconst.TableName_LoInputBetM
}

func (this *LoInputBetM) Read(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.Read(this)
	return e
}
func (this *LoInputBetM) Delete(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	_, e := o.Delete(this)
	return e
}

func (this *LoInputBetM) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}

	this.CreatedAt = time.Now()
	this.UpdatedAt = this.CreatedAt
	id, e := o.Insert(this)
	this.Id = id
	return e
}

func (this *LoInputBetM) Update(o orm.Ormer, cols ...string) error {
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

func (this *LoInputBetM) AddUpdate(o orm.Ormer, cols ...string) error {
	if o == nil {
		o = orm.NewOrm()
	}
	aData := &LoInputBetM{}
	e := o.QueryTable(this.TableName()).Filter("SysId", this.SysId).Filter("Num", this.Num).
		One(aData)
	if e == nil {
		this.Id = aData.Id
		e = this.Update(o, cols...)
	} else {
		e = this.Add(o)
	}
	return e
}
