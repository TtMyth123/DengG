package cBo

import (
	"fmt"
	"testing"
)

func Test_GetOrderBy(t *testing.T) {
	a := GetListParamBox()
	a.OrderBy = "num DESC"
	bb := a.GetOrderBy(nil)
	fmt.Println(bb)
}

func Test_GetOrderBy1(t *testing.T) {
	a := GetListParamBox()
	a.OrderBy = "num,   Bum"
	bb := a.GetOrderBy(nil)
	fmt.Println(bb)
}

func Test_GetOrderBy2(t *testing.T) {
	a := GetListParamBox()
	a.OrderBy = "num,,,    Bum"
	bb := a.GetOrderBy(nil)
	fmt.Println(bb)
}
