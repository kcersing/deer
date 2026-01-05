package schema

import (
	"entgo.io/ent/schema/index"
	"member/biz/dal/db/ent/schema/mixins"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type MemberContractContent struct {
	ent.Schema
}

func (MemberContractContent) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("member_contract_id").Comment("合同ID").Optional(),
		field.String("content").Optional().Comment("content | 内容"),
		field.Text("sign_img").Optional().Comment("sign_img | 会员签字b64 预处理"),
	}
}

func (MemberContractContent) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (MemberContractContent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("contract", MemberContract.Type).Ref("content").Field("member_contract_id").Unique(),
	}
}

func (MemberContractContent) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("member_contract_id"),
	}
}

func (MemberContractContent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "member_contract_content",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
	}
}
