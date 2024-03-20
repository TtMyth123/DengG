package Manager

import (
	"fmt"
	"github.com/TtMyth123/DengG/Lottery61BS/WebsocketConns"
	"github.com/TtMyth123/DengG/Lottery61BS/WebsocketConns/bo"

	"github.com/TtMyth123/kit/stringKit"
	"github.com/gorilla/websocket"
	"sync"
)

type WebsocketManager struct {
	connMap sync.Map
}

func NewWebsocketManager() *WebsocketManager {
	aWebsocketManager := &WebsocketManager{}
	return aWebsocketManager
}

func (this *WebsocketManager) AddConn(aWebsocketUser *WebsocketConns.WebsocketUser) error {
	if v, ok := this.connMap.Load(aWebsocketUser.UserName); !ok {
		this.connMap.Store(aWebsocketUser.UserName, aWebsocketUser)
	} else {
		oldWebsocketUser := v.(*WebsocketConns.WebsocketUser)
		this.delConn(oldWebsocketUser)

		this.connMap.Store(aWebsocketUser.UserName, aWebsocketUser)
	}
	return nil
}
func (this *WebsocketManager) DelConn(aWebsocketUser *WebsocketConns.WebsocketUser) error {
	return this.delConn(aWebsocketUser)
}

func (this *WebsocketManager) delConn(aWebsocketUser *WebsocketConns.WebsocketUser) error {
	this.connMap.Delete(aWebsocketUser.UserName)
	aWebsocketUser.Close()
	aWebsocketUser = nil
	return nil
}

func (this *WebsocketManager) SendMsg(UserName string, Msg string) error {
	if UserName == "" {
		this.broadcast(Msg)
		return nil
	}

	if v, ok := this.connMap.Load(UserName); ok {
		aWebsocketUser := v.(*WebsocketConns.WebsocketUser)
		data := []byte(Msg)
		wsData := bo.WsData{MsgType: websocket.TextMessage, Data: data}
		e := aWebsocketUser.Write(wsData)
		return e
	}
	return fmt.Errorf("没有对应的用户")
}

func (this *WebsocketManager) broadcast(Msg string) error {
	data := []byte(Msg)
	this.connMap.Range(func(key, v any) bool {
		aWebsocketUser := v.(*WebsocketConns.WebsocketUser)
		wsData := bo.WsData{MsgType: websocket.TextMessage, Data: data}
		aWebsocketUser.Write(wsData)
		return true
	})
	return nil
}

func (this *WebsocketManager) SendPrintInfo(UserName string, info bo.BetPrintInfo) error {
	Msg := stringKit.GetJsonStr(info)
	if UserName == "" {
		this.broadcast(Msg)
		return nil
	}

	if v, ok := this.connMap.Load(UserName); ok {
		aWebsocketUser := v.(*WebsocketConns.WebsocketUser)
		data := []byte(Msg)
		wsData := bo.WsData{MsgType: websocket.TextMessage, Data: data}
		e := aWebsocketUser.Write(wsData)
		return e
	}
	return fmt.Errorf("没有对应的用户")
}
