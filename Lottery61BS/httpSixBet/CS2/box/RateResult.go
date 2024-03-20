package box

type RateResult struct {
	Code string          `json:"code"`
	Msg  map[int]float64 `json:"msg"`
}
