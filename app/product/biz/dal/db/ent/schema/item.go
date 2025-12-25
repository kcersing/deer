package schema

import (
	"product/biz/dal/db/ent/schema/mixins"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Item struct {
	ent.Schema
}

func (Item) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称").Optional(),
		field.String("code").Comment("标识").Optional(),
		field.String("pic").Default("").Comment("主图").Optional(),
		field.String("desc").Default("").Comment("概述").Optional(),
		field.String("type").Comment("类型").Optional(),
		field.Int64("duration").Comment("总时长").Optional(),
		field.Int64("length").Comment("单次时长").Optional(),
		field.Int64("count").Comment("次数").Optional(),
		field.Int64("price").Comment("定价").Optional(),
		//field.Time("active_at").Comment("激活时间").Optional(),
		//field.Time("expired_at").Comment("到期时间").Optional(),
		field.JSON("tag_id", []int64{}).Comment("标签").Optional(),
	}
}

func (Item) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (Item) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.To("fields", Fields.Type),
		edge.From("product", Product.Type).Ref("items"),
	}
}

func (Item) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("name", "code"),
	}
}

func (Item) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "product_items"},
		entsql.WithComments(true),
	}
}
