package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"product/biz/dal/db/ent/schema/mixins"
)

type Product struct {
	ent.Schema
}

func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("商品名").Optional(),
		field.String("pic").Comment("主图").Optional(),
		field.String("desc").Comment("详情").Optional(),
		field.Float("price").Comment("价格").Optional(),
		field.Int64("stock").Comment("库存").Optional(),
		field.JSON("is_sales", []int64{}).Comment("销售方式 1会员端 2PC端").Optional(),
		field.Time("sign_sales_at").Comment("开始售卖时间").Optional(),
		field.Time("end_sales_at").Comment("结束售卖时间").Optional(),
	}
}

func (Product) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("propertys", ProductProperty.Type),
	}
}

func (Product) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
	}
}

func (Product) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "product"},
		entsql.WithComments(true),
	}
}
