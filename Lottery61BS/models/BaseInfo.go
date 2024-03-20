package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type BaseInfo struct {
	Id        int64
	CreatedAt time.Time `orm:"auto_now_add;type(datetime); description(创建时间)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime); description(更新时间)"`
}

func (this *BaseInfo) Read(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.Read(this)
	return e
}

func (this *BaseInfo) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}

	this.CreatedAt = time.Now()
	this.UpdatedAt = this.CreatedAt
	id, e := o.Insert(this)
	this.Id = id
	return e

}

func (this *BaseInfo) Update(o orm.Ormer, cols ...string) error {
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

func (this *BaseInfo) AddUpdate(o orm.Ormer, cols ...string) error {
	if o == nil {
		o = orm.NewOrm()
	}
	//_, e := o.InsertOrUpdate(this, "id")
	data := &BaseInfo{}
	data.Id = this.Id
	e := o.Read(data)
	if e == nil {
		e = this.Update(o, cols...)
	} else {
		e = this.Add(o)
	}

	return e
}
