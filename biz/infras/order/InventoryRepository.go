package order

import "context"

type InventoryRepository interface {
	// 预留库存（返回错误表示库存不足）
	Reserve(ctx context.Context, items []OrderItem) error
	// 释放库存
	Release(ctx context.Context, orderID int64) error
	// 检查库存是否充足
	CheckAvailability(ctx context.Context, items []OrderItem) (bool, error)
}
