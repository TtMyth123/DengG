package CacheData

import (
	"fmt"
	"github.com/TtMyth123/DengG/Lottery61BS/models"
)

const key_LoLotteryType = "LoLotteryType:"

func getLoLotteryTypeKey(t int64) string {
	return fmt.Sprintf("%s%d", key_LoLotteryType, t)
}
func GetLoLotteryType(t int64) models.LoLotteryType {
	aData := &models.LoLotteryType{}
	key := getLoLotteryTypeKey(t)
	data, e := GetICache().Get(key, aData)
	if e != nil {
		aData.Id = t
		return *aData
	}
	aData = data.(*models.LoLotteryType)
	return *aData
}

func SetLoLotteryTypeAnd2DB(aLoLotteryType models.LoLotteryType) error {
	key := getLoLotteryTypeKey(aLoLotteryType.Id)
	e := GetICache().SetCache(key, &aLoLotteryType, -1)
	if e != nil {
		return e
	}
	e = LoadLoLotteryType2DB(aLoLotteryType.Id)
	return e
}
func SaveLotteryType_CycleId(t int64, CycleId string) {
	aLoLotteryType := GetLoLotteryType(t)
	if aLoLotteryType.CycleId != CycleId {
		aLoLotteryType.CycleId = CycleId
		SetLoLotteryTypeAnd2DB(aLoLotteryType)
	}
}

func LoadLoLotteryType2DB(t int64) error {
	aData := GetLoLotteryType(t)
	e := aData.Update(nil)
	if e != nil {
		return e
	}
	return nil
}
