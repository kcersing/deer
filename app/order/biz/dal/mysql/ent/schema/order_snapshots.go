package schema

import (
	"deer/app/order/biz/dal/mysql/ent/schema/mixins"
	"entgo.io/ent/schema/edge"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type OrderSnapshots struct {
	ent.Schema
}

func (OrderSnapshots) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("aggregate_id").Comment("聚合根ID").Optional(),
		field.Int64("aggregate_version").Comment("快照版本").Optional(),
		field.Bytes("aggregate_data").Comment("快照数据").Optional(),
	}
}

func (OrderSnapshots) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (OrderSnapshots) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).
			Ref("snapshots").
			Field("aggregate_id").Unique(),
	}
}

func (OrderSnapshots) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("aggregate_id", "aggregate_version"),
	}
}

func (OrderSnapshots) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order_snapshots", Collation: "utf8mb4_unicode_ci"},
		entsql.WithComments(true),
	}
}
