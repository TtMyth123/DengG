package models

import (
	"github.com/TtMyth123/DengG/Lottery61BS/models/mconst"
	"github.com/astaxie/beego/orm"
	"time"
)

type LoAgentUser struct {
	LoBaseInfo
	TypeId   int64
	Title    string `orm:"size(200)"`
	SiteUrl  string `orm:"size(400)"`
	UserName string `orm:"size(200)"`
	Pwd      string `orm:"size(200)"`
	IsLogin  int

	State   int    `orm:"column(state);default(1)"`
	LogInfo string `orm:"size(2000)"`
}

func (a *LoAgentUser) TableName() string {
	return mconst.TableName_LoAgentUser
}

func (this *LoAgentUser) Read(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.Read(this)
	return e
}
func (this *LoAgentUser) Delete(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	_, e := o.Delete(this)
	return e
}

func (this *LoAgentUser) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}

	this.CreatedAt = time.Now()
	this.UpdatedAt = this.CreatedAt
	id, e := o.Insert(this)
	this.Id = id
	return e
}

func (this *LoAgentUser) Update(o orm.Ormer, cols ...string) error {
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
	AgentId_Cs2       = 1
	AgentUserName_Cs2 = "ddff123"
	AgentPwd_Cs2      = "aaa222"
	AgentUrl_Cs2      = `https://my45jee96-cs2.cs111.vip/`
)

func InitLoAgentUser(o orm.Ormer) {
	c, _ := o.QueryTable(mconst.TableName_LoAgentUser).Count()
	if c == 0 {
		arrData := make([]LoAgentUser, 0)
		arrData = append(arrData, LoAgentUser{LoBaseInfo: LoBaseInfo{Id: AgentId_Cs2, SysId: mconst.SysUser_AdminId}, Title: "CS2", State: 1,
			SiteUrl: AgentUrl_Cs2,
			TypeId:  mconst.AgentSiteType_Id_CS2, UserName: AgentUserName_Cs2, Pwd: AgentPwd_Cs2})

		_, e := o.InsertMulti(len(arrData), &arrData)
		if e != nil {
			panic(e)
		}
	}
}
