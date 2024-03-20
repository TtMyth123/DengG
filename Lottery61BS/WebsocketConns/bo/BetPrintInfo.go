package bo

type BetPrintInfo struct {
	Title      string
	LotteryStr string
	CreatedAt  string
	BetSn      string
	DetailData []BetDetailInfo
}

type BetDetailInfo struct {
	OddsName string
	M        int
	Win      float64
}
