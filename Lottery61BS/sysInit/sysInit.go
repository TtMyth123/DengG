package sysInit

import (
	"github.com/TtMyth123/DengG/Lottery61BS/CacheData"
	"github.com/TtMyth123/DengG/Lottery61BS/models"
	"github.com/TtMyth123/DengG/Lottery61BS/routers"
	"github.com/TtMyth123/kit/ttLog"
)

func Init() {
	ttLog.InitLogs()

	models.Init()

	CacheData.Init()

	routers.Init()
}
