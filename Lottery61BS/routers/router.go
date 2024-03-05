package routers

import (
	"github.com/TtMyth123/DengG/Lottery61BS/controllers"
	"github.com/astaxie/beego"
)

func Init() {
	beego.AutoRouter(&controllers.MainController{})
	beego.Router("/api/Get", &controllers.MainController{}, "*:Get")

	beego.Run()
}
