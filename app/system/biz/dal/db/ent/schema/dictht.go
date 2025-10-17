package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"system/biz/dal/db/ent/schema/mixins"
)

type Dictht struct {
	ent.Schema
}

func (Dictht) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Comment("the title shown in the ui | 展示名称 （建议配合i18n）"),
		field.String("key").Comment("key | 键"),
		field.String("value").Comment("value | 值"),
		field.Int64("dict_id").Optional().Comment("Dictionary ID | 字典ID"),
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
		index.Fields("title").Unique(),
		index.Fields("key").Unique(),
		index.Fields("value").Unique(),
		index.Fields("dict_id").Unique(),
	}
}

func (Dictht) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_dictht"},
		entsql.WithComments(true),
	}
}
