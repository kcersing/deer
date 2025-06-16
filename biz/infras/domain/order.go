package domain

import (
	"errors"
	"time"
)

// Aggregate 订单聚合根
type Aggregate struct {
	Id        int64
	Version   int64
	MemberId  int64
	ProductId int64
	status    int64
	Events    []interface{}
}

// CreatedEvent 订单创建事件
type CreatedEvent struct {
	OrderId     int64
	MemberId    int64
	ProductId   int64
	TotalAmount float64
	At          time.Time
}

type ItemAddedEvent struct {
	OrderId   int64
	ProductId int64
	Date      interface{}
}
type CreateOrderCommand struct {
	MemberId    int64
	ProductId   int64
	TotalAmount float64
}

func (a *Aggregate) CreateOrder(cmd CreateOrderCommand) error {
	// 业务规则验证
	if cmd.MemberId <= 0 {
		return errors.New("无效的会员ID")
	}
	//生成领域事件
	event := CreatedEvent{}
	// 应用领域事件
	a.apply(event)
	return nil
}
func (a *Aggregate) apply(event interface{}) {
	a.Events = append(a.Events, event)
	switch e := event.(type) {
	case CreatedEvent:
		a.MemberId = e.MemberId
	}
}
