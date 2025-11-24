package eventbus

import "time"

// PersistenEvent 用于储存的事件类型
type PersistenEvent struct {
	Id          int64
	topic       string
	Payload     []byte
	IsProcessed bool
	CreatedAt   time.Time
}

// EventStore 储存接口
type EventStore interface {
	SaveEvent(event PersistenEvent) error
	GetUnprocessedEvents() ([]PersistenEvent, error)
	MarkEventAsProcessed(eventId int64) error
}
