package util

import (
	"errors"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	wsConnect *websocket.Conn
	inChan    chan []byte
	outChan   chan []byte
	closeChan chan byte
	Name      string //客户的名称
	Id        string //客户id，唯一
	OpenId    string //客户openid，唯一

	mutex    sync.Mutex // 对closeChan关闭上锁
	IsClosed bool       // 防止closeChan被关闭多次
}

func InitConnection(wsConn *websocket.Conn) (conn *Client, err error) {
	conn = &Client{
		wsConnect: wsConn,
		inChan:    make(chan []byte, 1000),
		outChan:   make(chan []byte, 1000),
		closeChan: make(chan byte, 1),
		IsClosed:  false,
	}
	// 启动读协程
	go conn.readLoop()
	// 启动写协程
	go conn.writeLoop()
	return
}
func (conn *Client) ReadMessage() (data []byte, err error) {
	select {
	case data = <-conn.inChan:
	case <-conn.closeChan:
		err = errors.New("connection is closeed")
	}
	return
}
func (conn *Client) WriteMessage(data []byte) (err error) {
	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		err = errors.New("connection is closeed")
	}
	return
}
func (conn *Client) Close() {
	// 线程安全，可多次调用
	conn.wsConnect.Close()
	// 利用标记，让closeChan只关闭一次
	conn.mutex.Lock()
	if !conn.IsClosed {
		close(conn.closeChan)
		conn.IsClosed = true
	}
	conn.mutex.Unlock()
}

func (conn *Client) readLoop() {
	var (
		data []byte
		err  error
	)
	for {
		if _, data, err = conn.wsConnect.ReadMessage(); err != nil {
			goto ERR
		}
		//阻塞在这里，等待inChan有空闲位置
		select {
		case conn.inChan <- data:
		case <-conn.closeChan: // closeChan 感知 conn断开
			goto ERR
		}

	}

ERR:
	conn.Close()
}

func (conn *Client) writeLoop() {
	var (
		data []byte
		err  error
	)

	for {
		select {
		case data = <-conn.outChan:
		case <-conn.closeChan:
			goto ERR
		}
		if err = conn.wsConnect.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}

ERR:
	conn.Close()

}
