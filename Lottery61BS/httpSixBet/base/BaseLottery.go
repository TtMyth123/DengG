package base

import (
	"fmt"
	"github.com/TtMyth123/DengG/Lottery61BS/models"
	"github.com/TtMyth123/kit/httpClientKit"
	"strings"
	"sync"
	"time"
)

type Lottery struct {
	aHttpClient1 *httpClientKit.HttpClient
	mTicker      *time.Ticker
	mTickerI     int
	mTickerI2    int
	UserInfo     models.LoBetUser

	UpdateOddsF     UpdateOdds
	UpdateCurIssueF UpdateCurIssue
	CheckTokenF     CheckToken

	CloseTime    time.Time
	TypeDataLock sync.RWMutex

	SaveNumDataF SaveNumData
	CheckLoginF  CheckLogin
}

var empServer *Lottery

func init() {
	empServer = &Lottery{}
}
func GetEmpLottery() *Lottery {
	return empServer
}

func NewLottery(UserInfo models.LoBetUser) *Lottery {
	aServer := &Lottery{}
	aServer.aHttpClient1 = httpClientKit.GetHttpClient("")
	aServer.aHttpClient1.RetryI = 1
	aServer.UserInfo = UserInfo

	aServer.UpdateLoginState()
	go aServer.run()

	return aServer
}

func (receiver *Lottery) Refresh() {
}

func (receiver *Lottery) GetHttp() *httpClientKit.HttpClient {
	return receiver.aHttpClient1
}

func (receiver *Lottery) GetCode(v map[string]interface{}) (map[string]interface{}, error) {
	mpData := make(map[string]interface{})
	return mpData, nil
}
func (receiver *Lottery) Login(v map[string]interface{}) (map[string]interface{}, error) {
	mpData := make(map[string]interface{})
	return mpData, nil
}
func (receiver *Lottery) LoginOut(v map[string]interface{}) (map[string]interface{}, error) {
	receiver.UserInfo.IsLogin = 0
	mpData := make(map[string]interface{})

	return mpData, nil
}
func (this *Lottery) GetCodeImage(v map[string]interface{}) []byte {
	data, e := this.GetCode(v)
	if e != nil {
		return nil
	}

	dataB := data["CodePic"].([]byte)
	return dataB
}

func (receiver *Lottery) GetState() int {
	return receiver.UserInfo.State
}

func (receiver *Lottery) UpdateBaseUrl(url string) (map[string]interface{}, error) {
	receiver.UserInfo.SiteUrl = url

	return nil, nil
}
func (this *Lottery) Stop() {
	if this.mTicker != nil {
		this.mTicker.Stop()
	}
}
func (this *Lottery) run() {
	this.mTicker = time.NewTicker(time.Second)
	for {
		select {
		case <-this.mTicker.C:
			this.mTickerI++
			if this.mTickerI > 1000000 {
				this.mTickerI = 0
			}

			if this.mTickerI%15 == 0 {
				this.UpdateLoginState()
			}

			if this.UserInfo.IsLogin == 0 {
				continue
			}
		}

	}
}

func (receiver *Lottery) ToMapNumData(v interface{}, ABType string) map[string]models.LoAgentBetNumData {
	return nil
}
func (receiver *Lottery) SaveNumData(v map[string]interface{}) (map[string]interface{}, error) {
	return nil, nil
}

func (receiver *Lottery) GetUserInfo() models.LoBetUser {
	return receiver.UserInfo
}

func (receiver *Lottery) TestBaseUrl(url string) (map[string]interface{}, error) {
	return nil, fmt.Errorf("TestBaseUrl未完成")
}

func (receiver *Lottery) Bet(v map[string]interface{}) (map[string]interface{}, error) {
	return nil, fmt.Errorf("Bet 未完成")
}

func (receiver *Lottery) GetBaseData(v map[string]interface{}) (map[string]interface{}, error) {
	return nil, fmt.Errorf("GetBaseData 未完成")
}

func (receiver *Lottery) GetHttpData(strUrl string) (map[string]interface{}, error) {
	mp := make(map[string]interface{})
	aHttp := receiver.GetHttp()
	if aHttp == nil {
		return mp, fmt.Errorf("没有GetHttp")
	}
	data, e := aHttp.GetString(strUrl)
	mp["Data"] = data
	mp["E"] = e

	return mp, nil
}

func (this *Lottery) UpdateLoginState2Db(b bool) {
	if b {
		this.UserInfo.IsLogin = 1
	} else {
		this.UserInfo.IsLogin = 0

	}
	this.UserInfo.Update(nil, "IsLogin")
}
func (this *Lottery) UpdateLogInfo(LogInfo string) {
	this.UserInfo.LogInfo = LogInfo
}

func (this *Lottery) FormatCycleId(CycleId string) string {
	newCycleId := strings.Replace(CycleId, "-", "", -1)
	return newCycleId
}

func (this *Lottery) UpdateLoginState() {
	if this.CheckLoginF != nil {
		isLoginB := this.CheckLoginF(nil)
		IsLogin := 0
		if isLoginB {
			IsLogin = 1
		}
		if this.UserInfo.IsLogin != IsLogin {
			this.UserInfo.IsLogin = IsLogin
			this.UserInfo.Update(nil, "IsLogin")
		}
	}
}
