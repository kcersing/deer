package ws

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/url"
	"sync"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/gorilla/websocket"
)

type Hub struct {
	serverId string
	//所有节点
	nodes map[string]*Node
	//进入
	//register chan *Node
	//离开
	//unregister chan *Node

	broadcast  chan *Message
	NodeReadCh chan *Message
	//获取列表
	//list chan struct{}

	nodeList chan []*Node

	lock sync.RWMutex

	websocketUpgrader  *websocket.Upgrader
	CheckOriginHandler CheckOriginHandler
}

func NewHub(serverId string) *Hub {
	return &Hub{
		serverId: serverId,
		nodes:    make(map[string]*Node),
		//register:   make(chan *Node),
		//unregister: make(chan *Node),

		broadcast: make(chan *Message, 1024),

		//list:     make(chan struct{}),
		nodeList: make(chan []*Node),
	}
}

// Register 进入
func (h *Hub) Register(n *Node) {
	//	h.register <- n
	n.lock.Lock()
	defer n.lock.Unlock()
	h.nodes[n.Id] = n
}

// Unregister 离开
func (h *Hub) Unregister(n *Node) {
	//h.unregister <- n

	if _, ok := h.nodes[n.Id]; ok {
		n.Close()
		delete(h.nodes, n.Id)
	}
}

// Broadcast	广播消息
func (h *Hub) Broadcast(m *Message) {
	//h.broadcast <- m
	// 广播消息
	for _, node := range h.nodes {
		if node.Id == m.Node.Id {
			continue
		}
		select {
		case node.BroadcastChan <- m:
		default:
			node.Close()
			delete(h.nodes, node.Id)
		}
	}
}

func (h *Hub) NodeReadChHandler() {
	for {
		select {
		case body, ok := <-h.NodeReadCh:
			if ok {
				h.decodeNodeMsg(body)
			}
		}
	}
}

func (h *Hub) decodeNodeMsg(m *Message) {
	// 解析协议
	hlog.Warnf("msg: %v", m)

}

// GetList 获取在线设备列表
func (h *Hub) GetList() []*Node {
	//h.list <- struct{}{}

	nodeList := make([]*Node, 0, len(h.nodes))
	for _, node := range h.nodes {
		nodeList = append(nodeList, node)
	}
	//h.nodeList <- nodeList
	//
	//return <-h.nodeList
	return nodeList
}
func (h *Hub) Close() {
	for _, node := range h.nodes {
		node.Close()
		delete(h.nodes, node.Id)
	}

}
func (h *Hub) ServeWs(w http.ResponseWriter, r *http.Request) {

	if h.websocketUpgrader == nil {
		h.websocketUpgrader = &upgrader
	}

	ws, err := h.websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := errors.AsType[websocket.HandshakeError](err); !ok {
			log.Println(err)
		}
		return
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		hlog.Infof("Error getting IP: %v", err)
		return
	}

	values, _ := url.ParseQuery(r.URL.RawQuery)
	id := values.Get("id")

	node := NewNode(ws, h, id, ip)
	node.Run()
	//加入广播列表
	ndl := new(*node)
	node.hub.Register(ndl)
	hlog.Info("New client Read: ", id, " from IP: ", ip) // 连接提示
	node.ReadChan <- NewMessage(node, "已连接")
	msg := NewMessage(node, node.Id+" 连接成功")
	node.hub.Broadcast(msg)

	//退出
	//msg = NewMessage(ndl, "断开连接")
	//node.hub.Broadcast(msg)
	//node.hub.Unregister(ndl)

}

//func (h *Hub) Run() {
//	for {
//		select {
//		case node := <-h.register:
//			//进入加入列表
//			hlog.Info("注册成功 node: ", node.Id)
//			h.nodes[node.Id] = node
//		case node := <-h.unregister:
//			// 离开删除
//			if _, ok := h.nodes[node.Id]; ok {
//				hlog.Info("注销成功 node: ", node.Id)
//
//				node.Close()
//				delete(h.nodes, node.Id)
//			}
//		case msg := <-h.broadcast:
//			// 广播消息
//			for _, node := range h.nodes {
//				if node.Id == msg.Node.Id {
//					continue
//				}
//				select {
//				case node.ReadChan <- msg:
//				default:
//					node.Close()
//					delete(h.nodes, node.Id)
//				}
//			}
//		case <-h.list:
//			nodeList := make([]*Node, 0, len(h.nodes))
//			for _, node := range h.nodes {
//				nodeList = append(nodeList, node)
//			}
//			h.nodeList <- nodeList
//
//		}
//	}
//}

//func (h *Hub) ServeWs(w http.ResponseWriter, r *http.Request) {
//	if h.websocketUpgrader == nil {
//		h.websocketUpgrader = &upgrader
//	}
//	ws, err := h.websocketUpgrader.Upgrade(w, r, nil)
//	if err != nil {
//		if _, ok := errors.AsType[websocket.HandshakeError](err); !ok {
//			log.Println(err)
//		}
//		return
//	}
//}
