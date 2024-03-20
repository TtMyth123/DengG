package models

import "github.com/astaxie/beego/orm"

func initRegisterModelBo() {
	orm.RegisterModel(new(LoNumInfo))
	orm.RegisterModel(new(LoConfig))
	orm.RegisterModel(new(LoLotteryType))

	orm.RegisterModel(new(LoAgentBetNumData))
	orm.RegisterModel(new(LoAgentType))
	orm.RegisterModel(new(LoAgentUser))

	orm.RegisterModel(new(LoBetType))
	orm.RegisterModel(new(LoBetUser))
	orm.RegisterModel(new(LoBetBetData))

}

func initDataBo(o orm.Ormer) {
	InitLoConfig(o)
	InitLoNumInfo(o)
	InitLoLotteryType(o)
	InitLoAgentType(o)
	InitLoAgentUser(o)

	InitLoBetType(o)
	InitLoBetUser(o)
}
