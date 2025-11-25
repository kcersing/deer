package eventbus

import (
	"context"
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

// EventHandler 是订阅者处理函数类型
type EventHandler func(ctx context.Context, event *Event)

type Handler interface {
	Handle(ctx context.Context, event *Event) error
}

func (h EventHandler) Handle(ctx context.Context, event *Event) error {
	return h.Handle(ctx, event)
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

func (e *Event) SetSource(source string) {
	e.Source = source
}
func (e *Event) SetPriority(priority int64) {
	e.Priority = priority
}
