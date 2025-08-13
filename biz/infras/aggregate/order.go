package aggregate

import (
	"kcers-order/biz/infras/events"
	"kcers-order/biz/infras/states"
	"sync"
)

type Order struct {
	Id      int64
	Sn      string
	Items   []Item
	Status  states.OrderStatus
	Version int64
	Events  []events.Event
	Mu      sync.RWMutex
}
type Item struct {
	ProductId int64
	Quantity  int64
	Price     float64
	Name      string
}

func (o *Order) AddEvent(event events.Event) {
	o.Mu.Lock()
	defer o.Mu.Unlock()
	o.Events = append(o.Events, event)
}
func (o *Order) GetEvents() []events.Event {
	o.Mu.RLock()
	defer o.Mu.RUnlock()
	es := make([]events.Event, len(o.Events))
	copy(es, o.Events)
	return es
}

func (o *Order) ClearEvents() {
	o.Mu.Lock()
	defer o.Mu.Unlock()
	o.Events = []events.Event{}
}
