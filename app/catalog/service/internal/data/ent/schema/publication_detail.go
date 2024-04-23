// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type PublicationDetail struct {
	ent.Schema
}

func (PublicationDetail) Fields() []ent.Field {
	return []ent.Field{field.Int("id").Comment("详情ID").StorageKey("detail_id"), field.String("isbn").Comment("国际标准书号"), field.String("detail_html").Optional().Comment("详情"), field.String("detail_img").Optional().Comment("详情"), field.Int("publication_id").Optional().Comment("出版物ID")}
}
func (PublicationDetail) Edges() []ent.Edge {
	return nil
}
func (PublicationDetail) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "publication_detail"}}
}
