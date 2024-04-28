// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type BookAuthor struct {
	ent.Schema
}

func (BookAuthor) Fields() []ent.Field {
	return []ent.Field{field.Int("id").StorageKey("book_author_id"), field.Int("author_id").Comment("作者ID"), field.String("isbn").Comment("国际标准书号"), field.Int("book_id").Optional().Comment("出版物ID"), field.String("author").Optional().Comment("著者姓名"), field.String("author_about").Optional().Comment("作者简介")}
}
func (BookAuthor) Edges() []ent.Edge {
	return nil
}
func (BookAuthor) Annotations() []schema.Annotation {
	return nil
}
