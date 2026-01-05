package schema

import (
	"system/biz/dal/db/ent/schema/mixins"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type MessagesSentRecords struct {
	ent.Schema
}

func (MessagesSentRecords) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Comment("消息标题").Optional().Nillable(),
		field.String("type").Comment("类型[1:用户user;2:会员member]").Optional(),
		field.String("to_user_id").Comment("该消息接受者ID").Optional().Nillable(),
		field.String("from_user_id").Comment("该消息发送者ID").Optional().Nillable(),
		field.String("content").Comment("消息内容").Optional().Nillable(),
		field.Enum("status").
			Comment("消息状态").
			NamedValues(
				"Draft", "DRAFT",
				"Published", "PUBLISHED",
				"Scheduled", "SCHEDULED",
				"Revoked", "REVOKED",
				"Archived", "ARCHIVED",
				"Deleted", "DELETED",
			).
			Default("DRAFT").
			Optional().
			Nillable(),
		field.Enum("type").
			Comment("消息类型").
			NamedValues(
				"Notification", "NOTIFICATION",
				"Private", "PRIVATE",
				"Group", "GROUP",
			).
			Default("NOTIFICATION").
			Optional().
			Nillable(),
	}
}

func (MessagesSentRecords) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (MessagesSentRecords) Edges() []ent.Edge {
	return nil
}

func (MessagesSentRecords) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("to_user_id", "from_user_id"),
	}
}

func (MessagesSentRecords) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "messages_sent_records",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
	}
}
