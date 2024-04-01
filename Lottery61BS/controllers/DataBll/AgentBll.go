package DataBll

import (
	"fmt"
	"github.com/TtMyth123/DengG/Lottery61BS/CacheData"
	"github.com/TtMyth123/DengG/Lottery61BS/GConfig"
	"github.com/TtMyth123/DengG/Lottery61BS/GInstance"
	"github.com/TtMyth123/DengG/Lottery61BS/controllers/AgentSiteBll"
	"github.com/TtMyth123/DengG/Lottery61BS/models"
	"github.com/TtMyth123/kit/strconvEx"
	"github.com/astaxie/beego/orm"
)

func GetAgentUser(Id int64) (interface{}, error) {
	if Id == 0 {
		Id = GConfig.GetGConfig().AgentUserId
	}
	o := orm.NewOrm()
	aUser := models.LoAgentUser{}
	aUser.Id = Id
	e := aUser.Read(o)
	return aUser, e
}

func UpdateAgentUser(p map[string]string) (interface{}, error) {
	id := strconvEx.StrTry2Int64(p["Id"], 0)
	if id == 0 {
		return nil, fmt.Errorf("有问题的数据.Id为0")
	} else {
		aUser := models.LoAgentUser{
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
		GInstance.DelAgentLottery(aUser.Id)
		aLottery, _ := AgentSiteBll.GetAddLottery(id)
		aLottery.Refresh()
		return nil, e
	}
	return nil, nil
}

type AgentBetData struct {
	Num    int
	BetM   int
	I      int
	OkBetM int
	V      int
}

func GetAgentBetData(SysId int64) (interface{}, error) {
	o := orm.NewOrm()
	arrData := make([]AgentBetData, 0)
	sqlArgs := make([]interface{}, 0)
	loConfig := CacheData.GetLoConfig()

	sql := fmt.Sprintf(
		` select a.num, b.bet_m, c.bet_m as ok_bet_m from lo_num_info a
LEFT JOIN 
(select aa.num, sum(aa.bet_m ) bet_m from lo_agent_bet_num_data aa, lo_agent_user bb where 
	(aa.agent_user_id = bb.id and aa.sys_id=? and aa.lottery_t=? and aa.stock=?) 
group by num) b on (a.num = b.num) 
LEFT JOIN 
(select aa.num, sum(aa.bet_m ) bet_m from lo_bet_bet_data aa, lo_bet_user bb, lo_lottery_type cc where 
	(aa.cycle_id=cc.cycle_id and aa.lottery_t=cc.id and aa.user_id = bb.id and aa.sys_id=? and aa.lottery_t=?  ) 
group by num) c on (a.num = c.num)
ORDER BY b.bet_m DESC, a.num `)
	sqlArgs = append(sqlArgs,
		SysId, loConfig.LotteryType, loConfig.Stock,
		SysId, loConfig.LotteryType,
	)
	_, e := o.Raw(sql, sqlArgs).QueryRows(&arrData)
	iLen := len(arrData)
	for i := 0; i < iLen; i++ {
		arrData[i].I = i
	}
	newData := make([]AgentBetData, 0)
	AgentBetData0 := AgentBetData{}
	AgentBetData50 := AgentBetData{}
	newData = append(newData, AgentBetData0)
	newData = append(newData, arrData...)
	newData = append(newData, AgentBetData50)
	return newData, e
}

func GetAndSaveNumData(AgentId, IsSave int64, mpP map[string]interface{}) (interface{}, error) {
	lottery, ok := AgentSiteBll.GetAddLottery(AgentId)
	if !ok {
		return nil, fmt.Errorf("没有对应的服务")
	}

	mpData := make(map[string]interface{})
	mpData["Msg"] = ""
	if IsSave != 0 {
		_, e := lottery.SaveNumData(mpP)

		if e != nil {
			mpData["Msg"] = e.Error()
		}
	}

	data, e := GetAgentBetData(lottery.GetUserInfo().SysId)
	mpData["Data"] = data
	return mpData, e
}
