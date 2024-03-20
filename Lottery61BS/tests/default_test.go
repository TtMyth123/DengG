package test

import (
	"fmt"
	"github.com/TtMyth123/DengG/Lottery61BS/models"
	"github.com/TtMyth123/kit/stringKit"
	"testing"
)

func TestB(t *testing.T) {
	aLoConfig := models.LoConfig{}
	aLoConfig.LotteryType = 1
	aLoConfig.Stock = 1
	str := stringKit.GetJsonStr(aLoConfig)
	fmt.Println(str)
}
