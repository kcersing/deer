package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"user/biz/dal/db/ent/schema/mixins"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique().Comment("user's login name | 登录名"),
		field.String("password").Comment("password | 密码"),
		field.String("avatar").
			SchemaType(map[string]string{dialect.MySQL: "varchar(512)"}).
			Optional().
			Comment("avatar | 头像路径"),
		field.String("mobile").Unique().Comment("mobile number | 手机号"),
		field.String("name").Optional().Comment("姓名"),
		field.Int64("gender").Default(3).Comment("性别 | [0:女性;1:男性;3:保密]").Optional(),
		field.Time("birthday").Comment("出生日期").Optional(),
		field.Time("last_at").Comment("最后一次登录时间").Optional(),
		field.String("last_ip").Comment("最后一次登录ip").Optional(),

		field.String("detail").Comment("详情").Optional(),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("username"),
		index.Fields("mobile"),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "users"},
		entsql.WithComments(true),
	}
}
