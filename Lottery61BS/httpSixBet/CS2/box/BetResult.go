package box

type BetResultS struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

type BetResult struct {
	Code string        `json:"code"`
	Msg  string        `json:"msg"`
	Data BetResultData `json:"data"`
}

type BetResultData struct {
	Orders []OrdersInfo `json:"orders"`
}

type OrdersInfo struct {
	ShowId string `json:"show_id"`
}
