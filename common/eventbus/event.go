package eventbus

import (
	"time"
)

// Event 事件
type Event struct {
	Id        string
	Topic     string // 主题
	Payload   any    // 事件负载
	Source    string
	Version   int64
	Timestamp time.Time
	Priority  int64
}

func NewEvent(topic string, payload any) *Event {
	return &Event{
		Id:        "",
		Topic:     topic,
		Payload:   payload,
		Version:   0,
		Timestamp: time.Time{},
		Priority:  0,
	}
}
