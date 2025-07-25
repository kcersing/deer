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

type Event struct {
	ent.Schema
}

func (Event) Fields() []ent.Field {
	return []ent.Field{

		field.Int64("event_id").Comment("事件id").Optional(),
		field.Int64("aggregate_id").Comment("聚合根ID").Optional(),
		field.String("aggregate_type").Comment("聚合根类型").Optional(),
		field.String("event_type").Comment("事件类型").Optional(),
		field.Text("event_data").Comment("事件数据").Optional(),
		field.Int64("event_version").Comment("聚合根版本号").Optional(),
	}
}

func (Event) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (Event) Edges() []ent.Edge {
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

func (Event) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("order_sn"),
		index.Fields("venue_id"),
		index.Fields("member_id"),
		index.Fields("completion_at"),
		index.Fields("member_product_id"),
	}
}

func (Event) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order_views"},
		entsql.WithComments(true),
	}
}
