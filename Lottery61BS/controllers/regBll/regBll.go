package regBll

import (
	"fmt"
	"github.com/TtMyth123/DengG/Lottery61BS/CacheData"
	"github.com/TtMyth123/DengG/Lottery61BS/controllers/SysBll"
	"github.com/TtMyth123/DengG/Lottery61BS/models"
	"github.com/TtMyth123/kit/pwdKit"
	"time"
)

type LoginInfo struct {
	Pwd      string
	UserName string
	LoginIp  string
}

func Login(aLoginInfo LoginInfo) (interface{}, error) {
	mpData := make(map[string]interface{})
	newUserpwd := pwdKit.Sha1ToStr(aLoginInfo.Pwd)
	user, e := models.SysUserOneByUserName(aLoginInfo.UserName, newUserpwd)
	if e != nil {
		return mpData, fmt.Errorf("用户名或者密码错误")
	}

	return loginOk(user, aLoginInfo.LoginIp)
}

func loginOk(aUser models.SysUser, LoginIp string) (interface{}, error) {
	mpData := make(map[string]interface{})
	aUser.LoginTime = time.Now()
	aUser.LoginIp = LoginIp
	aUser.Update(nil, "LoginTime", "LoginIp")

	strToken, e := CacheData.SetToken(aUser)
	if e != nil {
		return mpData, e
	}

	menuList := SysBll.GetSysMenuList(aUser.Id)

	mpData["Token"] = strToken
	mpData["MenuList"] = menuList
	return mpData, e
}
