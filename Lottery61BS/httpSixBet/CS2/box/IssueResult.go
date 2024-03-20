package box

type IssueResult struct {
	CycleId   string      `json:"cycleid"`
	Msg       string      `json:"msg"`
	CloseTime interface{} `json:"close_time"`
	//NCloseTime  string `json:"nclosttime"`
	//LCloseTime  string `json:"L_close_time"`
	//LCloseTime2 string `json:"L_close_time2"`
	//LOpenTime   string `json:"L_open_time"`
}

//type IssueData struct {
//	CycleId   string `json:"cycleid"`
//	LotteryId string `json:"lottery_id"`
//}
