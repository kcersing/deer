package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type MessagesSentRecords struct {
	ent.Schema
}

func (MessagesSentRecords) Fields() []ent.Field {
	return []ent.Field{

		field.Time("created_at").
			Immutable().
			Default(time.Now).
			Comment("created time").
			Optional(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("last update time").
			Optional(),
		field.Int64("delete").
			Default(0).
			Comment("last delete  1:已删除 0:未删除").
			Optional(),
		field.Int64("created_id").
			Default(0).
			Comment("created").
			Optional(),

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

		field.Int64("type").
			Comment("消息类型[1会员;2员工]").
			Default(1).
			Optional().
			Nillable(),
	}
}

func (MessagesSentRecords) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}

func (MessagesSentRecords) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("messages", Messages.Type).Ref("sent_records").Field("message_id").Unique(),
	}
}

func (MessagesSentRecords) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("to_user_id"),
	}
}

func (MessagesSentRecords) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "msg_messages_sent_records",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
	}
}
