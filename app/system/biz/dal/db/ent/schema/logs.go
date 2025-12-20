package schema

import (
	"system/biz/dal/db/ent/schema/mixins"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Logs struct {
	ent.Schema
}

func (Logs) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").Comment("type of log | 日志类型").Optional(),
		field.String("method").Comment("method of log | 日志请求方法").Optional(),
		field.String("api").Comment("api of log | 日志请求api").Optional(),
		field.Int64("success").Comment("success of log | 日志请求是否成功").Optional(),
		field.Text("req_content").Optional().Comment("content of request log | 日志请求内容").Optional(),
		field.Text("resp_content").Optional().Comment("content of response log | 日志返回内容").Optional(),
		field.String("ip").Optional().Comment("ip of log | 日志IP").Optional(),
		field.String("user_agent").Optional().Comment("user_agent of log | 日志用户客户端").Optional(),
		field.Int64("identity").Optional().Comment("操作者").Optional(),
		field.Int64("time").Optional().Comment("time of log(millisecond) | 日志时间(毫秒)").Optional(),
	}
}

func (Logs) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (Logs) Edges() []ent.Edge {
	return nil
}

func (Logs) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("api"),
	}
}

func (Logs) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_logs_" + time.Now().Format("20060102")},
		entsql.WithComments(true),
	}
}
