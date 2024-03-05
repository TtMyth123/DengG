package DataSet

import (
	"fmt"
	"testing"
)

func TestDataSysUser(t *testing.T) {
	aDataSysUser := DataSysUser{}
	fileName := fmt.Sprintf("D:/t/%s", FileName_SysUser)
	e := aDataSysUser.Reload(fileName)
	fmt.Println(e)
	e = aDataSysUser.Save()
	fmt.Println(e)
}
