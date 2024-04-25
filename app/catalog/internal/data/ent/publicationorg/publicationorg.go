// Code generated by ent, DO NOT EDIT.

package publicationorg

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the publicationorg type in the database.
	Label = "publication_org"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "org_id"
	// FieldOrgCode holds the string denoting the org_code field in the database.
	FieldOrgCode = "org_code"
	// FieldOrgName holds the string denoting the org_name field in the database.
	FieldOrgName = "org_name"
	// FieldOrgAddress holds the string denoting the org_address field in the database.
	FieldOrgAddress = "org_address"
	// FieldModifiedAt holds the string denoting the modified_at field in the database.
	FieldModifiedAt = "modified_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the publicationorg in the database.
	Table = "publication_org"
)

// Columns holds all SQL columns for publicationorg fields.
var Columns = []string{
	FieldID,
	FieldOrgCode,
	FieldOrgName,
	FieldOrgAddress,
	FieldModifiedAt,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the PublicationOrg queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByOrgCode orders the results by the org_code field.
func ByOrgCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrgCode, opts...).ToFunc()
}

// ByOrgName orders the results by the org_name field.
func ByOrgName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrgName, opts...).ToFunc()
}

// ByOrgAddress orders the results by the org_address field.
func ByOrgAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrgAddress, opts...).ToFunc()
}

// ByModifiedAt orders the results by the modified_at field.
func ByModifiedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModifiedAt, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}