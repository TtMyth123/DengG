package controllers

import (
	"github.com/TtMyth123/DengG/Lottery61BS/controllers/DataBll"
	"github.com/TtMyth123/kit/httpKit"
)

func (c *DataApiController) GetAgentUser() {
	Id, _ := c.GetInt64("Id")
	data, e := DataBll.GetAgentUser(Id)
	c.JsonResultEx(data, e)
}
func (c *DataApiController) UpdateAgentUser() {
	Param := httpKit.GetParamValue(c.Ctx.Request)

	data, e := DataBll.UpdateAgentUser(Param)
	c.JsonResultEx(data, e)
}

func (c *DataApiController) GetAgentBetData() {
	data, e := DataBll.GetAgentBetData(c.SUser.Id)
	c.JsonResultEx(data, e)
}
