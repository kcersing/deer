package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"product/biz/dal/db/ent/schema/mixins"
)

type FieldValue struct {
	ent.Schema
}

func (FieldValue) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("field_id").Comment("字段ID").Optional(),
		field.Int64("field_attribute_id").Comment("属性ID").Optional(),
		field.String("name").Comment("显示名称").Optional(),
		field.JSON("value", map[string]any{}).Comment("值").Optional(),
	}
}

func (FieldValue) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (FieldValue) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (FieldValue) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("name"),
	}
}

func (FieldValue) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Field_value"},
		entsql.WithComments(true),
	}
}
