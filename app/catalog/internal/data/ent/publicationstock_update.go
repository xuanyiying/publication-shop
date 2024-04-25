// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/publication-shop/app/catalog/ent/predicate"
	"github.com/publication-shop/app/catalog/ent/publicationstock"
)

// PublicationStockUpdate is the builder for updating PublicationStock entities.
type PublicationStockUpdate struct {
	config
	hooks    []Hook
	mutation *PublicationStockMutation
}

// Where appends a list predicates to the PublicationStockUpdate builder.
func (psu *PublicationStockUpdate) Where(ps ...predicate.PublicationStock) *PublicationStockUpdate {
	psu.mutation.Where(ps...)
	return psu
}

// SetIsbn sets the "isbn" field.
func (psu *PublicationStockUpdate) SetIsbn(s string) *PublicationStockUpdate {
	psu.mutation.SetIsbn(s)
	return psu
}

// SetNillableIsbn sets the "isbn" field if the given value is not nil.
func (psu *PublicationStockUpdate) SetNillableIsbn(s *string) *PublicationStockUpdate {
	if s != nil {
		psu.SetIsbn(*s)
	}
	return psu
}

// SetQuantity sets the "quantity" field.
func (psu *PublicationStockUpdate) SetQuantity(i int) *PublicationStockUpdate {
	psu.mutation.ResetQuantity()
	psu.mutation.SetQuantity(i)
	return psu
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (psu *PublicationStockUpdate) SetNillableQuantity(i *int) *PublicationStockUpdate {
	if i != nil {
		psu.SetQuantity(*i)
	}
	return psu
}

// AddQuantity adds i to the "quantity" field.
func (psu *PublicationStockUpdate) AddQuantity(i int) *PublicationStockUpdate {
	psu.mutation.AddQuantity(i)
	return psu
}

// ClearQuantity clears the value of the "quantity" field.
func (psu *PublicationStockUpdate) ClearQuantity() *PublicationStockUpdate {
	psu.mutation.ClearQuantity()
	return psu
}

// SetPublicationID sets the "publication_id" field.
func (psu *PublicationStockUpdate) SetPublicationID(i int) *PublicationStockUpdate {
	psu.mutation.ResetPublicationID()
	psu.mutation.SetPublicationID(i)
	return psu
}

// SetNillablePublicationID sets the "publication_id" field if the given value is not nil.
func (psu *PublicationStockUpdate) SetNillablePublicationID(i *int) *PublicationStockUpdate {
	if i != nil {
		psu.SetPublicationID(*i)
	}
	return psu
}

// AddPublicationID adds i to the "publication_id" field.
func (psu *PublicationStockUpdate) AddPublicationID(i int) *PublicationStockUpdate {
	psu.mutation.AddPublicationID(i)
	return psu
}

// ClearPublicationID clears the value of the "publication_id" field.
func (psu *PublicationStockUpdate) ClearPublicationID() *PublicationStockUpdate {
	psu.mutation.ClearPublicationID()
	return psu
}

// Mutation returns the PublicationStockMutation object of the builder.
func (psu *PublicationStockUpdate) Mutation() *PublicationStockMutation {
	return psu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (psu *PublicationStockUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, psu.sqlSave, psu.mutation, psu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (psu *PublicationStockUpdate) SaveX(ctx context.Context) int {
	affected, err := psu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (psu *PublicationStockUpdate) Exec(ctx context.Context) error {
	_, err := psu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psu *PublicationStockUpdate) ExecX(ctx context.Context) {
	if err := psu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (psu *PublicationStockUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(publicationstock.Table, publicationstock.Columns, sqlgraph.NewFieldSpec(publicationstock.FieldID, field.TypeInt))
	if ps := psu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := psu.mutation.Isbn(); ok {
		_spec.SetField(publicationstock.FieldIsbn, field.TypeString, value)
	}
	if value, ok := psu.mutation.Quantity(); ok {
		_spec.SetField(publicationstock.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := psu.mutation.AddedQuantity(); ok {
		_spec.AddField(publicationstock.FieldQuantity, field.TypeInt, value)
	}
	if psu.mutation.QuantityCleared() {
		_spec.ClearField(publicationstock.FieldQuantity, field.TypeInt)
	}
	if value, ok := psu.mutation.PublicationID(); ok {
		_spec.SetField(publicationstock.FieldPublicationID, field.TypeInt, value)
	}
	if value, ok := psu.mutation.AddedPublicationID(); ok {
		_spec.AddField(publicationstock.FieldPublicationID, field.TypeInt, value)
	}
	if psu.mutation.PublicationIDCleared() {
		_spec.ClearField(publicationstock.FieldPublicationID, field.TypeInt)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, psu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{publicationstock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	psu.mutation.done = true
	return n, nil
}

// PublicationStockUpdateOne is the builder for updating a single PublicationStock entity.
type PublicationStockUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PublicationStockMutation
}

// SetIsbn sets the "isbn" field.
func (psuo *PublicationStockUpdateOne) SetIsbn(s string) *PublicationStockUpdateOne {
	psuo.mutation.SetIsbn(s)
	return psuo
}

// SetNillableIsbn sets the "isbn" field if the given value is not nil.
func (psuo *PublicationStockUpdateOne) SetNillableIsbn(s *string) *PublicationStockUpdateOne {
	if s != nil {
		psuo.SetIsbn(*s)
	}
	return psuo
}

// SetQuantity sets the "quantity" field.
func (psuo *PublicationStockUpdateOne) SetQuantity(i int) *PublicationStockUpdateOne {
	psuo.mutation.ResetQuantity()
	psuo.mutation.SetQuantity(i)
	return psuo
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (psuo *PublicationStockUpdateOne) SetNillableQuantity(i *int) *PublicationStockUpdateOne {
	if i != nil {
		psuo.SetQuantity(*i)
	}
	return psuo
}

// AddQuantity adds i to the "quantity" field.
func (psuo *PublicationStockUpdateOne) AddQuantity(i int) *PublicationStockUpdateOne {
	psuo.mutation.AddQuantity(i)
	return psuo
}

// ClearQuantity clears the value of the "quantity" field.
func (psuo *PublicationStockUpdateOne) ClearQuantity() *PublicationStockUpdateOne {
	psuo.mutation.ClearQuantity()
	return psuo
}

// SetPublicationID sets the "publication_id" field.
func (psuo *PublicationStockUpdateOne) SetPublicationID(i int) *PublicationStockUpdateOne {
	psuo.mutation.ResetPublicationID()
	psuo.mutation.SetPublicationID(i)
	return psuo
}

// SetNillablePublicationID sets the "publication_id" field if the given value is not nil.
func (psuo *PublicationStockUpdateOne) SetNillablePublicationID(i *int) *PublicationStockUpdateOne {
	if i != nil {
		psuo.SetPublicationID(*i)
	}
	return psuo
}

// AddPublicationID adds i to the "publication_id" field.
func (psuo *PublicationStockUpdateOne) AddPublicationID(i int) *PublicationStockUpdateOne {
	psuo.mutation.AddPublicationID(i)
	return psuo
}

// ClearPublicationID clears the value of the "publication_id" field.
func (psuo *PublicationStockUpdateOne) ClearPublicationID() *PublicationStockUpdateOne {
	psuo.mutation.ClearPublicationID()
	return psuo
}

// Mutation returns the PublicationStockMutation object of the builder.
func (psuo *PublicationStockUpdateOne) Mutation() *PublicationStockMutation {
	return psuo.mutation
}

// Where appends a list predicates to the PublicationStockUpdate builder.
func (psuo *PublicationStockUpdateOne) Where(ps ...predicate.PublicationStock) *PublicationStockUpdateOne {
	psuo.mutation.Where(ps...)
	return psuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (psuo *PublicationStockUpdateOne) Select(field string, fields ...string) *PublicationStockUpdateOne {
	psuo.fields = append([]string{field}, fields...)
	return psuo
}

// Save executes the query and returns the updated PublicationStock entity.
func (psuo *PublicationStockUpdateOne) Save(ctx context.Context) (*PublicationStock, error) {
	return withHooks(ctx, psuo.sqlSave, psuo.mutation, psuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (psuo *PublicationStockUpdateOne) SaveX(ctx context.Context) *PublicationStock {
	node, err := psuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (psuo *PublicationStockUpdateOne) Exec(ctx context.Context) error {
	_, err := psuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psuo *PublicationStockUpdateOne) ExecX(ctx context.Context) {
	if err := psuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (psuo *PublicationStockUpdateOne) sqlSave(ctx context.Context) (_node *PublicationStock, err error) {
	_spec := sqlgraph.NewUpdateSpec(publicationstock.Table, publicationstock.Columns, sqlgraph.NewFieldSpec(publicationstock.FieldID, field.TypeInt))
	id, ok := psuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PublicationStock.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := psuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, publicationstock.FieldID)
		for _, f := range fields {
			if !publicationstock.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != publicationstock.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := psuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := psuo.mutation.Isbn(); ok {
		_spec.SetField(publicationstock.FieldIsbn, field.TypeString, value)
	}
	if value, ok := psuo.mutation.Quantity(); ok {
		_spec.SetField(publicationstock.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := psuo.mutation.AddedQuantity(); ok {
		_spec.AddField(publicationstock.FieldQuantity, field.TypeInt, value)
	}
	if psuo.mutation.QuantityCleared() {
		_spec.ClearField(publicationstock.FieldQuantity, field.TypeInt)
	}
	if value, ok := psuo.mutation.PublicationID(); ok {
		_spec.SetField(publicationstock.FieldPublicationID, field.TypeInt, value)
	}
	if value, ok := psuo.mutation.AddedPublicationID(); ok {
		_spec.AddField(publicationstock.FieldPublicationID, field.TypeInt, value)
	}
	if psuo.mutation.PublicationIDCleared() {
		_spec.ClearField(publicationstock.FieldPublicationID, field.TypeInt)
	}
	_node = &PublicationStock{config: psuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, psuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{publicationstock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	psuo.mutation.done = true
	return _node, nil
}
