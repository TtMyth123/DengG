package models

import (
	"github.com/TtMyth123/DengG/Lottery61BS/models/mconst"
	"github.com/astaxie/beego/orm"
	"time"
)

type LoConfig struct {
	LoBaseInfo
	IsAutoGetData int //自动获得数据
	RateWay       int //比例方向 .0:反向,1正面
	BetWay        int //1:比例,2:金额
	Stock         int //占成方式:0:全部,1:实占
	Rate          int //比例：100为1
	BetTime1      string
	BetTime2      string
	LotteryType   int64 //彩种.1:澳门六合彩,2:新澳门六合彩,3:超级6+1六合彩,4:闪电6+1六合彩
}

func (a *LoConfig) TableName() string {
	return mconst.TableName_LoConfig
}

func (this *LoConfig) Read(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.Read(this)
	return e
}
func (this *LoConfig) Delete(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	_, e := o.Delete(this)
	return e
}

func (this *LoConfig) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}

	this.CreatedAt = time.Now()
	this.UpdatedAt = this.CreatedAt
	id, e := o.Insert(this)
	this.Id = id
	return e
}

func (this *LoConfig) Update(o orm.Ormer, cols ...string) error {
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

func InitLoConfig(o orm.Ormer) {
	c, _ := o.QueryTable(mconst.TableName_LoConfig).Count()
	if c == 0 {
		arrData := make([]LoConfig, 0)
		arrData = append(arrData, LoConfig{LoBaseInfo: LoBaseInfo{Id: 1},
			LotteryType: 1, Rate: 10, BetWay: 1, Stock: 0,
			IsAutoGetData: 0, BetTime1: "09:15", BetTime2: "09:20"})
		_, e := o.InsertMulti(len(arrData), &arrData)
		if e != nil {
			panic(e)
		}
	}
}
