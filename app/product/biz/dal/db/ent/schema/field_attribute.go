package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"product/biz/dal/db/ent/schema/mixins"
)

type FieldAttribute struct {
	ent.Schema
}

func (FieldAttribute) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("field_id").Comment("类型").Optional(),
		field.String("name").Comment("属性名称").Optional(),
	}
}

func (FieldAttribute) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (FieldAttribute) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (FieldAttribute) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("name"),
	}
}

func (FieldAttribute) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Field_attribute"},
		entsql.WithComments(true),
	}
}
