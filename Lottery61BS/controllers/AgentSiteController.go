package controllers

import (
	"fmt"
	"github.com/TtMyth123/DengG/Lottery61BS/controllers/AgentSiteBll"
	"github.com/TtMyth123/DengG/Lottery61BS/controllers/DataBll"
	"github.com/TtMyth123/DengG/Lottery61BS/controllers/base"
)

type AgentSiteController struct {
	base.ABaseController
}

func (c *AgentSiteController) GetCode() {
	Id, _ := c.GetInt64("Id")
	lottery, ok := AgentSiteBll.GetAddLottery(Id)
	if !ok {
		c.JsonResultEx("", fmt.Errorf("没有对应的服务"))
	}
	mpP := make(map[string]interface{})
	data, e := lottery.GetCode(mpP)
	c.JsonResultEx(data, e)
}

func (c *AgentSiteController) GetCodeImage() {
	Id, _ := c.GetInt64("Id")
	lottery, ok := AgentSiteBll.GetAddLottery(Id)
	if !ok {
		c.Ctx.WriteString(string(""))
		return
	}
	mpP := make(map[string]interface{})
	data := lottery.GetCodeImage(mpP)
	c.Ctx.WriteString(string(data))
}

func (c *AgentSiteController) Login() {
	Id, _ := c.GetInt64("Id")
	lottery, ok := AgentSiteBll.GetAddLottery(Id)
	if !ok {
		c.JsonResultEx("", fmt.Errorf("没有对应的服务"))
	}
	UserName := c.GetString("UserName")
	Pwd := c.GetString("Pwd")
	Code := c.GetString("Code")

	mpP := make(map[string]interface{})
	mpP["UserName"] = UserName
	mpP["Pwd"] = Pwd
	mpP["Code"] = Code
	data, e := lottery.Login(mpP)
	c.JsonResultEx(data, e)
}

func (c *AgentSiteController) LoginOut() {
	Id, _ := c.GetInt64("Id")
	lottery, ok := AgentSiteBll.GetAddLottery(Id)
	if !ok {
		c.JsonResultEx("", fmt.Errorf("没有对应的服务"))
	}
	mpP := make(map[string]interface{})
	data, e := lottery.LoginOut(mpP)
	c.JsonResultEx(data, e)
}

func (c *AgentSiteController) GetOriginalData() {
	Id, _ := c.GetInt64("Id")
	lottery, ok := AgentSiteBll.GetAddLottery(Id)
	if !ok {
		c.JsonResultEx("", fmt.Errorf("没有对应的服务"))
	}
	mpP := make(map[string]interface{})
	data, e := lottery.GetOriginalData(mpP)
	c.JsonResultEx(data, e)
}

func (c *AgentSiteController) GetAndSaveNumData() {
	AgentId, _ := c.GetInt64("AgentId")
	IsSave, _ := c.GetInt64("IsSave")
	mpP := make(map[string]interface{})
	data, e := DataBll.GetAndSaveNumData(AgentId, IsSave, mpP)
	c.JsonResultEx(data, e)
}
