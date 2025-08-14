package aggregate

import (
	"kcers-order/biz/infras/events"
	"kcers-order/biz/infras/service"
	"kcers-order/biz/infras/states"
	"sync"
)

type Order struct {
	Id int64
	Sn string

	MemberId int64
	Items    []Item

	Status       states.OrderStatus
	stateMachine *service.StateMachine

	Mu sync.RWMutex
}

type Item struct {
	ProductId int64
	Quantity  int64
	Price     float64
	Name      string
}

//func (o *Order) When(evt events.Event) error {
//
//	switch evt.GetType() {
//	case states.Created:
//		var eventDate events.CreatedOrderEvent
//		if err := sonic.Unmarshal(evt.GetData(), &eventDate); err != nil {
//			return err
//		}
//		o.Sn = eventDate.Sn
//
//	default:
//		return errors.New("unknown event type")
//	}
//
//}

func (o *Order) AddEvent(event events.Event) error {
	o.Mu.Lock()
	defer o.Mu.Unlock()

	if err := o.when(event); err != nil {
		return err
	}
	o.UncommittedEvents = append(o.UncommittedEvents, event)
	return nil
}
func (o *Order) GetEvents() []events.Event {
	o.Mu.RLock()
	defer o.Mu.RUnlock()
	es := make([]events.Event, len(o.UncommittedEvents))
	copy(es, o.UncommittedEvents)
	return es
}

func (o *Order) ClearEvents() {
	o.Mu.Lock()
	defer o.Mu.Unlock()
	o.UncommittedEvents = []events.Event{}
}
