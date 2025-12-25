package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	_ "entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"system/biz/dal/db/ent/schema/mixins"
)

type Role struct {
	ent.Schema
}

func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("角色名").Optional(),
		field.String("code").Comment("角色标识").Unique(),

		field.String("desc").Default("").Comment("描述").Optional(),
		field.Int64("order_no").Default(0).Comment("排序编号").Optional(),

		field.JSON("menus", []int64{}).
			Comment("分配的菜单列表").
			Optional(),

		field.JSON("apis", []int64{}).
			Comment("分配的API列表").
			Optional(),
	}
}
func (Role) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
		index.Fields("code").Unique(),
	}
}
func (Role) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (Role) Edges() []ent.Edge {
	return []ent.Edge{}
}
func (Role) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_roles"},
		entsql.WithComments(true),
	}
}
