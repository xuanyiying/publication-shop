// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/xuanyiying/publication-shop/app/transaction/internal/data/ent/predicate"
	"github.com/xuanyiying/publication-shop/app/transaction/internal/data/ent/txitem"
)

// TxItemUpdate is the builder for updating TxItem entities.
type TxItemUpdate struct {
	config
	hooks    []Hook
	mutation *TxItemMutation
}

// Where appends a list predicates to the TxItemUpdate builder.
func (tiu *TxItemUpdate) Where(ps ...predicate.TxItem) *TxItemUpdate {
	tiu.mutation.Where(ps...)
	return tiu
}

// SetTxType sets the "tx_type" field.
func (tiu *TxItemUpdate) SetTxType(tt txitem.TxType) *TxItemUpdate {
	tiu.mutation.SetTxType(tt)
	return tiu
}

// SetNillableTxType sets the "tx_type" field if the given value is not nil.
func (tiu *TxItemUpdate) SetNillableTxType(tt *txitem.TxType) *TxItemUpdate {
	if tt != nil {
		tiu.SetTxType(*tt)
	}
	return tiu
}

// SetTxID sets the "tx_id" field.
func (tiu *TxItemUpdate) SetTxID(i int64) *TxItemUpdate {
	tiu.mutation.ResetTxID()
	tiu.mutation.SetTxID(i)
	return tiu
}

// SetNillableTxID sets the "tx_id" field if the given value is not nil.
func (tiu *TxItemUpdate) SetNillableTxID(i *int64) *TxItemUpdate {
	if i != nil {
		tiu.SetTxID(*i)
	}
	return tiu
}

// AddTxID adds i to the "tx_id" field.
func (tiu *TxItemUpdate) AddTxID(i int64) *TxItemUpdate {
	tiu.mutation.AddTxID(i)
	return tiu
}

// SetBookID sets the "book_id" field.
func (tiu *TxItemUpdate) SetBookID(i int64) *TxItemUpdate {
	tiu.mutation.ResetBookID()
	tiu.mutation.SetBookID(i)
	return tiu
}

// SetNillableBookID sets the "book_id" field if the given value is not nil.
func (tiu *TxItemUpdate) SetNillableBookID(i *int64) *TxItemUpdate {
	if i != nil {
		tiu.SetBookID(*i)
	}
	return tiu
}

// AddBookID adds i to the "book_id" field.
func (tiu *TxItemUpdate) AddBookID(i int64) *TxItemUpdate {
	tiu.mutation.AddBookID(i)
	return tiu
}

// SetQuantity sets the "quantity" field.
func (tiu *TxItemUpdate) SetQuantity(i int) *TxItemUpdate {
	tiu.mutation.ResetQuantity()
	tiu.mutation.SetQuantity(i)
	return tiu
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (tiu *TxItemUpdate) SetNillableQuantity(i *int) *TxItemUpdate {
	if i != nil {
		tiu.SetQuantity(*i)
	}
	return tiu
}

// AddQuantity adds i to the "quantity" field.
func (tiu *TxItemUpdate) AddQuantity(i int) *TxItemUpdate {
	tiu.mutation.AddQuantity(i)
	return tiu
}

// SetPrice sets the "price" field.
func (tiu *TxItemUpdate) SetPrice(f float64) *TxItemUpdate {
	tiu.mutation.ResetPrice()
	tiu.mutation.SetPrice(f)
	return tiu
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (tiu *TxItemUpdate) SetNillablePrice(f *float64) *TxItemUpdate {
	if f != nil {
		tiu.SetPrice(*f)
	}
	return tiu
}

// AddPrice adds f to the "price" field.
func (tiu *TxItemUpdate) AddPrice(f float64) *TxItemUpdate {
	tiu.mutation.AddPrice(f)
	return tiu
}

// SetIsbn sets the "isbn" field.
func (tiu *TxItemUpdate) SetIsbn(s string) *TxItemUpdate {
	tiu.mutation.SetIsbn(s)
	return tiu
}

// SetNillableIsbn sets the "isbn" field if the given value is not nil.
func (tiu *TxItemUpdate) SetNillableIsbn(s *string) *TxItemUpdate {
	if s != nil {
		tiu.SetIsbn(*s)
	}
	return tiu
}

// SetTitle sets the "title" field.
func (tiu *TxItemUpdate) SetTitle(s string) *TxItemUpdate {
	tiu.mutation.SetTitle(s)
	return tiu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (tiu *TxItemUpdate) SetNillableTitle(s *string) *TxItemUpdate {
	if s != nil {
		tiu.SetTitle(*s)
	}
	return tiu
}

// SetAuthor sets the "author" field.
func (tiu *TxItemUpdate) SetAuthor(s string) *TxItemUpdate {
	tiu.mutation.SetAuthor(s)
	return tiu
}

// SetNillableAuthor sets the "author" field if the given value is not nil.
func (tiu *TxItemUpdate) SetNillableAuthor(s *string) *TxItemUpdate {
	if s != nil {
		tiu.SetAuthor(*s)
	}
	return tiu
}

// SetPublisherID sets the "publisher_id" field.
func (tiu *TxItemUpdate) SetPublisherID(i int64) *TxItemUpdate {
	tiu.mutation.ResetPublisherID()
	tiu.mutation.SetPublisherID(i)
	return tiu
}

// SetNillablePublisherID sets the "publisher_id" field if the given value is not nil.
func (tiu *TxItemUpdate) SetNillablePublisherID(i *int64) *TxItemUpdate {
	if i != nil {
		tiu.SetPublisherID(*i)
	}
	return tiu
}

// AddPublisherID adds i to the "publisher_id" field.
func (tiu *TxItemUpdate) AddPublisherID(i int64) *TxItemUpdate {
	tiu.mutation.AddPublisherID(i)
	return tiu
}

// ClearPublisherID clears the value of the "publisher_id" field.
func (tiu *TxItemUpdate) ClearPublisherID() *TxItemUpdate {
	tiu.mutation.ClearPublisherID()
	return tiu
}

// SetImageURL sets the "image_url" field.
func (tiu *TxItemUpdate) SetImageURL(s string) *TxItemUpdate {
	tiu.mutation.SetImageURL(s)
	return tiu
}

// SetNillableImageURL sets the "image_url" field if the given value is not nil.
func (tiu *TxItemUpdate) SetNillableImageURL(s *string) *TxItemUpdate {
	if s != nil {
		tiu.SetImageURL(*s)
	}
	return tiu
}

// ClearImageURL clears the value of the "image_url" field.
func (tiu *TxItemUpdate) ClearImageURL() *TxItemUpdate {
	tiu.mutation.ClearImageURL()
	return tiu
}

// Mutation returns the TxItemMutation object of the builder.
func (tiu *TxItemUpdate) Mutation() *TxItemMutation {
	return tiu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tiu *TxItemUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tiu.sqlSave, tiu.mutation, tiu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tiu *TxItemUpdate) SaveX(ctx context.Context) int {
	affected, err := tiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tiu *TxItemUpdate) Exec(ctx context.Context) error {
	_, err := tiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tiu *TxItemUpdate) ExecX(ctx context.Context) {
	if err := tiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tiu *TxItemUpdate) check() error {
	if v, ok := tiu.mutation.TxType(); ok {
		if err := txitem.TxTypeValidator(v); err != nil {
			return &ValidationError{Name: "tx_type", err: fmt.Errorf(`ent: validator failed for field "TxItem.tx_type": %w`, err)}
		}
	}
	return nil
}

func (tiu *TxItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tiu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(txitem.Table, txitem.Columns, sqlgraph.NewFieldSpec(txitem.FieldID, field.TypeInt64))
	if ps := tiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tiu.mutation.TxType(); ok {
		_spec.SetField(txitem.FieldTxType, field.TypeEnum, value)
	}
	if value, ok := tiu.mutation.TxID(); ok {
		_spec.SetField(txitem.FieldTxID, field.TypeInt64, value)
	}
	if value, ok := tiu.mutation.AddedTxID(); ok {
		_spec.AddField(txitem.FieldTxID, field.TypeInt64, value)
	}
	if value, ok := tiu.mutation.BookID(); ok {
		_spec.SetField(txitem.FieldBookID, field.TypeInt64, value)
	}
	if value, ok := tiu.mutation.AddedBookID(); ok {
		_spec.AddField(txitem.FieldBookID, field.TypeInt64, value)
	}
	if value, ok := tiu.mutation.Quantity(); ok {
		_spec.SetField(txitem.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := tiu.mutation.AddedQuantity(); ok {
		_spec.AddField(txitem.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := tiu.mutation.Price(); ok {
		_spec.SetField(txitem.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := tiu.mutation.AddedPrice(); ok {
		_spec.AddField(txitem.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := tiu.mutation.Isbn(); ok {
		_spec.SetField(txitem.FieldIsbn, field.TypeString, value)
	}
	if value, ok := tiu.mutation.Title(); ok {
		_spec.SetField(txitem.FieldTitle, field.TypeString, value)
	}
	if value, ok := tiu.mutation.Author(); ok {
		_spec.SetField(txitem.FieldAuthor, field.TypeString, value)
	}
	if value, ok := tiu.mutation.PublisherID(); ok {
		_spec.SetField(txitem.FieldPublisherID, field.TypeInt64, value)
	}
	if value, ok := tiu.mutation.AddedPublisherID(); ok {
		_spec.AddField(txitem.FieldPublisherID, field.TypeInt64, value)
	}
	if tiu.mutation.PublisherIDCleared() {
		_spec.ClearField(txitem.FieldPublisherID, field.TypeInt64)
	}
	if value, ok := tiu.mutation.ImageURL(); ok {
		_spec.SetField(txitem.FieldImageURL, field.TypeString, value)
	}
	if tiu.mutation.ImageURLCleared() {
		_spec.ClearField(txitem.FieldImageURL, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{txitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tiu.mutation.done = true
	return n, nil
}

// TxItemUpdateOne is the builder for updating a single TxItem entity.
type TxItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TxItemMutation
}

// SetTxType sets the "tx_type" field.
func (tiuo *TxItemUpdateOne) SetTxType(tt txitem.TxType) *TxItemUpdateOne {
	tiuo.mutation.SetTxType(tt)
	return tiuo
}

// SetNillableTxType sets the "tx_type" field if the given value is not nil.
func (tiuo *TxItemUpdateOne) SetNillableTxType(tt *txitem.TxType) *TxItemUpdateOne {
	if tt != nil {
		tiuo.SetTxType(*tt)
	}
	return tiuo
}

// SetTxID sets the "tx_id" field.
func (tiuo *TxItemUpdateOne) SetTxID(i int64) *TxItemUpdateOne {
	tiuo.mutation.ResetTxID()
	tiuo.mutation.SetTxID(i)
	return tiuo
}

// SetNillableTxID sets the "tx_id" field if the given value is not nil.
func (tiuo *TxItemUpdateOne) SetNillableTxID(i *int64) *TxItemUpdateOne {
	if i != nil {
		tiuo.SetTxID(*i)
	}
	return tiuo
}

// AddTxID adds i to the "tx_id" field.
func (tiuo *TxItemUpdateOne) AddTxID(i int64) *TxItemUpdateOne {
	tiuo.mutation.AddTxID(i)
	return tiuo
}

// SetBookID sets the "book_id" field.
func (tiuo *TxItemUpdateOne) SetBookID(i int64) *TxItemUpdateOne {
	tiuo.mutation.ResetBookID()
	tiuo.mutation.SetBookID(i)
	return tiuo
}

// SetNillableBookID sets the "book_id" field if the given value is not nil.
func (tiuo *TxItemUpdateOne) SetNillableBookID(i *int64) *TxItemUpdateOne {
	if i != nil {
		tiuo.SetBookID(*i)
	}
	return tiuo
}

// AddBookID adds i to the "book_id" field.
func (tiuo *TxItemUpdateOne) AddBookID(i int64) *TxItemUpdateOne {
	tiuo.mutation.AddBookID(i)
	return tiuo
}

// SetQuantity sets the "quantity" field.
func (tiuo *TxItemUpdateOne) SetQuantity(i int) *TxItemUpdateOne {
	tiuo.mutation.ResetQuantity()
	tiuo.mutation.SetQuantity(i)
	return tiuo
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (tiuo *TxItemUpdateOne) SetNillableQuantity(i *int) *TxItemUpdateOne {
	if i != nil {
		tiuo.SetQuantity(*i)
	}
	return tiuo
}

// AddQuantity adds i to the "quantity" field.
func (tiuo *TxItemUpdateOne) AddQuantity(i int) *TxItemUpdateOne {
	tiuo.mutation.AddQuantity(i)
	return tiuo
}

// SetPrice sets the "price" field.
func (tiuo *TxItemUpdateOne) SetPrice(f float64) *TxItemUpdateOne {
	tiuo.mutation.ResetPrice()
	tiuo.mutation.SetPrice(f)
	return tiuo
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (tiuo *TxItemUpdateOne) SetNillablePrice(f *float64) *TxItemUpdateOne {
	if f != nil {
		tiuo.SetPrice(*f)
	}
	return tiuo
}

// AddPrice adds f to the "price" field.
func (tiuo *TxItemUpdateOne) AddPrice(f float64) *TxItemUpdateOne {
	tiuo.mutation.AddPrice(f)
	return tiuo
}

// SetIsbn sets the "isbn" field.
func (tiuo *TxItemUpdateOne) SetIsbn(s string) *TxItemUpdateOne {
	tiuo.mutation.SetIsbn(s)
	return tiuo
}

// SetNillableIsbn sets the "isbn" field if the given value is not nil.
func (tiuo *TxItemUpdateOne) SetNillableIsbn(s *string) *TxItemUpdateOne {
	if s != nil {
		tiuo.SetIsbn(*s)
	}
	return tiuo
}

// SetTitle sets the "title" field.
func (tiuo *TxItemUpdateOne) SetTitle(s string) *TxItemUpdateOne {
	tiuo.mutation.SetTitle(s)
	return tiuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (tiuo *TxItemUpdateOne) SetNillableTitle(s *string) *TxItemUpdateOne {
	if s != nil {
		tiuo.SetTitle(*s)
	}
	return tiuo
}

// SetAuthor sets the "author" field.
func (tiuo *TxItemUpdateOne) SetAuthor(s string) *TxItemUpdateOne {
	tiuo.mutation.SetAuthor(s)
	return tiuo
}

// SetNillableAuthor sets the "author" field if the given value is not nil.
func (tiuo *TxItemUpdateOne) SetNillableAuthor(s *string) *TxItemUpdateOne {
	if s != nil {
		tiuo.SetAuthor(*s)
	}
	return tiuo
}

// SetPublisherID sets the "publisher_id" field.
func (tiuo *TxItemUpdateOne) SetPublisherID(i int64) *TxItemUpdateOne {
	tiuo.mutation.ResetPublisherID()
	tiuo.mutation.SetPublisherID(i)
	return tiuo
}

// SetNillablePublisherID sets the "publisher_id" field if the given value is not nil.
func (tiuo *TxItemUpdateOne) SetNillablePublisherID(i *int64) *TxItemUpdateOne {
	if i != nil {
		tiuo.SetPublisherID(*i)
	}
	return tiuo
}

// AddPublisherID adds i to the "publisher_id" field.
func (tiuo *TxItemUpdateOne) AddPublisherID(i int64) *TxItemUpdateOne {
	tiuo.mutation.AddPublisherID(i)
	return tiuo
}

// ClearPublisherID clears the value of the "publisher_id" field.
func (tiuo *TxItemUpdateOne) ClearPublisherID() *TxItemUpdateOne {
	tiuo.mutation.ClearPublisherID()
	return tiuo
}

// SetImageURL sets the "image_url" field.
func (tiuo *TxItemUpdateOne) SetImageURL(s string) *TxItemUpdateOne {
	tiuo.mutation.SetImageURL(s)
	return tiuo
}

// SetNillableImageURL sets the "image_url" field if the given value is not nil.
func (tiuo *TxItemUpdateOne) SetNillableImageURL(s *string) *TxItemUpdateOne {
	if s != nil {
		tiuo.SetImageURL(*s)
	}
	return tiuo
}

// ClearImageURL clears the value of the "image_url" field.
func (tiuo *TxItemUpdateOne) ClearImageURL() *TxItemUpdateOne {
	tiuo.mutation.ClearImageURL()
	return tiuo
}

// Mutation returns the TxItemMutation object of the builder.
func (tiuo *TxItemUpdateOne) Mutation() *TxItemMutation {
	return tiuo.mutation
}

// Where appends a list predicates to the TxItemUpdate builder.
func (tiuo *TxItemUpdateOne) Where(ps ...predicate.TxItem) *TxItemUpdateOne {
	tiuo.mutation.Where(ps...)
	return tiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tiuo *TxItemUpdateOne) Select(field string, fields ...string) *TxItemUpdateOne {
	tiuo.fields = append([]string{field}, fields...)
	return tiuo
}

// Save executes the query and returns the updated TxItem entity.
func (tiuo *TxItemUpdateOne) Save(ctx context.Context) (*TxItem, error) {
	return withHooks(ctx, tiuo.sqlSave, tiuo.mutation, tiuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tiuo *TxItemUpdateOne) SaveX(ctx context.Context) *TxItem {
	node, err := tiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tiuo *TxItemUpdateOne) Exec(ctx context.Context) error {
	_, err := tiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tiuo *TxItemUpdateOne) ExecX(ctx context.Context) {
	if err := tiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tiuo *TxItemUpdateOne) check() error {
	if v, ok := tiuo.mutation.TxType(); ok {
		if err := txitem.TxTypeValidator(v); err != nil {
			return &ValidationError{Name: "tx_type", err: fmt.Errorf(`ent: validator failed for field "TxItem.tx_type": %w`, err)}
		}
	}
	return nil
}

func (tiuo *TxItemUpdateOne) sqlSave(ctx context.Context) (_node *TxItem, err error) {
	if err := tiuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(txitem.Table, txitem.Columns, sqlgraph.NewFieldSpec(txitem.FieldID, field.TypeInt64))
	id, ok := tiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TxItem.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, txitem.FieldID)
		for _, f := range fields {
			if !txitem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != txitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tiuo.mutation.TxType(); ok {
		_spec.SetField(txitem.FieldTxType, field.TypeEnum, value)
	}
	if value, ok := tiuo.mutation.TxID(); ok {
		_spec.SetField(txitem.FieldTxID, field.TypeInt64, value)
	}
	if value, ok := tiuo.mutation.AddedTxID(); ok {
		_spec.AddField(txitem.FieldTxID, field.TypeInt64, value)
	}
	if value, ok := tiuo.mutation.BookID(); ok {
		_spec.SetField(txitem.FieldBookID, field.TypeInt64, value)
	}
	if value, ok := tiuo.mutation.AddedBookID(); ok {
		_spec.AddField(txitem.FieldBookID, field.TypeInt64, value)
	}
	if value, ok := tiuo.mutation.Quantity(); ok {
		_spec.SetField(txitem.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := tiuo.mutation.AddedQuantity(); ok {
		_spec.AddField(txitem.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := tiuo.mutation.Price(); ok {
		_spec.SetField(txitem.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := tiuo.mutation.AddedPrice(); ok {
		_spec.AddField(txitem.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := tiuo.mutation.Isbn(); ok {
		_spec.SetField(txitem.FieldIsbn, field.TypeString, value)
	}
	if value, ok := tiuo.mutation.Title(); ok {
		_spec.SetField(txitem.FieldTitle, field.TypeString, value)
	}
	if value, ok := tiuo.mutation.Author(); ok {
		_spec.SetField(txitem.FieldAuthor, field.TypeString, value)
	}
	if value, ok := tiuo.mutation.PublisherID(); ok {
		_spec.SetField(txitem.FieldPublisherID, field.TypeInt64, value)
	}
	if value, ok := tiuo.mutation.AddedPublisherID(); ok {
		_spec.AddField(txitem.FieldPublisherID, field.TypeInt64, value)
	}
	if tiuo.mutation.PublisherIDCleared() {
		_spec.ClearField(txitem.FieldPublisherID, field.TypeInt64)
	}
	if value, ok := tiuo.mutation.ImageURL(); ok {
		_spec.SetField(txitem.FieldImageURL, field.TypeString, value)
	}
	if tiuo.mutation.ImageURLCleared() {
		_spec.ClearField(txitem.FieldImageURL, field.TypeString)
	}
	_node = &TxItem{config: tiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{txitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tiuo.mutation.done = true
	return _node, nil
}