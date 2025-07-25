package new

import (
	"kcers/biz/dal/db/mysql/ent/schema/mixins"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type OrderEvents struct {
	ent.Schema
}

func (OrderEvents) Fields() []ent.Field {
	return []ent.Field{
		field.String("event_id").Comment("事件标识").Optional(),
		field.String("order_sn").Comment("订单编号").Optional(),
		field.String("event_type").Comment("事件类型").Optional(),
		field.String("event_data").Comment("事件数据").Optional(),
		field.Int64("version").Comment("事件版本号").Optional(),
	}
}

func (OrderEvents) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (OrderEvents) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("amount", OrderAmount.Type),
		edge.To("item", OrderItem.Type),
		edge.To("pay", OrderPay.Type),
		edge.To("order_contents", MemberContract.Type),
		edge.To("sales", OrderSales.Type),

		edge.From("order_venues", Venue.Type).Ref("venue_orders").Field("venue_id").Unique(),
		edge.From("order_members", Member.Type).Ref("member_orders").Field("member_id").Unique(),
		edge.From("order_creates", User.Type).Ref("created_orders").Field("created_id").Unique(),
	}
}

func (OrderEvents) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("order_sn"),
		index.Fields("event_id"),
	}
}

func (OrderEvents) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order_events"},
		entsql.WithComments(true),
	}
}
