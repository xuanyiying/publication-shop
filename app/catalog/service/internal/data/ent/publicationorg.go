// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"github.com/go-kratos/publication-shop/app/catalog/service/internal/data/ent/publicationorg"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// PublicationOrg is the model entity for the PublicationOrg schema.
type PublicationOrg struct {
	config `json:"-"`
	// ID of the ent.
	// 出版单位ID
	ID int `json:"id,omitempty"`
	// 出版单位编码
	OrgCode string `json:"org_code,omitempty"`
	// 出版单位名称
	OrgName string `json:"org_name,omitempty"`
	// 出版单位住所
	OrgAddress string `json:"org_address,omitempty"`
	// 数据更新日期
	ModifiedAt time.Time `json:"modified_at,omitempty"`
	// 数据创建日期
	CreatedAt    time.Time `json:"created_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PublicationOrg) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case publicationorg.FieldID:
			values[i] = new(sql.NullInt64)
		case publicationorg.FieldOrgCode, publicationorg.FieldOrgName, publicationorg.FieldOrgAddress:
			values[i] = new(sql.NullString)
		case publicationorg.FieldModifiedAt, publicationorg.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PublicationOrg fields.
func (po *PublicationOrg) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case publicationorg.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			po.ID = int(value.Int64)
		case publicationorg.FieldOrgCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field org_code", values[i])
			} else if value.Valid {
				po.OrgCode = value.String
			}
		case publicationorg.FieldOrgName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field org_name", values[i])
			} else if value.Valid {
				po.OrgName = value.String
			}
		case publicationorg.FieldOrgAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field org_address", values[i])
			} else if value.Valid {
				po.OrgAddress = value.String
			}
		case publicationorg.FieldModifiedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field modified_at", values[i])
			} else if value.Valid {
				po.ModifiedAt = value.Time
			}
		case publicationorg.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				po.CreatedAt = value.Time
			}
		default:
			po.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PublicationOrg.
// This includes values selected through modifiers, order, etc.
func (po *PublicationOrg) Value(name string) (ent.Value, error) {
	return po.selectValues.Get(name)
}

// Update returns a builder for updating this PublicationOrg.
// Note that you need to call PublicationOrg.Unwrap() before calling this method if this PublicationOrg
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *PublicationOrg) Update() *PublicationOrgUpdateOne {
	return NewPublicationOrgClient(po.config).UpdateOne(po)
}

// Unwrap unwraps the PublicationOrg entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (po *PublicationOrg) Unwrap() *PublicationOrg {
	_tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("ent: PublicationOrg is not a transactional entity")
	}
	po.config.driver = _tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *PublicationOrg) String() string {
	var builder strings.Builder
	builder.WriteString("PublicationOrg(")
	builder.WriteString(fmt.Sprintf("id=%v, ", po.ID))
	builder.WriteString("org_code=")
	builder.WriteString(po.OrgCode)
	builder.WriteString(", ")
	builder.WriteString("org_name=")
	builder.WriteString(po.OrgName)
	builder.WriteString(", ")
	builder.WriteString("org_address=")
	builder.WriteString(po.OrgAddress)
	builder.WriteString(", ")
	builder.WriteString("modified_at=")
	builder.WriteString(po.ModifiedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(po.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// PublicationOrgs is a parsable slice of PublicationOrg.
type PublicationOrgs []*PublicationOrg
