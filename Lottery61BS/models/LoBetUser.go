package models

import (
	"github.com/TtMyth123/DengG/Lottery61BS/models/mconst"
	"github.com/astaxie/beego/orm"
	"time"
)

type LoBetUser struct {
	LoBaseInfo
	TypeId   int64
	Title    string `orm:"size(200)"`
	SiteUrl  string `orm:"size(400)"`
	UserName string `orm:"size(200)"`
	Pwd      string `orm:"size(200)"`

	State   int    `orm:"column(state);default(1)"`
	LogInfo string `orm:"size(2000)"`
	IsLogin int
}

func (a *LoBetUser) TableName() string {
	return mconst.TableName_LoBetUser
}

func (this *LoBetUser) Read(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.Read(this)
	return e
}
func (this *LoBetUser) Delete(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	_, e := o.Delete(this)
	return e
}

func (this *LoBetUser) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}

	this.CreatedAt = time.Now()
	this.UpdatedAt = this.CreatedAt
	id, e := o.Insert(this)
	this.Id = id
	return e
}

func (this *LoBetUser) Update(o orm.Ormer, cols ...string) error {
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

const (
	Bet_Id_Cs2       = 1
	Bet_UserName_Cs2 = "dkd333"
	Bet_Pwd_Cs2      = "aaa111"
	Bet_Url_Cs2      = `https://wy4996m42-cs2.cs555.vip/`
)

func InitLoBetUser(o orm.Ormer) {
	c, _ := o.QueryTable(mconst.TableName_LoBetUser).Count()
	if c == 0 {
		arrData := make([]LoBetUser, 0)
		arrData = append(arrData, LoBetUser{LoBaseInfo: LoBaseInfo{SysId: mconst.SysUser_AdminId}, Title: "CS2", State: 1,
			SiteUrl: Bet_Url_Cs2,
			TypeId:  mconst.BetSiteType_Id_CS2, UserName: Bet_UserName_Cs2, Pwd: Bet_Pwd_Cs2})

		_, e := o.InsertMulti(len(arrData), &arrData)
		if e != nil {
			panic(e)
		}
	}
}
