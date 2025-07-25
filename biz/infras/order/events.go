package order

import "time"

type Event interface {
	GetEventId() string
	SetId() int64
	GetCreatedAt() time.Time
	SetEventType() string
}
