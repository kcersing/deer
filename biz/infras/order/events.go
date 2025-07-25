package order

import "time"

type Event interface {
	GetEventId() string
	GetId() int64
	GetCreatedAt() time.Time
	GetEventType() string
}
