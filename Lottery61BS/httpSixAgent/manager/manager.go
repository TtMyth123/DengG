package manager

import (
	"github.com/TtMyth123/DengG/Lottery61BS/httpSixAgent/CS2"
	"github.com/TtMyth123/DengG/Lottery61BS/models"
	"github.com/TtMyth123/DengG/Lottery61BS/models/mconst"
)

func NewSixLottery(aSiteUserInfo models.LoAgentUser, isRun bool) interface{} {
	switch aSiteUserInfo.TypeId {
	case mconst.AgentSiteType_Id_CS2:
		lottery := CS2.NewLottery(aSiteUserInfo)
		return lottery
	}

	return nil
}
