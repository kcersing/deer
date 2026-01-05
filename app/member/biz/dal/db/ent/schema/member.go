package schema

import (
	"member/biz/dal/db/ent/schema/mixins"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Member struct {
	ent.Schema
}

func (Member) Fields() []ent.Field {
	return []ent.Field{
		field.String("password").Optional().Comment("password | 密码"),
		field.String("username").Unique().Comment("账号 ").Optional(),
		field.String("mobile").Optional().Comment("mobile number | 手机号"),
		field.String("name").Optional().Comment("name | 名称"),
		field.String("avatar").
			SchemaType(map[string]string{dialect.MySQL: "varchar(512)"}).
			Optional().
			Default("").
			Comment("avatar | 头像路径"),
		field.Int64("condition").
			Default(0).
			Optional().
			Comment("状态[0:潜在;1:正式;3:冻结;4:到期]"),
	}
}

func (Member) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (Member) Edges() []ent.Edge {
	return []ent.Edge{

		edge.To("member_profile", MemberProfile.Type),

		edge.To("member_notes", MemberNote.Type),

		edge.To("member_products", MemberProduct.Type),

		edge.To("member_contents", MemberContract.Type),
	}
}

func (Member) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("username"),
		index.Fields("mobile"),
	}
}

func (Member) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "member",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
	}
}
