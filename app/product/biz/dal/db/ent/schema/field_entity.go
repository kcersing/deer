package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"product/biz/dal/db/ent/schema/mixins"
)

type FieldEntity struct {
	ent.Schema
}

func (FieldEntity) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("显示名称").Optional(),
		field.String("code").Comment("标识").Optional(),
		field.Int64("type").Comment("类型").Optional(),
		field.Int64("required").Comment("是否必填").Optional(),
		field.Int64("order_on").Comment("排序").Optional(),
		field.JSON("value", map[string]any{}).Comment("值").Optional(),
	}
}

func (FieldEntity) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (FieldEntity) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (FieldEntity) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("name"),
	}
}

func (FieldEntity) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Field_entity"},
		entsql.WithComments(true),
	}
}
