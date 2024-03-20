package GInstance

import (
	"github.com/TtMyth123/DengG/Lottery61BS/httpSixAgent/base"
	"sync"
)

var mpIAgent sync.Map

func GetAgentLottery(id int64) base.ILottery {
	if i, ok := mpIAgent.Load(id); ok {
		ILottery := i.(base.ILottery)
		return ILottery
	} else {
		aILottery := base.GetEmpLottery()
		return aILottery
	}
}

func AddAgentLottery(id int64, aLottery interface{}) error {
	mpIAgent.Store(id, aLottery)

	return nil
}

func DelAgentLottery(id int64) error {
	if i, ok := mpIAgent.Load(id); ok {
		ILottery := i.(base.ILottery)
		ILottery.Stop()
		mpIAgent.Delete(id)
		return nil
	}

	return nil
}
