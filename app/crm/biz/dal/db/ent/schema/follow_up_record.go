package schema

import (
	"crm/biz/dal/db/ent/schema/mixins"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type FollowUpRecord struct {
	ent.Schema
}

func (FollowUpRecord) Fields() []ent.Field {
	return []ent.Field{

		field.String("content").Comment("跟进记录内容").Optional(),
		field.Int64("follow_up_id").Comment("跟进计划").Optional(),
		field.Int64("method").Comment("跟进方式").Optional(),
		field.Int64("user_id").Comment("跟进人").Optional(),
		field.Int64("division").Comment("部门").Optional(),
		field.Int64("opportunities_id").Comment("跟进商机").Optional(),
		field.String("record").Comment("跟进记录").Optional(),
	}
}

func (FollowUpRecord) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (FollowUpRecord) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (FollowUpRecord) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
	}
}

func (FollowUpRecord) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "crm_follow_up_record",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment(""),
	}
}
