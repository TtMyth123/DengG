package Singleton

import (
	"github.com/TtMyth123/DengG/Lottery61BS/WebsocketConns/Manager"
)

var (
	mWebsocketManager *Manager.WebsocketManager
)

func Init() {
	mWebsocketManager = Manager.NewWebsocketManager()
}
func GetWebsocketManager() *Manager.WebsocketManager {
	return mWebsocketManager
}
