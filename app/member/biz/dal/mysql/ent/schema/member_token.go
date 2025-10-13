package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"kcers/biz/dal/db/mysql/ent/schema/mixins"
)

type MemberToken struct {
	ent.Schema
}

func (MemberToken) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("member_id").Comment(" User's ID | 用户的ID").Unique(),
		field.String("token").Comment("Token string | Token 字符串"),
		field.String("source").Comment("Log in source such as GitHub | Token 来源 （本地为core, 第三方如github等）"),
		field.Time("expired_at").Comment(" Expire time | 过期时间"),
	}
}

func (MemberToken) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}
func (MemberToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Member.Type).Ref("token").Unique(),
	}
}
func (MemberToken) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("member_id"),
		index.Fields("expired_at"),
	}
}

func (MemberToken) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "member_token"},
		entsql.WithComments(true),
	}
}
