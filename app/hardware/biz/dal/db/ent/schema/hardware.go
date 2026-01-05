package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"hardware/biz/dal/db/ent/schema/mixins"
)

type Hardware struct {
	ent.Schema
}

func (Hardware) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional().Comment("name | 名称"),
		field.String("pic").
			SchemaType(map[string]string{dialect.MySQL: "varchar(512)"}).
			Optional().
			Default("").
			Comment("pic"),
	}
}

func (Hardware) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (Hardware) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Hardware) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
	}
}

func (Hardware) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "hardware",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
	}
}
