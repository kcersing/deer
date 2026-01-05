package schema

import (
	"deer/app/order/biz/dal/mysql/ent/schema/mixins"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type OrderRefund struct {
	ent.Schema
}

func (OrderRefund) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("order_id").Comment("订单id").Optional(),
		field.Time("refund_at").Comment("订单退费时间").Optional(),
		field.Float("refund").Default(0).Comment("退费金额").Optional(),
		field.String("refund_nature").Comment("退费原因").Optional(),
	}
}

func (OrderRefund) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (OrderRefund) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).
			Ref("refund").
			Field("order_id").Unique(),
	}
}

func (OrderRefund) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("order_id"),
	}
}

func (OrderRefund) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order_refund",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
	}
}
