package server

import (
	"encoding/json"
	"fmt"
	"learngo/websocket/util"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	clients = make(map[string]*util.Client) // 用户组映射 key 为openid val为Client
	join    = make(chan *util.Client, 10)   // 用户加入通道
	leave   = make(chan *util.Client, 10)   // 用户退出通道
	message = make(chan Message, 10)        // 消息通道
)

type Message struct {
	EventType byte       `json:"type"`    // 0表示用户发布消息；1表示用户进入；2表示用户退出
	Name      string     `json:"name"`    // 用户名称
	Message   ReceiveMsg `json:"message"` // 消息内容
	OpenId    string     `json:"openid"`
}

//客户端消息结构体 event 和 消息
type ReceiveMsg struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

//处理每个连接，每个连接，go都会重新起一个协程来处理
func wsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		wsConn *websocket.Conn
		err    error
		client *util.Client
		data   []byte
	)
	r.ParseForm() //返回一个map,并且赋值给r.Form
	name := r.Form["name"][0]
	id := r.Form["id"][0]
	openid := r.Form["openid"][0]

	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}

	if client, err = util.InitConnection(wsConn); err != nil {
		goto ERR
	}
	client.Id = id
	client.Name = name
	client.OpenId = openid

	// 如果用户列表中没有该用户
	if _, ok := clients[openid]; !ok {
		join <- client
	}
	for {
		if data, err = client.ReadMessage(); err != nil { //一直读消息，没有消息就阻塞
			goto ERR
		}
		fmt.Println(string(data)) //{"event":"ping","data":1601435905425}
		//data 就是客户端发来的消息 判断是不是ping 变成 对象
		var msgData ReceiveMsg
		err := json.Unmarshal(data, &msgData)
		if err != nil {
			log.Println("解析消息失败")
		}
		//服务端传递的消息结构
		var msg Message
		if msgData.Event == "ping" {
			msg.EventType = 4
		} else {
			msg.EventType = 0
		}
		msg.Name = client.Name
		msg.OpenId = client.OpenId
		msg.Message = msgData
		message <- msg
	}

ERR:
	leave <- client //这个客户断开
	client.Close()

}

//建立通道 用于侦察客户端返回的消息
func broadcaster() {
	for {
		select {
		// 消息通道中有消息则执行，否则堵塞
		case msg := <-message:
			// 将数据编码成json形式，data是[]byte类型
			// json.Marshal()只会编码结构体中公开的属性(即大写字母开头的属性)
			data, err := json.Marshal(msg) //
			if err != nil {
				return
			}
			if msg.EventType == 4 {
				//只返回给自己
				fmt.Println(msg.OpenId + "的心跳来了")
				clients[msg.OpenId].WriteMessage(data)
				continue
			}
			//给其他用户发送 这里就可以限制房间号
			for openid, client := range clients {
				if client.IsClosed == true {
					leave <- client //这个客户断开
					continue
				}
				fmt.Println("=======the json message is", string(data)) // 转换成字符串类型便于查看
				fmt.Println("=======现在的openid", openid)                 // 转换成字符串类型便于查看
				if client.WriteMessage(data) != nil {
					continue //发送失败就跳过
				}
			}

		// 有用户加入
		case client := <-join:
			clients[client.OpenId] = client // 将用户加入映射
			// 将用户加入消息放入消息通道
			var msg Message
			var msgData ReceiveMsg
			msgData.Event = "join"
			msgData.Data = fmt.Sprintf("%s join in, there are %d preson in room", client.Name, len(clients))
			msg.Name = client.Name
			msg.OpenId = client.OpenId
			msg.EventType = 1
			msg.Message = msgData
			message <- msg

		// 有用户退出
		case client := <-leave:
			// 如果该用户已经被删除
			if _, ok := clients[client.OpenId]; !ok {
				fmt.Println("没有该用户")
				break
			}
			delete(clients, client.OpenId) // 将用户从映射中删除
			// 将用户退出消息放入消息通道
			var msg Message
			var msgData ReceiveMsg
			msgData.Event = "leave"
			msgData.Data = fmt.Sprintf("%s leave, there are %d preson in room", client.Name, len(clients))
			msg.Name = client.Name
			msg.OpenId = client.OpenId
			msg.EventType = 2
			msg.Message = msgData
			message <- msg
		}
	}
}

func Main() {
	go broadcaster()
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe("0.0.0.0:7777", nil)
}
