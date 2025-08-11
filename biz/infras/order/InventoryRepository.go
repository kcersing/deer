package order

import (
	"context"
	"kcers-order/biz/dal/db/mysql/ent"
)

type InventoryRepository interface {
	// 预留库存（返回错误表示库存不足）
	Reserve(ctx context.Context, items []OrderItem) error
	// 释放库存
	Release(ctx context.Context, orderID int64) error
	// 检查库存是否充足
	CheckAvailability(ctx context.Context, items []OrderItem) (bool, error)
}

// InventoryRepositoryImpl 订单仓储实现
type InventoryRepositoryImpl struct {
	db              *ent.Client
	ctx             context.Context
	snapshotFreq    int                 // 快照频率
	subscriptionSvc SubscriptionService // 新增：订阅服务
}

func (o InventoryRepositoryImpl) Reserve(ctx context.Context, items []OrderItem) error {

	return nil
}
func (o InventoryRepositoryImpl) Release(ctx context.Context, orderID int64) error {
	return nil
}
func (o InventoryRepositoryImpl) CheckAvailability(ctx context.Context, items []OrderItem) (bool, error) {
	return true, nil
}
