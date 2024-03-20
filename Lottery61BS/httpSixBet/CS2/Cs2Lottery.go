package CS2

import (
	"fmt"
	"github.com/TtMyth123/DengG/Lottery61BS/CacheData"
	"github.com/TtMyth123/DengG/Lottery61BS/httpSixBet/CS2/box"
	"github.com/TtMyth123/DengG/Lottery61BS/httpSixBet/base"
	"github.com/TtMyth123/DengG/Lottery61BS/httpSixBet/base/BetBo"
	"github.com/TtMyth123/DengG/Lottery61BS/models"
	"github.com/TtMyth123/DengG/Lottery61BS/models/mconst"
	"github.com/TtMyth123/kit/stringKit"
	"github.com/astaxie/beego/orm"
	"net/url"
	"strings"
	"time"
)

const (
	controller_str = "web_login"

	//澳门
	itemlids_OM = "1975"
	toplotid_OM = "23"
	open_id_OM  = "69787190"
	lotcode_OM  = "bjamlh"

	//新澳门
	itemlids_NOM = "1975"
	toplotid_NOM = "23"
	open_id_NOM  = "73250008"
	lotcode_NOM  = "xamlh"

	//超级6+1
	itemlids_CJ = "1974"
	toplotid_CJ = "217"
	open_id_CJ  = "72342151"
	lotcode_CJ  = "splh"

	//闪电6+1
	itemlids_SD = "1974"
	toplotid_SD = "231"
	open_id_SD  = "72343380"
	lotcode_SD  = "lglh"

	C_region = "A"
)

/*
*
https://mye29wvi-cs2.cs111.vip
cs2
账号：uc01
密码：qqq111
*/
type Cs2Lottery struct {
	*base.Lottery
	UUID              string
	greycdn_recaptcha string
	//CodeId           string
	//CurLotteryResult box.LotteryResult
	//
	//mpNum2Tid  map[int]int
	//mpTid2Rate map[int]float64
	//
	//CycleId string
	//ItemLIds string
	//TopLotId string
	//Region   string
	//Open_id  string
	//LotCode  string

	mpLotteryData map[int64]*Cs2LotteryData
}

type Cs2LotteryData struct {
	MpNum2Rate map[int]float64

	CodeId      string
	SettingCode string
	LotteryType int64

	BetLotCode string
	LotCode    string
	Open_id    string

	ItemLIds string
	TopLotId string
	Region   string

	MpNum2Tid  map[int]int
	MpTid2Rate map[int]float64

	CycleId string
}

const (
	key_Num2Tid  = "Num2Tid"
	key_Tid2Rate = "Tid2Rate"
)
const default_oddsd = 47.4

func NewLottery(aSiteUserInfo models.LoBetUser, isRun bool) *Cs2Lottery {
	aServer := &Cs2Lottery{}
	aServer.Lottery = base.NewLottery(aSiteUserInfo)

	aServer.Lottery.UpdateCurIssueF = aServer.UpdateCurIssue
	aServer.Lottery.UpdateOddsF = aServer.UpdateOdds
	aServer.Lottery.CheckLoginF = aServer.CheckLogin

	aServer.mpLotteryData = map[int64]*Cs2LotteryData{}

	return aServer
}

func (this *Cs2Lottery) getInitData_CJ61() map[string]interface{} {
	mpResult := make(map[string]interface{})
	mpNum2Tid := make(map[int]int)
	mpTid2Rate := make(map[int]float64)

	iTid := 64080
	for i := 1; i <= 49; i++ {
		mpNum2Tid[i] = iTid
		mpTid2Rate[iTid] = default_oddsd
		iTid++
	}
	mpResult[key_Num2Tid] = mpNum2Tid
	mpResult[key_Tid2Rate] = mpTid2Rate

	return mpResult
}
func (this *Cs2Lottery) getInitData_SD61() map[string]interface{} {
	return this.getInitData_CJ61()
}

func (this *Cs2Lottery) getInitDataOM() map[string]interface{} {
	mpResult := make(map[string]interface{})
	mpNum2Tid := make(map[int]int)
	mpTid2Rate := make(map[int]float64)

	iTid := 64129
	for i := 1; i <= 49; i++ {
		mpNum2Tid[i] = iTid
		mpTid2Rate[iTid] = default_oddsd
		iTid++
	}
	mpResult[key_Num2Tid] = mpNum2Tid
	mpResult[key_Tid2Rate] = mpTid2Rate

	return mpResult
}

func (this *Cs2Lottery) getInitDataNewOM() map[string]interface{} {
	mpResult := make(map[string]interface{})
	mpNum2Tid := make(map[int]int)
	mpTid2Rate := make(map[int]float64)

	iTid := 64129
	for i := 1; i <= 49; i++ {
		mpNum2Tid[i] = iTid
		mpTid2Rate[iTid] = default_oddsd
		iTid++
	}
	mpResult[key_Num2Tid] = mpNum2Tid
	mpResult[key_Tid2Rate] = mpTid2Rate

	return mpResult
}

func (this *Cs2Lottery) getInitData(t int64) map[string]interface{} {
	return this.getInitDataOM()
}

func (this *Cs2Lottery) getCode(url1 string, v map[string]interface{}) (map[string]interface{}, error) {
	mpData := make(map[string]interface{})
	//strUrl := `https://wy34n50v57-cs2.cs111.vip/index.php?controller=web_login&action=ImgCode&0.25038928799333116`
	d := time.Now().UnixMicro()
	strUrl := fmt.Sprintf(`%s/index.php?controller=%s&action=ImgCode&0.%d`, url1, controller_str, d)
	datas, e := this.GetHttp().GetBytes(strUrl)
	mpData["CodePic"] = datas
	strData := string(datas)
	ii := strings.Index(strData, "I am not robot")
	if ii > 0 {
		greycdn_recaptcha := `03AFcWeA5sFU5XJCTAijlOFC1GEYa9UZFo5y0DeYB6a7nFEK1s8NDCTPSMB8MsB1XWSD4McqszuFAqMnMGPvxfMCDugAy1WIJea8R4fektiryVyvU-ngaq5oRCLPJI_S_KxgJk_3zAqLoZJeXWj4PCuJXYN0PyiyP5M-1TM8fIeFzQvz-n8lTimIUNq2Te1HZpicr_2EEWY2kBknP6sqoJ0ioBOY3sK2M7u-zEyDm9oc7Ce_vdF5hTyf_DqR4VgY0b0Igo4uilZfdHEaZ5BM4XI5RExoET2YMRR67uzzm-jL28WBObAbeZoyESh-BH2zYag8Q3i2noNMuS4Iyqgas6T_GjJ_0U9KmHQZWLG4ft1gwaXlrHFQ9VKAZXXzdRQTKjL21rpk0FkbwqmteyXbUkySgs5KnWchIcOSZRlRNssc9S181QfWh7G9ZS2hzkmL7fPiYxIwicFdqfcVGHc4wBB0IRjxJNN5TSbUTc8Zjrlc3_rBP4YhzVv8utHuXLeqB2exnO3-znYNSUyFCyS5p3hJWXgntBo4LKdB-FWsd82LirrE8r14raYM2tDMcmutuJK8eeU6oF6zvj2DhyM3UIwab1Ad8z5400BFF0Wr-vwRc_G4dVksI1m-XMPE4BAOMBa45gJpIyj3_CfYwnJ1LpnlLFRKbJ3MI5mcORHyeFbNG8d1TEpYp6rzbwRReutXYBNDezHnVPOQB6sLhUQpNaEpsWO0xOHCAy8EyM7DPhQMX8A8Qk0MtO96pSSWh4y4L8j44C6BzKh5kC1sasmaFC1xoa0EKyD60QB4R7rvjt_cuF8qlSF1L6_0_ICctTUTX61BdLt-5J3evOYxLgZp0TnwjofSd2gCKl3IHbpG7riHW3wWLtfGuMVFPYglfzNFvb8quJoOO8vSKeaarWm_S9YnX6on2LR0e3CaQKmNS4FH4JyD9yv4WKCb8U84qHkOFvx_8P2NbE_Z4MSrE4OcWstNVJsglzrJ8VbI1F-odj-HemKFGRDFbMKGNppxaHEa62U6Tw-YdO3Jy5Jum-a3U2S01Z4tdr8yz4_bukLsgoyfeEcWWvP_NUMWeYZ9yfxuAJ9bgFnS2ZGKTm13miv9jDbamUAuGZ-2Os66JxpFPIozEM5IKW2QfmOwRsgKC_M40HMFEqfLvCmrraBA0zWbxKqzdXunCqLj8VxL7BwswKVn44tRccqfLX03KD2sbp7NJtfk-55_NOGt3HBF3MMf8czhx378JISxVTqK6dkQObcHN2TV_jAJTHofCvMa0SITP1TUemlZX0xux_euPpGPS2ug18bUCI1RAjCHyHjAys4_eOhV7fKUnaU0tHVsBuStUP6fALdLM_KjwhiZCQXoCUbzaTNSljIRt9R_96yQros79f7k1akxK-PGzbb9pO9V-t791r4Wl8zYJycbxCuCudRVnZQPJcKtIMuoG3pEToQwaVRqXVBR7eyQZiAwvV-jL4DIk1vJZ0m0aD9qFWbyG-M0CczbX_AAbBaxeHV3q6OAFsCqjO_xAsijRATm871BZReKJJKSqsAGHXAhEAfFE7VOJZEwEuwCgcsWYQq8kErjGQbeUnVAWIuDS4JfDdxK_sKyWOKVWbFPWpJH9i1EoLE96X3PMDU6wKd7lksy0TaWk48WSGCHw2pks3rc_-m3reqmtUEOE-G01VcrYIh1uiI1qKEIIUFDQYsM2EY4xeJWmSnOwvPROsvOsXlTISiXHyxyYh6RiMemfq7IJKcoPYgcXy01rTXfKsOR4bfBIxgnkoBsd-Rpg3tbFXDYxjXnRZirqXCo3JlpKjzYX421JxaVG9F0KM5p7xY93fz6zwjo_lHwL_4O2l8oD7hbm9pAg_rIge640RJQMcF_aivEXYbZNQ3aBKRjJS8ZVzQFRtQnymF3KDmVjD5dWxFwK2HMzsLhWSMluCLTTyepCPwsgsVvyc3ZeJbYg-Yu3_Gr3tTbm56G1DhZD6YGQ_H5zYDw51dgoDZXZwq-xT`
		strUrl1 := fmt.Sprintf(
			`%s/index.php?__greycdn_recaptcha=%s&controller=%s&action=ImgCode&0.%d`,
			url1, greycdn_recaptcha, controller_str, d)

		datas1, _ := this.GetHttp().GetBytes(strUrl1)
		this.greycdn_recaptcha = greycdn_recaptcha
		mpData["CodePic"] = datas1
	} else {
		this.greycdn_recaptcha = ""
	}

	return mpData, e
}

func (this *Cs2Lottery) getUrlRecaptcha(strUrl string) string {
	if this.greycdn_recaptcha != "" {
		if strings.Index(strUrl, "__greycdn_recaptcha") < 0 {
			strUrl = strUrl + "&__greycdn_recaptcha=" + this.greycdn_recaptcha
		}
	}
	return strUrl
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
	strUrl := fmt.Sprintf(`%s/index.php?controller=%s&action=login`, this.UserInfo.SiteUrl, controller_str)
	strUrl = this.getUrlRecaptcha(strUrl)
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

	this.UpdateCurIssue(nil)
	this.UpdateOdds(nil)

	return mpData, e
}
func (this *Cs2Lottery) LoginOut(v map[string]interface{}) (map[string]interface{}, error) {
	mpData := make(map[string]interface{})
	//strUrl := `https://mye29wvi-cs2.cs111.vip/index.php?%s`
	strUrl := fmt.Sprintf(`%s/index.php?index.php?%s`, this.UserInfo.SiteUrl, this.UUID)
	strUrl = this.getUrlRecaptcha(strUrl)
	_, e := this.GetHttp().GetString(strUrl)

	this.UpdateLoginState2Db(false)

	return mpData, e
}

func (this *Cs2Lottery) getRate(itemlids, toplotid, Region string) (map[int]float64, error) {

	mpResult := make(map[int]float64)

	//strUrl := `https://wy34n50v57-cs2.cs111.vip/index.php?controller=web_lotsgam&action=ItemsRate&itemlids=1975&toplotid=23&region=A`
	strUrl := fmt.Sprintf(`%s/index.php?controller=web_lotsgam&action=ItemsRate&itemlids=%s&toplotid=%s&region=%s`,
		this.UserInfo.SiteUrl, itemlids, toplotid, Region)
	strUrl = this.getUrlRecaptcha(strUrl)
	datas, e := this.GetHttp().GetString(strUrl)
	if e != nil {
		return mpResult, e
	}
	aRateResult := box.RateResult{}
	e = stringKit.GetJsonObj(datas, &aRateResult)
	if e != nil {
		return mpResult, e
	}

	return aRateResult.Msg, e
}

func (this *Cs2Lottery) getLoConfig(v map[string]interface{}) models.LoConfig {
	aLoConfig := models.LoConfig{}
	if v != nil {
		if v["LoConfig"] != nil {
			aLoConfig = v["LoConfig"].(models.LoConfig)
		}
	}
	aLoConfig = CacheData.GetLoConfig()
	return aLoConfig
}

func (this *Cs2Lottery) UpdateOdds(v map[string]interface{}) (map[string]interface{}, error) {
	aLoConfig := this.getLoConfig(v)
	aLotteryData := this.GetLotteryData(aLoConfig.LotteryType)
	mpRate, e := this.getRate(aLotteryData.ItemLIds, aLotteryData.TopLotId, aLotteryData.Region)
	if e != nil {
		return nil, e
	}
	for tid, rate := range mpRate {
		aLotteryData.MpTid2Rate[tid] = rate
	}

	return nil, nil
}
func (this *Cs2Lottery) UpdateCurIssue(v map[string]interface{}) (map[string]interface{}, error) {
	config := this.getLoConfig(v)
	r, e := this.getIssueResult(config)
	if e != nil {
		return nil, e
	}
	mpResult := make(map[string]interface{})
	mpResult["Data"] = r

	return mpResult, nil
}
func (this *Cs2Lottery) FormatCycleId(CycleId string) string {
	newCycleId := strings.Replace(CycleId, "-", "", -1)
	return newCycleId
}
func (this *Cs2Lottery) loginOKinit() {
	this.UpdateCurIssue(nil)
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

func (this *Cs2Lottery) GetBaseData(v map[string]interface{}) (map[string]interface{}, error) {
	config := this.getLoConfig(v)
	aLotteryData := this.GetLotteryData(config.LotteryType)

	mpData := make(map[string]interface{})
	mpData["CycleId"] = aLotteryData.CycleId
	mpData["mpTid2Rate"] = aLotteryData.MpTid2Rate
	return mpData, nil
}

func (this *Cs2Lottery) Bet(v map[string]interface{}) (map[string]interface{}, error) {
	arrBetData := make([]BetBo.BetData, 0)
	aLoConfig := this.getLoConfig(v)
	if v["StrBetData"] != "" {
		StrBetData := v["StrBetData"].(string)
		e := stringKit.GetJsonObj(StrBetData, &arrBetData)
		if e != nil {
			return nil, e
		}
	}

	data, e := this.bet(arrBetData, aLoConfig, 0)
	return data, e
}

func (this *Cs2Lottery) getIssueResult(config models.LoConfig) (box.IssueResult, error) {
	aResult := box.IssueResult{}
	aLotteryData := this.GetLotteryData(config.LotteryType)
	strUrl := fmt.Sprintf(
		`%s/index.php?controller=dsn_index&action=GetNextOpeninfo&is_ajax=1&lotcode=%s&open_id=%s`,
		this.UserInfo.SiteUrl, aLotteryData.LotCode, aLotteryData.Open_id)
	strUrl = this.getUrlRecaptcha(strUrl)
	datas, e := this.GetHttp().GetString(strUrl)
	if e != nil {
		return aResult, e
	}

	e = stringKit.GetJsonObj(datas, &aResult)
	if e == nil {
		aLotteryData.CycleId = aResult.CycleId

		CycleId := this.FormatCycleId(aResult.CycleId)
		CacheData.SaveLotteryType_CycleId(config.LotteryType, CycleId)
	}
	return aResult, e
}
func (this *Cs2Lottery) bet(arrData []BetBo.BetData, config models.LoConfig, ii int) (map[string]interface{}, error) {
	ii++
	aLotteryData := this.GetLotteryData(config.LotteryType)

	IssueResult, e := this.getIssueResult(config)
	if e != nil {
		return nil, e
	}

	rows, e := this.getBetData(arrData, config)
	if e != nil {
		return nil, e
	}
	mpParam := url.Values{}
	mpParam.Set("cycleid", IssueResult.CycleId)
	mpParam.Set("rows", rows)
	mpParam.Set("lotcode", aLotteryData.BetLotCode)
	mpParam.Set("lastwebos", "")

	strUrl := fmt.Sprintf(
		`%s/index.php?controller=Mob_CreateOrderM&action=Suborder&is_ajax=json`,
		this.UserInfo.SiteUrl)
	strUrl = this.getUrlRecaptcha(strUrl)
	datas, e := this.GetHttp().PostForm(strUrl, mpParam)
	if e != nil {
		return nil, e
	}
	aBetResult := box.BetResult{}
	e = stringKit.GetJsonObj(datas, &aBetResult)

	if e != nil {
		aBetResultS := box.BetResultS{}
		e = stringKit.GetJsonObj(datas, &aBetResultS)
		if e == nil {
			if aBetResultS.Code != "ok" {
				e = fmt.Errorf("%s", aBetResultS.Msg)
			}
		}
		return nil, e
	}
	if aBetResult.Code == "205" {
		this.UpdateOdds(nil)
		return this.bet(arrData, config, ii)
	}
	if aBetResult.Code != "ok" {
		return nil, fmt.Errorf("%s", aBetResult.Msg)
	}
	CycleId := this.FormatCycleId(IssueResult.CycleId)
	this.AddLoBetBetData(CycleId, arrData, config)
	return nil, e
}
func (this *Cs2Lottery) AddLoBetBetData(CycleId string, arrData []BetBo.BetData, config models.LoConfig) {
	o := orm.NewOrm()
	for _, betI := range arrData {
		aLoBetBetData := models.LoBetBetData{
			CycleId:  CycleId,
			UserId:   this.UserInfo.Id,
			Num:      betI.Num,
			BetM:     betI.V,
			LotteryT: config.LotteryType,
		}
		aLoBetBetData.SysId = this.UserInfo.SysId
		e := aLoBetBetData.AddUpdate(o)
		if e != nil {
			fmt.Println(e)
		}
	}
}

func (this *Cs2Lottery) getBetData(arrData []BetBo.BetData, config models.LoConfig) (string, error) {
	aLotteryData := this.GetLotteryData(config.LotteryType)
	rowData := make([]box.BetRowInfo, 0)
	for _, betItem := range arrData {
		if betItem.Num < 1 || betItem.Num > 49 {
			continue
		}
		if betItem.V < 1 {
			continue
		}
		tid := aLotteryData.MpNum2Tid[betItem.Num]
		Rate := aLotteryData.MpTid2Rate[tid]
		aBetRowInfo := box.BetRowInfo{
			Tid:   fmt.Sprintf("%d", tid),
			Cash:  fmt.Sprintf("%d", betItem.V),
			Balls: fmt.Sprintf("%02d", betItem.Num),
			Rate:  fmt.Sprintf("%g", Rate),
		}
		rowData = append(rowData, aBetRowInfo)
	}
	strRow := stringKit.GetJsonStr(rowData)
	return strRow, nil
}

//func (this *Cs2Lottery) bet1(v map[string]interface{}, aBetSiteData *models.BetSiteData, ii int, err error) (map[string]interface{}, error) {
//	mpData := make(map[string]interface{})
//	aLoConfig := CacheData.GetLoConfig()
//	//if this.SiteUserInfo.IsBet != mconst.SiteUserInfo_IsBet || this.SiteUserInfo.State != mconst.SiteUserInfo_State_1_OK {
//	//	return mpData, nil
//	//}
//
//	aLotteryData := this.GetLotteryData(aLoConfig.LotteryType)
//
//	//mpBet := v["BetData"].(map[int]int)
//	//BetMainId := v["BetMainId"].(int64)
//	//NumOrder := v["NumOrder"].(int)
//	//ii++
//	//if aBetSiteData == nil {
//	//	aBetSiteData = &models.BetSiteData{}
//	//}
//	//if ii > this.BetCount {
//	//	aBetSiteData.State = mconst.BetSiteData_State_Err
//	//	if err != nil {
//	//		aBetSiteData.ErrMsg = err.Error()
//	//	}
//	//	aBetSiteData.Add(nil)
//	//	this.UpdateDataState(fmt.Sprintf("单号(%05d)投注失败：%s", NumOrder, aBetSiteData.ErrMsg))
//	//	return mpData, err
//	//}
//	//
//	//aBetSiteData.MainId = BetMainId
//	//aBetSiteData.SysId = this.SiteUserInfo.SysId
//	//aBetSiteData.LotteryT = aLotteryData.LotteryType
//	//aBetSiteData.SiteId = this.SiteUserInfo.Id
//	//aBetSiteData.State = mconst.BetSiteData_State_OK
//	//aBetSiteData.Rate = this.SiteUserInfo.Rate
//	//aBetSiteData.CycleId = this.SiteUserInfo.CycleId
//	//
//	////strUrl := `https://wy34n50v57-cs2.cs111.vip/index.php?controller=Mob_CreateOrderM&action=Suborder&is_ajax=json`
//	//strUrl := fmt.Sprintf(`%s/index.php?controller=Mob_CreateOrderM&action=Suborder&is_ajax=json`, this.UserInfo.SiteUrl)
//	//
//	//if aLotteryData.CycleId == "" {
//	//	aBetSiteData.State = mconst.BetSiteData_State_Err
//	//	aBetSiteData.ErrMsg = "期号未获取成功"
//	//	aBetSiteData.Add(nil)
//	//	this.UpdateDataState(fmt.Sprintf("单号(%05d)投注失败：%s", NumOrder, aBetSiteData.ErrMsg))
//	//	return mpData, fmt.Errorf("期号未获取成功")
//	//}
//	//
//	//rows, strBetInfo, OldBetMoney, BetMoney, e := this.getBetData(mpBet, *aBetSiteData)
//	//aBetSiteData.BetInfo = strBetInfo
//	//aBetSiteData.BetMoney = BetMoney
//	//aBetSiteData.OldBetMoney = OldBetMoney
//	//if e != nil {
//	//	aBetSiteData.State = mconst.BetSiteData_State_Err
//	//	aBetSiteData.ErrMsg = e.Error()
//	//	aBetSiteData.Add(nil)
//	//	this.UpdateDataState(fmt.Sprintf("单号(%05d)投注失败：%s", NumOrder, aBetSiteData.ErrMsg))
//	//	return mpData, e
//	//}
//	//
//	//mpParam := url.Values{}
//	//mpParam.Set("cycleid", aLotteryData.CycleId)
//	//mpParam.Set("rows", rows)
//	//mpParam.Set("lotcode", fmt.Sprintf("%s_tm", aLotteryData.LotCode))
//	//mpParam.Set("lastwebos", "")
//	////{"code":"205","msg":"该项赔率已变动，请核对后再下注!","data":{"64129":{"post_rate":"47","rate":48},"qstr":""}}
//	////{"code":"302","msg":"下注失败:账户余额不足!","data":[]}
//	////{"code":"ok","msg":"操作成功","data":{"orders":[{"order_id":"752416748","toplotid":"23","lotcode":"bjamlh_tm","cycleid":"2023-006","open_id":"69684424","show_id":"r2006151519785145","depiction":"02","user_id":"226637","item_id":"64130","lottery_id":"1975","rate":48,"cash":"1.00","balls":"02","balls_type":"T","is_win":"0","region":"A","created":"1672989319","stype":"0","state":"0","outcome":"0","stk6":"-100","rbt6":"1.93","rbs6":"0.0193","lotname":"特码B","tname":"02","username":"ao88","ushowname":"本级","group_id":"6","result":0,"result2":47.0193,"zfresult":0,"ustock":-100,"urebate":0.02,"ucash":-100,"desc":"澳门六合彩 2023-006 期  <span class=\"red\">特码B<\/span>  02","toplotname":"澳门六合彩","zfrebate":-1.93,"created2":"01-06 15:15:19","cycleid_ds":"d","isautozf":0},{"order_id":"752416751","toplotid":"23","lotcode":"bjamlh_tm","cycleid":"2023-006","open_id":"69684424","show_id":"r200615151959061715","depiction":"03","user_id":"226637","item_id":"64131","lottery_id":"1975","rate":48,"cash":"1.00","balls":"03","balls_type":"T","is_win":"0","region":"A","created":"1672989319","stype":"0","state":"0","outcome":"0","stk6":"-100","rbt6":"1.93","rbs6":"0.0193","lotname":"特码B","tname":"03","username":"ao88","ushowname":"本级","group_id":"6","result":0,"result2":47.0193,"zfresult":0,"ustock":-100,"urebate":0.02,"ucash":-100,"desc":"澳门六合彩 2023-006 期  <span class=\"red\">特码B<\/span>  03","toplotname":"澳门六合彩","zfrebate":-1.93,"created2":"01-06 15:15:19","cycleid_ds":"d","isautozf":0}],"sum":{"cnt":2,"sum":2},"cycleid":"澳门六合彩[2023-006]"}}
//	//datas, e := this.GetHttp().PostForm(strUrl, mpParam)
//	//if e != nil {
//	//	models.AddBetErrInfo(aBetSiteData.MainId, aBetSiteData.SiteId, ii, aLotteryData.CycleId,
//	//		stringKit.GetJsonStr(v), stringKit.GetJsonStr(datas), stringKit.GetJsonStr(e))
//	//
//	//	aBetSiteData.State = mconst.BetSiteData_State_Err
//	//	aBetSiteData.ErrMsg = e.Error()
//	//	aBetSiteData.Add(nil)
//	//	this.UpdateDataState(fmt.Sprintf("单号(%05d)投注失败：%s", NumOrder, aBetSiteData.ErrMsg))
//	//	return mpData, e
//	//	//return this.bet(v, aBetSiteData, ii, e)
//	//}
//	//
//	//aBetResult := box.BetResult{}
//	//e = stringKit.GetJsonObj(datas, &aBetResult)
//	//if e != nil {
//	//	aBetResultS := box.BetResultS{}
//	//	e = stringKit.GetJsonObj(datas, &aBetResultS)
//	//	if e == nil {
//	//		if aBetResultS.Code != "ok" {
//	//			e = fmt.Errorf("%s", aBetResultS.Msg)
//	//		}
//	//	}
//	//
//	//	models.AddBetErrInfo(aBetSiteData.MainId, aBetSiteData.SiteId, ii, aLotteryData.CycleId,
//	//		stringKit.GetJsonStr(v), stringKit.GetJsonStr(datas), stringKit.GetJsonStr(e))
//	//
//	//	aBetSiteData.State = mconst.BetSiteData_State_Err
//	//	aBetSiteData.ErrMsg = e.Error()
//	//	aBetSiteData.Add(nil)
//	//	this.UpdateDataState(fmt.Sprintf("单号(%05d)投注失败：%s", NumOrder, aBetSiteData.ErrMsg))
//	//	return mpData, e
//	//}
//	//if aBetResult.Code == "205" {
//	//	models.AddBetErrInfo(aBetSiteData.MainId, aBetSiteData.SiteId, ii, aLotteryData.CycleId,
//	//		stringKit.GetJsonStr(v), stringKit.GetJsonStr(aBetResult), `aBetResult.Code == "205"`)
//	//
//	//	this.UpdateOdds(nil)
//	//	return this.bet(v, aBetSiteData, ii, fmt.Errorf("%s", aBetSiteData.ErrMsg))
//	//}
//	//if aBetResult.Code != "ok" {
//	//	models.AddBetErrInfo(aBetSiteData.MainId, aBetSiteData.SiteId, ii, aLotteryData.CycleId,
//	//		stringKit.GetJsonStr(v), stringKit.GetJsonStr(aBetResult), `aBetResult.Code != "ok"`)
//	//
//	//	aBetSiteData.State = mconst.BetSiteData_State_Err
//	//	aBetSiteData.ErrMsg = aBetResult.Msg
//	//	aBetSiteData.Add(nil)
//	//	this.UpdateDataState(fmt.Sprintf("单号(%05d)投注失败：%s", NumOrder, aBetSiteData.ErrMsg))
//	//	return mpData, fmt.Errorf("%s", aBetResult.Msg)
//	//}
//	//
//	//aBetSiteData.State = mconst.BetSiteData_State_OK
//	//aBetSiteData.OkMsg = stringKit.GetJsonStr(aBetResult.Data)
//	//this.UpdateDataState(fmt.Sprintf("单号(%05d)投注成功", NumOrder))
//	//e = aBetSiteData.Add(nil)
//	return mpData, nil
//}

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

func (this *Cs2Lottery) GetInitLotteryData(t int64) *Cs2LotteryData {
	aLotteryData := &Cs2LotteryData{}
	aLotteryData.LotteryType = t
	switch t {
	case mconst.LotteryType_1:
		aLotteryData.ItemLIds = itemlids_OM
		aLotteryData.TopLotId = toplotid_OM
		aLotteryData.Region = C_region
		aLotteryData.Open_id = open_id_OM
		aLotteryData.LotCode = lotcode_OM
		aLotteryData.BetLotCode = fmt.Sprintf("%s_tm", aLotteryData.LotCode)
		mpData := this.getInitData(t)
		aLotteryData.MpNum2Tid = mpData[key_Num2Tid].(map[int]int)
		aLotteryData.MpTid2Rate = mpData[key_Tid2Rate].(map[int]float64)
	case mconst.LotteryType_2:
		aLotteryData.ItemLIds = itemlids_NOM
		aLotteryData.TopLotId = toplotid_NOM
		aLotteryData.Region = C_region
		aLotteryData.Open_id = open_id_NOM
		aLotteryData.LotCode = lotcode_NOM
		aLotteryData.BetLotCode = fmt.Sprintf("%s_tm", aLotteryData.LotCode)

		mpData := this.getInitData(t)
		aLotteryData.MpNum2Tid = mpData[key_Num2Tid].(map[int]int)
		aLotteryData.MpTid2Rate = mpData[key_Tid2Rate].(map[int]float64)

	case mconst.LotteryType_3:
		aLotteryData.ItemLIds = itemlids_CJ
		aLotteryData.TopLotId = toplotid_CJ
		aLotteryData.Region = C_region
		aLotteryData.Open_id = open_id_CJ
		aLotteryData.LotCode = lotcode_CJ
		aLotteryData.BetLotCode = fmt.Sprintf("%s_tm", aLotteryData.LotCode)

		mpData := this.getInitData(t)
		aLotteryData.MpNum2Tid = mpData[key_Num2Tid].(map[int]int)
		aLotteryData.MpTid2Rate = mpData[key_Tid2Rate].(map[int]float64)

	case mconst.LotteryType_4:
		aLotteryData.ItemLIds = itemlids_SD
		aLotteryData.TopLotId = toplotid_SD
		aLotteryData.Region = C_region
		aLotteryData.Open_id = open_id_SD
		aLotteryData.LotCode = lotcode_SD
		aLotteryData.BetLotCode = fmt.Sprintf("%s_tm", aLotteryData.LotCode)

		mpData := this.getInitData(t)
		aLotteryData.MpNum2Tid = mpData[key_Num2Tid].(map[int]int)
		aLotteryData.MpTid2Rate = mpData[key_Tid2Rate].(map[int]float64)
	}

	return aLotteryData
}

func (this *Cs2Lottery) CheckLogin(v map[string]interface{}) bool {
	config := CacheData.GetLoConfig()
	_, e := this.getIssueResult(config)
	if e != nil {
		return false
	}
	return true
}
