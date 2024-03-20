package cBo

import (
	"fmt"
	"github.com/TtMyth123/kit"
)

const (
	FieldType_0_Int     = 0
	FieldType_1_Str     = 1
	FieldType_2_Float64 = 2
	FieldType_3_Int64   = 3
)

type BaseListParam struct {
	SysUserId int64
	PageIndex int
	PageSize  int
	MaxId     int64
	Other     map[string]interface{}
}

func (this BaseListParam) GetString(str string) string {
	if v, ok := this.Other[str].(string); ok {
		return v
	} else {
		strV := fmt.Sprint(this.Other[str])
		return strV
	}
}

func (this BaseListParam) GetInt(str string) int {
	if v, ok := this.Other[str].(int); ok {
		return v
	} else {
		newV := kit.GetInterface2Int(this.Other[str], 0)
		return newV
	}
}

func (this BaseListParam) GetFloat(str string) float64 {
	if v, ok := this.Other[str].(float64); ok {
		return v
	} else {
		newV := kit.GetInterface2Float64(this.Other[str], 0)
		return newV
	}
}

func (this BaseListParam) GetInt64(str string) int64 {
	if v, ok := this.Other[str].(int64); ok {
		return v
	} else {
		newV := kit.GetInterface2Int64(this.Other[str], 0)
		return newV
	}
}
