// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/xuanyiying/publication-shop/app/transaction/internal/data/ent/order"
)

// Order is the model entity for the Order schema.
type Order struct {
	config `json:"-"`
	// ID of the ent.
	// 订单ID，主键
	ID int64 `json:"id,omitempty"`
	// 订单编号
	OrderNo string `json:"order_no,omitempty"`
	// 交易ID
	TxID int64 `json:"tx_id,omitempty"`
	// 订单状态
	OrderStatus order.OrderStatus `json:"order_status,omitempty"`
	// 配送地址
	DeliveredAddress string `json:"delivered_address,omitempty"`
	// 运费
	ShippingCost float64 `json:"shipping_cost,omitempty"`
	// 总金额
	TotalAmount float64 `json:"total_amount,omitempty"`
	// 下单用户ID
	PlacedUserID int64 `json:"placed_user_id,omitempty"`
	// 下单时间
	PlacedAt time.Time `json:"placed_at,omitempty"`
	// 发货地址
	ShippedAddress string `json:"shipped_address,omitempty"`
	// 发货时间
	ShippedAt time.Time `json:"shipped_at,omitempty"`
	// 支付ID
	PaymentID int64 `json:"payment_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderQuery when eager-loading is set.
	Edges                OrderEdges `json:"edges"`
	transaction_tx_order *int64
	selectValues         sql.SelectValues
}

// OrderEdges holds the relations/edges for other nodes in the graph.
type OrderEdges struct {
	// Transaction holds the value of the transaction edge.
	Transaction []*Transaction `json:"transaction,omitempty"`
	// TxItems holds the value of the txItems edge.
	TxItems []*TxItem `json:"txItems,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TransactionOrErr returns the Transaction value or an error if the edge
// was not loaded in eager-loading.
func (e OrderEdges) TransactionOrErr() ([]*Transaction, error) {
	if e.loadedTypes[0] {
		return e.Transaction, nil
	}
	return nil, &NotLoadedError{edge: "transaction"}
}

// TxItemsOrErr returns the TxItems value or an error if the edge
// was not loaded in eager-loading.
func (e OrderEdges) TxItemsOrErr() ([]*TxItem, error) {
	if e.loadedTypes[1] {
		return e.TxItems, nil
	}
	return nil, &NotLoadedError{edge: "txItems"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Order) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case order.FieldShippingCost, order.FieldTotalAmount:
			values[i] = new(sql.NullFloat64)
		case order.FieldID, order.FieldTxID, order.FieldPlacedUserID, order.FieldPaymentID:
			values[i] = new(sql.NullInt64)
		case order.FieldOrderNo, order.FieldOrderStatus, order.FieldDeliveredAddress, order.FieldShippedAddress:
			values[i] = new(sql.NullString)
		case order.FieldPlacedAt, order.FieldShippedAt:
			values[i] = new(sql.NullTime)
		case order.ForeignKeys[0]: // transaction_tx_order
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Order fields.
func (o *Order) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case order.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			o.ID = int64(value.Int64)
		case order.FieldOrderNo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field order_no", values[i])
			} else if value.Valid {
				o.OrderNo = value.String
			}
		case order.FieldTxID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tx_id", values[i])
			} else if value.Valid {
				o.TxID = value.Int64
			}
		case order.FieldOrderStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field order_status", values[i])
			} else if value.Valid {
				o.OrderStatus = order.OrderStatus(value.String)
			}
		case order.FieldDeliveredAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field delivered_address", values[i])
			} else if value.Valid {
				o.DeliveredAddress = value.String
			}
		case order.FieldShippingCost:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field shipping_cost", values[i])
			} else if value.Valid {
				o.ShippingCost = value.Float64
			}
		case order.FieldTotalAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field total_amount", values[i])
			} else if value.Valid {
				o.TotalAmount = value.Float64
			}
		case order.FieldPlacedUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field placed_user_id", values[i])
			} else if value.Valid {
				o.PlacedUserID = value.Int64
			}
		case order.FieldPlacedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field placed_at", values[i])
			} else if value.Valid {
				o.PlacedAt = value.Time
			}
		case order.FieldShippedAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field shipped_address", values[i])
			} else if value.Valid {
				o.ShippedAddress = value.String
			}
		case order.FieldShippedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field shipped_at", values[i])
			} else if value.Valid {
				o.ShippedAt = value.Time
			}
		case order.FieldPaymentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field payment_id", values[i])
			} else if value.Valid {
				o.PaymentID = value.Int64
			}
		case order.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field transaction_tx_order", value)
			} else if value.Valid {
				o.transaction_tx_order = new(int64)
				*o.transaction_tx_order = int64(value.Int64)
			}
		default:
			o.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Order.
// This includes values selected through modifiers, order, etc.
func (o *Order) Value(name string) (ent.Value, error) {
	return o.selectValues.Get(name)
}

// QueryTransaction queries the "transaction" edge of the Order entity.
func (o *Order) QueryTransaction() *TransactionQuery {
	return NewOrderClient(o.config).QueryTransaction(o)
}

// QueryTxItems queries the "txItems" edge of the Order entity.
func (o *Order) QueryTxItems() *TxItemQuery {
	return NewOrderClient(o.config).QueryTxItems(o)
}

// Update returns a builder for updating this Order.
// Note that you need to call Order.Unwrap() before calling this method if this Order
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Order) Update() *OrderUpdateOne {
	return NewOrderClient(o.config).UpdateOne(o)
}

// Unwrap unwraps the Order entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Order) Unwrap() *Order {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Order is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Order) String() string {
	var builder strings.Builder
	builder.WriteString("Order(")
	builder.WriteString(fmt.Sprintf("id=%v, ", o.ID))
	builder.WriteString("order_no=")
	builder.WriteString(o.OrderNo)
	builder.WriteString(", ")
	builder.WriteString("tx_id=")
	builder.WriteString(fmt.Sprintf("%v", o.TxID))
	builder.WriteString(", ")
	builder.WriteString("order_status=")
	builder.WriteString(fmt.Sprintf("%v", o.OrderStatus))
	builder.WriteString(", ")
	builder.WriteString("delivered_address=")
	builder.WriteString(o.DeliveredAddress)
	builder.WriteString(", ")
	builder.WriteString("shipping_cost=")
	builder.WriteString(fmt.Sprintf("%v", o.ShippingCost))
	builder.WriteString(", ")
	builder.WriteString("total_amount=")
	builder.WriteString(fmt.Sprintf("%v", o.TotalAmount))
	builder.WriteString(", ")
	builder.WriteString("placed_user_id=")
	builder.WriteString(fmt.Sprintf("%v", o.PlacedUserID))
	builder.WriteString(", ")
	builder.WriteString("placed_at=")
	builder.WriteString(o.PlacedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("shipped_address=")
	builder.WriteString(o.ShippedAddress)
	builder.WriteString(", ")
	builder.WriteString("shipped_at=")
	builder.WriteString(o.ShippedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("payment_id=")
	builder.WriteString(fmt.Sprintf("%v", o.PaymentID))
	builder.WriteByte(')')
	return builder.String()
}

// Orders is a parsable slice of Order.
type Orders []*Order
