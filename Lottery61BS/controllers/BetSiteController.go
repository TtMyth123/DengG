package controllers

import (
	"fmt"
	"github.com/TtMyth123/DengG/Lottery61BS/controllers/BetSiteBll"
	"github.com/TtMyth123/DengG/Lottery61BS/controllers/base"
	"github.com/TtMyth123/DengG/Lottery61BS/models"
	"github.com/TtMyth123/kit/stringKit"
)

type BetSiteController struct {
	base.ABaseController
}

func (c *BetSiteController) GetCode() {
	Id, _ := c.GetInt64("Id")
	lottery, ok := BetSiteBll.GetAddLottery(Id)
	if !ok {
		c.JsonResultEx("", fmt.Errorf("没有对应的服务"))
	}
	mpP := make(map[string]interface{})
	data, e := lottery.GetCode(mpP)
	c.JsonResultEx(data, e)
}

func (c *BetSiteController) GetCodeImage() {
	Id, _ := c.GetInt64("Id")
	lottery, ok := BetSiteBll.GetAddLottery(Id)
	if !ok {
		c.Ctx.WriteString(string(""))
		return
	}
	mpP := make(map[string]interface{})
	data := lottery.GetCodeImage(mpP)
	c.Ctx.WriteString(string(data))
}

func (c *BetSiteController) Login() {
	Id, _ := c.GetInt64("Id")
	lottery, ok := BetSiteBll.GetAddLottery(Id)
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

func (c *BetSiteController) LoginOut() {
	Id, _ := c.GetInt64("Id")
	lottery, ok := BetSiteBll.GetAddLottery(Id)
	if !ok {
		c.JsonResultEx("", fmt.Errorf("没有对应的服务"))
	}
	mpP := make(map[string]interface{})
	data, e := lottery.LoginOut(mpP)
	c.JsonResultEx(data, e)
}

/**
[{"Num":0,"BetM":0,"I":0,"OkBetM":0,"V":null},{"Num":49,"BetM":0,"I":0,"OkBetM":0,"V":432},{"Num":48,"BetM":0,"I":1,"OkBetM":0,"V":232},{"Num":47,"BetM":0,"I":2,"OkBetM":0,"V":2},{"Num":46,"BetM":0,"I":3,"OkBetM":0,"V":111},{"Num":45,"BetM":0,"I":4,"OkBetM":0,"V":333},{"Num":44,"BetM":0,"I":5,"OkBetM":0,"V":null},{"Num":43,"BetM":0,"I":6,"OkBetM":0,"V":null},{"Num":42,"BetM":0,"I":7,"OkBetM":0,"V":null},{"Num":41,"BetM":0,"I":8,"OkBetM":0,"V":null},{"Num":40,"BetM":0,"I":9,"OkBetM":0,"V":null},{"Num":39,"BetM":0,"I":10,"OkBetM":0,"V":null},{"Num":38,"BetM":0,"I":11,"OkBetM":0,"V":null},{"Num":37,"BetM":0,"I":12,"OkBetM":0,"V":null},{"Num":36,"BetM":0,"I":13,"OkBetM":0,"V":null},{"Num":35,"BetM":0,"I":14,"OkBetM":0,"V":null},{"Num":34,"BetM":0,"I":15,"OkBetM":0,"V":null},{"Num":33,"BetM":0,"I":16,"OkBetM":0,"V":null},{"Num":32,"BetM":0,"I":17,"OkBetM":0,"V":null},{"Num":31,"BetM":0,"I":18,"OkBetM":0,"V":null},{"Num":30,"BetM":0,"I":19,"OkBetM":0,"V":null},{"Num":29,"BetM":0,"I":20,"OkBetM":0,"V":null},{"Num":28,"BetM":0,"I":21,"OkBetM":0,"V":null},{"Num":27,"BetM":0,"I":22,"OkBetM":0,"V":null},{"Num":26,"BetM":0,"I":23,"OkBetM":0,"V":null},{"Num":25,"BetM":0,"I":24,"OkBetM":0,"V":null},{"Num":24,"BetM":0,"I":25,"OkBetM":0,"V":null},{"Num":23,"BetM":0,"I":26,"OkBetM":0,"V":null},{"Num":22,"BetM":0,"I":27,"OkBetM":0,"V":null},{"Num":21,"BetM":0,"I":28,"OkBetM":0,"V":null},{"Num":20,"BetM":0,"I":29,"OkBetM":0,"V":null},{"Num":19,"BetM":0,"I":30,"OkBetM":0,"V":null},{"Num":18,"BetM":0,"I":31,"OkBetM":0,"V":null},{"Num":17,"BetM":0,"I":32,"OkBetM":0,"V":null},{"Num":16,"BetM":0,"I":33,"OkBetM":0,"V":null},{"Num":15,"BetM":0,"I":34,"OkBetM":0,"V":null},{"Num":14,"BetM":0,"I":35,"OkBetM":0,"V":null},{"Num":13,"BetM":0,"I":36,"OkBetM":0,"V":null},{"Num":12,"BetM":0,"I":37,"OkBetM":0,"V":null},{"Num":11,"BetM":0,"I":38,"OkBetM":0,"V":null},{"Num":10,"BetM":0,"I":39,"OkBetM":0,"V":null},{"Num":9,"BetM":0,"I":40,"OkBetM":0,"V":null},{"Num":8,"BetM":0,"I":41,"OkBetM":0,"V":null},{"Num":7,"BetM":0,"I":42,"OkBetM":0,"V":null},{"Num":6,"BetM":0,"I":43,"OkBetM":0,"V":null},{"Num":5,"BetM":0,"I":44,"OkBetM":0,"V":null},{"Num":4,"BetM":0,"I":45,"OkBetM":0,"V":null},{"Num":3,"BetM":0,"I":46,"OkBetM":0,"V":null},{"Num":2,"BetM":0,"I":47,"OkBetM":0,"V":null},{"Num":1,"BetM":0,"I":48,"OkBetM":0,"V":null},{"Num":0,"BetM":0,"I":0,"OkBetM":0,"V":null}]
*/

func (c *BetSiteController) Bet() {
	Id, _ := c.GetInt64("Id")
	lottery, ok := BetSiteBll.GetAddLottery(Id)
	if !ok {
		c.JsonResultEx("", fmt.Errorf("没有对应的服务"))
	}

	BetData := c.GetString("BetData")
	mpP := make(map[string]interface{})
	mpP["StrBetData"] = BetData
	StrLoConfig := c.GetString("LoConfig")
	if StrLoConfig != "" {
		aLoConfig := models.LoConfig{}
		e := stringKit.GetJsonObj(StrLoConfig, &aLoConfig)
		if e != nil {
			c.JsonResultEx("data", e)
		} else {
			mpP["LoConfig"] = aLoConfig
		}
	}

	data, e := lottery.Bet(mpP)
	c.JsonResultEx(data, e)
}
