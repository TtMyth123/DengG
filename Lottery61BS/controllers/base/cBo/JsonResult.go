package cBo

type JsonResult struct {
	Code JsonResultCode `json:"code"`
	Msg  string         `json:"msg"`
	Obj  interface{}    `json:"obj"`
}

type ListResult struct {
	LastId   int
	ListData interface{}
}

type JsonListResult struct {
	Code JsonResultCode `json:"code"`
	Msg  string         `json:"msg"`
	Obj  ListResult     `json:"obj"`
}

type PageResult struct {
	PageTotal int
	DataEx    interface{}
	ListData  interface{}
	GroupData interface{}
}
type JsonPageResult struct {
	Code JsonResultCode `json:"code"`
	Msg  string         `json:"msg"`
	Obj  PageResult     `json:"obj"`
}
