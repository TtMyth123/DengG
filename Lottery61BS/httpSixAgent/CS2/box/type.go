package box

//type NumData struct {
//	StrNum string
//	Num    int
//	Rate   float64
//	Cash   float64
//	Result float64
//}

type GetCurrentInstall struct {
	Code string          `json:"code"`
	Msg  string          `json:"msg"`
	Data []LotteryResult `json:"data"`
}

type LotteryResult struct {
	Code      string        `json:"code"`
	Cycleid   string        `json:"cycleid"` //期号
	CloseTime interface{}   `json:"close_time"`
	PreData   PreDataResult `json:"predata"`
}
type PreDataResult struct {
	Cycleid   string      `json:"cycleid"` //期号
	CloseTime interface{} `json:"close_time"`
}

/*
{
	"code": "ok",
	"msg": "操作成功",
	"data": [{
		"cycleid": "2022-213",
		"code": "bjamlh",
		"att": "2",
		"class": "lhc",
		"showname": "澳门六合彩",
		"balls": "05,12,07,01,08,19,42",
		"ballsHtml": "<li><p class=\"greenball b05\">05<\/p><span>狗<\/span><li><p class=\"redball b12\">12<\/p><span>兔<\/span><li><p class=\"redball b07\">07<\/p><span>猴<\/span><li><p class=\"redball b01\">01<\/p><span>虎<\/span><li><p class=\"redball b08\">08<\/p><span>羊<\/span><li><p class=\"redball b19\">19<\/p><span>猴<\/span><li class=\"ballsjia\"><p>+<\/p><span>&nbsp;&nbsp;<\/span><li><li><p class=\"blueball b42\">42<\/p><span>鸡<\/span>",
		"close_time": "1659360600",
		"optime": 1659360780,
		"opdate": "2022-08-01 21:33:00",
		"predata": {
			"cycleid": "2022-214",
			"code": "bjamlh",
			"att": "0",
			"class": "lhc",
			"showname": "澳门六合彩",
			"balls": ",,,,,,",
			"ballsHtml": "<li><p class=\"emptyball b&nbsp;&nbsp;\">&nbsp;&nbsp;<\/p><span>肖<\/span><li><p class=\"emptyball b&nbsp;&nbsp;\">&nbsp;&nbsp;<\/p><span>肖<\/span><li><p class=\"emptyball b&nbsp;&nbsp;\">&nbsp;&nbsp;<\/p><span>肖<\/span><li><p class=\"emptyball b&nbsp;&nbsp;\">&nbsp;&nbsp;<\/p><span>肖<\/span><li><p class=\"emptyball b&nbsp;&nbsp;\">&nbsp;&nbsp;<\/p><span>肖<\/span><li><p class=\"emptyball b&nbsp;&nbsp;\">&nbsp;&nbsp;<\/p><span>肖<\/span><li class=\"ballsjia\"><p>+<\/p><span>&nbsp;&nbsp;<\/span><li><li><p class=\"emptyball b&nbsp;&nbsp;\">&nbsp;&nbsp;<\/p><span>肖<\/span>",
			"close_time": 1659360000,
			"optime": 1659447180,
			"opdate": "2022-08-02 21:33:00"
		}
	}],
	"stime": 1659372955
}
*/

type BetDetailData struct {
	Total int           `json:"total"`
	Data  []BetDataInfo `json:"data"`
}

type BetDataInfo struct {
	Id        string  `json:"Id"`
	BetAmount float64 `json:"betAmount"`
	Note      string  `json:"Note"`
}
