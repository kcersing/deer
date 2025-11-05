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

type Dict struct {
	ent.Schema
}

func (Dict) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Comment("the title shown in the ui | 展示名称 （建议配合i18n）").Optional(),
		field.String("name").Unique().Comment("the name of dictionary for search | 字典搜索名称").Optional(),
		field.String("desc").Comment("the desc of dictionary | 字典描述").Optional(),
	}
}

func (Dict) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (Dict) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("dictht", Dictht.Type),
	}
}
func (Dict) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("title").Unique(),
		index.Fields("name").Unique(),
	}
}

func (Dict) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_dict"},
		entsql.WithComments(true),
	}
}
