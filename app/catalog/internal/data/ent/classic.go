// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/xuanyiying/publication-shop/app/catalog/ent/classic"
)

// Classic is the model entity for the Classic schema.
type Classic struct {
	config `json:"-"`
	// ID of the ent.
	// 出版物性质ID
	ID int `json:"id,omitempty"`
	// 出版物性质名称
	ClassicName  string `json:"classic_name,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Classic) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case classic.FieldID:
			values[i] = new(sql.NullInt64)
		case classic.FieldClassicName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Classic fields.
func (c *Classic) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case classic.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case classic.FieldClassicName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field classic_name", values[i])
			} else if value.Valid {
				c.ClassicName = value.String
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Classic.
// This includes values selected through modifiers, order, etc.
func (c *Classic) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// Update returns a builder for updating this Classic.
// Note that you need to call Classic.Unwrap() before calling this method if this Classic
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Classic) Update() *ClassicUpdateOne {
	return NewClassicClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Classic entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Classic) Unwrap() *Classic {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Classic is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Classic) String() string {
	var builder strings.Builder
	builder.WriteString("Classic(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("classic_name=")
	builder.WriteString(c.ClassicName)
	builder.WriteByte(')')
	return builder.String()
}

// Classics is a parsable slice of Classic.
type Classics []*Classic
