package manager

import (
	"github.com/TtMyth123/DengG/Lottery61BS/httpSixBet/CS2"
	"github.com/TtMyth123/DengG/Lottery61BS/models"
	"github.com/TtMyth123/DengG/Lottery61BS/models/mconst"
)

func NewSixLottery(aSiteUserInfo models.LoBetUser, isRun bool) interface{} {
	switch aSiteUserInfo.TypeId {
	case mconst.AgentSiteType_Id_CS2:
		lottery := CS2.NewLottery(aSiteUserInfo, isRun)
		return lottery
	}

	return nil
}
