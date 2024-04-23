// Code generated by ent, DO NOT EDIT.

package catalog

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the catalog type in the database.
	Label = "catalog"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "catalog_id"
	// FieldCatalogCode holds the string denoting the catalog_code field in the database.
	FieldCatalogCode = "catalog_code"
	// FieldCatalogName holds the string denoting the catalog_name field in the database.
	FieldCatalogName = "catalog_name"
	// FieldParentCatalogID holds the string denoting the parent_catalog_id field in the database.
	FieldParentCatalogID = "parent_catalog_id"
	// Table holds the table name of the catalog in the database.
	Table = "catalog"
)

// Columns holds all SQL columns for catalog fields.
var Columns = []string{
	FieldID,
	FieldCatalogCode,
	FieldCatalogName,
	FieldParentCatalogID,
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

// OrderOption defines the ordering options for the Catalog queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCatalogCode orders the results by the catalog_code field.
func ByCatalogCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCatalogCode, opts...).ToFunc()
}

// ByCatalogName orders the results by the catalog_name field.
func ByCatalogName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCatalogName, opts...).ToFunc()
}

// ByParentCatalogID orders the results by the parent_catalog_id field.
func ByParentCatalogID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParentCatalogID, opts...).ToFunc()
}
