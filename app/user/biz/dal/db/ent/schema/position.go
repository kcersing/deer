package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"user/biz/dal/db/ent/schema/mixins"
)

// Position holds the schema definition for the Position entity.
type Position struct {
	ent.Schema
}

func (Position) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "user_positions",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("职位表"),
	}
}

// Fields of the Position.
func (Position) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("职位名称").
			//Unique().
			NotEmpty().
			Optional().
			Nillable(),

		field.String("code").
			Comment("唯一编码").
			//Unique().
			NotEmpty().
			Optional().
			Nillable(),

		field.Int64("department_id").
			Comment("所属部门ID").
			Nillable(),

		field.Int64("parent_id").
			Optional().
			Comment("父ID"),

		field.String("description").
			Comment("职能描述").
			Optional().
			Nillable(),

		field.Int64("quota").
			Comment("编制人数").
			Optional().
			Nillable(),
	}
}

// Mixin of the Position.
func (Position) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

// Indexes of the Position.
func (Position) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code").Unique().StorageKey("idx_sys_position_code"),
		index.Fields("name").StorageKey("idx_sys_position_name"),
	}
}
