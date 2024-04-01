package models

import (
	"github.com/TtMyth123/DengG/Lottery61BS/models/mconst"
	"github.com/astaxie/beego/orm"
	"time"
)

type LoBetBetData struct {
	LoBaseInfo
	CycleId  string
	UserId   int64
	Num      int
	BetM     int
	LotteryT int64
}

func (a *LoBetBetData) TableName() string {
	return mconst.TableName_LoBetBetData
}

func (this *LoBetBetData) Read(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.Read(this)
	return e
}

func (this *LoBetBetData) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}

	this.CreatedAt = time.Now()
	this.UpdatedAt = this.CreatedAt
	id, e := o.Insert(this)
	this.Id = id
	return e

}

func (this *LoBetBetData) Update(o orm.Ormer, cols ...string) error {
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

func (this *LoBetBetData) AddUpdate(o orm.Ormer, cols ...string) error {
	if o == nil {
		o = orm.NewOrm()
	}
	//_, e := o.InsertOrUpdate(this, "id")
	aData := make([]LoBetBetData, 0)
	c, e := o.QueryTable(this).Filter("SysId", this.SysId).Filter("Num", this.Num).
		Filter("UserId", this.UserId).Filter("LotteryT", this.LotteryT).Filter("CycleId", this.CycleId).
		All(&aData)
	if e == nil {
		if c == 0 {
			e = this.Add(o)
		} else {
			this.Id = aData[0].Id
			this.BetM = this.BetM + aData[0].BetM
			e = this.Update(o, cols...)
		}
	}

	return e
}
