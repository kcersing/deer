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

type API struct {
	ent.Schema
}

func (API) Fields() []ent.Field {
	return []ent.Field{
		field.String("path").Comment("路径").Optional(),
		field.String("title").Comment("API 名称").Optional(),
		field.String("desc").Comment("API 描述").Optional(),
		field.String("group").Comment("API 分组").Optional(),
		field.String("method").Default("POST").Comment("HTTP 请求类型").Optional(),
	}
}

func (API) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (API) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("roles", Role.Type).Ref("api"),
	}
}

func (API) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("path", "method").Unique(),
	}
}

func (API) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_apis"},
		entsql.WithComments(true),
	}
}
