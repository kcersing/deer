package schema

import (
	"member/biz/dal/db/ent/schema/mixins"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type MemberProfile struct {
	ent.Schema
}

func (MemberProfile) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("member_id").Comment("会员id").Optional(),
		field.Int64("intention").Comment("意向").Optional(),
		field.Int64("source").Default(0).Comment("来源").Optional(),

		field.Int64("gender").Default(3).Comment("性别 | [1:女性;2:男性;3:保密]").Optional(),
		field.Time("birthday").Comment("出生日期").Optional(),
		field.String("email").Optional().Comment("email | 邮箱号"),
		field.String("wecom").Optional().Comment("wecom | 微信号"),

		field.Int64("relation_mid").Default(0).Comment("关联会员").Optional(),
		field.String("relation_mame").Comment("关联会员").Optional(),
	}
}

func (MemberProfile) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (MemberProfile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("member", Member.Type).
			Ref("member_profile").
			Field("member_id").
			Unique(),
	}
}

func (MemberProfile) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("member_id"),
	}
}

func (MemberProfile) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "member_profile"},
		entsql.WithComments(true),
	}
}
