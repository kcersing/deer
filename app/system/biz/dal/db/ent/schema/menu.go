package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"system/biz/dal/db/ent/schema/mixins"
)

type Menu struct {
	ent.Schema
}

func (Menu) Fields() []ent.Field {
	return []ent.Field{
		field.String("path").Optional().Default("").Comment("菜单路由路径"),
		field.String("name").Comment("菜单名称").Optional(),
		field.String("component").Optional().Default("").Comment(" 组件路径"),
		field.String("redirect").Optional().Default("").Comment("跳转路径 （外链）"),
		field.String("icon").Comment("菜单图标").Optional(),

		field.Int64("parent_id").Optional().Comment("父菜单ID"),

		field.Int64("order_no").Comment("排序编号").Optional(),

		field.Int64("ignore").Optional().Comment("当前路由是否渲染菜单项，为 1 的话不会在菜单中显示，但可通过路由地址访问"),

		field.Int64("menu_type").Comment("菜单类型 0 目录 1 菜单 2 按钮").Optional(),
		field.Int64("level").Comment("菜单层级").Optional(),
	}
}

func (Menu) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (Menu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Menu.Type).From("parent").Unique().Field("parent_id"),
	}
}
func (Menu) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("path").Unique(),
	}
}
func (Menu) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_menus",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("菜单表"),
	}
}
