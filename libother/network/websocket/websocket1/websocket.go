package websocket1

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
)

//消息结构体
type Message struct {
	messageType int //这里的类型对应1文本，2二进制，8关闭，9ping,10字节数组
	data        []byte
}

// 客户端连接数据包结构体
type wsConnection struct {
	Conn    *websocket.Conn
	inChan  chan *Message // 读队列
	outChan chan *Message // 写队列

	mutex     sync.Mutex // 互斥锁，避免重复关闭管道
	isClosed  bool
	closeChan chan byte // 关闭通知
}

func (wsConn *wsConnection) wsWriteLoop() {
	for {
		//当没有可执行的IO语句时，也没有default语句，会阻塞在这里
		select {
		case msg := <-wsConn.outChan:
			if err := wsConn.Conn.WriteMessage(msg.messageType, msg.data); err != nil {
				wsConn.wsClose()
				break
			}

		case <-wsConn.closeChan:
			break
		}
	}
}

func (wsConn *wsConnection) wsWrite(msg *Message) error {
	select {
	case wsConn.outChan <- msg:
	case <-wsConn.closeChan:
		return errors.New("webSocket 关闭")
	}

	return nil
}

func (wsConn *wsConnection) wsRead() (*Message, error) {
	select {
	case msg := <-wsConn.inChan:
		return msg, nil
	case <-wsConn.closeChan:
	}

	return nil, errors.New("webSocket 关闭")
}

func (wsConn *wsConnection) wsReadLoop() {
	for {
		msgType, data, err := wsConn.Conn.ReadMessage()
		if err != nil {
			wsConn.wsClose()
			break
		}

		req := &Message{msgType, data}

		select {
		case wsConn.inChan <- req:
		case <-wsConn.closeChan:
			break
		}
	}
}

//关闭连接和通道
func (wsConn *wsConnection) wsClose() {
	wsConn.Conn.Close()

	wsConn.mutex.Lock()
	defer wsConn.mutex.Unlock()

	if !wsConn.isClosed {
		wsConn.isClosed = true
		close(wsConn.closeChan)
	}
}

//正常处理流程
func (wsConn *wsConnection) processLoop() {
	for {
		msg, err := wsConn.wsRead()
		if err != nil {
			fmt.Println("读取失败")
			break
		}

		req := string(msg.data)
		fmt.Println("收到的客户端信息：", req)

		response := &Message{websocket.TextMessage, []byte("三生三世十里桃花")}

		err = wsConn.wsWrite(response)
		if err != nil {
			fmt.Println("写失败")
			break
		}
	}
}

//检测心跳
func (wsConn *wsConnection) SendHeartBeat() {
	heartBeatText := &Message{websocket.TextMessage, []byte("服务端心跳检测")}

	go func() {
		for {
			time.Sleep(2 * time.Second)
			if err := wsConn.wsWrite(heartBeatText); err != nil {
				fmt.Println("心跳已死")
				wsConn.wsClose()
				break
			}
		}
	}()
}

//建立连接握手
func WsHandShake(resp http.ResponseWriter, req *http.Request) *websocket.Conn {
	wsUpgrader := websocket.Upgrader{
		// 允许所有CORS跨域请求
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	wsSocket, err := wsUpgrader.Upgrade(resp, req, nil)
	if err != nil {
		return nil
	}

	return wsSocket
}

//构造一个数据包结构
func DataBody(wsSocket *websocket.Conn) *wsConnection {
	return &wsConnection{
		Conn:      wsSocket,
		inChan:    make(chan *Message, 1000),
		outChan:   make(chan *Message, 1000),
		closeChan: make(chan byte),
		isClosed:  false,
	}
}

//处理请求的方法
func wsHandler(resp http.ResponseWriter, req *http.Request) {
	//①握手，由http协议转为websocket协议
	wsSocket := WsHandShake(resp, req)
	//②构造接收数据的数据包结构体
	wsConn := DataBody(wsSocket)
	//③发送心跳到chan
	wsConn.SendHeartBeat()
	//④循环读取chan数据接收
	go wsConn.processLoop()
	//⑤循环读取websocket数据送入chan
	go wsConn.wsReadLoop()
	//⑥循环输出数据
	go wsConn.wsWriteLoop()
}

func Entry() {

	go http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe(":8889", nil)
}
