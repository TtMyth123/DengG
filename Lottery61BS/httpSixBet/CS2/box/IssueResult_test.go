package box

import (
	"fmt"
	"github.com/TtMyth123/kit/stringKit"
	"testing"
)

func Test_IssueResult1(t *testing.T) {

	strJson := `{"open_id":"69787190","pack_id":"0","part_id":"0","table_id":"0","lotcode":"bjamlh","cycleid":"2023-011","user_id":"61258","open_time":1673402400,"close_time":1673443200,"close_time2":1673443200,"is_auto":"1","is_auto2":"1","b1":"06","b2":"20","b3":"27","b4":"25","b5":"44","b6":"05","b7":"16","b8":null,"b9":null,"b10":null,"bstr":null,"blotstr":null,"created":"1672925772","updated":"1673444157","att":"2","auto_open":"1","cycleid2":null,"code2":null,"limitval":"0","L_open_time":"1673402400","L_close_time":"1673443800","L_close_time2":"1673443800","week":"\u4e09","lfopentime":0,"nopeinfo":{"open_id":"69787347","open_time":1673488800,"close_time":1673529600,"close_time2":1673529600,"att":"0","lotcode":"bjamlh","cycleid":"2023-012","user_id":"61258","L_open_time":"1673488800","L_close_time":"1673530200","L_close_time2":"1673530200","is_auto":"1","is_auto2":"1"},"nclosttime":1673529600,"servertime":1673447231}`

	aIssueResult := IssueResult{}
	e := stringKit.GetJsonObj(strJson, &aIssueResult)
	if e != nil {
		t.Fail()
	}
	fmt.Println(aIssueResult.CycleId)
}
