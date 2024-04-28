// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/xuanyiying/publication-shop/app/book/service/internal/data/ent/bookimage"
)

// BookImage is the model entity for the BookImage schema.
type BookImage struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 图片地址
	ImgURL string `json:"img_url,omitempty"`
	// 图书编码
	Isbn int `json:"isbn,omitempty"`
	// ImgEncode holds the value of the "img_encode" field.
	ImgEncode string `json:"img_encode,omitempty"`
	// 出版物ID
	BookID int `json:"book_id,omitempty"`
	// MainFlag holds the value of the "mainFlag" field.
	MainFlag     int32 `json:"mainFlag,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BookImage) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case bookimage.FieldID, bookimage.FieldIsbn, bookimage.FieldBookID, bookimage.FieldMainFlag:
			values[i] = new(sql.NullInt64)
		case bookimage.FieldImgURL, bookimage.FieldImgEncode:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BookImage fields.
func (bi *BookImage) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case bookimage.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			bi.ID = int(value.Int64)
		case bookimage.FieldImgURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field img_url", values[i])
			} else if value.Valid {
				bi.ImgURL = value.String
			}
		case bookimage.FieldIsbn:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field isbn", values[i])
			} else if value.Valid {
				bi.Isbn = int(value.Int64)
			}
		case bookimage.FieldImgEncode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field img_encode", values[i])
			} else if value.Valid {
				bi.ImgEncode = value.String
			}
		case bookimage.FieldBookID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field book_id", values[i])
			} else if value.Valid {
				bi.BookID = int(value.Int64)
			}
		case bookimage.FieldMainFlag:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field mainFlag", values[i])
			} else if value.Valid {
				bi.MainFlag = int32(value.Int64)
			}
		default:
			bi.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the BookImage.
// This includes values selected through modifiers, order, etc.
func (bi *BookImage) Value(name string) (ent.Value, error) {
	return bi.selectValues.Get(name)
}

// Update returns a builder for updating this BookImage.
// Note that you need to call BookImage.Unwrap() before calling this method if this BookImage
// was returned from a transaction, and the transaction was committed or rolled back.
func (bi *BookImage) Update() *BookImageUpdateOne {
	return NewBookImageClient(bi.config).UpdateOne(bi)
}

// Unwrap unwraps the BookImage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bi *BookImage) Unwrap() *BookImage {
	_tx, ok := bi.config.driver.(*txDriver)
	if !ok {
		panic("ent: BookImage is not a transactional entity")
	}
	bi.config.driver = _tx.drv
	return bi
}

// String implements the fmt.Stringer.
func (bi *BookImage) String() string {
	var builder strings.Builder
	builder.WriteString("BookImage(")
	builder.WriteString(fmt.Sprintf("id=%v, ", bi.ID))
	builder.WriteString("img_url=")
	builder.WriteString(bi.ImgURL)
	builder.WriteString(", ")
	builder.WriteString("isbn=")
	builder.WriteString(fmt.Sprintf("%v", bi.Isbn))
	builder.WriteString(", ")
	builder.WriteString("img_encode=")
	builder.WriteString(bi.ImgEncode)
	builder.WriteString(", ")
	builder.WriteString("book_id=")
	builder.WriteString(fmt.Sprintf("%v", bi.BookID))
	builder.WriteString(", ")
	builder.WriteString("mainFlag=")
	builder.WriteString(fmt.Sprintf("%v", bi.MainFlag))
	builder.WriteByte(')')
	return builder.String()
}

// BookImages is a parsable slice of BookImage.
type BookImages []*BookImage