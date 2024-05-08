// Code generated by ent, DO NOT EDIT.

package order

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the order type in the database.
	Label = "order"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "order_id"
	// FieldOrderNo holds the string denoting the order_no field in the database.
	FieldOrderNo = "order_no"
	// FieldTxID holds the string denoting the tx_id field in the database.
	FieldTxID = "tx_id"
	// FieldOrderStatus holds the string denoting the order_status field in the database.
	FieldOrderStatus = "order_status"
	// FieldDeliveredAddress holds the string denoting the delivered_address field in the database.
	FieldDeliveredAddress = "delivered_address"
	// FieldShippingCost holds the string denoting the shipping_cost field in the database.
	FieldShippingCost = "shipping_cost"
	// FieldTotalAmount holds the string denoting the total_amount field in the database.
	FieldTotalAmount = "total_amount"
	// FieldPlacedUserID holds the string denoting the placed_user_id field in the database.
	FieldPlacedUserID = "placed_user_id"
	// FieldPlacedAt holds the string denoting the placed_at field in the database.
	FieldPlacedAt = "placed_at"
	// FieldShippedAddress holds the string denoting the shipped_address field in the database.
	FieldShippedAddress = "shipped_address"
	// FieldShippedAt holds the string denoting the shipped_at field in the database.
	FieldShippedAt = "shipped_at"
	// FieldPaymentID holds the string denoting the payment_id field in the database.
	FieldPaymentID = "payment_id"
	// EdgeTransaction holds the string denoting the transaction edge name in mutations.
	EdgeTransaction = "transaction"
	// EdgeTxItems holds the string denoting the txitems edge name in mutations.
	EdgeTxItems = "txItems"
	// TransactionFieldID holds the string denoting the ID field of the Transaction.
	TransactionFieldID = "tx_id"
	// TxItemFieldID holds the string denoting the ID field of the TxItem.
	TxItemFieldID = "tx_item_id"
	// Table holds the table name of the order in the database.
	Table = "orders"
	// TransactionTable is the table that holds the transaction relation/edge.
	TransactionTable = "transactions"
	// TransactionInverseTable is the table name for the Transaction entity.
	// It exists in this package in order to avoid circular dependency with the "transaction" package.
	TransactionInverseTable = "transactions"
	// TransactionColumn is the table column denoting the transaction relation/edge.
	TransactionColumn = "order_transaction"
	// TxItemsTable is the table that holds the txItems relation/edge.
	TxItemsTable = "tx_items"
	// TxItemsInverseTable is the table name for the TxItem entity.
	// It exists in this package in order to avoid circular dependency with the "txitem" package.
	TxItemsInverseTable = "tx_items"
	// TxItemsColumn is the table column denoting the txItems relation/edge.
	TxItemsColumn = "order_tx_items"
)

// Columns holds all SQL columns for order fields.
var Columns = []string{
	FieldID,
	FieldOrderNo,
	FieldTxID,
	FieldOrderStatus,
	FieldDeliveredAddress,
	FieldShippingCost,
	FieldTotalAmount,
	FieldPlacedUserID,
	FieldPlacedAt,
	FieldShippedAddress,
	FieldShippedAt,
	FieldPaymentID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "orders"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"transaction_tx_order",
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

// OrderStatus defines the type for the "order_status" enum field.
type OrderStatus string

// OrderStatus values.
const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusShipped    OrderStatus = "shipped"
	OrderStatusDelivered  OrderStatus = "delivered"
	OrderStatusCancelled  OrderStatus = "cancelled"
)

func (os OrderStatus) String() string {
	return string(os)
}

// OrderStatusValidator is a validator for the "order_status" field enum values. It is called by the builders before save.
func OrderStatusValidator(os OrderStatus) error {
	switch os {
	case OrderStatusPending, OrderStatusProcessing, OrderStatusShipped, OrderStatusDelivered, OrderStatusCancelled:
		return nil
	default:
		return fmt.Errorf("order: invalid enum value for order_status field: %q", os)
	}
}

// OrderOption defines the ordering options for the Order queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByOrderNo orders the results by the order_no field.
func ByOrderNo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrderNo, opts...).ToFunc()
}

// ByTxID orders the results by the tx_id field.
func ByTxID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTxID, opts...).ToFunc()
}

// ByOrderStatus orders the results by the order_status field.
func ByOrderStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrderStatus, opts...).ToFunc()
}

// ByDeliveredAddress orders the results by the delivered_address field.
func ByDeliveredAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeliveredAddress, opts...).ToFunc()
}

// ByShippingCost orders the results by the shipping_cost field.
func ByShippingCost(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldShippingCost, opts...).ToFunc()
}

// ByTotalAmount orders the results by the total_amount field.
func ByTotalAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalAmount, opts...).ToFunc()
}

// ByPlacedUserID orders the results by the placed_user_id field.
func ByPlacedUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlacedUserID, opts...).ToFunc()
}

// ByPlacedAt orders the results by the placed_at field.
func ByPlacedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlacedAt, opts...).ToFunc()
}

// ByShippedAddress orders the results by the shipped_address field.
func ByShippedAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldShippedAddress, opts...).ToFunc()
}

// ByShippedAt orders the results by the shipped_at field.
func ByShippedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldShippedAt, opts...).ToFunc()
}

// ByPaymentID orders the results by the payment_id field.
func ByPaymentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaymentID, opts...).ToFunc()
}

// ByTransactionCount orders the results by transaction count.
func ByTransactionCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTransactionStep(), opts...)
	}
}

// ByTransaction orders the results by transaction terms.
func ByTransaction(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTransactionStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTxItemsCount orders the results by txItems count.
func ByTxItemsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTxItemsStep(), opts...)
	}
}

// ByTxItems orders the results by txItems terms.
func ByTxItems(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTxItemsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newTransactionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TransactionInverseTable, TransactionFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TransactionTable, TransactionColumn),
	)
}
func newTxItemsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TxItemsInverseTable, TxItemFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TxItemsTable, TxItemsColumn),
	)
}
