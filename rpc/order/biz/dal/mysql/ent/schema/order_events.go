package schema

import (
	"deer/rpc/order/biz/dal/mysql/ent/schema/mixins"
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

		field.String("event_id").Comment("事件id").Optional(),
		field.Int64("aggregate_id").Comment("聚合根ID").Optional(),
		field.String("aggregate_type").Comment("聚合根类型").Optional(),
		field.String("event_type").Comment("事件类型").Optional(),

		field.Bytes("event_data").Comment("事件数据").Optional(),
		field.Int64("event_version").Comment("聚合根版本号").Optional(),
	}
}

func (OrderEvents) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (OrderEvents) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).
			Ref("events").
			Field("aggregate_id").Unique(),
	}
}

func (OrderEvents) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("event_type"),
		index.Fields("aggregate_id", "event_version"),
	}
}

func (OrderEvents) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order_events", Collation: "utf8mb4_unicode_ci"},
		entsql.WithComments(true),
	}
}
