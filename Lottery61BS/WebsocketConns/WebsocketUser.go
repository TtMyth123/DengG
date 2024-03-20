package WebsocketConns

import (
	"errors"
	"fmt"
	"github.com/TtMyth123/DengG/Lottery61BS/WebsocketConns/bo"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

type UserBaseInfo struct {
	UserName string
	Pwd      string
	DataEx   interface{}
	LoginIp  string
}
type WebsocketUser struct {
	UserBaseInfo
	mConn      *websocket.Conn
	inChan     chan *bo.WsData // 读队列
	outChan    chan *bo.WsData // 写队列
	mutex      sync.Mutex      // 避免重复关闭管道
	isClosed   bool            // 管道是否已经关闭
	closeChan  chan byte       // 关闭通知
	fClientMsg ClientMsg
}

type ClientMsg func(*WebsocketUser, bo.WsData) (interface{}, error)

func NewWebsocketUser(UserName, Pwd, LoginIp string, Conn *websocket.Conn, fClientMsg ClientMsg) *WebsocketUser {
	wsConn := &WebsocketUser{
		UserBaseInfo: UserBaseInfo{
			UserName: UserName,
			Pwd:      Pwd,
			LoginIp:  LoginIp,
		},

		inChan:    make(chan *bo.WsData, 1000),
		outChan:   make(chan *bo.WsData, 1000),
		closeChan: make(chan byte),
		isClosed:  false,
		mConn:     Conn,
	}
	wsConn.fClientMsg = fClientMsg
	go wsConn.procLoop()
	// 读协程
	go wsConn.wsReadLoop()
	// 写协程
	go wsConn.wsWriteLoop()
	return wsConn
}

// 写入消息
func (wsConn *WebsocketUser) wsWrite(messageType int, data []byte) error {
	select {
	case wsConn.outChan <- &bo.WsData{MsgType: messageType, Data: data}:
	case <-wsConn.closeChan:
		return errors.New("websocket closed")
	}
	return nil
}
func (wsConn *WebsocketUser) Write(data bo.WsData) error {
	return wsConn.wsWrite(data.MsgType, data.Data)
}

// 读取消息
func (wsConn *WebsocketUser) Read() (*bo.WsData, error) {
	select {
	case msg := <-wsConn.inChan:
		return msg, nil
	case <-wsConn.closeChan:
		return nil, errors.New("websocket closed")
	}
}

// 关闭websocket连接
func (wsConn *WebsocketUser) Close() {
	_ = wsConn.mConn.Close()

	wsConn.mutex.Lock()
	defer wsConn.mutex.Unlock()
	if !wsConn.isClosed {
		wsConn.isClosed = true
		close(wsConn.closeChan)
	}
}

// 循环读取
func (wsConn *WebsocketUser) wsReadLoop() {
	for {
		// 读一个message
		msgType, data, err := wsConn.mConn.ReadMessage()
		if err != nil {
			goto error
		}
		req := &bo.WsData{
			MsgType: msgType,
			Data:    data,
		}

		// 请求放入队列
		select {
		case wsConn.inChan <- req:
		case <-wsConn.closeChan:
			goto closed
		}
	}
error:
	wsConn.Close()
closed:
}

// 循环写入
func (wsConn *WebsocketUser) wsWriteLoop() {
	for {
		select {
		// 取一个应答
		case msg := <-wsConn.outChan:
			// 写给websocket
			if err := wsConn.mConn.WriteMessage(msg.MsgType, msg.Data); err != nil {
				goto error
			}
		case <-wsConn.closeChan:
			goto closed
		}
	}
error:
	wsConn.Close()
closed:
}

// 发送存活心跳
func (wsConn *WebsocketUser) procLoop() {
	// 启动一个gouroutine发送心跳
	go func() {
		for {
			time.Sleep(2 * time.Second)
			data := bo.WsData{
				MsgType: websocket.TextMessage,
				Data:    []byte("heartbeat from server"),
			}
			if err := wsConn.Write(data); err != nil {
				fmt.Println("heartbeat fail")
				wsConn.Close()
				break
			}
		}
	}()

	// 这是一个同步处理模型（只是一个例子），如果希望并行处理可以每个请求一个gorutine，注意控制并发goroutine的数量!!!
	for {
		msg, err := wsConn.Read()
		if err == nil {
			if wsConn.fClientMsg != nil {
				wsConn.fClientMsg(wsConn, *msg)
			}
		}
	}
}
