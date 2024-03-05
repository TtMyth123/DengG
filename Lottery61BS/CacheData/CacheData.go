package CacheData

import "github.com/TtMyth123/DengG/Lottery61BS/CacheData/DataSet"

var gData CacheManage

type CacheManage struct {
	DataAgentUser DataSet.DataAgentUser
}

func Init() {
	gData = CacheManage{}
}

func GetICache() {

}
