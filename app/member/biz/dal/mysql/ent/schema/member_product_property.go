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

type MemberProductProperty struct {
	ent.Schema
}

func (MemberProductProperty) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("member_id").Comment("会员id").Optional(),
		field.Int64("member_product_id").Comment("会员产品ID").Optional(),
		field.String("sn").Comment("编号").Optional(),
		field.Int64("property_id").Comment("属性ID").Optional(),
		field.String("type").Comment("类型").Optional(),
		field.String("name").Comment("名称").Optional(),
		field.Int64("duration").Comment("总时长").Optional(),
		field.Int64("length").Comment("单次时长").Optional(),
		field.Int64("count").Default(0).Comment("总次数").Optional(),
		field.Int64("count_surplus").Default(0).Comment("剩余次数").Optional(),
		field.Float("price").Comment("定价").Optional(),
		field.Time("validity_at").Comment("生效时间").Optional(),
		field.Time("cancel_at").Comment("作废时间").Optional(),
	}
}

func (MemberProductProperty) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (MemberProductProperty) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", MemberProduct.Type).Ref("member_product_propertys").
			Field("member_product_id").Unique(),
		edge.To("venues", Venue.Type),
	}
}

func (MemberProductProperty) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("property_id"),
		index.Fields("member_id"),
		index.Fields("member_product_id"),
		index.Fields("validity_at"),
		index.Fields("cancel_at"),
	}
}

func (MemberProductProperty) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "member_product_property"},
		entsql.WithComments(true),
	}
}
