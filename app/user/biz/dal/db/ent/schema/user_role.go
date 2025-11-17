package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	_ "entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type UserRole struct {
	ent.Schema
}

func (UserRole) Fields() []ent.Field {
	return []ent.Field{

		field.Int64("user_id").Comment("user id"),
		field.Int64("role_id").Default(0).Comment("role id"),
	}
}

func (UserRole) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}

func (UserRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("user_role").
			Field("user_id").Required().
			Unique(),
	}
}
func (UserRole) Indexes() []ent.Index {
	return []ent.Index{
		// 避免用户重复分配同一角色
		index.Fields("user_id", "role_id").Unique().StorageKey("idx_sys_user_role_user_id_role_id"),
		index.Fields("user_id").StorageKey("idx_sys_user_role_user_id"),
		index.Fields("role_id").StorageKey("idx_sys_user_role_role_id"),
	}
}

func (UserRole) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "user_roles",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("用户角色表"),
	}
}
