// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/xuanyiying/publication-shop/app/book/service/internal/data/ent/bookxclassic"
)

// BookXClassic is the model entity for the BookXClassic schema.
type BookXClassic struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 详情ID
	ClassicID int `json:"classic_id,omitempty"`
	// 国际标准书号
	Isbn string `json:"isbn,omitempty"`
	// 详情
	ClassicName string `json:"classic_name,omitempty"`
	// 出版物ID
	BookID       int `json:"book_id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BookXClassic) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case bookxclassic.FieldID, bookxclassic.FieldClassicID, bookxclassic.FieldBookID:
			values[i] = new(sql.NullInt64)
		case bookxclassic.FieldIsbn, bookxclassic.FieldClassicName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BookXClassic fields.
func (bx *BookXClassic) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case bookxclassic.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			bx.ID = int(value.Int64)
		case bookxclassic.FieldClassicID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field classic_id", values[i])
			} else if value.Valid {
				bx.ClassicID = int(value.Int64)
			}
		case bookxclassic.FieldIsbn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field isbn", values[i])
			} else if value.Valid {
				bx.Isbn = value.String
			}
		case bookxclassic.FieldClassicName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field classic_name", values[i])
			} else if value.Valid {
				bx.ClassicName = value.String
			}
		case bookxclassic.FieldBookID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field book_id", values[i])
			} else if value.Valid {
				bx.BookID = int(value.Int64)
			}
		default:
			bx.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the BookXClassic.
// This includes values selected through modifiers, order, etc.
func (bx *BookXClassic) Value(name string) (ent.Value, error) {
	return bx.selectValues.Get(name)
}

// Update returns a builder for updating this BookXClassic.
// Note that you need to call BookXClassic.Unwrap() before calling this method if this BookXClassic
// was returned from a transaction, and the transaction was committed or rolled back.
func (bx *BookXClassic) Update() *BookXClassicUpdateOne {
	return NewBookXClassicClient(bx.config).UpdateOne(bx)
}

// Unwrap unwraps the BookXClassic entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bx *BookXClassic) Unwrap() *BookXClassic {
	_tx, ok := bx.config.driver.(*txDriver)
	if !ok {
		panic("ent: BookXClassic is not a transactional entity")
	}
	bx.config.driver = _tx.drv
	return bx
}

// String implements the fmt.Stringer.
func (bx *BookXClassic) String() string {
	var builder strings.Builder
	builder.WriteString("BookXClassic(")
	builder.WriteString(fmt.Sprintf("id=%v, ", bx.ID))
	builder.WriteString("classic_id=")
	builder.WriteString(fmt.Sprintf("%v", bx.ClassicID))
	builder.WriteString(", ")
	builder.WriteString("isbn=")
	builder.WriteString(bx.Isbn)
	builder.WriteString(", ")
	builder.WriteString("classic_name=")
	builder.WriteString(bx.ClassicName)
	builder.WriteString(", ")
	builder.WriteString("book_id=")
	builder.WriteString(fmt.Sprintf("%v", bx.BookID))
	builder.WriteByte(')')
	return builder.String()
}

// BookXClassics is a parsable slice of BookXClassic.
type BookXClassics []*BookXClassic
