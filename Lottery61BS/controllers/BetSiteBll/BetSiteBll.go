package BetSiteBll

import (
	"github.com/TtMyth123/DengG/Lottery61BS/GConfig"
	"github.com/TtMyth123/DengG/Lottery61BS/GInstance"

	"github.com/TtMyth123/DengG/Lottery61BS/httpSixBet/base"
	"github.com/TtMyth123/DengG/Lottery61BS/httpSixBet/manager"
	"github.com/TtMyth123/DengG/Lottery61BS/models"
	"github.com/astaxie/beego/orm"
)

func GetAddLottery(Id int64) (base.ILottery, bool) {
	if Id == 0 {
		Id = GConfig.GetGConfig().BetUserId
	}

	l := GInstance.GetBetLottery(Id)
	if l == nil || l.GetUserInfo().Id == 0 {
		o := orm.NewOrm()
		aSiteUserInfo := models.LoBetUser{LoBaseInfo: models.LoBaseInfo{Id: Id}}
		e := aSiteUserInfo.Read(o)
		if e != nil {
			return l, false
		}
		aSiteUserInfo.IsLogin = 0
		aSiteUserInfo.Update(nil, "IsLogin")

		lottery := manager.NewSixLottery(aSiteUserInfo, true)
		GInstance.AddBetLottery(Id, lottery)
		bL, ok := lottery.(base.ILottery)
		return bL, ok
	}

	return l, true
}
