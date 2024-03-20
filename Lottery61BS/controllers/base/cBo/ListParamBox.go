package cBo

import (
	"fmt"
	"github.com/TtMyth123/kit"
	"strings"
)

type ListParamBox struct {
	PageIndex int
	PageSize  int
	MaxId     int64
	Other     map[string]interface{}
	OrderBy   string
}

func (this *ListParamBox) GetFloat64(k string, def ...float64) float64 {
	v := float64(0)
	if len(def) > 0 {
		v = def[0]
	}

	if v1, ok := this.Other[k]; !ok {
		return v
	} else {
		return kit.GetInterface2Float64(v1, v)
	}
}

func (this *ListParamBox) GetInt(k string, def ...int) int {
	v := int(0)
	if len(def) > 0 {
		v = def[0]
	}
	if v1, ok := this.Other[k]; !ok {
		return v
	} else {
		return kit.GetInterface2Int(v1, v)
	}
}

func (this *ListParamBox) GetInt32(k string, def ...int32) int32 {
	v := int32(0)
	if len(def) > 0 {
		v = def[0]
	}
	if v1, ok := this.Other[k]; !ok {
		return v
	} else {
		return kit.GetInterface2Int32(v1, v)
	}
}

func (this *ListParamBox) GetInt64(k string, def ...int64) int64 {
	v := int64(0)
	if len(def) > 0 {
		v = def[0]
	}
	if v1, ok := this.Other[k]; !ok {
		return v
	} else {
		return kit.GetInterface2Int64(v1, v)
	}
}

func (this *ListParamBox) GetString(k string, def ...string) string {
	v := ""
	if len(def) > 0 {
		v = def[0]
	}
	if v1, ok := this.Other[k]; !ok {
		return v
	} else {
		return kit.GetInterface2Str(v1, v)
	}
}

func (this *ListParamBox) AddParam(p ArgsParam, value interface{}) {
	this.Other[p.PName] = value
}
func (this *ListParamBox) AddParamValue(key string, value interface{}) {
	this.Other[key] = value
}

func GetListParamBox() ListParamBox {
	aListParamBox := ListParamBox{}
	aListParamBox.PageSize = 50
	aListParamBox.PageIndex = 1
	aListParamBox.Other = make(map[string]interface{})
	return aListParamBox
}

func (this *ListParamBox) GetOrderBy(mp map[string]string) string {
	strOrder := strings.Replace(this.OrderBy, " ", ",", -1)
	strOrder = strings.Replace(strOrder, ",,", ",", -1)
	for {
		strOrder1 := strings.Replace(strOrder, ",,", ",", -1)
		if strOrder1 == strOrder {
			break
		}
		strOrder = strOrder1
	}

	arrOrder := strings.Split(strOrder, ",")
	iLen := len(arrOrder)
	if iLen >= 2 {
		if arrOrder[iLen-1] == "" {
			arrOrder = arrOrder[0 : iLen-1]
			iLen--
		}
	}

	for i := 0; i < iLen; i++ {
		key := arrOrder[i]
		if v, ok := mp[key]; ok {
			arrOrder[i] = v
		}
	}

	ii := iLen
	if iLen >= 2 {
		desc := strings.ToLower(arrOrder[iLen-1])
		if desc == "desc" || desc == "asc" {
			ii = iLen - 1
		}

	}

	if ii == iLen {
		strOrder = strings.Join(arrOrder[0:ii], ",")
	} else {
		strOrder = strings.Join(arrOrder[0:ii], ",")
		strOrder += fmt.Sprintf(" %s", arrOrder[ii])
	}

	return strOrder
}
