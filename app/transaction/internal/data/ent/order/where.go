// Code generated by ent, DO NOT EDIT.

package order

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/xuanyiying/publication-shop/app/transaction/internal/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldID, id))
}

// OrderNo applies equality check predicate on the "order_no" field. It's identical to OrderNoEQ.
func OrderNo(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldOrderNo, v))
}

// TxID applies equality check predicate on the "tx_id" field. It's identical to TxIDEQ.
func TxID(v int64) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldTxID, v))
}

// DeliveredAddress applies equality check predicate on the "delivered_address" field. It's identical to DeliveredAddressEQ.
func DeliveredAddress(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldDeliveredAddress, v))
}

// ShippingCost applies equality check predicate on the "shipping_cost" field. It's identical to ShippingCostEQ.
func ShippingCost(v float64) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldShippingCost, v))
}

// TotalAmount applies equality check predicate on the "total_amount" field. It's identical to TotalAmountEQ.
func TotalAmount(v float64) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldTotalAmount, v))
}

// PlacedUserID applies equality check predicate on the "placed_user_id" field. It's identical to PlacedUserIDEQ.
func PlacedUserID(v int64) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldPlacedUserID, v))
}

// PlacedAt applies equality check predicate on the "placed_at" field. It's identical to PlacedAtEQ.
func PlacedAt(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldPlacedAt, v))
}

// ShippedAddress applies equality check predicate on the "shipped_address" field. It's identical to ShippedAddressEQ.
func ShippedAddress(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldShippedAddress, v))
}

// ShippedAt applies equality check predicate on the "shipped_at" field. It's identical to ShippedAtEQ.
func ShippedAt(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldShippedAt, v))
}

// PaymentID applies equality check predicate on the "payment_id" field. It's identical to PaymentIDEQ.
func PaymentID(v int64) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldPaymentID, v))
}

// OrderNoEQ applies the EQ predicate on the "order_no" field.
func OrderNoEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldOrderNo, v))
}

// OrderNoNEQ applies the NEQ predicate on the "order_no" field.
func OrderNoNEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldOrderNo, v))
}

// OrderNoIn applies the In predicate on the "order_no" field.
func OrderNoIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldOrderNo, vs...))
}

// OrderNoNotIn applies the NotIn predicate on the "order_no" field.
func OrderNoNotIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldOrderNo, vs...))
}

// OrderNoGT applies the GT predicate on the "order_no" field.
func OrderNoGT(v string) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldOrderNo, v))
}

// OrderNoGTE applies the GTE predicate on the "order_no" field.
func OrderNoGTE(v string) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldOrderNo, v))
}

// OrderNoLT applies the LT predicate on the "order_no" field.
func OrderNoLT(v string) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldOrderNo, v))
}

// OrderNoLTE applies the LTE predicate on the "order_no" field.
func OrderNoLTE(v string) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldOrderNo, v))
}

// OrderNoContains applies the Contains predicate on the "order_no" field.
func OrderNoContains(v string) predicate.Order {
	return predicate.Order(sql.FieldContains(FieldOrderNo, v))
}

// OrderNoHasPrefix applies the HasPrefix predicate on the "order_no" field.
func OrderNoHasPrefix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasPrefix(FieldOrderNo, v))
}

// OrderNoHasSuffix applies the HasSuffix predicate on the "order_no" field.
func OrderNoHasSuffix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasSuffix(FieldOrderNo, v))
}

// OrderNoIsNil applies the IsNil predicate on the "order_no" field.
func OrderNoIsNil() predicate.Order {
	return predicate.Order(sql.FieldIsNull(FieldOrderNo))
}

// OrderNoNotNil applies the NotNil predicate on the "order_no" field.
func OrderNoNotNil() predicate.Order {
	return predicate.Order(sql.FieldNotNull(FieldOrderNo))
}

// OrderNoEqualFold applies the EqualFold predicate on the "order_no" field.
func OrderNoEqualFold(v string) predicate.Order {
	return predicate.Order(sql.FieldEqualFold(FieldOrderNo, v))
}

// OrderNoContainsFold applies the ContainsFold predicate on the "order_no" field.
func OrderNoContainsFold(v string) predicate.Order {
	return predicate.Order(sql.FieldContainsFold(FieldOrderNo, v))
}

// TxIDEQ applies the EQ predicate on the "tx_id" field.
func TxIDEQ(v int64) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldTxID, v))
}

// TxIDNEQ applies the NEQ predicate on the "tx_id" field.
func TxIDNEQ(v int64) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldTxID, v))
}

// TxIDIn applies the In predicate on the "tx_id" field.
func TxIDIn(vs ...int64) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldTxID, vs...))
}

// TxIDNotIn applies the NotIn predicate on the "tx_id" field.
func TxIDNotIn(vs ...int64) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldTxID, vs...))
}

// TxIDGT applies the GT predicate on the "tx_id" field.
func TxIDGT(v int64) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldTxID, v))
}

// TxIDGTE applies the GTE predicate on the "tx_id" field.
func TxIDGTE(v int64) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldTxID, v))
}

// TxIDLT applies the LT predicate on the "tx_id" field.
func TxIDLT(v int64) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldTxID, v))
}

// TxIDLTE applies the LTE predicate on the "tx_id" field.
func TxIDLTE(v int64) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldTxID, v))
}

// OrderStatusEQ applies the EQ predicate on the "order_status" field.
func OrderStatusEQ(v OrderStatus) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldOrderStatus, v))
}

// OrderStatusNEQ applies the NEQ predicate on the "order_status" field.
func OrderStatusNEQ(v OrderStatus) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldOrderStatus, v))
}

// OrderStatusIn applies the In predicate on the "order_status" field.
func OrderStatusIn(vs ...OrderStatus) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldOrderStatus, vs...))
}

// OrderStatusNotIn applies the NotIn predicate on the "order_status" field.
func OrderStatusNotIn(vs ...OrderStatus) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldOrderStatus, vs...))
}

// DeliveredAddressEQ applies the EQ predicate on the "delivered_address" field.
func DeliveredAddressEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldDeliveredAddress, v))
}

// DeliveredAddressNEQ applies the NEQ predicate on the "delivered_address" field.
func DeliveredAddressNEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldDeliveredAddress, v))
}

// DeliveredAddressIn applies the In predicate on the "delivered_address" field.
func DeliveredAddressIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldDeliveredAddress, vs...))
}

// DeliveredAddressNotIn applies the NotIn predicate on the "delivered_address" field.
func DeliveredAddressNotIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldDeliveredAddress, vs...))
}

// DeliveredAddressGT applies the GT predicate on the "delivered_address" field.
func DeliveredAddressGT(v string) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldDeliveredAddress, v))
}

// DeliveredAddressGTE applies the GTE predicate on the "delivered_address" field.
func DeliveredAddressGTE(v string) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldDeliveredAddress, v))
}

// DeliveredAddressLT applies the LT predicate on the "delivered_address" field.
func DeliveredAddressLT(v string) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldDeliveredAddress, v))
}

// DeliveredAddressLTE applies the LTE predicate on the "delivered_address" field.
func DeliveredAddressLTE(v string) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldDeliveredAddress, v))
}

// DeliveredAddressContains applies the Contains predicate on the "delivered_address" field.
func DeliveredAddressContains(v string) predicate.Order {
	return predicate.Order(sql.FieldContains(FieldDeliveredAddress, v))
}

// DeliveredAddressHasPrefix applies the HasPrefix predicate on the "delivered_address" field.
func DeliveredAddressHasPrefix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasPrefix(FieldDeliveredAddress, v))
}

// DeliveredAddressHasSuffix applies the HasSuffix predicate on the "delivered_address" field.
func DeliveredAddressHasSuffix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasSuffix(FieldDeliveredAddress, v))
}

// DeliveredAddressIsNil applies the IsNil predicate on the "delivered_address" field.
func DeliveredAddressIsNil() predicate.Order {
	return predicate.Order(sql.FieldIsNull(FieldDeliveredAddress))
}

// DeliveredAddressNotNil applies the NotNil predicate on the "delivered_address" field.
func DeliveredAddressNotNil() predicate.Order {
	return predicate.Order(sql.FieldNotNull(FieldDeliveredAddress))
}

// DeliveredAddressEqualFold applies the EqualFold predicate on the "delivered_address" field.
func DeliveredAddressEqualFold(v string) predicate.Order {
	return predicate.Order(sql.FieldEqualFold(FieldDeliveredAddress, v))
}

// DeliveredAddressContainsFold applies the ContainsFold predicate on the "delivered_address" field.
func DeliveredAddressContainsFold(v string) predicate.Order {
	return predicate.Order(sql.FieldContainsFold(FieldDeliveredAddress, v))
}

// ShippingCostEQ applies the EQ predicate on the "shipping_cost" field.
func ShippingCostEQ(v float64) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldShippingCost, v))
}

// ShippingCostNEQ applies the NEQ predicate on the "shipping_cost" field.
func ShippingCostNEQ(v float64) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldShippingCost, v))
}

// ShippingCostIn applies the In predicate on the "shipping_cost" field.
func ShippingCostIn(vs ...float64) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldShippingCost, vs...))
}

// ShippingCostNotIn applies the NotIn predicate on the "shipping_cost" field.
func ShippingCostNotIn(vs ...float64) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldShippingCost, vs...))
}

// ShippingCostGT applies the GT predicate on the "shipping_cost" field.
func ShippingCostGT(v float64) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldShippingCost, v))
}

// ShippingCostGTE applies the GTE predicate on the "shipping_cost" field.
func ShippingCostGTE(v float64) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldShippingCost, v))
}

// ShippingCostLT applies the LT predicate on the "shipping_cost" field.
func ShippingCostLT(v float64) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldShippingCost, v))
}

// ShippingCostLTE applies the LTE predicate on the "shipping_cost" field.
func ShippingCostLTE(v float64) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldShippingCost, v))
}

// ShippingCostIsNil applies the IsNil predicate on the "shipping_cost" field.
func ShippingCostIsNil() predicate.Order {
	return predicate.Order(sql.FieldIsNull(FieldShippingCost))
}

// ShippingCostNotNil applies the NotNil predicate on the "shipping_cost" field.
func ShippingCostNotNil() predicate.Order {
	return predicate.Order(sql.FieldNotNull(FieldShippingCost))
}

// TotalAmountEQ applies the EQ predicate on the "total_amount" field.
func TotalAmountEQ(v float64) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldTotalAmount, v))
}

// TotalAmountNEQ applies the NEQ predicate on the "total_amount" field.
func TotalAmountNEQ(v float64) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldTotalAmount, v))
}

// TotalAmountIn applies the In predicate on the "total_amount" field.
func TotalAmountIn(vs ...float64) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldTotalAmount, vs...))
}

// TotalAmountNotIn applies the NotIn predicate on the "total_amount" field.
func TotalAmountNotIn(vs ...float64) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldTotalAmount, vs...))
}

// TotalAmountGT applies the GT predicate on the "total_amount" field.
func TotalAmountGT(v float64) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldTotalAmount, v))
}

// TotalAmountGTE applies the GTE predicate on the "total_amount" field.
func TotalAmountGTE(v float64) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldTotalAmount, v))
}

// TotalAmountLT applies the LT predicate on the "total_amount" field.
func TotalAmountLT(v float64) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldTotalAmount, v))
}

// TotalAmountLTE applies the LTE predicate on the "total_amount" field.
func TotalAmountLTE(v float64) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldTotalAmount, v))
}

// TotalAmountIsNil applies the IsNil predicate on the "total_amount" field.
func TotalAmountIsNil() predicate.Order {
	return predicate.Order(sql.FieldIsNull(FieldTotalAmount))
}

// TotalAmountNotNil applies the NotNil predicate on the "total_amount" field.
func TotalAmountNotNil() predicate.Order {
	return predicate.Order(sql.FieldNotNull(FieldTotalAmount))
}

// PlacedUserIDEQ applies the EQ predicate on the "placed_user_id" field.
func PlacedUserIDEQ(v int64) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldPlacedUserID, v))
}

// PlacedUserIDNEQ applies the NEQ predicate on the "placed_user_id" field.
func PlacedUserIDNEQ(v int64) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldPlacedUserID, v))
}

// PlacedUserIDIn applies the In predicate on the "placed_user_id" field.
func PlacedUserIDIn(vs ...int64) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldPlacedUserID, vs...))
}

// PlacedUserIDNotIn applies the NotIn predicate on the "placed_user_id" field.
func PlacedUserIDNotIn(vs ...int64) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldPlacedUserID, vs...))
}

// PlacedUserIDGT applies the GT predicate on the "placed_user_id" field.
func PlacedUserIDGT(v int64) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldPlacedUserID, v))
}

// PlacedUserIDGTE applies the GTE predicate on the "placed_user_id" field.
func PlacedUserIDGTE(v int64) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldPlacedUserID, v))
}

// PlacedUserIDLT applies the LT predicate on the "placed_user_id" field.
func PlacedUserIDLT(v int64) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldPlacedUserID, v))
}

// PlacedUserIDLTE applies the LTE predicate on the "placed_user_id" field.
func PlacedUserIDLTE(v int64) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldPlacedUserID, v))
}

// PlacedAtEQ applies the EQ predicate on the "placed_at" field.
func PlacedAtEQ(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldPlacedAt, v))
}

// PlacedAtNEQ applies the NEQ predicate on the "placed_at" field.
func PlacedAtNEQ(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldPlacedAt, v))
}

// PlacedAtIn applies the In predicate on the "placed_at" field.
func PlacedAtIn(vs ...time.Time) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldPlacedAt, vs...))
}

// PlacedAtNotIn applies the NotIn predicate on the "placed_at" field.
func PlacedAtNotIn(vs ...time.Time) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldPlacedAt, vs...))
}

// PlacedAtGT applies the GT predicate on the "placed_at" field.
func PlacedAtGT(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldPlacedAt, v))
}

// PlacedAtGTE applies the GTE predicate on the "placed_at" field.
func PlacedAtGTE(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldPlacedAt, v))
}

// PlacedAtLT applies the LT predicate on the "placed_at" field.
func PlacedAtLT(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldPlacedAt, v))
}

// PlacedAtLTE applies the LTE predicate on the "placed_at" field.
func PlacedAtLTE(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldPlacedAt, v))
}

// PlacedAtIsNil applies the IsNil predicate on the "placed_at" field.
func PlacedAtIsNil() predicate.Order {
	return predicate.Order(sql.FieldIsNull(FieldPlacedAt))
}

// PlacedAtNotNil applies the NotNil predicate on the "placed_at" field.
func PlacedAtNotNil() predicate.Order {
	return predicate.Order(sql.FieldNotNull(FieldPlacedAt))
}

// ShippedAddressEQ applies the EQ predicate on the "shipped_address" field.
func ShippedAddressEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldShippedAddress, v))
}

// ShippedAddressNEQ applies the NEQ predicate on the "shipped_address" field.
func ShippedAddressNEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldShippedAddress, v))
}

// ShippedAddressIn applies the In predicate on the "shipped_address" field.
func ShippedAddressIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldShippedAddress, vs...))
}

// ShippedAddressNotIn applies the NotIn predicate on the "shipped_address" field.
func ShippedAddressNotIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldShippedAddress, vs...))
}

// ShippedAddressGT applies the GT predicate on the "shipped_address" field.
func ShippedAddressGT(v string) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldShippedAddress, v))
}

// ShippedAddressGTE applies the GTE predicate on the "shipped_address" field.
func ShippedAddressGTE(v string) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldShippedAddress, v))
}

// ShippedAddressLT applies the LT predicate on the "shipped_address" field.
func ShippedAddressLT(v string) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldShippedAddress, v))
}

// ShippedAddressLTE applies the LTE predicate on the "shipped_address" field.
func ShippedAddressLTE(v string) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldShippedAddress, v))
}

// ShippedAddressContains applies the Contains predicate on the "shipped_address" field.
func ShippedAddressContains(v string) predicate.Order {
	return predicate.Order(sql.FieldContains(FieldShippedAddress, v))
}

// ShippedAddressHasPrefix applies the HasPrefix predicate on the "shipped_address" field.
func ShippedAddressHasPrefix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasPrefix(FieldShippedAddress, v))
}

// ShippedAddressHasSuffix applies the HasSuffix predicate on the "shipped_address" field.
func ShippedAddressHasSuffix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasSuffix(FieldShippedAddress, v))
}

// ShippedAddressIsNil applies the IsNil predicate on the "shipped_address" field.
func ShippedAddressIsNil() predicate.Order {
	return predicate.Order(sql.FieldIsNull(FieldShippedAddress))
}

// ShippedAddressNotNil applies the NotNil predicate on the "shipped_address" field.
func ShippedAddressNotNil() predicate.Order {
	return predicate.Order(sql.FieldNotNull(FieldShippedAddress))
}

// ShippedAddressEqualFold applies the EqualFold predicate on the "shipped_address" field.
func ShippedAddressEqualFold(v string) predicate.Order {
	return predicate.Order(sql.FieldEqualFold(FieldShippedAddress, v))
}

// ShippedAddressContainsFold applies the ContainsFold predicate on the "shipped_address" field.
func ShippedAddressContainsFold(v string) predicate.Order {
	return predicate.Order(sql.FieldContainsFold(FieldShippedAddress, v))
}

// ShippedAtEQ applies the EQ predicate on the "shipped_at" field.
func ShippedAtEQ(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldShippedAt, v))
}

// ShippedAtNEQ applies the NEQ predicate on the "shipped_at" field.
func ShippedAtNEQ(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldShippedAt, v))
}

// ShippedAtIn applies the In predicate on the "shipped_at" field.
func ShippedAtIn(vs ...time.Time) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldShippedAt, vs...))
}

// ShippedAtNotIn applies the NotIn predicate on the "shipped_at" field.
func ShippedAtNotIn(vs ...time.Time) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldShippedAt, vs...))
}

// ShippedAtGT applies the GT predicate on the "shipped_at" field.
func ShippedAtGT(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldShippedAt, v))
}

// ShippedAtGTE applies the GTE predicate on the "shipped_at" field.
func ShippedAtGTE(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldShippedAt, v))
}

// ShippedAtLT applies the LT predicate on the "shipped_at" field.
func ShippedAtLT(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldShippedAt, v))
}

// ShippedAtLTE applies the LTE predicate on the "shipped_at" field.
func ShippedAtLTE(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldShippedAt, v))
}

// ShippedAtIsNil applies the IsNil predicate on the "shipped_at" field.
func ShippedAtIsNil() predicate.Order {
	return predicate.Order(sql.FieldIsNull(FieldShippedAt))
}

// ShippedAtNotNil applies the NotNil predicate on the "shipped_at" field.
func ShippedAtNotNil() predicate.Order {
	return predicate.Order(sql.FieldNotNull(FieldShippedAt))
}

// PaymentIDEQ applies the EQ predicate on the "payment_id" field.
func PaymentIDEQ(v int64) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldPaymentID, v))
}

// PaymentIDNEQ applies the NEQ predicate on the "payment_id" field.
func PaymentIDNEQ(v int64) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldPaymentID, v))
}

// PaymentIDIn applies the In predicate on the "payment_id" field.
func PaymentIDIn(vs ...int64) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldPaymentID, vs...))
}

// PaymentIDNotIn applies the NotIn predicate on the "payment_id" field.
func PaymentIDNotIn(vs ...int64) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldPaymentID, vs...))
}

// PaymentIDGT applies the GT predicate on the "payment_id" field.
func PaymentIDGT(v int64) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldPaymentID, v))
}

// PaymentIDGTE applies the GTE predicate on the "payment_id" field.
func PaymentIDGTE(v int64) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldPaymentID, v))
}

// PaymentIDLT applies the LT predicate on the "payment_id" field.
func PaymentIDLT(v int64) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldPaymentID, v))
}

// PaymentIDLTE applies the LTE predicate on the "payment_id" field.
func PaymentIDLTE(v int64) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldPaymentID, v))
}

// PaymentIDIsNil applies the IsNil predicate on the "payment_id" field.
func PaymentIDIsNil() predicate.Order {
	return predicate.Order(sql.FieldIsNull(FieldPaymentID))
}

// PaymentIDNotNil applies the NotNil predicate on the "payment_id" field.
func PaymentIDNotNil() predicate.Order {
	return predicate.Order(sql.FieldNotNull(FieldPaymentID))
}

// HasTransaction applies the HasEdge predicate on the "transaction" edge.
func HasTransaction() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TransactionTable, TransactionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTransactionWith applies the HasEdge predicate on the "transaction" edge with a given conditions (other predicates).
func HasTransactionWith(preds ...predicate.Transaction) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := newTransactionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTxItems applies the HasEdge predicate on the "txItems" edge.
func HasTxItems() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TxItemsTable, TxItemsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTxItemsWith applies the HasEdge predicate on the "txItems" edge with a given conditions (other predicates).
func HasTxItemsWith(preds ...predicate.TxItem) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := newTxItemsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Order) predicate.Order {
	return predicate.Order(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Order) predicate.Order {
	return predicate.Order(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Order) predicate.Order {
	return predicate.Order(sql.NotPredicates(p))
}