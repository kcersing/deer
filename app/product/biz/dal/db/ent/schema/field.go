package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"product/biz/dal/db/ent/schema/mixins"
)

type Field struct {
	ent.Schema
}

func (Field) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("显示名称").Optional(),
		field.String("code").Comment("标识").Optional(),
		field.Int64("type").Comment("类型").Optional(),
		field.Int64("required").Comment("是否必填").Optional(),
		field.Int64("order_on").Comment("排序").Optional(),
		field.JSON("value", map[string]any{}).Comment("值").Optional(),
	}
}

func (Field) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (Field) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Field) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("name"),
	}
}

func (Field) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Field"},
		entsql.WithComments(true),
	}
}
