package schema

import (
	"product/biz/dal/db/ent/schema/mixins"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type ProductField struct {
	ent.Schema
}

func (ProductField) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("product_item_id").Comment("项Id").Optional(),
		field.String("name").Comment("显示名称").Optional(),
		field.String("code").Comment("标识").Optional(),
		field.Enum("type").Values("text", "textarea", "select", "radio", "checkbox", "date", "number").
			Default("text").Comment("类型").Optional(),
		field.Int64("required").Comment("是否必填").Optional(),
		field.Int64("order_on").Comment("排序").Optional(),
		field.Any("value").Comment("值").Optional(),
	}
}

func (ProductField) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (ProductField) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.From("item", ProductItem.Type).Ref("fields"),
	}
}

func (ProductField) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("name"),
	}
}

func (ProductField) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "product_fields",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
	}
}
