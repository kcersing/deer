package aggregate

import (
	"github.com/pkg/errors"
	"kcers-order/biz/infras/common"
	"kcers-order/biz/infras/status"
)

type Order struct {
	Id int64
	Sn string

	MemberId  int64
	CreatedId int64

	Items       []Item
	TotalAmount float64

	Status       status.OrderStatus
	stateMachine *StateMachine

	common.AggregateBase
}

type Item struct {
	ProductId int64
	Quantity  int64
	Price     float64
	Name      string
}

func NewOrder(sn string, items []Item, amount float64, memberId int64, userId int64) (order *Order) {
	order = &Order{
		Sn:          sn,
		MemberId:    memberId,
		CreatedId:   userId,
		Items:       items,
		TotalAmount: amount,
		Status:      status.Created,
	}
	order.stateMachine = NewStateMachine(order)

	return order
}

//func (o *Order) When(evt Event) error {
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

func (o *Order) AddEvent(event common.Event) error {
	o.Mu.Lock()
	defer o.Mu.Unlock()
	if event.GetAggregateID() != o.AggregateID {
		return errors.New("aggregate id not match")
	}
	if err := o.When(event); err != nil {
		return err
	}
	o.Version++
	//event.Version = o.GetVersion()

	o.UncommittedEvents = append(o.UncommittedEvents, event)
	return nil
}
func (o *Order) GetEvents() []common.Event {
	o.Mu.RLock()
	defer o.Mu.RUnlock()
	es := make([]common.Event, len(o.UncommittedEvents))
	copy(es, o.UncommittedEvents)
	return es
}

func (o *Order) ClearEvents() {
	o.Mu.Lock()
	defer o.Mu.Unlock()
	o.UncommittedEvents = []common.Event{}
}
