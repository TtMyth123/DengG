package controllers

import (
	"github.com/TtMyth123/DengG/Lottery61BS/controllers/DataBll"
	"github.com/TtMyth123/kit/httpKit"
)

func (c *DataApiController) GetBetUser() {
	Id, _ := c.GetInt64("Id")
	data, e := DataBll.GetBetUser(Id)
	c.JsonResultEx(data, e)
}

func (c *DataApiController) UpdateBetUser() {
	Param := httpKit.GetParamValue(c.Ctx.Request)
	data, e := DataBll.UpdateBetUser(Param)
	c.JsonResultEx(data, e)
}
