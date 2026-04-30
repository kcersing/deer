package ws

import (
	"bytes"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/google/uuid"

	"sync"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 1024
)

// Node 节点
type Node struct {
	Id            string
	Ip            string
	Nid           string
	ReadChan      chan *Message
	WriteChan     chan []byte
	lock          sync.RWMutex
	conn          *websocket.Conn
	hub           *Hub
	BroadcastChan chan *Message
}

var idBase uint64 = 10000

func NewNode(conn *websocket.Conn, hub *Hub, Id, ip string) (eq *Node) {
	eq = &Node{
		Nid:           fmt.Sprintf("%s-%s-%d", uuid.New().String(), hub.serverId, atomic.AddUint64(&idBase, 1)),
		Id:            Id,
		Ip:            ip,
		hub:           hub,
		ReadChan:      hub.broadcast,
		WriteChan:     make(chan []byte, 1024),
		conn:          conn,
		BroadcastChan: make(chan *Message, 1024),
	}
	return
}

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// 读
func (n *Node) read() {
	n.lock.RLock()

	defer func() {
		n.lock.RUnlock()
		n.hub.Unregister(n)
		n.Close()
	}()
	n.conn.SetReadLimit(maxMessageSize)
	err := n.conn.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		return
	}

	for {
		msgType, message, err := n.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				hlog.Errorf("error: %v", err)
			}
			break
		}

		// 只处理文本消息和二进制消息
		if msgType == websocket.BinaryMessage {
			if n.ReadChan != nil {
				n.ReadChan <- &Message{
					Node:    n,
					Content: message,
				}
			} else {
				hlog.Warnf("Node %s ReadChan channel is nil, cannot send message", n.Id)
			}
		} else if msgType == websocket.TextMessage {
			message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
			hlog.Infof("Received message from node %s: %s", n.Id, message)

			n.hub.Broadcast(&Message{
				Content: message,
				Node:    n,
			})
		} else {
			hlog.Warnf("Unsupported message type: %d", msgType)
			continue
		}

	}

}

func (n *Node) write() {
	ticker := time.NewTicker(pingPeriod)
	for {
		select {
		case msg, ok := <-n.WriteChan:
			if !ok {
				// The hub closed the channel.
				err := n.conn.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					hlog.Errorf("error: %v", err)
				}
				return
			}
			if err := n.conn.WriteMessage(websocket.BinaryMessage, msg); err != nil {
				hlog.Errorf("error: %v", err)
			}
		case msg, ok := <-n.BroadcastChan:
			if !ok {
				// The hub closed the channel.
				err := n.conn.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					hlog.Errorf("error: %v", err)
				}
				return
			}
			w, err := n.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				hlog.Errorf("error: %v", err)
				return
			}

			w.Write(msg.Content)
			//n.conn.WriteMessage(websocket.TextMessage, msg.Content)
			if len(n.BroadcastChan) > 0 {
				hlog.Infof("Node %s has %d more messages to send", n.Id, len(n.BroadcastChan))
				for m := range n.BroadcastChan {
					w.Write(newline)
					w.Write(m.Content)
				}
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			err := n.conn.SetWriteDeadline(time.Now().Add(writeWait))
			hlog.Info("ping...")
			if err != nil {
				hlog.Error(err)
				return
			}
			if err := n.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				hlog.Error(err)
				return
			}
		}
	}
}

func (n *Node) Close() {
	if n.conn != nil {
		n.conn.Close()
	}
}
func (n *Node) Run() {
	go n.write()
	go n.read()
	n.conn.SetPongHandler(n.pongHandler)
}

func (n *Node) pongHandler(data string) error {
	hlog.Info("pong...")
	if err := n.conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
		hlog.Infof("error: %v", err)
		return err
	}
	return nil

}
