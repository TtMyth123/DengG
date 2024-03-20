package box

type BetRowInfo struct {
	Tid   string `json:"tid"`   //投注ID
	Balls string `json:"balls"` //投注内容
	Cash  string `json:"cash"`  //投注金额
	Rate  string `json:"rate"`  //投注赔率
}
