// Code generated by ent, DO NOT EDIT.

package bookxcatalog

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the bookxcatalog type in the database.
	Label = "book_xcatalog"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "book_catalog_id"
	// FieldCatalogID holds the string denoting the catalog_id field in the database.
	FieldCatalogID = "catalog_id"
	// FieldIsbn holds the string denoting the isbn field in the database.
	FieldIsbn = "isbn"
	// FieldCatalogName holds the string denoting the catalog_name field in the database.
	FieldCatalogName = "catalog_name"
	// FieldBookID holds the string denoting the book_id field in the database.
	FieldBookID = "book_id"
	// Table holds the table name of the bookxcatalog in the database.
	Table = "book_xcatalogs"
)

// Columns holds all SQL columns for bookxcatalog fields.
var Columns = []string{
	FieldID,
	FieldCatalogID,
	FieldIsbn,
	FieldCatalogName,
	FieldBookID,
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

// OrderOption defines the ordering options for the BookXCatalog queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCatalogID orders the results by the catalog_id field.
func ByCatalogID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCatalogID, opts...).ToFunc()
}

// ByIsbn orders the results by the isbn field.
func ByIsbn(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsbn, opts...).ToFunc()
}

// ByCatalogName orders the results by the catalog_name field.
func ByCatalogName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCatalogName, opts...).ToFunc()
}

// ByBookID orders the results by the book_id field.
func ByBookID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBookID, opts...).ToFunc()
}