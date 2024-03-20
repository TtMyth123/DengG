package SysBll

import (
	"errors"
	"fmt"
	"github.com/TtMyth123/DengG/Lottery61BS/CacheData"
	"github.com/TtMyth123/DengG/Lottery61BS/controllers/base/cBo"
	"github.com/TtMyth123/DengG/Lottery61BS/models"
	"github.com/TtMyth123/DengG/Lottery61BS/models/mconst"
	"github.com/TtMyth123/kit/sqlKit"
	"github.com/astaxie/beego/orm"
	"time"

	"github.com/TtMyth123/kit/pwdKit"
	"github.com/TtMyth123/kit/ttLog"
)

func GetSysMenuList(userId int64) []*models.GroupMenu {
	IsSuper := false
	data := make([]*models.GroupMenu, 0)
	key := fmt.Sprintf("GetMenuList:%d", userId)
	e := CacheData.GetICache().GetCache(key, &data)
	if e != nil {
		data = models.GetSysMenuList(userId, IsSuper)
		e := CacheData.GetICache().SetCache(key, data, 10)

		if e != nil {
			ttLog.LogDebug(e)
		}
	}

	return data
}
func GetAllMenuList(userId int64) ([]*models.GroupMenu, error) {
	arrGroupMenu := models.GetGroupMenu(0)
	mpMenu := models.GetSysUserMenuRel(userId)
	newG, e := setGroupMenuPower(arrGroupMenu, mpMenu)

	return newG, e
}
func setGroupMenuPower(arrGroupMenu []*models.GroupMenu, mpMenu map[int]int) ([]*models.GroupMenu, error) {
	for i := 0; i < len(arrGroupMenu); i++ {
		arrGroupMenu[i].Power = mpMenu[arrGroupMenu[i].Id]
		setGroupMenuPower(arrGroupMenu[i].SubMenuList, mpMenu)
	}

	return arrGroupMenu, nil
}

type MenuPower struct {
	Id    int64
	Power int
}

func SaveMenuPower(userId int64, arrMenuPower []MenuPower) (interface{}, error) {
	if userId == mconst.SysUserAdminId {
		return "", fmt.Errorf("不能修改超级管理员的权限")
	}
	o := orm.NewOrm()
	sql := fmt.Sprintf(`delete from %s where sys_user_id=?`, mconst.TableName_SysUserMenuRel)
	_, e := o.Raw(sql, userId).Exec()
	if e != nil {
		return "", e
	}

	arrData := make([]models.SysUserMenuRel, 0)
	for _, power := range arrMenuPower {
		aSysUserMenuRel := models.SysUserMenuRel{
			SysUserId: userId,
			SysMenuId: power.Id,
			Power:     power.Power,
		}

		arrData = append(arrData, aSysUserMenuRel)
	}

	_, e = o.InsertMulti(len(arrData), &arrData)

	return "", e
}

func GetSysUserList(aParam cBo.BaseListParam) (cBo.BaseListResult, error) {
	type SysUser struct {
		Id  int64
		Pid int64

		NickName  string `orm:"size(256)"`
		UserName  string `orm:"size(256)"`
		IsSuper   bool
		Status    int
		LoginTime time.Time
	}
	type GroupInfo struct {
		C     int
		MaxId int
	}
	aResult := cBo.BaseListResult{}
	aGroup := GroupInfo{}

	PageTotal := 0
	o := orm.NewOrm()
	arrData := make([]SysUser, 0)
	sqlArgs := make([]interface{}, 0)

	maxId := aParam.MaxId
	if maxId == 0 {
		sqlCount := fmt.Sprintf(`select max(a.id) as maxid from %s a `, mconst.TableName_SysUser)
		e := o.Raw(sqlCount, sqlArgs).QueryRow(&maxId)
		if e != nil {
			ttLog.LogError(e)
		}
	}

	SysUserId := aParam.SysUserId
	sqlWhere := " where (a.pid=? or a.id = ?) "
	sqlArgs = append(sqlArgs, SysUserId, SysUserId)

	UserId := aParam.GetInt("UserId")
	if UserId != 0 {
		sqlWhere += ` and a.id=? `
		sqlArgs = append(sqlArgs, UserId)
	}

	userName := aParam.GetString("UserName")
	if userName != "" {
		sqlWhere = sqlWhere + ` and locate(?,a.user_name)>0`
		sqlArgs = append(sqlArgs, userName)
	}

	sqlCount := fmt.Sprintf(`select count(1) c from sys_user a %s`, sqlWhere)
	e := o.Raw(sqlCount, sqlArgs).QueryRow(&aGroup)
	if e != nil {
		return aResult, e
	}

	offset, PageTotal := sqlKit.GetOffset(aGroup.C, aParam.PageSize, aParam.PageIndex)
	sqlWhere = sqlWhere + ` LIMIT ?,?`
	sqlArgs = append(sqlArgs, offset, aParam.PageSize)

	sql := fmt.Sprintf(`select a.* from  sys_user a %s `,
		sqlWhere)
	_, e = o.Raw(sql, sqlArgs).QueryRows(&arrData)
	if e != nil {
		return aResult, e
	}

	aResult.DataList = arrData
	aResult.C = aGroup.C
	aResult.MaxId = maxId
	aResult.PageTotal = PageTotal
	aResult.GroupInfo = aGroup
	return aResult, e
}

func UpdateSysUser(aSysUser models.SysUser, cols ...string) error {
	o := orm.NewOrm()
	sqlArgs := make([]interface{}, 0)

	if aSysUser.Id == 1 && aSysUser.UserName != "admin" {
		return errors.New(fmt.Sprintf("不允许修改管理员用户名"))
	}
	sqlArgs = append(sqlArgs, aSysUser.UserName, aSysUser.Id)
	sql := fmt.Sprintf(`select count(1) c from sys_user a where a.user_name=? and a.id<>?`)
	c := 0
	e := o.Raw(sql, sqlArgs).QueryRow(&c)
	if e != nil {
		return e
	}
	if c > 0 {
		return fmt.Errorf("用户名已存在")
	}

	hasPwd := false
	for _, col := range cols {
		if col == "Password" {
			hasPwd = true
			break
		}
	}
	if hasPwd {
		aSysUser.Password = pwdKit.Sha1ToStr(aSysUser.Password)
	}

	e = aSysUser.Update(nil, cols...)
	return e
}
func AddSysUser(aSysUser models.SysUser) error {
	e := aSysUser.AddSysUser()
	if e == nil {
		//aBoSysUserGameUserRel := models.BoSysUserGameUserRel{
		//	Id:     aSysUser.Id,
		//	GameId: mconst2.RootGameUserId,
		//}
		//o := orm.NewOrm()
		//_, e = o.Insert(&aBoSysUserGameUserRel)
	}
	return e
}

func DelSysUser(id int64) error {
	o := orm.NewOrm()
	if id == mconst.SysUserAdminId {
		return errors.New("不能删除管理员用户")
	}
	aSysUser := models.SysUser{Id: id}
	c, e := o.Delete(&aSysUser)
	if c == 0 {
		return fmt.Errorf("没有对应用用户")
	}

	return e
}

type FunMenuPermission struct {
	Id       int64
	TypeName string
	Title    string
	Power    int
}

func GetFunMenuPermission(userId int) []FunMenuPermission {
	arr := make([]FunMenuPermission, 0)
	o := orm.NewOrm()
	sql := `select a.id, b.title as type_name, a.title, IFNULL(e.power,0) as power from sys_menu a
LEFT JOIN sys_menu b on (a.pid = b.id)
LEFT JOIN 
(select c.sys_menu_id, c.power from sys_user_menu_rel c, sys_user d where d.id = c.sys_user_id and d.id=?) e on (e.sys_menu_id=a.id)
where a.menu_type = ? order by b.seq
`
	_, e := o.Raw(sql, userId, mconst.MenuType_Fun).QueryRows(&arr)
	if e != nil {
		ttLog.LogError(e)
	}

	return arr
}

func SaveFunMenuPermission(PUserId, userId int64, arr []FunMenuPermission) error {
	o := orm.NewOrm()
	if len(arr) == 0 {
		return errors.New("数据为空无法保存。")
	}
	sql := fmt.Sprintf(`delete from %s where sys_user_id=?`, mconst.TableName_SysUserMenuRel)
	_, e := o.Raw(sql, userId).Exec()
	if e != nil {
		return e
	}

	sqlP := fmt.Sprintf(`select * from %s where sys_user_id=? `, mconst.TableName_SysUserMenuRel)
	arrPData := make([]models.SysUserMenuRel, 0)
	o.Raw(sqlP, PUserId).QueryRows(&arrPData)
	mpMenu := make(map[int64]models.SysUserMenuRel)
	for _, r := range arrPData {
		mpMenu[r.SysMenuId] = r
	}

	curTime := time.Now()

	arrData := make([]models.SysUserMenuRel, 0)
	for _, menu := range arr {
		Power := menu.Power
		if mpMenu[menu.Id].Power < menu.Power {
			Power = mpMenu[menu.Id].Power
		}
		arrData = append(arrData, models.SysUserMenuRel{SysUserId: userId, SysMenuId: menu.Id, Power: Power, CreatedAt: curTime})
	}
	_, e = o.InsertMulti(len(arrData), &arrData)

	return e
}
func ModifyPwd(userId int64, OldPwd, NewPwd string) error {
	sysU, e := models.GetSysUser(userId)
	if e != nil {
		return e
	}
	aNewPwd := pwdKit.Sha1ToStr(OldPwd)
	if sysU.Password != aNewPwd {
		return errors.New("原密码不正确.")
	}

	sysU.Password = pwdKit.Sha1ToStr(NewPwd)
	e = sysU.Update(nil, "Password")
	return e
}

func ModifySysUserInfo(userId int64, FieldName, FieldValue string) error {
	sysU, e := models.GetSysUser(userId)
	if e != nil {
		return e
	}

	switch FieldName {
	case "NickName":
		sysU.NickName = FieldValue
		e = sysU.Update(nil, "NickName")
		return e
	case "Mobile":
		sysU.Mobile = FieldValue
		e = sysU.Update(nil, "Mobile")
		return e
	case "Email":
		sysU.Email = FieldValue
		e = sysU.Update(nil, "Email")
		return e
	}

	return errors.New("有问题的参数")
}

func GetSysUserInfo(userId int64) (interface{}, error) {
	sysU, e := models.GetSysUser(userId)

	return sysU, e
}
