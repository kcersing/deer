package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"system/biz/dal/db/ent/schema/mixins"
)

type Sms struct {
	ent.Schema
}

func (Sms) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("notice_count").Default(0).Comment("通知短信数量").Optional(),
		field.Int64("used_notice").Default(0).Comment("已用通知").Optional(),
	}
}

func (Sms) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (Sms) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Sms) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Sms) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_sms"},
		entsql.WithComments(true),
	}
}
