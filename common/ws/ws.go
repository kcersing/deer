package ws

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// Publisher 发布
type Publisher interface {
	Publish()
}

// Subscriber 订阅
type Subscriber interface {
	Subscribe() (ch chan struct{}, cleanUp func(), err error)
}

type Write interface {
	write()
}

type Read interface {
	read()
}

type Close interface {
	Close()
}

type CheckOriginHandler = func(r *http.Request) bool

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Message struct {
	Node    *Node
	Content []byte
}

func NewMessage(n *Node, content string) *Message {
	return &Message{
		Node:    n,
		Content: []byte(content),
	}
}

// 全局指标变量
var (
	activeConnections  int32 //	当前活跃连接数
	totalConnections   int64 // 总连接数
	messagesReceived   int64 // 收到的消息总数
	messagesSent       int64 // 消息
	bytesReceived      int64 //
	bytesSent          int64 //
	lastConnectionTime int64 // Unix 时间戳
)
