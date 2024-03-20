package controllers

import (
	"github.com/TtMyth123/DengG/Lottery61BS/WebsocketConns"
	"github.com/TtMyth123/DengG/Lottery61BS/WebsocketConns/Singleton"
	"github.com/TtMyth123/DengG/Lottery61BS/controllers/base"
	"github.com/TtMyth123/DengG/Lottery61BS/models"
	"github.com/astaxie/beego"
	"net/http"

	"github.com/gorilla/websocket"
)

type SocketController struct {
	base.ABaseController
}

func (this *SocketController) Join() {
	UserName := this.GetString("UserName")
	Pwd := this.GetString("Pwd")
	_, e := models.SysUserOneByUserName(UserName, Pwd)
	if e != nil {
		http.Error(this.Ctx.ResponseWriter, e.Error(), 400)
		return
	}

	//ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	upgrader := websocket.Upgrader{
		// 支持跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a SocketController handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup SocketController connection:", err)
		return
	}
	LoginIp := this.Ctx.Request.RemoteAddr
	aWebsocketUser := WebsocketConns.NewWebsocketUser(UserName, Pwd, LoginIp, ws, nil)
	Singleton.GetWebsocketManager().AddConn(aWebsocketUser)
}

func (this *SocketController) Broadcast() {
	msg := this.GetString("Msg")
	e := Singleton.GetWebsocketManager().SendMsg("", msg)
	this.JsonResultEx("", e)
}
