package models

import (
	"github.com/TtMyth123/DengG/Lottery61BS/models/mconst"
	"github.com/astaxie/beego/orm"
	"time"
)

type LoAgentBetNumData struct {
	LoBaseInfo
	AgentUserId int64 `orm:"description(站点用户ID)"` //站点用户ID
	Num         int   `orm:"default(0);description(号码)"`
	BetM        int   `orm:"default(0);description(投注金额)"`
	Stock       int   `orm:"default(0);description(占成类型.0全部,1实占)"`
	LotteryT    int64 `orm:"description(彩票类型)"`
}

func (a *LoAgentBetNumData) TableName() string {
	return mconst.TableName_LoAgentBetNumData
}

func (this *LoAgentBetNumData) Read(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.Read(this)
	return e
}
func (this *LoAgentBetNumData) Delete(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	_, e := o.Delete(this)
	return e
}

func (this *LoAgentBetNumData) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}

	this.CreatedAt = time.Now()
	this.UpdatedAt = this.CreatedAt
	id, e := o.Insert(this)
	this.Id = id
	return e
}

func (this *LoAgentBetNumData) Update(o orm.Ormer, cols ...string) error {
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

func (this *LoAgentBetNumData) AddUpdate(o orm.Ormer, cols ...string) error {
	if o == nil {
		o = orm.NewOrm()
	}
	//aData := &LoAgentBetNumData{}
	aData := make([]LoAgentBetNumData, 0)
	c, e := o.QueryTable(this).Filter("SysId", this.SysId).Filter("Num", this.Num).
		Filter("AgentUserId", this.AgentUserId).Filter("Stock", this.Stock).Filter("LotteryT", this.LotteryT).
		All(&aData)
	if e == nil {
		if c == 0 {
			e = this.Add(o)
		} else {
			this.Id = aData[0].Id
			e = this.Update(o, cols...)
		}
	}

	return e
}
