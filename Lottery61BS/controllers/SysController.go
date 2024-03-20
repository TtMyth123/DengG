package controllers

import (
	"fmt"
	"github.com/TtMyth123/DengG/Lottery61BS/controllers/SysBll"
	"github.com/TtMyth123/DengG/Lottery61BS/controllers/base"
	"github.com/TtMyth123/DengG/Lottery61BS/controllers/base/enums"
	"github.com/TtMyth123/DengG/Lottery61BS/models"
	"github.com/TtMyth123/kit/stringKit"
)

type SysController struct {
	base.AuthorBaseController
}

func (c *SysController) GetMenuList() {
	data := SysBll.GetSysMenuList(c.SUser.Id)
	c.JsonResult(enums.JRCodeSucc, "", data)
}

func (c *SysController) GetAllMenuList() {
	Id, _ := c.GetInt64("Id")
	data, e := SysBll.GetAllMenuList(Id)
	c.JsonResultEx(data, e)
}
func (c *SysController) SaveMenuPower() {
	Id, _ := c.GetInt64("Id")
	strPower := c.GetString("StrPower")
	fmt.Println("Id", Id, "strPower:", strPower)

	arr := make([]SysBll.MenuPower, 0)
	e := stringKit.GetJsonObj(strPower, &arr)

	data, e := SysBll.SaveMenuPower(Id, arr)
	c.JsonResultEx(data, e)
}

func (c *SysController) GetSysUserList() {
	mpParam := make(map[string]int)
	mpParam["UserName"] = 1
	mpParam["UserId"] = 0
	mpParam["BeginDay"] = 1
	mpParam["EndDay"] = 1
	mpParam["State"] = 0

	aParam := c.GetBaseListParam(mpParam)
	data, e := SysBll.GetSysUserList(aParam)
	c.JsonResultEx(data, e)
}

func (c *SysController) UpdateSysUser() {
	Id, _ := c.GetInt64("Id")

	UserName := c.GetString("UserName")
	NickName := c.GetString("NickName")
	aSysUser, e := models.GetSysUser(Id)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	aSysUser.UserName = UserName
	aSysUser.NickName = NickName

	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "", "")
}

func (c *SysController) UpdateSysUserPwd() {
	Id, _ := c.GetInt64("Id")
	Password := c.GetString("Password")
	aSysUser, e := models.GetSysUser(Id)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	aSysUser.Password = Password
	e = SysBll.UpdateSysUser(aSysUser, "Password")

	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "", "")
}

func (c *SysController) AddSysUser() {
	UserName := c.GetString("UserName")
	NickName := c.GetString("NickName")
	Password := c.GetString("Password")
	if UserName == "" {
		c.JsonResult(enums.JRCodeFailed, "用户名不能为空", "")
	}
	if Password == "" {
		c.JsonResult(enums.JRCodeFailed, "密码不能为空", "")
	}

	aSysUser := models.SysUser{UserName: UserName, NickName: NickName, Password: Password, Pid: c.SUser.Id}
	e := SysBll.AddSysUser(aSysUser)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "", "")
}

func (c *SysController) DelSysUser() {
	id, _ := c.GetInt64("Id", 0)
	e := SysBll.DelSysUser(id)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "", "")
}

func (c *SysController) GetPermissionTree() {
	UserId, _ := c.GetInt("UserId")
	permissionTree := SysBll.GetFunMenuPermission(UserId)
	c.JsonResult(enums.JRCodeSucc, "", permissionTree)
}

func (c *SysController) SavePermissionTree() {
	type TmpA struct {
		UserId int64
		Data   []SysBll.FunMenuPermission
	}
	aTmp := TmpA{}
	e := c.GetJsonData(&aTmp)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}

	permissionTree := SysBll.SaveFunMenuPermission(c.SUser.Id, aTmp.UserId, aTmp.Data)
	c.JsonResult(enums.JRCodeSucc, "", permissionTree)
}

func (c *SysController) ModifyPwd() {
	OldPwd := c.GetString("OldPwd")
	NewPwd := c.GetString("NewPwd")
	e := SysBll.ModifyPwd(c.SUser.Id, OldPwd, NewPwd)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "", "")
}
func (c *SysController) ModifySysUserInfo() {
	FieldName := c.GetString("FieldName")
	FieldValue := c.GetString("FieldValue")
	e := SysBll.ModifySysUserInfo(c.SUser.Id, FieldName, FieldValue)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "", "")
}

func (c *SysController) GetSysUserInfo() {
	data, e := SysBll.GetSysUserInfo(c.SUser.Id)

	c.JsonResultEx(data, e)
}
