package CS2

import (
	"fmt"
	"github.com/TtMyth123/DengG/Lottery61BS/CacheData"
	"github.com/TtMyth123/DengG/Lottery61BS/httpSixAgent/CS2/box"
	"github.com/TtMyth123/DengG/Lottery61BS/httpSixAgent/base"
	"github.com/TtMyth123/DengG/Lottery61BS/models"
	"github.com/TtMyth123/kit/strconvEx"
	"github.com/TtMyth123/kit/stringKit"
	"github.com/TtMyth123/kit/ttLog"
	"github.com/astaxie/beego/orm"
	"net/url"
	"strings"
	"time"
)

const (
	lotteryType_MCHK6 = "MCHK6"
	settingCode_MCHK6 = "TMS"
)

/*
*
https://mye29wvi-cs2.cs111.vip
账号：dc2026
密码：aaa111
*/
type Cs2Lottery struct {
	*base.Lottery
	CodeId string
	UUID   string

	mpLotteryData map[int64]*Cs2LotteryData
	isSaveNumData bool
}

type Cs2LotteryData struct {
	LotteryType    int64
	StrLotteryType string

	MpNum2ItemId map[int]string

	LotCode string

	ItemLids         string
	BetLotCode       string
	Is_stock         int
	OriginalDataGUID string
	CountGUID        string
	CycleId          string
}

const (
	key_Num2ItemId = "Num2ItemId"
)

func NewLottery(aSiteUserInfo models.LoAgentUser) *Cs2Lottery {
	aServer := &Cs2Lottery{}
	aServer.Lottery = base.NewLottery(aSiteUserInfo)
	aServer.Lottery.CheckLoginF = aServer.CheckLogin

	aServer.mpLotteryData = map[int64]*Cs2LotteryData{}
	return aServer
}
func (this *Cs2Lottery) getInitData(t int64) map[string]interface{} {
	mpResult := make(map[string]interface{})
	mpNum2ItemId := make(map[int]string)
	iT := 64080
	iTid := 64129
	for i := 1; i <= 49; i++ {
		mpNum2ItemId[i] = fmt.Sprintf("%d,%d", iT, iTid)
		iTid++
		iT++
	}

	mpResult[key_Num2ItemId] = mpNum2ItemId
	return mpResult
}

func (this *Cs2Lottery) getCode(url string, v map[string]interface{}) (map[string]interface{}, error) {
	mpData := make(map[string]interface{})
	d := time.Now().UnixMicro()
	strUrl := fmt.Sprintf(`%s/index.php?controller=privilege&action=ImgCode&0.%d`, url, d)
	datas, e := this.GetHttp().GetBytes(strUrl)
	mpData["CodePic"] = datas
	return mpData, e
}

func (this *Cs2Lottery) GetCode(v map[string]interface{}) (map[string]interface{}, error) {
	return this.getCode(this.UserInfo.SiteUrl, v)
}

func (this *Cs2Lottery) GetCodeImage(v map[string]interface{}) []byte {
	data, e := this.GetCode(v)
	if e != nil {
		return nil
	}

	dataB := data["CodePic"].([]byte)
	return dataB
}

func (this *Cs2Lottery) Login(v map[string]interface{}) (map[string]interface{}, error) {
	UserName := v["UserName"].(string)
	Pwd := v["Pwd"].(string)
	Code := v["Code"].(string)
	if UserName == "" {
		if this.UserInfo.UserName == "" {
			this.UserInfo.Read(nil)
		}
		UserName = this.UserInfo.UserName
	}
	if Pwd == "" {
		Pwd = this.UserInfo.Pwd
	}

	mpData := make(map[string]interface{})
	//strUrl := `https://mye29wvi-cs2.cs111.vip/index.php?controller=privilege&action=login`
	strUrl := fmt.Sprintf(`%s/index.php?controller=privilege&action=login`, this.UserInfo.SiteUrl)
	mpParam := url.Values{}
	mpParam.Set("username", UserName)
	mpParam.Set("password", Pwd)
	mpParam.Set("imgcode", Code)
	this.UUID = ""
	datas, Location, e := this.GetHttp().Post302LocationFormHeader(strUrl, nil, mpParam)

	strScript := `</script>`
	lenScript := len(strScript)
	iIndex := strings.LastIndex(datas, strScript)
	lenS := len(datas)
	if iIndex+lenScript == lenS {
		hint := stringKit.GetBetweenStr(datas, `alert('`, `')`)
		if hint == "" {
			hint = "登录失败。"
		}
		return mpData, fmt.Errorf(hint)
	}

	if Location == "" {
		hint := "登录失败。"
		return mpData, fmt.Errorf(hint)
	}
	items := strings.Split(Location, "?")
	if len(items) != 2 {
		hint := "登录失败。"
		return mpData, fmt.Errorf(hint)
	}
	this.UUID = strings.Trim(items[1], `/`)

	this.UpdateLoginState2Db(true)
	this.loginOKinit()

	return mpData, e
}
func (this *Cs2Lottery) LoginOut(v map[string]interface{}) (map[string]interface{}, error) {
	mpData := make(map[string]interface{})
	//strUrl := `https://mye29wvi-cs2.cs111.vip/index.php?%s`
	strUrl := fmt.Sprintf(`%s/index.php?index.php?%s`, this.UserInfo.SiteUrl, this.UUID)

	_, e := this.GetHttp().GetString(strUrl)

	this.UpdateLoginState2Db(false)

	return mpData, e
}

func (this *Cs2Lottery) updateCurIssue(v map[string]interface{}) (map[string]interface{}, error) {
	//aLotteryData := this.GetLotteryData(this.GetCurLotteryType())
	//if aLotteryData.StrLotteryType == "" {
	//	return nil, nil
	//}
	////strUrl := `https://mye29wvi-cs2.cs111.vip/index.php?controller=admin_index&action=FindOpenInfo&code=bjamlh`
	//strUrl := fmt.Sprintf(
	//	`%s/index.php?controller=admin_index&action=FindOpenInfo&code=%s`,
	//	this.UserInfo.SiteUrl, aLotteryData.LotCode)
	//
	//datas, e := this.GetHttp().GetString(strUrl)
	//if e != nil {
	//	ttLog.LogError("获取数据当前结果出错。E：", e)
	//	return nil, e
	//}
	//if datas == "" {
	//	this.UpdateState(true)
	//}
	//
	//aTmp := &box.GetCurrentInstall{}
	//e = stringKit.GetJsonObj(datas, aTmp)
	//if e != nil {
	//	//<meta http-equiv='Content-Type' content='text/html; charset=utf-8' /><script language='javascript'> alert('登陆超时或其他原因请重新登陆!');top.top.location.href="/index.php?Y29udHJvbGxlcj1wcml2aWxlZ2UmYWN0aW9uPWluZGV4/";</script>
	//	strHint := stringKit.GetBetweenStr(datas, `alert('`, `);top.top.location.href=`)
	//	ii := strings.Index(strHint, `登`)
	//	if ii >= 0 {
	//		this.UpdateState(false)
	//		return nil, fmt.Errorf(strHint)
	//	}
	//
	//	this.UpdateState(true)
	//	return nil, e
	//}
	////this.UpdateState(false)
	//if len(aTmp.Data) > 0 {
	//	aLotteryData.CycleId = aTmp.Data[0].Cycleid
	//
	//	FormatCycleId := this.FormatCycleId(aLotteryData.CycleId)
	//	this.SiteUserInfo.CycleId = FormatCycleId
	//	this.SiteUserInfo.LotteryType = aLotteryData.LotteryType
	//	this.SiteUserInfo.Update(nil, "CycleId", "LotteryType")
	//
	//	iCloseTime := kit.GetInterface2Int64(aTmp.Data[0].PreData.CloseTime, 0) * 1000
	//	CloseTime := time.UnixMilli(iCloseTime)
	//	this.CloseTime = CloseTime
	//	CacheData.SetCurCycleCloseTime(this.GetCurLotteryType(), this.CloseTime, false)
	//}

	return nil, nil
}
func (this *Cs2Lottery) loginOKinit() {
	this.updateCurIssue(nil)
}

type OriginalQueryDataP struct {
	LotteryT         int64
	Stock            int
	OriginalDataGUID string
	ItemLids         string
	BetLotCode       string
}

func (this *Cs2Lottery) getOriginalQueryDataP(v map[string]interface{}) *OriginalQueryDataP {
	aOriginalQueryDataP := &OriginalQueryDataP{}
	if v["OriginalQueryDataP"] == nil {
		aLoConfig := CacheData.GetLoConfig()
		aLotteryData := this.GetLotteryData(aLoConfig.LotteryType)
		aOriginalQueryDataP.LotteryT = aLoConfig.LotteryType
		aOriginalQueryDataP.Stock = aLoConfig.Stock
		aOriginalQueryDataP.OriginalDataGUID = aLotteryData.OriginalDataGUID
		aOriginalQueryDataP.ItemLids = aLotteryData.ItemLids
		aOriginalQueryDataP.BetLotCode = aLotteryData.BetLotCode
	} else {
		aOriginalQueryDataP = v["OriginalQueryDataP"].(*OriginalQueryDataP)
	}
	return aOriginalQueryDataP
}

func (this *Cs2Lottery) GetOriginalData(v map[string]interface{}) (map[string]interface{}, error) {
	mpResult := make(map[string]interface{})
	aOriginalQueryDataP := this.getOriginalQueryDataP(v)

	//strUrl := `https://mye29wvi-cs2.cs111.vip/index.php?Y29udHJvbGxlcj1nYW1yYXRlJmFjdGlvbj1zaG93cmF0ZQ==/&lid=1974&code=bjamlh_tm&&1=1&stypes=&selectsid=&is_stock=1&region=all&orderby=balls&is_ajax=1`
	strUrl := fmt.Sprintf(`%s/index.php?%s/&lid=%s&code=%s&&1=1&stypes=&selectsid=&is_stock=%d&region=all&orderby=balls&is_ajax=1`,
		this.UserInfo.SiteUrl, aOriginalQueryDataP.OriginalDataGUID,
		aOriginalQueryDataP.ItemLids, aOriginalQueryDataP.BetLotCode, aOriginalQueryDataP.Stock)
	data, e := this.GetHttp().GetString(strUrl)
	if e != nil {
		return mpResult, e
	}
	if data == "" {
		this.UpdateLoginState2Db(false)
		return mpResult, e
	} else {
		type Tmp struct {
			Code string `json:"code"`
			Msg  string `json:"msg"`
		}
		aTmp := Tmp{}
		//{"code":"页面出错了，可能是权限不足，请核实后重新访问！","msg":"logout"}
		e = stringKit.GetJsonObj(data, &aTmp)
		if e == nil {
			if aTmp.Msg == "logout" {
				this.UpdateLoginState2Db(false)
				return mpResult, fmt.Errorf(aTmp.Code)
			}
		}
	}

	mpResult["Data"] = data

	return mpResult, nil
}

/**
v := make(map[string]interface{}
v["Data"]
v["OriginalQueryDataP"]

*/
func (this *Cs2Lottery) ToMapNumData(v interface{}) map[string]models.LoAgentBetNumData {
	mpP := v.(map[string]interface{})
	mpData := make(map[string]models.LoAgentBetNumData)
	str, ok := mpP["Data"].(string)
	if !ok {
		ttLog.LogError("类型不对")
		return mpData
	}
	aOriginalQueryDataP := this.getOriginalQueryDataP(mpP)
	Stock := aOriginalQueryDataP.Stock
	LotteryT := aOriginalQueryDataP.LotteryT

	arr := strings.Split(str, `<tr`)
	for _, item := range arr {
		//fmt.Println(item)
		num := stringKit.GetBetweenStr(item, `code='`, `'`)
		if num == "" {
			continue
		}

		//iIdRate := strings.Index(item, `id='rate_`)
		//tmp := item[iIdRate:]
		//rate := stringKit.GetBetweenStr(tmp, `>`, `<`)

		iId_cash := strings.Index(item, `id='cash_`)
		cash := stringKit.GetBetweenStr(item[iId_cash:], `>`, `<`)

		//iId_td_result := strings.Index(item, `td_result`)
		//result := stringKit.GetBetweenStr(item[iId_td_result:], `>`, `<`)
		aNumData := models.LoAgentBetNumData{
			LoBaseInfo:  models.LoBaseInfo{SysId: this.UserInfo.SysId},
			AgentUserId: this.UserInfo.Id,
			Num:         strconvEx.StrTry2Int(num, 0),
			BetM:        strconvEx.StrTry2Int(cash, 0),
			Stock:       Stock,
			LotteryT:    LotteryT,
		}
		StrNum := fmt.Sprintf("%d", aNumData.Num)
		mpData[StrNum] = aNumData
	}
	return mpData
}

func (this *Cs2Lottery) SaveNumData(v map[string]interface{}) (map[string]interface{}, error) {
	if this.isSaveNumData {
		return nil, nil
	}
	this.isSaveNumData = true
	defer func() {
		this.isSaveNumData = false
	}()
	aOriginalQueryDataP := this.getOriginalQueryDataP(v)
	mpGetOriginalData := make(map[string]interface{})
	if this.UserInfo.IsLogin == 0 {
		return nil, fmt.Errorf("请登录")
	}
	mpGetOriginalData["OriginalQueryDataP"] = aOriginalQueryDataP
	mpOriginalData, e := this.GetOriginalData(mpGetOriginalData)
	if e != nil {
		if !this.CheckLogin(v) {
			this.UpdateLoginState2Db(false)
		}
		return nil, e
	}

	mpP := make(map[string]interface{})
	mpP["Data"] = mpOriginalData["Data"]
	mpP["OriginalQueryDataP"] = aOriginalQueryDataP
	mpResultData := this.ToMapNumData(mpP)
	o := orm.NewOrm()
	for _, data := range mpResultData {
		data.AddUpdate(o)
	}
	return nil, nil
}

func (this *Cs2Lottery) CheckToken(v map[string]interface{}) (map[string]interface{}, error) {
	mpData := make(map[string]interface{})
	//aLotteryData := this.GetLotteryData(this.GetCurLotteryType())
	//
	////strUrl := `https://mye29wvi-cs2.cs111.vip/index.php?controller=admin_index&action=FindOpenInfo&code=bjamlh`
	//
	//strUrl := fmt.Sprintf(
	//	`%s/index.php?controller=admin_index&action=FindOpenInfo&code=%s`,
	//	this.UserInfo.SiteUrl, aLotteryData.LotCode)
	//data, e := this.GetHttp().GetString(strUrl)
	//if e != nil {
	//	return mpData, e
	//}
	//strHint := stringKit.GetBetweenStr(data, "alert('", "')")
	//if strHint != "" {
	//	return mpData, fmt.Errorf("%s", strHint)
	//}

	return mpData, nil
}

func (this *Cs2Lottery) TestBaseUrl(url string) (map[string]interface{}, error) {
	mpData, e := this.getCode(url, nil)
	if e != nil {
		return nil, fmt.Errorf("Url有问题1。")
	}

	data, ok := mpData["CodePic"].([]byte)
	if !ok {
		return nil, fmt.Errorf("Url有问题2。")
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("Url有问题3。")
	}

	return nil, nil
}

func (this *Cs2Lottery) GetInitLotteryData(t int64) *Cs2LotteryData {
	aLotteryData := &Cs2LotteryData{}
	aLotteryData.LotteryType = t
	mpData := this.getInitData(t)
	aLotteryData.MpNum2ItemId = mpData[key_Num2ItemId].(map[int]string)

	switch t {
	case base.LotteryType_1:
		aLotteryData.OriginalDataGUID = `Y29udHJvbGxlcj1nYW1yYXRlJmFjdGlvbj1zaG93cmF0ZQ==`
		aLotteryData.CountGUID = `Y29udHJvbGxlcj1yZXBvcnQmYWN0aW9uPWRldGFpbGJ5aXRlbSZvcGVuX2lkPTc3OTM2MDgxJmlzX3N0b2NrPTEmbG90c2hvdz0xJmxvdGNvZGU9YmphbWxoX3Rt`
		aLotteryData.LotCode = "bjamlh"
		aLotteryData.BetLotCode = "bjamlh_tm2"
		aLotteryData.ItemLids = "1974"
		aLotteryData.Is_stock = 0
		aLotteryData.StrLotteryType = aLotteryData.LotCode
	case base.LotteryType_2:
		aLotteryData.OriginalDataGUID = `Y29udHJvbGxlcj1nYW1yYXRlJmFjdGlvbj1zaG93cmF0ZQ==`
		aLotteryData.CountGUID = `Y29udHJvbGxlcj1yZXBvcnQmYWN0aW9uPWRldGFpbGJ5aXRlbSZvcGVuX2lkPTc3OTU3NDM3JmlzX3N0b2NrPTEmbG90c2hvdz0xJmxvdGNvZGU9eGFtbGhfdG0=`
		aLotteryData.LotCode = "xamlh"
		aLotteryData.BetLotCode = "xamlh_tm2"
		aLotteryData.ItemLids = "1974"
		aLotteryData.Is_stock = 0
		aLotteryData.StrLotteryType = aLotteryData.LotCode
	case base.LotteryType_3:
		aLotteryData.OriginalDataGUID = `Y29udHJvbGxlcj1nYW1yYXRlJmFjdGlvbj1zaG93cmF0ZQ==`
		aLotteryData.CountGUID = `Y29udHJvbGxlcj1yZXBvcnQmYWN0aW9uPWRldGFpbGJ5aXRlbSZvcGVuX2lkPTc4MDY1NjI2JmlzX3N0b2NrPTEmbG90c2hvdz0xJmxvdGNvZGU9c3BsaF90bQ==`
		aLotteryData.LotCode = "splh"
		aLotteryData.BetLotCode = "splh_tm2"
		aLotteryData.ItemLids = "1974"
		aLotteryData.Is_stock = 0
		aLotteryData.StrLotteryType = aLotteryData.LotCode
	case base.LotteryType_4:
		aLotteryData.OriginalDataGUID = `Y29udHJvbGxlcj1nYW1yYXRlJmFjdGlvbj1zaG93cmF0ZQ==`
		aLotteryData.CountGUID = `Y29udHJvbGxlcj1yZXBvcnQmYWN0aW9uPWRldGFpbGJ5aXRlbSZvcGVuX2lkPTc4MDY1OTc5JmlzX3N0b2NrPTEmbG90c2hvdz0xJmxvdGNvZGU9bGdsaF90bQ==`
		aLotteryData.LotCode = "lglh"
		aLotteryData.BetLotCode = "lglh_tm2"
		aLotteryData.ItemLids = "1974"
		aLotteryData.Is_stock = 0
		aLotteryData.StrLotteryType = aLotteryData.LotCode
	}

	return aLotteryData
}

func (this *Cs2Lottery) GetLotteryData(t int64) *Cs2LotteryData {
	this.TypeDataLock.Lock()
	defer this.TypeDataLock.Unlock()

	aLotteryData, ok := this.mpLotteryData[t]
	if ok {
		return aLotteryData
	} else {
		aLotteryData = this.GetInitLotteryData(t)
		this.mpLotteryData[t] = aLotteryData
		return aLotteryData
	}
}

func (this *Cs2Lottery) getIssueResult(config models.LoConfig) (box.GetCurrentInstall, error) {
	aResult := box.GetCurrentInstall{}
	aLotteryData := this.GetLotteryData(config.LotteryType)

	strUrl := fmt.Sprintf(
		`%s/index.php?controller=admin_index&action=FindOpenInfo&code=%s`,
		this.UserInfo.SiteUrl, aLotteryData.LotCode)

	datas, e := this.GetHttp().GetString(strUrl)
	if e != nil {
		ttLog.LogError("获取数据当前结果出错。E：", e)
		return aResult, e
	}

	e = stringKit.GetJsonObj(datas, &aResult)
	return aResult, e
}

func (this *Cs2Lottery) CheckLogin(v map[string]interface{}) bool {
	config := CacheData.GetLoConfig()
	_, e := this.getIssueResult(config)
	if e != nil {
		return false
	}
	return true
}
