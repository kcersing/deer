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

type MemberProduct struct {
	ent.Schema
}

func (MemberProduct) Fields() []ent.Field {

	return []ent.Field{
		field.String("sn").Comment("编号").Optional(),
		field.Int64("member_id").Comment("会员id").Optional(),
		field.Int64("product_id").Comment("产品ID").Optional(),
		field.Int64("venue_id").Comment("场馆ID").Optional(),
		field.Int64("order_id").Comment("订单ID").Optional(),
		field.String("name").Comment("产品名称").Optional(),
		field.Float("price").Comment("产品价格").Optional(),
	}
}

func (MemberProduct) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (MemberProduct) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("members", Member.Type).
			Ref("member_products").
			Field("member_id").Unique(),
		edge.To("member_product_propertys", MemberProductProperty.Type),
		edge.To("member_product_entry", VenueEntry.Type),
		edge.To("member_product_contents", MemberContract.Type),
	}
}

func (MemberProduct) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("venue_id"),
		index.Fields("member_id"),
		index.Fields("product_id"),
		index.Fields("order_id"),
	}
}

func (MemberProduct) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "member_product",
			//OnDelete: entsql.Cascade,
			Options: "AUTO_INCREMENT = 100000",
		},
		entsql.WithComments(true),
	}
}
