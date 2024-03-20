package controllers

import (
	"github.com/TtMyth123/DengG/Lottery61BS/controllers/base"
	"github.com/TtMyth123/DengG/Lottery61BS/controllers/base/enums"
	"github.com/TtMyth123/DengG/Lottery61BS/controllers/regBll"
	"github.com/TtMyth123/kit"
	"strings"
)

type LoginRegController struct {
	base.ABaseController
}

func (this *LoginRegController) DoLogin() {
	username := this.GetString("username")
	password := this.GetString("password")

	username = strings.ToLower(strings.TrimSpace(username))
	userpwd := strings.ToLower(strings.TrimSpace(password))

	if len(username) == 0 {
		this.JsonResult(enums.JRCodeFailed, "用户名和密码不正确", "")
	}
	aLoginInfo := regBll.LoginInfo{Pwd: userpwd, UserName: username, LoginIp: this.Ctx.Request.RemoteAddr}
	data, e := regBll.Login(aLoginInfo)
	this.JsonResultEx(data, e)
}

func (this *LoginRegController) DoLogout() {
	//strToken := this.Ctx.Input.Header(GConfig.GetGConfig().TokenKey)
	//CacheData.DelToken(strToken)
	//this.JsonResultEx("", nil)
}

func (c *LoginRegController) GetGuid() {
	a := kit.GetGuid()
	c.JsonResult(enums.JRCodeSucc, "", a)
}
