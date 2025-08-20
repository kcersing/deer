package infras

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"kcers-order/biz/dal/db/mysql/ent"
	"kcers-order/biz/infras/common"
	"kcers-order/biz/infras/order/events"
)

// InventoryHandler 库存处理器
type InventoryHandler struct {
	repo InventoryRepository
	ctx  context.Context
	// 库存服务依赖
}

func (h *InventoryHandler) Handle(ctx context.Context, event common.Event) error {
	switch e := event.(type) {
	case *events.CreatedOrderEvent:
		// 处理库存预留
		return h.Reserve(e.AggregateID, e.Items)
	case *events.CancelledOrderEvent:
		// 处理库存释放
		return h.Release(e.AggregateID)
	}
	return nil
}

func (h *InventoryHandler) Reserve(aggregateID int64, items []common.Item) error {
	// 实现库存预留逻辑
	// 1. 校验库存是否充足
	// 2. 预留库存
	// 3. 记录预留信息
	//err := h.repo.Reserve(h.ctx, items)
	//if err != nil {
	//	return errors.New("库存不足")
	//}
	//return h.repo.Reserve(h.ctx, items)
	klog.Info("库存预留123")
	return nil
}

func (h *InventoryHandler) Release(aggregateID int64) error {
	// 实现库存释放逻辑
	klog.Info("库存释放123")
	return nil
}

type InventoryRepository interface {
	// 预留库存（返回错误表示库存不足）
	Reserve(ctx context.Context, items []common.Item) error
	// 释放库存
	Release(ctx context.Context, aggregateID int64) error
	// 检查库存是否充足
	CheckAvailability(ctx context.Context, items []common.Item) (bool, error)
}

// InventoryRepositoryImpl 订单仓储实现
type InventoryRepositoryImpl struct {
	db              *ent.Client
	ctx             context.Context
	snapshotFreq    int                 // 快照频率
	subscriptionSvc SubscriptionService // 新增：订阅服务
}

func (o InventoryRepositoryImpl) Reserve(ctx context.Context, items []common.Item) error {
	klog.Info("库存预留")
	return nil
}
func (o InventoryRepositoryImpl) Release(ctx context.Context, aggregateID int64) error {
	klog.Info("库存释放")
	return nil
}
func (o InventoryRepositoryImpl) CheckAvailability(ctx context.Context, items []common.Item) (bool, error) {
	return true, nil
}

//
//// PayService 支付服务接口
//type PayService interface {
//	Refund(id int64, amount float64, reason string, createdId int64) error
//}
//type PayHandler struct {
//	PayService PayService
//}
//
//func (h *PayHandler) Handle(event Event) error {
//	if h.PayService == nil {
//		return errors.New("支付服务未初始化")
//	}
//	switch e := event.(type) {
//	case *OrderRefundedEvent:
//		return h.PayService.Refund(e.Id, e.RefundedAmount, e.Reason, e.CreatedId)
//	default:
//		return nil
//	}
//
//}
//
//// NotificationServer 通知服务接口
//type NotificationServer interface {
//	SendOrderCreatedNotification(id int64) error
//	SendOrderPaidNotification(id int64) error
//	SendOrderShippedNotification(id int64) error
//	SendOrderCompletedNotification(id int64) error
//	SendOrderCancelledNotification(id int64) error
//	SendOrderRefundNotification(id int64) error
//}
//type NotificationHandler struct {
//	NotificationServer NotificationServer
//}
//
//func (h *NotificationHandler) Handler(event Event) error {
//	if h.NotificationServer == nil {
//		return errors.New("通知服务未初始化")
//	}
//	switch e := event.(type) {
//	case *OrderCreatedEvent:
//		// 处理订单创建事件
//		return h.NotificationServer.SendOrderCreatedNotification(e.Id)
//	case *OrderPaidEvent:
//		// 处理订单支付事件
//		return h.NotificationServer.SendOrderPaidNotification(e.Id)
//	case *OrderShippedEvent:
//		// 处理订单发货事件
//		return h.NotificationServer.SendOrderShippedNotification(e.Id)
//	case *OrderCompletedEvent:
//		// 处理订单完成事件
//		return h.NotificationServer.SendOrderCompletedNotification(e.Id)
//	case *OrderCancelledEvent:
//		// 处理订单取消事件
//		return h.NotificationServer.SendOrderCancelledNotification(e.Id)
//	case *OrderRefundedEvent:
//		// 处理订单退款事件
//		return h.NotificationServer.SendOrderRefundNotification(e.Id)
//	default:
//		return nil
//	}
//}
