package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"kcers-order/biz/dal/db/mysql/ent/schema/mixins"
)

type OrderStatusHistory struct {
	ent.Schema
}

func (OrderStatusHistory) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("order_id").Comment("订单ID").Optional(),
		field.Int64("old_status").Comment("旧状态").Optional(),
		field.Int64("new_status").Comment("新状态").Optional(),
		field.String("change_source").Comment("变更来源").Optional(),
		field.String("change_reason").Comment("变更原因").Optional(),
		field.Time("change_at").Comment("变更时间").Optional(),
	}
}

func (OrderStatusHistory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (OrderStatusHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).
			Ref("status_history").
			Field("order_id").Unique(),
	}
}

func (OrderStatusHistory) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
	}
}

func (OrderStatusHistory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order_status_history"},
		entsql.WithComments(true),
	}
}
