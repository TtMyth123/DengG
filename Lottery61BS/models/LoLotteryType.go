package models

import (
	"github.com/TtMyth123/DengG/Lottery61BS/models/mconst"
	"github.com/astaxie/beego/orm"
	"time"
)

type LoLotteryType struct {
	LoBaseInfo
	CycleId string
}

func (a *LoLotteryType) TableName() string {
	return mconst.TableName_LoLotteryType
}

func (this *LoLotteryType) Read(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.Read(this)
	return e
}

func (this *LoLotteryType) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}

	this.CreatedAt = time.Now()
	this.UpdatedAt = this.CreatedAt
	id, e := o.Insert(this)
	this.Id = id
	return e

}

func (this *LoLotteryType) Update(o orm.Ormer, cols ...string) error {
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

func InitLoLotteryType(o orm.Ormer) {
	c, _ := o.QueryTable(mconst.TableName_LoLotteryType).Count()
	if c == 0 {
		arrData := make([]LoLotteryType, 0)
		arrData = append(arrData, LoLotteryType{LoBaseInfo: LoBaseInfo{Id: 1}})
		arrData = append(arrData, LoLotteryType{LoBaseInfo: LoBaseInfo{Id: 2}})
		arrData = append(arrData, LoLotteryType{LoBaseInfo: LoBaseInfo{Id: 3}})
		arrData = append(arrData, LoLotteryType{LoBaseInfo: LoBaseInfo{Id: 4}})
		arrData = append(arrData, LoLotteryType{LoBaseInfo: LoBaseInfo{Id: 5}})
		_, e := o.InsertMulti(len(arrData), &arrData)
		if e != nil {
			panic(e)
		}
	}

}
