package schema

import (
	"system/biz/dal/db/ent/schema/mixins"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type File struct {
	ent.Schema
}

// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{

		field.String("bucket_name").
			Comment("存储桶名称").
			Optional().
			Nillable(),

		field.String("file_directory").
			Comment("文件目录").
			Optional().
			Nillable(),

		field.String("file_guid").
			Comment("文件Guid").
			Optional().
			Nillable(),

		field.String("save_file_name").
			Comment("保存文件名").
			Optional().
			Nillable(),

		field.String("file_name").
			Comment("文件名").
			Optional().
			Nillable(),

		field.String("extension").
			Comment("文件扩展名").
			Optional().
			Nillable(),

		field.Uint64("size").
			Comment("文件字节长度").
			Optional().
			Nillable(),

		field.String("size_format").
			Comment("文件大小格式化").
			Optional().
			Nillable(),

		field.String("link_url").
			Comment("链接地址").
			Optional().
			Nillable(),

		field.String("md5").
			Comment("md5码，防止上传重复文件").
			Optional().
			Nillable(),
	}
}

func (File) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}
func (File) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "files",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("文件表"),
	}
}
