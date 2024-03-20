package base

import "github.com/TtMyth123/DengG/Lottery61BS/models"

type ILottery interface {
	GetCode(v map[string]interface{}) (map[string]interface{}, error)
	Login(v map[string]interface{}) (map[string]interface{}, error)
	LoginOut(v map[string]interface{}) (map[string]interface{}, error)
	GetCodeImage(v map[string]interface{}) []byte
	GetOriginalData(v map[string]interface{}) (map[string]interface{}, error)
	ToMapNumData(v interface{}) map[string]models.LoAgentBetNumData
	SaveNumData(v map[string]interface{}) (map[string]interface{}, error)
	GetState() int
	//GetId() int64
	Refresh()
	GetUserInfo() models.LoAgentUser
	UpdateBaseUrl(url string) (map[string]interface{}, error)
	TestBaseUrl(url string) (map[string]interface{}, error)
	GetBaseData(v map[string]interface{}) (map[string]interface{}, error)

	GetHttpData(strUrl string) (map[string]interface{}, error)
	Stop()
}

type UpdateOdds func(v map[string]interface{}) (map[string]interface{}, error)
type UpdateCurIssue func(v map[string]interface{}) (map[string]interface{}, error)
type SaveNumData func(v map[string]interface{}) (map[string]interface{}, error)
type CheckToken func(v map[string]interface{}) (map[string]interface{}, error)
type GetStrLotteryType func(t int) string
type CheckLogin func(v map[string]interface{}) bool
