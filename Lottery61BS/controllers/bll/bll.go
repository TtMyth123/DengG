package bll

func GetC(SysId int64) (interface{}, error) {
	//tt := time.Now()
	//if LotteryType == 0 {
	//	LotteryType = CacheData.GetLotteryType(SysId)
	//}
	//aCurCycleCloseTime, _ := CacheData.GetCurCycleCloseTime(LotteryType)
	//CloseT := aCurCycleCloseTime.Sub(tt) / time.Second
	//if CloseT < 0 {
	//	CloseT = 0
	//}
	aa := make(map[string]interface{})
	//aa["S"] = CloseT
	//aa["L"] = CacheData.GetCurCycleId(LotteryType)
	//aa["NumOrder"] = fmt.Sprintf("%05d", CacheData.GetNumOrder(LotteryType, SysId))
	//aa["HasLogin"], _ = CacheData.GetHasLogin(SysId)
	//aa["LotteryType"] = CacheData.GetLotteryType(SysId)
	return aa, nil
}
