package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"user/biz/dal/db/ent/schema/mixins"
)

// Department holds the schema definition for the Department entity.
type Department struct {
	ent.Schema
}

func (Department) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "user_departments",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("部门表"),
	}
}

// Fields of the Department.
func (Department) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("部门名称").
			NotEmpty().
			Optional().
			Nillable(),

		field.Int64("manager_id").
			Comment("负责人ID").
			Optional().
			Nillable(),

		field.Int64("parent_id").
			Optional().
			Comment("父ID"),

		field.String("description").
			Comment("职能描述").
			Optional().
			Nillable(),
	}
}

// Mixin of the Department.
func (Department) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

// Indexes of the Department.
func (Department) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").StorageKey("idx_sys_department_name"),
	}
}
