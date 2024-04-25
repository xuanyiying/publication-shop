// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type PublicationStock struct {
	ent.Schema
}

func (PublicationStock) Fields() []ent.Field {
	return []ent.Field{field.Int("id").Comment("库存Id").StorageKey("stock_id"), field.String("isbn").Comment("国际标准书号"), field.Int("quantity").Optional().Comment("数量"), field.Int("publication_id").Optional().Comment("出版物ID")}
}
func (PublicationStock) Edges() []ent.Edge {
	return nil
}
func (PublicationStock) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "publication_stock"}}
}
