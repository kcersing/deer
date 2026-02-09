package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"order/biz/dal/db/ent/schema/mixins"
)

type Order struct {
	ent.Schema
}

func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.String("sn").Comment("订单编号").Unique(),
		field.Int64("member_id").Comment("会员id").Optional(),
		field.Enum("status").
			Values(
				"created",
				"paid",
				"shipped",
				"cancelled",
				"completed",
				"refunded",
			).
			Default("created").Comment("状态").Optional(),

		field.Int64("nature").Comment("业务类型").Optional(),
		field.Time("completion_at").Comment("订单完成时间").Optional(),
		field.Time("close_at").Comment("订单关闭时间").Optional(),

		field.Int64("version").Default(1).Comment("乐观锁版本号").Optional(),

		field.Int64("total_amount").Default(0).Comment("总金额").Optional(),
		field.Int64("actual").Default(0).Comment("实际已付款").Optional(),
		field.Int64("remission").Default(0).Comment("减免").Optional(),
		field.String("close_nature").Comment("关闭原因").Optional(),
	}
}

func (Order) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (Order) Edges() []ent.Edge {
	return []ent.Edge{

		edge.To("items", OrderItem.Type),
		edge.To("pay", OrderPay.Type),
		edge.To("eventbus", OrderEvents.Type),
		edge.To("snapshots", OrderSnapshots.Type),
		edge.To("status_history", OrderStatusHistory.Type),

		//edge.From("order_venues", Venue.Type).Ref("venue_orders").Field("venue_id").Unique(),
		edge.To("refund", OrderRefund.Type),
		//edge.From("order_creates", User.Type).Ref("created_orders").Field("created_id").Unique(),
	}
}

func (Order) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("sn"),
		//index.Fields("venue_id"),
		index.Fields("member_id"),
		index.Fields("status"),
		index.Fields("completion_at"),
		//index.Fields("member_product_id"),
	}
}

func (Order) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment(""),
	}
}
