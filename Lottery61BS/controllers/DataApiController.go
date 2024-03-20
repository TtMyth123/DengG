package controllers

import (
	"github.com/TtMyth123/DengG/Lottery61BS/controllers/DataBll"
	"github.com/TtMyth123/DengG/Lottery61BS/controllers/base"
	"github.com/TtMyth123/DengG/Lottery61BS/controllers/bll"
	"github.com/TtMyth123/kit/httpKit"
)

type DataApiController struct {
	base.AuthorBaseController
}

func (c *DataApiController) GetC() {
	data, e := bll.GetC(c.SUser.Id)
	c.JsonResultEx(data, e)
}

func (c *DataApiController) GetConfig() {
	data, e := DataBll.GetConfig()

	c.JsonResultEx(data, e)
}
func (c *DataApiController) SaveConfig() {
	Param := httpKit.GetParamValue(c.Ctx.Request)

	data, e := DataBll.SaveConfig(Param)
	c.JsonResultEx(data, e)
}
