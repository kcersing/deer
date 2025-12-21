package schema

import (
	"product/biz/dal/db/ent/schema/mixins"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Fields struct {
	ent.Schema
}

func (Fields) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("product_item_id").Comment("项Id").Optional(),
		field.String("name").Comment("显示名称").Optional(),
		field.String("code").Comment("标识").Optional(),
		field.Enum("type").Values("text", "textarea", "select", "radio", "checkbox", "date", "number").
			Default("text").Comment("类型").Optional(),
		field.Int64("required").Comment("是否必填").Optional(),
		field.Int64("order_on").Comment("排序").Optional(),
		field.Any("value").Comment("值").Optional(),
	}
}

func (Fields) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (Fields) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("item", Item.Type).Ref("fields"),
	}
}

func (Fields) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("name"),
	}
}

func (Fields) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "fields"},
		entsql.WithComments(true),
	}
}
