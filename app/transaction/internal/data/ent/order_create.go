// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/xuanyiying/publication-shop/app/transaction/internal/data/ent/order"
	"github.com/xuanyiying/publication-shop/app/transaction/internal/data/ent/transaction"
	"github.com/xuanyiying/publication-shop/app/transaction/internal/data/ent/txitem"
)

// OrderCreate is the builder for creating a Order entity.
type OrderCreate struct {
	config
	mutation *OrderMutation
	hooks    []Hook
}

// SetOrderNo sets the "order_no" field.
func (oc *OrderCreate) SetOrderNo(s string) *OrderCreate {
	oc.mutation.SetOrderNo(s)
	return oc
}

// SetNillableOrderNo sets the "order_no" field if the given value is not nil.
func (oc *OrderCreate) SetNillableOrderNo(s *string) *OrderCreate {
	if s != nil {
		oc.SetOrderNo(*s)
	}
	return oc
}

// SetTxID sets the "tx_id" field.
func (oc *OrderCreate) SetTxID(i int64) *OrderCreate {
	oc.mutation.SetTxID(i)
	return oc
}

// SetOrderStatus sets the "order_status" field.
func (oc *OrderCreate) SetOrderStatus(os order.OrderStatus) *OrderCreate {
	oc.mutation.SetOrderStatus(os)
	return oc
}

// SetDeliveredAddress sets the "delivered_address" field.
func (oc *OrderCreate) SetDeliveredAddress(s string) *OrderCreate {
	oc.mutation.SetDeliveredAddress(s)
	return oc
}

// SetNillableDeliveredAddress sets the "delivered_address" field if the given value is not nil.
func (oc *OrderCreate) SetNillableDeliveredAddress(s *string) *OrderCreate {
	if s != nil {
		oc.SetDeliveredAddress(*s)
	}
	return oc
}

// SetShippingCost sets the "shipping_cost" field.
func (oc *OrderCreate) SetShippingCost(f float64) *OrderCreate {
	oc.mutation.SetShippingCost(f)
	return oc
}

// SetNillableShippingCost sets the "shipping_cost" field if the given value is not nil.
func (oc *OrderCreate) SetNillableShippingCost(f *float64) *OrderCreate {
	if f != nil {
		oc.SetShippingCost(*f)
	}
	return oc
}

// SetTotalAmount sets the "total_amount" field.
func (oc *OrderCreate) SetTotalAmount(f float64) *OrderCreate {
	oc.mutation.SetTotalAmount(f)
	return oc
}

// SetNillableTotalAmount sets the "total_amount" field if the given value is not nil.
func (oc *OrderCreate) SetNillableTotalAmount(f *float64) *OrderCreate {
	if f != nil {
		oc.SetTotalAmount(*f)
	}
	return oc
}

// SetPlacedUserID sets the "placed_user_id" field.
func (oc *OrderCreate) SetPlacedUserID(i int64) *OrderCreate {
	oc.mutation.SetPlacedUserID(i)
	return oc
}

// SetPlacedAt sets the "placed_at" field.
func (oc *OrderCreate) SetPlacedAt(t time.Time) *OrderCreate {
	oc.mutation.SetPlacedAt(t)
	return oc
}

// SetNillablePlacedAt sets the "placed_at" field if the given value is not nil.
func (oc *OrderCreate) SetNillablePlacedAt(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetPlacedAt(*t)
	}
	return oc
}

// SetShippedAddress sets the "shipped_address" field.
func (oc *OrderCreate) SetShippedAddress(s string) *OrderCreate {
	oc.mutation.SetShippedAddress(s)
	return oc
}

// SetNillableShippedAddress sets the "shipped_address" field if the given value is not nil.
func (oc *OrderCreate) SetNillableShippedAddress(s *string) *OrderCreate {
	if s != nil {
		oc.SetShippedAddress(*s)
	}
	return oc
}

// SetShippedAt sets the "shipped_at" field.
func (oc *OrderCreate) SetShippedAt(t time.Time) *OrderCreate {
	oc.mutation.SetShippedAt(t)
	return oc
}

// SetNillableShippedAt sets the "shipped_at" field if the given value is not nil.
func (oc *OrderCreate) SetNillableShippedAt(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetShippedAt(*t)
	}
	return oc
}

// SetPaymentID sets the "payment_id" field.
func (oc *OrderCreate) SetPaymentID(i int64) *OrderCreate {
	oc.mutation.SetPaymentID(i)
	return oc
}

// SetNillablePaymentID sets the "payment_id" field if the given value is not nil.
func (oc *OrderCreate) SetNillablePaymentID(i *int64) *OrderCreate {
	if i != nil {
		oc.SetPaymentID(*i)
	}
	return oc
}

// SetID sets the "id" field.
func (oc *OrderCreate) SetID(i int64) *OrderCreate {
	oc.mutation.SetID(i)
	return oc
}

// AddTransactionIDs adds the "transaction" edge to the Transaction entity by IDs.
func (oc *OrderCreate) AddTransactionIDs(ids ...int64) *OrderCreate {
	oc.mutation.AddTransactionIDs(ids...)
	return oc
}

// AddTransaction adds the "transaction" edges to the Transaction entity.
func (oc *OrderCreate) AddTransaction(t ...*Transaction) *OrderCreate {
	ids := make([]int64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return oc.AddTransactionIDs(ids...)
}

// AddTxItemIDs adds the "txItems" edge to the TxItem entity by IDs.
func (oc *OrderCreate) AddTxItemIDs(ids ...int64) *OrderCreate {
	oc.mutation.AddTxItemIDs(ids...)
	return oc
}

// AddTxItems adds the "txItems" edges to the TxItem entity.
func (oc *OrderCreate) AddTxItems(t ...*TxItem) *OrderCreate {
	ids := make([]int64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return oc.AddTxItemIDs(ids...)
}

// Mutation returns the OrderMutation object of the builder.
func (oc *OrderCreate) Mutation() *OrderMutation {
	return oc.mutation
}

// Save creates the Order in the database.
func (oc *OrderCreate) Save(ctx context.Context) (*Order, error) {
	return withHooks(ctx, oc.sqlSave, oc.mutation, oc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OrderCreate) SaveX(ctx context.Context) *Order {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OrderCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OrderCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OrderCreate) check() error {
	if _, ok := oc.mutation.TxID(); !ok {
		return &ValidationError{Name: "tx_id", err: errors.New(`ent: missing required field "Order.tx_id"`)}
	}
	if _, ok := oc.mutation.OrderStatus(); !ok {
		return &ValidationError{Name: "order_status", err: errors.New(`ent: missing required field "Order.order_status"`)}
	}
	if v, ok := oc.mutation.OrderStatus(); ok {
		if err := order.OrderStatusValidator(v); err != nil {
			return &ValidationError{Name: "order_status", err: fmt.Errorf(`ent: validator failed for field "Order.order_status": %w`, err)}
		}
	}
	if _, ok := oc.mutation.PlacedUserID(); !ok {
		return &ValidationError{Name: "placed_user_id", err: errors.New(`ent: missing required field "Order.placed_user_id"`)}
	}
	return nil
}

func (oc *OrderCreate) sqlSave(ctx context.Context) (*Order, error) {
	if err := oc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	oc.mutation.id = &_node.ID
	oc.mutation.done = true
	return _node, nil
}

func (oc *OrderCreate) createSpec() (*Order, *sqlgraph.CreateSpec) {
	var (
		_node = &Order{config: oc.config}
		_spec = sqlgraph.NewCreateSpec(order.Table, sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt64))
	)
	if id, ok := oc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := oc.mutation.OrderNo(); ok {
		_spec.SetField(order.FieldOrderNo, field.TypeString, value)
		_node.OrderNo = value
	}
	if value, ok := oc.mutation.TxID(); ok {
		_spec.SetField(order.FieldTxID, field.TypeInt64, value)
		_node.TxID = value
	}
	if value, ok := oc.mutation.OrderStatus(); ok {
		_spec.SetField(order.FieldOrderStatus, field.TypeEnum, value)
		_node.OrderStatus = value
	}
	if value, ok := oc.mutation.DeliveredAddress(); ok {
		_spec.SetField(order.FieldDeliveredAddress, field.TypeString, value)
		_node.DeliveredAddress = value
	}
	if value, ok := oc.mutation.ShippingCost(); ok {
		_spec.SetField(order.FieldShippingCost, field.TypeFloat64, value)
		_node.ShippingCost = value
	}
	if value, ok := oc.mutation.TotalAmount(); ok {
		_spec.SetField(order.FieldTotalAmount, field.TypeFloat64, value)
		_node.TotalAmount = value
	}
	if value, ok := oc.mutation.PlacedUserID(); ok {
		_spec.SetField(order.FieldPlacedUserID, field.TypeInt64, value)
		_node.PlacedUserID = value
	}
	if value, ok := oc.mutation.PlacedAt(); ok {
		_spec.SetField(order.FieldPlacedAt, field.TypeTime, value)
		_node.PlacedAt = value
	}
	if value, ok := oc.mutation.ShippedAddress(); ok {
		_spec.SetField(order.FieldShippedAddress, field.TypeString, value)
		_node.ShippedAddress = value
	}
	if value, ok := oc.mutation.ShippedAt(); ok {
		_spec.SetField(order.FieldShippedAt, field.TypeTime, value)
		_node.ShippedAt = value
	}
	if value, ok := oc.mutation.PaymentID(); ok {
		_spec.SetField(order.FieldPaymentID, field.TypeInt64, value)
		_node.PaymentID = value
	}
	if nodes := oc.mutation.TransactionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.TransactionTable,
			Columns: []string{order.TransactionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.TxItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.TxItemsTable,
			Columns: []string{order.TxItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(txitem.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrderCreateBulk is the builder for creating many Order entities in bulk.
type OrderCreateBulk struct {
	config
	err      error
	builders []*OrderCreate
}

// Save creates the Order entities in the database.
func (ocb *OrderCreateBulk) Save(ctx context.Context) ([]*Order, error) {
	if ocb.err != nil {
		return nil, ocb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Order, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OrderCreateBulk) SaveX(ctx context.Context) []*Order {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OrderCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OrderCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}
