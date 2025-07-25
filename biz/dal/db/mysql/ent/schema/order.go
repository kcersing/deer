package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"kcers-order/biz/dal/db/mysql/ent/schema/mixins"
)

type Order struct {
	ent.Schema
}

func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.String("order_sn").Comment("订单编号").Optional(),
		//field.Int64("venue_id").Comment("场馆id").Optional(),
		//field.Int64("member_id").Comment("会员id").Optional(),
		//field.Int64("member_product_id").Comment("会员产品id").Optional(),
		field.Int64("status").Default(0).Optional(),
		field.String("source").Default("").Comment("订单来源").Optional(),
		field.String("device").Default("").Comment("设备来源").Optional(),
		field.Int64("nature").Comment("业务类型").Optional(),
		field.Time("completion_at").Comment("订单完成时间").Optional(),
		field.Time("close_at").Comment("订单关闭时间").Optional(),
		field.Time("refund_at").Comment("订单退费时间").Optional(),
	}
}

func (Order) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.To("amount", OrderAmount.Type),
		//edge.To("item", OrderItem.Type),
		//edge.To("pay", OrderPay.Type),
		//edge.To("order_contents", MemberContract.Type),
		//edge.To("sales", OrderSales.Type),

		//edge.From("order_venues", Venue.Type).Ref("venue_orders").Field("venue_id").Unique(),
		//edge.From("order_members", Member.Type).Ref("member_orders").Field("member_id").Unique(),
		//edge.From("order_creates", User.Type).Ref("created_orders").Field("created_id").Unique(),
	}
}

func (Order) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("order_sn"),
		//index.Fields("venue_id"),
		//index.Fields("member_id"),
		index.Fields("completion_at"),
		//index.Fields("member_product_id"),
	}
}

func (Order) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order"},
		entsql.WithComments(true),
	}
}
