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
		field.String("title").Comment("名称").Optional(),
		field.String("code").Unique().Comment("字典唯一代码").Optional(),
		field.String("desc").Comment("字典描述").Optional(),
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
		index.Fields("code").Unique(),
	}
}

func (Dict) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_dict",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("字典表"),
	}
}
