package schema

import (
	"system/biz/dal/db/ent/schema/mixins"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Messages struct {
	ent.Schema
}

func (Messages) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Comment("消息标题").Optional().Nillable(),
		field.Int64("from_user_id").Comment("发送者").Optional().Nillable(),
		field.String("content").Comment("消息内容").Optional().Nillable(),
		field.Int64("status").
			Comment("消息状态").
			Default(0).
			Optional().
			Nillable(),
		field.String("type").
			Comment("消息类型").
			Default("").
			Optional().
			Nillable(),
	}
}

func (Messages) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (Messages) Edges() []ent.Edge {
	return nil
}

func (Messages) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("from_user_id"),
	}
}

func (Messages) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "msg_messages",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
	}
}
