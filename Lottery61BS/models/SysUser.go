package models

import (
	"errors"
	"fmt"
	"github.com/TtMyth123/DengG/Lottery61BS/models/mconst"
	"github.com/TtMyth123/kit/pwdKit"
	"github.com/TtMyth123/kit/sqlKit"
	"github.com/TtMyth123/kit/ttLog"
	"github.com/astaxie/beego/orm"
	"time"
)

type SysUser struct {
	Id            int64
	Pid           int64
	PortraitIndex int
	PortraitUrl   string `orm:"size(1024)"`
	NickName      string `orm:"size(256)"`
	UserName      string `orm:"size(256)"`           //UserName
	Password      string `orm:"size(256)"; json:"-"` //pwdKit.Sha1ToStr("admin")="0DPiKuNIrrVmD8IUCuw1hQxNqZc="
	WechatId      string `orm:"size(256)"`
	Mobile        string `orm:"size(16)"`
	Email         string `orm:"size(256)"`
	IsSuper       bool
	Status        int
	CreatedAt     time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt     time.Time `orm:"auto_now;type(datetime)"`
	LoginIp       string    `orm:"size(20)"`
	LoginTime     time.Time `orm:"auto_now;type(datetime);null"`
	//GameId        int64
	//UserType      int
}

func (a *SysUser) TableName() string {
	return mconst.TableName_SysUser
}

// 根据用户名密码获取单条
func SysUserOneByUserName(username, userpwd string) (SysUser, error) {
	m := SysUser{}
	err := orm.NewOrm().QueryTable(mconst.TableName_SysUser).Filter("UserName", username).Filter("Password", userpwd).One(&m)
	if err != nil {
		return m, err
	}
	return m, nil
}
func GetSysUser(Id int64) (SysUser, error) {
	o := orm.NewOrm()
	aSysUser := SysUser{Id: Id}
	e := o.Read(&aSysUser)

	return aSysUser, e
}

func (user *SysUser) AddSysUser() error {
	o := orm.NewOrm()
	query := orm.NewOrm().QueryTable(mconst.TableName_SysUser)
	c, e := query.Filter("UserName", user.UserName).Count()

	if e != nil {
		return e
	}
	if c > 0 {
		return errors.New("用户名已存在")
	}
	user.Password = pwdKit.Sha1ToStr(user.Password)
	id, err := o.Insert(user)
	user.Id = id

	return err
}
func (user *SysUser) Update(o orm.Ormer, cols ...string) error {
	if o == nil {
		o = orm.NewOrm()
	}
	_, e := o.Update(user, cols...)
	return e
}

func iniSysUser(o orm.Ormer) error {
	c, _ := o.QueryTable(mconst.TableName_SysUser).Count()
	if c == 0 {
		//pwdKit.Sha1ToStr("admin")//"0DPiKuNIrrVmD8IUCuw1hQxNqZc="
		arrData := make([]SysUser, 0)
		//arrData = append(arrData, SysUser{Id: mconst.SysUserAdminId, Pid: mconst.SysUserRootId, UserName: "admin", Password: pwdKit.Sha1ToStr("admin"), IsSuper: false, Status: 1, GameId: mconst2.RootGameUserId, UserType: mconst.SysUserUserType_1_Agent})
		arrData = append(arrData, SysUser{Id: mconst.SysUserAdminId, Pid: mconst.SysUserRootId, UserName: "admin", NickName: "admin", Password: pwdKit.Sha1ToStr("admin"), IsSuper: false, Status: 1})
		_, e := o.InsertMulti(len(arrData), &arrData)
		return e
	}
	return nil
}

func GetSysUserListByPage(sysUserId int64, userName string, userId int64, pageIndex, pageSize int) (int, []SysUser) {
	PageTotal := 0
	o := orm.NewOrm()
	arrData := make([]SysUser, 0)
	sqlArgs := make([]interface{}, 0)

	sqlWhere := " where a.pid=? "
	sqlArgs = append(sqlArgs, sysUserId)
	if userName != "" {
		sqlWhere = sqlWhere + ` and locate(?,a.user_name)>0`
		sqlArgs = append(sqlArgs, userName)
	}

	if userId != 0 {
		sqlWhere += " and a.id=?"
		sqlArgs = append(sqlArgs, userId)
	}

	sqlCount := fmt.Sprintf(`select count(1) c from sys_user a %s`, sqlWhere)
	e := o.Raw(sqlCount, sqlArgs).QueryRow(&PageTotal)
	if e != nil {
		return PageTotal, arrData
	}

	offset, _ := sqlKit.GetOffset(PageTotal, pageSize, pageIndex)
	sqlWhere = sqlWhere + ` LIMIT ?,?`
	sqlArgs = append(sqlArgs, offset, pageSize)

	sql := fmt.Sprintf(`select a.* from  sys_user a %s `,
		sqlWhere)
	o.Raw(sql, sqlArgs).QueryRows(&arrData)

	return PageTotal, arrData
}

func (user *SysUser) SetMenu(arr []SysMenuR) {
	arrR := "0"
	arrW := "0"

	for _, v := range arr {
		if v.Modify == 1 {
			arrR = arrR + fmt.Sprintf(",%d", v.Id)
		} else if v.Modify == 2 {
			arrW = arrW + fmt.Sprintf(",%d", v.Id)
		}
	}
	o := orm.NewOrm()
	sql := `delete from sys_user_menu_rel where sys_user_id=?`
	_, e := o.Raw(sql, user.Id).Exec()
	if e != nil {
		ttLog.LogError(e)
	}

	sql = fmt.Sprintf(`insert into sys_user_menu_rel(sys_user_id,sys_menu_id, allow_modify, created_at)
select ?, a.id, ?, now() from sys_menu a where a.id in (%s)
`, arrR)
	_, e = o.Raw(sql, user.Id, 0).Exec()
	if e != nil {
		ttLog.LogError(e)
	}

	sql = fmt.Sprintf(`insert into sys_user_menu_rel(sys_user_id,sys_menu_id, allow_modify, created_at)
select ?, a.id, ?, now() from sys_menu a where a.id in (%s)
`, arrW)
	_, e = o.Raw(sql, user.Id, 1).Exec()
	if e != nil {
		ttLog.LogError(e)
	}
	return
}
