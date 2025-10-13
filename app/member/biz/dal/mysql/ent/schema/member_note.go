package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"kcers/biz/dal/db/mysql/ent/schema/mixins"
)

type MemberNote struct {
	ent.Schema
}

func (MemberNote) Fields() []ent.Field {
	return []ent.Field{

		field.Int64("member_id").Comment("会员id").Optional(),

		field.String("note").
			SchemaType(map[string]string{dialect.MySQL: "varchar(512)"}).
			Optional().
			Default("").
			Comment("内部备注"),
	}
}

func (MemberNote) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (MemberNote) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("notes", Member.Type).
			Ref("member_notes").Unique().
			Field("member_id").Unique(),
	}
}

func (MemberNote) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("member_id"),
	}
}

func (MemberNote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "member_note"},
		entsql.WithComments(true),
	}
}
