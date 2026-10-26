package schema

import (
	"order/biz/dal/db/ent/schema/mixins"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// OrderCache 缓存表
type OrderCache struct {
	ent.Schema
}

func (OrderCache) Fields() []ent.Field {
	return []ent.Field{
		field.Text("key").Comment("key").Optional(),
		field.JSON("value", map[string]interface{}{}).Comment("value").Optional(),
	}
}

func (OrderCache) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (OrderCache) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (OrderCache) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("key"),
	}
}

func (OrderCache) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "order_cache",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},

		entsql.WithComments(true),
		schema.Comment(""),
	}
}
