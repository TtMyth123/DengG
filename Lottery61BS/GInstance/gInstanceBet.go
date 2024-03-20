package GInstance

import (
	"github.com/TtMyth123/DengG/Lottery61BS/httpSixBet/base"
	"sync"
)

var mpIBet sync.Map

func GetBetLottery(id int64) base.ILottery {
	if i, ok := mpIBet.Load(id); ok {
		ILottery := i.(base.ILottery)
		return ILottery
	} else {
		aILottery := base.GetEmpLottery()
		return aILottery
	}
}

func AddBetLottery(id int64, aLottery interface{}) error {
	mpIBet.Store(id, aLottery)

	return nil
}

func DelBetLottery(id int64) error {
	if i, ok := mpIBet.Load(id); ok {
		ILottery := i.(base.ILottery)
		ILottery.Stop()
		mpIBet.Delete(id)
		return nil
	}

	return nil
}
