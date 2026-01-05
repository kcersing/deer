package schema

import (
	"system/biz/dal/db/ent/schema/mixins"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type SmsSentRecords struct {
	ent.Schema
}

func (SmsSentRecords) Fields() []ent.Field {
	return []ent.Field{
		field.String("mobile").Comment("手机号").Optional(),
		field.String("biz_id").Comment("BizId").Optional(),
		field.String("code").Comment("验证码").Optional(),
		field.String("content").Comment("内容").Optional(),
		field.Int64("notify_type").Comment("通知类型[1会员;2员工]").Default(1).Optional(),
		field.String("template").Comment("短信模板").Optional(),
	}
}

func (SmsSentRecords) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (SmsSentRecords) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (SmsSentRecords) Indexes() []ent.Index {
	return []ent.Index{}
}

func (SmsSentRecords) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sms_sent_records",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment(""),
	}
}
