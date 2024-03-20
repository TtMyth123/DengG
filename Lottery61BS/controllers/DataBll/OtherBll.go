package DataBll

import (
	"github.com/TtMyth123/DengG/Lottery61BS/CacheData"
	"github.com/TtMyth123/kit/strconvEx"
)

func GetConfig() (interface{}, error) {
	aConfig := CacheData.GetLoConfig()
	return aConfig, nil
}

func SaveConfig(p map[string]string) (interface{}, error) {
	aConfig := CacheData.GetLoConfig()
	if p["IsAutoGetData"] != "" {
		aConfig.IsAutoGetData = strconvEx.StrTry2Int(p["IsAutoGetData"], 0)
	}
	if p["BetWay"] != "" {
		aConfig.BetWay = strconvEx.StrTry2Int(p["BetWay"], 0)
	}
	if p["Stock"] != "" {
		aConfig.Stock = strconvEx.StrTry2Int(p["Stock"], 0)
	}
	if p["Rate"] != "" {
		aConfig.Rate = strconvEx.StrTry2Int(p["Rate"], 0)
	}
	if p["LotteryType"] != "" {
		aConfig.LotteryType = strconvEx.StrTry2Int64(p["LotteryType"], 0)
	}
	aConfig.BetTime1 = p["BetTime1"]
	aConfig.BetTime2 = p["BetTime2"]

	CacheData.SetLoConfigAnd2DB(aConfig)
	return aConfig, nil
}
