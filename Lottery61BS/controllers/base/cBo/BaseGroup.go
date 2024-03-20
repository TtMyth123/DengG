package cBo

type BaseGroup struct {
	C         int
	MaxId     int64
	PageTotal int
	PageIndex int
}

type BaseListResult struct {
	BaseGroup
	GroupInfo interface{}
	DataList  interface{}
}
