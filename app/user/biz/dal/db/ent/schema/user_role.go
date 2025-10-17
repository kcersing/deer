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
func (UserRole) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_roles"},
		entsql.WithComments(true),
	}
}
