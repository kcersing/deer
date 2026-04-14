package schema

import (
	"contents/biz/dal/db/ent/schema/mixins"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Article struct {
	ent.Schema
}

func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Comment("标题").Optional(),
		field.Text("content").Comment("内容").Optional(),
		field.JSON("tag_id", []int64{}).Comment("标签").Optional(),
		field.JSON("pic", []string{}).Comment("主图").Optional(),
	}
}

func (Article) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (Article) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Article) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
	}
}

func (Article) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "articles",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("文章"),
	}
}
