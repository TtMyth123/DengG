package controllers

import (
	"github.com/TtMyth123/DengG/Lottery61BS/controllers/base"
)

type ApiController struct {
	base.ABaseController
}

func (c *ApiController) Hello() {
	c.JsonResultEx("Hello", nil)
}

func (c *ApiController) GetCode() {
	//Id, _ := c.GetInt64("Id")
	//lottery, ok := bll.GetAddLottery(Id)
	//if !ok {
	//	c.JsonResultEx("", fmt.Errorf("没有对应的服务"))
	//}
	//mpP := make(map[string]interface{})
	//data, e := lottery.GetCode(mpP)
	//c.JsonResultEx(data, e)
}
func (c *ApiController) GetCodeImage() {
	//Id, _ := c.GetInt64("Id")
	//lottery, ok := bll.GetAddLottery(Id)
	//if !ok {
	//	c.Ctx.WriteString(string(""))
	//	return
	//}
	//mpP := make(map[string]interface{})
	//data := lottery.GetCodeImage(mpP)
	//
	//c.Ctx.WriteString(string(data))
}

func (c *ApiController) Login() {
	//Id, _ := c.GetInt64("Id")
	//lottery, ok := bll.GetAddLottery(Id)
	//if !ok {
	//	c.JsonResultEx("", fmt.Errorf("没有对应的服务"))
	//}
	//UserName := c.GetString("UserName")
	//Pwd := c.GetString("Pwd")
	//Code := c.GetString("Code")
	//IsBet, _ := c.GetInt("IsBet")
	//
	//mpP := make(map[string]interface{})
	//mpP["UserName"] = UserName
	//mpP["Pwd"] = Pwd
	//mpP["Code"] = Code
	//mpP["IsBet"] = IsBet
	//data, e := lottery.Login(mpP)
	//c.JsonResultEx(data, e)
}
func (c *ApiController) LoginOut() {
	//Id, _ := c.GetInt64("Id")
	//lottery, ok := bll.GetAddLottery(Id)
	//if !ok {
	//	c.JsonResultEx("", fmt.Errorf("没有对应的服务"))
	//}
	//mpP := make(map[string]interface{})
	//data, e := lottery.LoginOut(mpP)
	//c.JsonResultEx(data, e)
}

func (c *ApiController) GetOriginalData() {
	//Id, _ := c.GetInt64("Id")
	//lottery, ok := bll.GetAddLottery(Id)
	//ABType := c.GetString("ABType")
	//DataType, _ := c.GetInt("DataType")
	//if !ok {
	//	c.JsonResultEx("", fmt.Errorf("没有对应的服务"))
	//}
	//
	//data, e := lottery.GetOriginalData(nil, ABType, DataType)
	//c.JsonResultEx(data, e)
}

func (c *ApiController) Bet() {
	//Id, _ := c.GetInt64("Id")
	//lottery, ok := bll.GetAddLottery(Id)
	//if !ok {
	//	c.JsonResultEx("", fmt.Errorf("没有对应的服务"))
	//}
	//
	//LotteryType, _ := c.GetInt("LotteryType")
	//BetData := c.GetString("BetData")
	//mpP := make(map[string]interface{})
	//mpP["BetData"] = BetData
	//mpP["LotteryType"] = LotteryType
	//
	//data, e := lottery.Bet(mpP)
	//c.JsonResultEx(data, e)
}
func (c *ApiController) GetBaseData() {
	//Id, _ := c.GetInt64("Id")
	//lottery, ok := bll.GetAddLottery(Id)
	//if !ok {
	//	c.JsonResultEx("", fmt.Errorf("没有对应的服务"))
	//}
	//
	//BetData := c.GetString("BetData")
	//mpP := make(map[string]interface{})
	//mpP["BetData"] = BetData
	//data, e := lottery.GetBaseData(mpP)
	//c.JsonResultEx(data, e)
}
