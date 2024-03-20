package base

import (
	"encoding/json"
	"fmt"
	"github.com/TtMyth123/DengG/Lottery61BS/controllers/base/cBo"
	"github.com/TtMyth123/DengG/Lottery61BS/controllers/base/enums"
	"github.com/TtMyth123/kit/httpKit"
	"github.com/astaxie/beego"
)

type UserS struct {
	Id int64
}

type ABaseController struct {
	beego.Controller
}

func (this *ABaseController) GetJsonData(v interface{}) error {
	strJson := this.GetString("jsonData")
	e := json.Unmarshal([]byte(strJson), v)
	return e
}

func (this *ABaseController) JsonResult(code enums.JsonResultCode, msg string, obj interface{}) {
	res := JsonResult{Code: code, Msg: msg, Obj: obj}
	this.Data["json"] = res
	this.ServeJSON()
	this.StopRun()
}

func (this *ABaseController) JsonResultEx(obj interface{}, err error) {
	if err == nil {
		this.JsonResult(enums.JRCodeSucc, "", obj)
	} else {
		this.JsonResult(enums.JRCodeFailed, err.Error(), obj)
	}
}
func (this *ABaseController) GetListParam(otherPs ...cBo.ArgsParam) cBo.ListParamBox {
	data := httpKit.GetParamValue(this.Ctx.Request)
	fmt.Println(data)
	param := cBo.GetListParamBox()
	param.PageIndex, _ = this.GetInt("PageIndex")
	param.PageSize, _ = this.GetInt("PageSize")
	param.MaxId, _ = this.GetInt64("MaxId")
	param.OrderBy = this.GetString("OrderBy")
	if param.PageSize == 0 {
		param.PageSize = 10
	}

	if param.PageIndex == 0 {
		param.PageIndex = 1
	}
	for k, v := range data {
		param.Other[k] = v
	}

	for _, p := range otherPs {
		switch p.T {
		case cBo.ParamTypeS:
			if def, ok := p.DefValue.(string); ok {
				param.Other[p.PName] = this.GetString(p.PName, def)
			} else {
				param.Other[p.PName] = this.GetString(p.PName)
			}
		case p.T:
			if def, ok := p.DefValue.(int); ok {
				param.Other[p.PName], _ = this.GetInt(p.PName, def)
			} else {
				param.Other[p.PName], _ = this.GetInt(p.PName)
			}
		default:
			param.Other[p.PName] = p.DefValue
		}
	}
	return param
}
