package routers

import (
	"github.com/TtMyth123/DengG/Lottery61BS/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func Init() {

	beego.Router("/", &controllers.MainController{})

	beego.AutoRouter(&controllers.MainController{})
	beego.Router("/api/Get", &controllers.MainController{}, "*:Get")
	beego.AutoRouter(&controllers.LoginRegController{})
	beego.AutoRouter(&controllers.DataApiController{})
	beego.AutoRouter(&controllers.SysController{})

	beego.AutoRouter(&controllers.SocketController{})
	beego.AutoRouter(&controllers.ApiController{})
	beego.AutoRouter(&controllers.AgentSiteController{})
	beego.AutoRouter(&controllers.BetSiteController{})

	beego.SetStaticPath("/js", "static/js")
	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/fonts", "static/fonts")
	beego.SetStaticPath("/icons", "static/icons")
	beego.SetStaticPath("/img", "static/img")
	beego.SetStaticPath("/profile", "static/profile")
	beego.SetStaticPath("/", "static")

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"token", "Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers"},
		AllowCredentials: true,
	}))
	beego.Run()
}
