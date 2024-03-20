package DataBll

import (
	"fmt"
	"github.com/TtMyth123/DengG/Lottery61BS/GConfig"
	"github.com/TtMyth123/DengG/Lottery61BS/GInstance"
	"github.com/TtMyth123/DengG/Lottery61BS/controllers/BetSiteBll"
	"github.com/TtMyth123/DengG/Lottery61BS/models"
	"github.com/TtMyth123/kit/strconvEx"
	"github.com/astaxie/beego/orm"
)

func GetBetUser(Id int64) (interface{}, error) {
	if Id == 0 {
		Id = GConfig.GetGConfig().BetUserId
	}
	o := orm.NewOrm()
	aUser := models.LoBetUser{}
	aUser.Id = Id
	e := aUser.Read(o)
	return aUser, e
}

func UpdateBetUser(p map[string]string) (interface{}, error) {
	id := strconvEx.StrTry2Int64(p["Id"], 0)
	if id == 0 {
		return nil, fmt.Errorf("有问题的数据.Id为0")
	} else {
		aUser := models.LoBetUser{
			LoBaseInfo: models.LoBaseInfo{Id: id},
		}
		e := aUser.Read(nil)
		if e != nil {
			return nil, e
		}
		//OldTypeId := aUser.TypeId

		//aUser.TypeId = strconvEx.StrTry2Int64(p["TypeId"], 0)
		aUser.Title = p["Title"]
		aUser.SiteUrl = p["SiteUrl"]
		aUser.UserName = p["UserName"]
		aUser.Pwd = p["Pwd"]
		aUser.IsLogin = 0

		e = aUser.Update(nil, "Title", "SiteUrl", "UserName", "Pwd", "IsLogin")

		GInstance.DelBetLottery(aUser.Id)
		aLottery, _ := BetSiteBll.GetAddLottery(id)
		aLottery.Refresh()
		return nil, e
	}
	return nil, nil
}
