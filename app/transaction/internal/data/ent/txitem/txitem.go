// Code generated by ent, DO NOT EDIT.

package txitem

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the txitem type in the database.
	Label = "tx_item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "tx_item_id"
	// FieldTxType holds the string denoting the tx_type field in the database.
	FieldTxType = "tx_type"
	// FieldTxID holds the string denoting the tx_id field in the database.
	FieldTxID = "tx_id"
	// FieldBookID holds the string denoting the book_id field in the database.
	FieldBookID = "book_id"
	// FieldQuantity holds the string denoting the quantity field in the database.
	FieldQuantity = "quantity"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldIsbn holds the string denoting the isbn field in the database.
	FieldIsbn = "isbn"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldAuthor holds the string denoting the author field in the database.
	FieldAuthor = "author"
	// FieldPublisherID holds the string denoting the publisher_id field in the database.
	FieldPublisherID = "publisher_id"
	// FieldImageURL holds the string denoting the image_url field in the database.
	FieldImageURL = "image_url"
	// Table holds the table name of the txitem in the database.
	Table = "tx_items"
)

// Columns holds all SQL columns for txitem fields.
var Columns = []string{
	FieldID,
	FieldTxType,
	FieldTxID,
	FieldBookID,
	FieldQuantity,
	FieldPrice,
	FieldIsbn,
	FieldTitle,
	FieldAuthor,
	FieldPublisherID,
	FieldImageURL,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "tx_items"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"order_tx_items",
	"transaction_tx_items",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// TxType defines the type for the "tx_type" enum field.
type TxType string

// TxType values.
const (
	TxTypeBuy  TxType = "buy"
	TxTypeSell TxType = "sell"
)

func (tt TxType) String() string {
	return string(tt)
}

// TxTypeValidator is a validator for the "tx_type" field enum values. It is called by the builders before save.
func TxTypeValidator(tt TxType) error {
	switch tt {
	case TxTypeBuy, TxTypeSell:
		return nil
	default:
		return fmt.Errorf("txitem: invalid enum value for tx_type field: %q", tt)
	}
}

// OrderOption defines the ordering options for the TxItem queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTxType orders the results by the tx_type field.
func ByTxType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTxType, opts...).ToFunc()
}

// ByTxID orders the results by the tx_id field.
func ByTxID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTxID, opts...).ToFunc()
}

// ByBookID orders the results by the book_id field.
func ByBookID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBookID, opts...).ToFunc()
}

// ByQuantity orders the results by the quantity field.
func ByQuantity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQuantity, opts...).ToFunc()
}

// ByPrice orders the results by the price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
}

// ByIsbn orders the results by the isbn field.
func ByIsbn(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsbn, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByAuthor orders the results by the author field.
func ByAuthor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthor, opts...).ToFunc()
}

// ByPublisherID orders the results by the publisher_id field.
func ByPublisherID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPublisherID, opts...).ToFunc()
}

// ByImageURL orders the results by the image_url field.
func ByImageURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldImageURL, opts...).ToFunc()
}