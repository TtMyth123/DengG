package CS2

//
//import (
//	"fmt"
//	"github.com/TtMyth123/SixLotteryServerHong/OM61/models"
//	"github.com/TtMyth123/SixLotteryServerHong/OM61/sixKit"
//	"github.com/TtMyth123/kit/stringKit"
//	"time"
//)
//
//func (this *Cs2Lottery) SaveNumData(v map[string]interface{}) (map[string]interface{}, error) {
//	t1 := time.Now()
//	isOK := true
//	strE := ""
//	arrABType := []string{"A", "B"}
//	arrDataType := []int{1, 2}
//	for _, abType := range arrABType {
//		for _, dataType := range arrDataType {
//			mpOriginalData, e := this.GetOriginalData(nil, abType, dataType)
//			if e != nil {
//				strE = e.Error()
//				isOK = false
//				continue
//			}
//			mpData := this.ToMapNumData(mpOriginalData["Data"], abType)
//			arrBetSumData := this.ToBetSumData(mpData, abType, dataType)
//			models.BatchBetSumData(arrBetSumData)
//		}
//	}
//	_, useT := sixKit.GetUseTime(t1, "")
//
//	if isOK {
//		this.UpdateDataState(fmt.Sprintf("获取数据成功(用时：%.2f)", useT))
//	} else {
//		this.UpdateDataState(fmt.Sprintf("获取数据有误(用时：%.2f):%s", useT, strE))
//	}
//
//	return nil, nil
//}
//
//func (this *Cs2Lottery) GetOriginalData(v map[string]interface{}, ABType string, DataType int) (map[string]interface{}, error) {
//	aLotteryData := this.GetLotteryData(this.GetCurLotteryType())
//
//	mpResult := make(map[string]interface{})
//	//lid := "1974"
//	//code := "bjamlh_tm2"
//	//if ABType == "A" {
//	//	lid = "1974"
//	//	code = "bjamlh_tm"
//	//} else if ABType == "B" {
//	//	lid = "1975"
//	//	code = "bjamlh_tm"
//	//} else if ABType == "O" {
//	//	lid = "1974"
//	//	code = "bjamlh_tm2"
//	//}
//	//oddstype：O,A,B,C,D
//
//	is_stock := 0
//	if DataType == 1 {
//		is_stock = 0
//	} else {
//		is_stock = 1
//	}
//
//	//strUrl := `https://mye29wvi-cs2.cs111.vip/index.php?Y29udHJvbGxlcj1nYW1yYXRlJmFjdGlvbj1zaG93cmF0ZQ==/&lid=1974&code=bjamlh_tm&&1=1&stypes=&selectsid=&is_stock=1&region=all&orderby=balls&is_ajax=1`
//	strUrl := fmt.Sprintf(`%s/index.php?Y29udHJvbGxlcj1nYW1yYXRlJmFjdGlvbj1zaG93cmF0ZQ==/&lid=%s&code=%s&&1=1&stypes=&selectsid=&is_stock=%d&region=all&orderby=balls&is_ajax=1`,
//		this.BaseUrl, aLotteryData.LotCode, aLotteryData.BetLotCode, is_stock)
//	data, e := this.GetHttp().GetString(strUrl)
//	if e != nil {
//		return mpResult, e
//	}
//	if data == "" {
//		this.UpdateState(false)
//		return mpResult, e
//	} else {
//		type Tmp struct {
//			Code string `json:"code"`
//			Msg  string `json:"msg"`
//		}
//		aTmp := Tmp{}
//		//{"code":"页面出错了，可能是权限不足，请核实后重新访问！","msg":"logout"}
//		e = stringKit.GetJsonObj(data, &aTmp)
//		if e == nil {
//			if aTmp.Msg == "logout" {
//				this.UpdateState(false)
//				return mpResult, fmt.Errorf(aTmp.Code)
//			}
//		}
//	}
//
//	//aResultData = GetStr2NumData(data)
//	mpResult["Data"] = data
//
//	return mpResult, nil
//}
