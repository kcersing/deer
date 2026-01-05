package schema

import (
	"crm/biz/dal/db/ent/schema/mixins"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type FollowUpPlan struct {
	ent.Schema
}

func (FollowUpPlan) Fields() []ent.Field {
	return []ent.Field{
		field.String("content").Comment("计划内容").Optional(),
		field.Time("time").Comment("计划时间").Optional(),
		field.Int64("member_id").Comment("跟进客户").Optional(),
		field.Int64("user_id").Comment("执行人").Optional(),
		field.Int64("division").Comment("部门").Optional(),
	}
}

func (FollowUpPlan) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (FollowUpPlan) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (FollowUpPlan) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
	}
}

func (FollowUpPlan) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "crm_follow_up_plan",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment(""),
	}
}
