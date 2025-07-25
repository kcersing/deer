package new

import (
	"entgo.io/ent/schema/index"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"kcers-order/biz/dal/db/mysql/ent/schema/mixins"
)

type OrderItem struct {
	ent.Schema
}

func (OrderItem) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("order_id").Comment("订单id").Optional(),
		field.Int64("number").Default(1).Comment("数量").Optional(),
		field.String("name").Comment("名称").Optional(),
		field.Int64("product_id").Comment("产品id").Optional(),
		field.Int64("related_user_product_id").Default(0).Comment("关联会员产品id").Optional(),
		field.Float("unit_price").Comment("单价").Optional(),
		field.String("sku").Comment("sku").Optional(),
	}
}

func (OrderItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (OrderItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).
			Ref("item").
			Field("order_id").Unique(),
	}
}

func (OrderItem) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("order_id"),
		index.Fields("product_id"),
	}
}

func (OrderItem) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order_item"},
		entsql.WithComments(true),
	}
}
