package domain

import (
	"context"
	"kcers/biz/dal/db/mysql/ent"
)

//事件处理器

type CreatedHandler struct {
	db *ent.Client
}

func (h *CreatedHandler) Handle(event CreatedEvent) error {
	// 创建订单记录
	_, err := h.db.Order.Create().Save(context.Background())
	return err
}

type ItemHandler struct {
	db *ent.Client
}

func (h *ItemHandler) Handle(event CreatedEvent) error {
	// 创建订单记录
	_, err := h.db.OrderItem.Create().Save(context.Background())
	return err
}
