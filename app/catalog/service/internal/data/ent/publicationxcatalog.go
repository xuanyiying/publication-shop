// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"github.com/go-kratos/publication-shop/app/catalog/service/internal/data/ent/publicationxcatalog"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// PublicationXCatalog is the model entity for the PublicationXCatalog schema.
type PublicationXCatalog struct {
	config `json:"-"`
	// ID of the ent.
	// 出版物分类关系ID
	ID int `json:"id,omitempty"`
	// 详情ID
	CatalogID int `json:"catalog_id,omitempty"`
	// 国际标准书号
	Isbn string `json:"isbn,omitempty"`
	// 详情
	CatalogName string `json:"catalog_name,omitempty"`
	// 出版物ID
	PublicationID int `json:"publication_id,omitempty"`
	selectValues  sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PublicationXCatalog) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case publicationxcatalog.FieldID, publicationxcatalog.FieldCatalogID, publicationxcatalog.FieldPublicationID:
			values[i] = new(sql.NullInt64)
		case publicationxcatalog.FieldIsbn, publicationxcatalog.FieldCatalogName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PublicationXCatalog fields.
func (px *PublicationXCatalog) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case publicationxcatalog.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			px.ID = int(value.Int64)
		case publicationxcatalog.FieldCatalogID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field catalog_id", values[i])
			} else if value.Valid {
				px.CatalogID = int(value.Int64)
			}
		case publicationxcatalog.FieldIsbn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field isbn", values[i])
			} else if value.Valid {
				px.Isbn = value.String
			}
		case publicationxcatalog.FieldCatalogName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field catalog_name", values[i])
			} else if value.Valid {
				px.CatalogName = value.String
			}
		case publicationxcatalog.FieldPublicationID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field publication_id", values[i])
			} else if value.Valid {
				px.PublicationID = int(value.Int64)
			}
		default:
			px.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PublicationXCatalog.
// This includes values selected through modifiers, order, etc.
func (px *PublicationXCatalog) Value(name string) (ent.Value, error) {
	return px.selectValues.Get(name)
}

// Update returns a builder for updating this PublicationXCatalog.
// Note that you need to call PublicationXCatalog.Unwrap() before calling this method if this PublicationXCatalog
// was returned from a transaction, and the transaction was committed or rolled back.
func (px *PublicationXCatalog) Update() *PublicationXCatalogUpdateOne {
	return NewPublicationXCatalogClient(px.config).UpdateOne(px)
}

// Unwrap unwraps the PublicationXCatalog entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (px *PublicationXCatalog) Unwrap() *PublicationXCatalog {
	_tx, ok := px.config.driver.(*txDriver)
	if !ok {
		panic("ent: PublicationXCatalog is not a transactional entity")
	}
	px.config.driver = _tx.drv
	return px
}

// String implements the fmt.Stringer.
func (px *PublicationXCatalog) String() string {
	var builder strings.Builder
	builder.WriteString("PublicationXCatalog(")
	builder.WriteString(fmt.Sprintf("id=%v, ", px.ID))
	builder.WriteString("catalog_id=")
	builder.WriteString(fmt.Sprintf("%v", px.CatalogID))
	builder.WriteString(", ")
	builder.WriteString("isbn=")
	builder.WriteString(px.Isbn)
	builder.WriteString(", ")
	builder.WriteString("catalog_name=")
	builder.WriteString(px.CatalogName)
	builder.WriteString(", ")
	builder.WriteString("publication_id=")
	builder.WriteString(fmt.Sprintf("%v", px.PublicationID))
	builder.WriteByte(')')
	return builder.String()
}

// PublicationXCatalogs is a parsable slice of PublicationXCatalog.
type PublicationXCatalogs []*PublicationXCatalog
