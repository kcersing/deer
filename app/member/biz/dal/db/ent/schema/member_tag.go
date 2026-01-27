package schema

import (
	"member/biz/dal/db/ent/schema/mixins"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type MemberTag struct {
	ent.Schema
}

func (MemberTag) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("member_id").Comment("会员id").Optional(),
		field.Int64("tag_id").Comment("标签").Optional(),
		field.Int64("weight").Comment("权重").Optional(),
	}
}

func (MemberTag) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (MemberTag) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (MemberTag) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("tag_id"),
		index.Fields("member_id"),
	}
}

func (MemberTag) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "member_tags",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment(""),
	}
}
