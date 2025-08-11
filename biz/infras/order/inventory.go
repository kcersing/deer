package order

import (
	"context"
	"github.com/pkg/errors"
)

// InventoryHandler 库存处理器
type InventoryHandler struct {
	repo InventoryRepository
	ctx  context.Context
	// 库存服务依赖
}

func (h *InventoryHandler) Handle(ctx context.Context, event Event) error {
	switch e := event.(type) {
	case *CreatedEvent:
		// 处理库存预留
		return h.Reserve(e.AggregateID, e.Items)
	case *CancelledEvent:
		// 处理库存释放
		return h.Release(e.AggregateID)
	}
	return nil
}

func (h *InventoryHandler) Reserve(orderID int64, items []OrderItem) error {
	// 实现库存预留逻辑
	// 1. 校验库存是否充足
	// 2. 预留库存
	// 3. 记录预留信息
	err := h.repo.Reserve(h.ctx, items)
	if err != nil {
		return errors.New("库存不足")
	}
	return h.repo.Reserve(h.ctx, items)

}

func (h *InventoryHandler) Release(orderID int64) error {
	// 实现库存释放逻辑
	return nil
}
