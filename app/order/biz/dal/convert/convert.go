package convert

import (
	"common/pkg/utils"
	"gen/kitex_gen/base"
	"order/biz/dal/db/ent"
	"time"
)

func EntToOrder(e *ent.Order) *base.Order {
	if e == nil {
		return nil
	}

	mapper := utils.NewCopierMapper[base.Order, ent.Order]()
	var dto = mapper.ToDTO(e)

	dto.CreatedAt = e.CreatedAt.Format(time.DateOnly)

	return dto

}
func EntToOrderItem(e *ent.OrderItem) *base.OrderItem {
	if e == nil {
		return nil
	}

	mapper := utils.NewCopierMapper[base.OrderItem, ent.OrderItem]()
	var dto = mapper.ToDTO(e)

	return dto

}

func EntToOrderPay(e *ent.OrderPay) *base.OrderPay {
	if e == nil {
		return nil
	}

	mapper := utils.NewCopierMapper[base.OrderPay, ent.OrderPay]()
	var dto = mapper.ToDTO(e)

	return dto

}
func EntToOrderRefund(e *ent.OrderRefund) *base.OrderRefund {
	if e == nil {
		return nil
	}

	mapper := utils.NewCopierMapper[base.OrderRefund, ent.OrderRefund]()
	var dto = mapper.ToDTO(e)

	return dto

}
