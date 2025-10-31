package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"system/biz/dal/db/ent/schema/mixins"
)

type Menu struct {
	ent.Schema
}

// path,component,routes,redirect,wrappers,name,icon
func (Menu) Fields() []ent.Field {
	return []ent.Field{
		field.String("path").Optional().Default("").Comment("index path | 菜单路由路径"),
		field.String("name").Comment("index name | 菜单名称").Optional(),
		field.String("component").Optional().Default("").Comment("the path of vue file | 组件路径"),
		field.String("redirect").Optional().Default("").Comment("redirect path | 跳转路径 （外链）"),
		field.String("icon").Comment("menu icon | 菜单图标").Optional(),

		field.Int64("parent_id").Optional().Comment("parent menu ID | 父菜单ID"),

		field.Int64("order_no").Comment("sorting numbers | 排序编号").Optional(),

		field.Int64("ignore").Optional().Comment("当前路由是否渲染菜单项，为 1 的话不会在菜单中显示，但可通过路由地址访问"),

		field.Int64("menu_type").Comment("menu type | 菜单类型 0 目录 1 菜单 2 按钮").Optional(),
		field.Int64("level").Comment("menu level | 菜单层级").Optional(),
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
		edge.From("roles", Role.Type).Ref("menus"),
		edge.To("children", Menu.Type).From("parent").Unique().Field("parent_id"),
	}
}

func (Menu) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_menus"},
		entsql.WithComments(true),
	}
}
