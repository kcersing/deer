package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"user/biz/dal/db/ent/schema/mixins"
)

type UserLoginRestriction struct {
	ent.Schema
}

func (UserLoginRestriction) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("target_id").
			Comment("目标用户ID").
			Optional().
			Nillable(),

		field.String("value").
			Comment("限制值（如IP地址、MAC地址或地区代码）").
			Optional().
			Nillable(),

		field.String("reason").
			Comment("限制原因").
			Optional().
			Nillable(),

		field.Enum("type").
			Comment("限制类型").
			NamedValues(
				"Blacklist", "BLACKLIST",
				"Whitelist", "WHITELIST",
			).
			Default("BLACKLIST").
			Optional().
			Nillable(),

		field.Enum("method").
			Comment("限制方式").
			NamedValues(
				"Ip", "IP",
				"Mac", "MAC",
				"Region", "REGION",
				"Time", "TIME",
				"Device", "DEVICE",
			).
			Default("IP").
			Optional().
			Nillable(),
	}
}

func (UserLoginRestriction) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("target_id", "type", "method").Unique().StorageKey("idx_sys_admin_login_restriction_target_type_method"),
	}
}

// Mixin of the UserLoginRestriction.
func (UserLoginRestriction) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (UserLoginRestriction) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "user_login_restrictions",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("后台登录限制表"),
	}
}
