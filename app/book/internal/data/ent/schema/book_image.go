// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type BookImage struct {
	ent.Schema
}

func (BookImage) Fields() []ent.Field {
	return []ent.Field{field.Int("id").StorageKey("img_id"), field.String("img_url").Optional().Comment("图片地址"), field.Int("isbn").Comment("图书编码"), field.String("img_encode").Optional(), field.Int("book_id").Optional().Comment("出版物ID"), field.Int32("mainFlag").Optional()}
}
func (BookImage) Edges() []ent.Edge {
	return nil
}
func (BookImage) Annotations() []schema.Annotation {
	return nil
}