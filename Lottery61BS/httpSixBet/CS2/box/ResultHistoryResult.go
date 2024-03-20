package box

type ResultHistoryResult struct {
	Data ResultData `json:"data"`
}
type ResultData struct {
	List []InstallmentsItem `json:"list"`
}

type InstallmentsItem struct {
	CycleId  string      `json:"cycleid"`
	OpenTime interface{} `json:"open_time"`
	B1       string      `json:"b1"`
	B2       string      `json:"b2"`
	B3       string      `json:"b3"`
	B4       string      `json:"b4"`
	B5       string      `json:"b5"`
	B6       string      `json:"b6"`
	B7       string      `json:"b7"`
}
