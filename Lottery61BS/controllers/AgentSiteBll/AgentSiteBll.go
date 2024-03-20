package AgentSiteBll

import (
	"github.com/TtMyth123/DengG/Lottery61BS/GConfig"
	"github.com/TtMyth123/DengG/Lottery61BS/GInstance"
	"github.com/TtMyth123/DengG/Lottery61BS/httpSixAgent/base"
	"github.com/TtMyth123/DengG/Lottery61BS/httpSixAgent/manager"
	"github.com/TtMyth123/DengG/Lottery61BS/models"
	"github.com/astaxie/beego/orm"
)

func GetAddLottery(Id int64) (base.ILottery, bool) {
	if Id == 0 {
		Id = GConfig.GetGConfig().AgentUserId
	}

	l := GInstance.GetAgentLottery(Id)
	if l == nil || l.GetUserInfo().Id == 0 {
		o := orm.NewOrm()
		aSiteUserInfo := models.LoAgentUser{LoBaseInfo: models.LoBaseInfo{Id: Id}}
		e := aSiteUserInfo.Read(o)
		if e != nil {
			return l, false
		}

		lottery := manager.NewSixLottery(aSiteUserInfo, true)
		GInstance.AddAgentLottery(Id, lottery)
		bL, ok := lottery.(base.ILottery)
		return bL, ok
	}

	return l, true
}

func LoadData(Id int64, v map[string]interface{}) (base.ILottery, bool) {
	if Id == 0 {
		Id = GConfig.GetGConfig().AgentUserId
	}

	l := GInstance.GetAgentLottery(Id)
	l.SaveNumData(v)
	return l, true
}
