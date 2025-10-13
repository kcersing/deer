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

type ProductProperty struct {
	ent.Schema
}

func (ProductProperty) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").Comment("类型").Optional(),
		field.String("name").Comment("名称").Optional(),
		field.Int64("duration").Comment("总时长").Optional(),
		field.Int64("length").Comment("单次时长").Optional(),
		field.Int64("count").Comment("次数").Optional(),
		field.Float("price").Comment("定价").Optional(),
		field.String("data").Comment("").Optional(),
		field.String("pic").Default("").Comment("主图").Optional(),
		field.Text("description").Default("").Comment("详情").Optional(),
	}
}

func (ProductProperty) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (ProductProperty) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("product", Product.Type).Ref("propertys"),
		edge.To("tags", DictionaryDetail.Type),
		edge.To("contracts", Contract.Type),
		edge.To("venues", Venue.Type),
	}
}

func (ProductProperty) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("name"),
	}
}

func (ProductProperty) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "product_property"},
		entsql.WithComments(true),
	}
}
