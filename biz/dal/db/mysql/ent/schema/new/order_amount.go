package new

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"kcers-order/biz/dal/db/mysql/ent/schema/mixins"
)

type OrderAmount struct {
	ent.Schema
}

func (OrderAmount) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("order_id").NonNegative().Comment("订单id").Optional(),
		field.Float("total").Comment("总金额").Optional(),
		field.Float("actual").Comment("实际已付款").Optional(),
		field.Float("residue").Comment("未支付金额").Optional(),
		field.Float("remission").Comment("减免").Optional(),
		field.Float("refund").Comment("退费金额").Optional(),
		field.String("refund_reason").Comment("退费原因").Optional(),
	}
}

func (OrderAmount) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (OrderAmount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).
			Ref("amount").
			Field("order_id").Unique(),
	}
}
func (OrderAmount) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("order_id"),
	}
}

func (OrderAmount) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order_amount"},
		entsql.WithComments(true),
	}
}
