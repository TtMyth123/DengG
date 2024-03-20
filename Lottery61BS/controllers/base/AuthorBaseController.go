package base

import (
	"fmt"
	"github.com/TtMyth123/DengG/Lottery61BS/CacheData"
	"github.com/TtMyth123/DengG/Lottery61BS/GConfig"
	"github.com/TtMyth123/DengG/Lottery61BS/controllers/base/cBo"
	"github.com/TtMyth123/DengG/Lottery61BS/controllers/base/enums"
	"github.com/TtMyth123/DengG/Lottery61BS/models"
)

type AuthorBaseController struct {
	ABaseController

	SUser models.SysUser
}

func (this *AuthorBaseController) Prepare() {
	switch GConfig.GetGConfig().IsDev {
	case 1:
		strToken := this.Ctx.Input.Header(GConfig.GetGConfig().TokenKey)
		aSimpleUser, e := CacheData.GetToken(strToken)
		if e == nil {
			this.SUser = aSimpleUser
		} else {
			this.SUser.Id, _ = this.GetInt64("SysId", 0)
		}

		if this.SUser.Id == 0 {
			this.SUser.Id = GConfig.GetGConfig().DevUserId
		}
	default:
		strToken := this.Ctx.Input.Header(GConfig.GetGConfig().TokenKey)
		aSimpleUser, e := CacheData.GetToken(strToken)
		if e != nil || aSimpleUser.Id == 0 {
			Ip := CacheData.Get401Ip(strToken)
			strHit := "授权过期"
			if Ip != "" {
				strHit = fmt.Sprintf("授权过期:在%s被登录", Ip)
			}
			this.JsonResult(enums.JRCode401, strHit, nil)
		}
		this.SUser = aSimpleUser
	}
}

func (this *AuthorBaseController) GetToken() string {
	strToken := this.Ctx.Input.Header("X-Token")
	return strToken
}

func (this *AuthorBaseController) GetBaseListParam(mpParam map[string]int) cBo.BaseListParam {
	aParam := cBo.BaseListParam{}
	aParam.Other = map[string]interface{}{}
	aParam.PageIndex, _ = this.GetInt("PageIndex", 1)
	aParam.PageSize, _ = this.GetInt("PageSize", 20)
	aParam.MaxId, _ = this.GetInt64("MaxId", 0)

	for fieldName, v := range mpParam {
		switch v {
		case cBo.FieldType_0_Int:
			aParam.Other[fieldName], _ = this.GetInt(fieldName)
		case cBo.FieldType_1_Str:
			aParam.Other[fieldName] = this.GetString(fieldName)
		case cBo.FieldType_2_Float64:
			aParam.Other[fieldName], _ = this.GetFloat(fieldName)
		case cBo.FieldType_3_Int64:
			aParam.Other[fieldName], _ = this.GetInt64(fieldName)
		default:
			aParam.Other[fieldName] = this.GetString(fieldName)
		}
	}

	return aParam
}
