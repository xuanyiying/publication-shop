// Code generated by ent, DO NOT EDIT.

package publicationdetail

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the publicationdetail type in the database.
	Label = "publication_detail"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "detail_id"
	// FieldIsbn holds the string denoting the isbn field in the database.
	FieldIsbn = "isbn"
	// FieldDetailHTML holds the string denoting the detail_html field in the database.
	FieldDetailHTML = "detail_html"
	// FieldDetailImg holds the string denoting the detail_img field in the database.
	FieldDetailImg = "detail_img"
	// FieldPublicationID holds the string denoting the publication_id field in the database.
	FieldPublicationID = "publication_id"
	// Table holds the table name of the publicationdetail in the database.
	Table = "publication_detail"
)

// Columns holds all SQL columns for publicationdetail fields.
var Columns = []string{
	FieldID,
	FieldIsbn,
	FieldDetailHTML,
	FieldDetailImg,
	FieldPublicationID,
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

// OrderOption defines the ordering options for the PublicationDetail queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByIsbn orders the results by the isbn field.
func ByIsbn(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsbn, opts...).ToFunc()
}

// ByDetailHTML orders the results by the detail_html field.
func ByDetailHTML(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDetailHTML, opts...).ToFunc()
}

// ByPublicationID orders the results by the publication_id field.
func ByPublicationID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPublicationID, opts...).ToFunc()
}
