package schema

import (
	"crm/biz/dal/db/ent/schema/mixins"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Opportunities struct {
	ent.Schema
}

func (Opportunities) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Comment("标题").Optional(),

		field.String("content").Comment("内容").Optional(),
		field.Int64("user_id").Comment("执行人").Optional(),
		field.Int64("member_id").Comment("会员ID").Optional(),

		field.Int64("period").Comment("阶段").Optional(),
		field.Int64("prediction_amount").Comment("预测成交金额").Optional(),

		field.Time("period_at").Comment("阶段变更时间").Optional(),
		field.Time("end_sales_at").Comment("结束售卖时间").Optional(),
	}
}

func (Opportunities) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (Opportunities) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Opportunities) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
	}
}

func (Opportunities) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "crm_opportunities",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
	}
}
