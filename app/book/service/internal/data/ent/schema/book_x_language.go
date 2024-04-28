// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type BookXLanguage struct {
	ent.Schema
}

func (BookXLanguage) Fields() []ent.Field {
	return []ent.Field{field.Int("id").Comment("出版物语言关系ID").StorageKey("book_language_id"), field.Int("language_id").Comment("详情ID"), field.String("isbn").Comment("国际标准书号"), field.String("language_name").Optional().Comment("详情"), field.Int("book_id").Optional().Comment("出版物ID")}
}
func (BookXLanguage) Edges() []ent.Edge {
	return nil
}
func (BookXLanguage) Annotations() []schema.Annotation {
	return nil
}
