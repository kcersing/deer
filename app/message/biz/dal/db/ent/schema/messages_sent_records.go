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
		field.Int64("message_id").
			Comment("").
			Optional().
			Nillable(),
		field.Int64("to_user_id").
			Comment("接收者用户ID").
			Optional().
			Nillable(),

		field.Enum("status").
			Comment("消息状态").
			NamedValues(
				"Sent", "SENT",
				"Received", "RECEIVED",
				"Read", "READ",
				"Revoked", "REVOKED",
				"Deleted", "DELETED",
			).
			Optional().
			Nillable(),

		field.Time("received_at").
			Comment("消息到达用户收件箱的时间").
			Optional().
			Nillable(),

		field.Time("read_at").
			Comment("用户阅读消息的时间").
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
		index.Fields("to_user_id"),
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
