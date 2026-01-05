package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"member/biz/dal/db/ent/schema/mixins"
)

type MemberContract struct {
	ent.Schema
}

func (MemberContract) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("member_id").Comment("会员id").Optional(),
		field.Int64("contract_id").Comment("原始合同id").Optional(),
		field.Int64("order_id").Comment("订单id").Optional(),
		field.Int64("venue_id").Comment("场馆id").Optional(),
		field.Int64("member_product_id").Comment("会员产品id").Optional(),
		field.String("name").Optional().Comment("name | 名称"),
		field.String("sign").Optional().Comment("sign | 签字"),
	}
}

func (MemberContract) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (MemberContract) Edges() []ent.Edge {
	return []ent.Edge{

		edge.To("content", MemberContractContent.Type),
		edge.From("member", Member.Type).Ref("member_contents").Field("member_id").Unique(),
		edge.From("member_product", MemberProduct.Type).Ref("member_product_contents").Field("member_product_id").Unique(),
	}
}

func (MemberContract) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("venue_id"),
		index.Fields("order_id"),
		index.Fields("member_id"),
		index.Fields("member_product_id"),
	}
}

func (MemberContract) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "member_contract",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
	}
}
