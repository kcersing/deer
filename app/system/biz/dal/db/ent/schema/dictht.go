package schema

import (
	"system/biz/dal/db/ent/schema/mixins"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Dictht struct {
	ent.Schema
}

func (Dictht) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Comment("展示名称").Optional(),
		field.String("value").Comment("值").Optional(),
		field.Int64("dict_id").Optional().Comment("字典ID").Optional(),
	}
}

func (Dictht) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (Dictht) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("dict", Dict.Type).
			Field("dict_id").
			Ref("dictht").
			Unique(),
	}
}

func (Dictht) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("title"),
		index.Fields("value"),
		index.Fields("dict_id"),
	}
}

func (Dictht) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_dictht"},
		entsql.WithComments(true),
	}
}
