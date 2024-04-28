// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/xuanyiying/publication-shop/app/book/service/internal/data/ent/bookxcatalog"
	"github.com/xuanyiying/publication-shop/app/book/service/internal/data/ent/predicate"
)

// BookXCatalogUpdate is the builder for updating BookXCatalog entities.
type BookXCatalogUpdate struct {
	config
	hooks    []Hook
	mutation *BookXCatalogMutation
}

// Where appends a list predicates to the BookXCatalogUpdate builder.
func (bxu *BookXCatalogUpdate) Where(ps ...predicate.BookXCatalog) *BookXCatalogUpdate {
	bxu.mutation.Where(ps...)
	return bxu
}

// SetCatalogID sets the "catalog_id" field.
func (bxu *BookXCatalogUpdate) SetCatalogID(i int) *BookXCatalogUpdate {
	bxu.mutation.ResetCatalogID()
	bxu.mutation.SetCatalogID(i)
	return bxu
}

// SetNillableCatalogID sets the "catalog_id" field if the given value is not nil.
func (bxu *BookXCatalogUpdate) SetNillableCatalogID(i *int) *BookXCatalogUpdate {
	if i != nil {
		bxu.SetCatalogID(*i)
	}
	return bxu
}

// AddCatalogID adds i to the "catalog_id" field.
func (bxu *BookXCatalogUpdate) AddCatalogID(i int) *BookXCatalogUpdate {
	bxu.mutation.AddCatalogID(i)
	return bxu
}

// SetIsbn sets the "isbn" field.
func (bxu *BookXCatalogUpdate) SetIsbn(s string) *BookXCatalogUpdate {
	bxu.mutation.SetIsbn(s)
	return bxu
}

// SetNillableIsbn sets the "isbn" field if the given value is not nil.
func (bxu *BookXCatalogUpdate) SetNillableIsbn(s *string) *BookXCatalogUpdate {
	if s != nil {
		bxu.SetIsbn(*s)
	}
	return bxu
}

// SetCatalogName sets the "catalog_name" field.
func (bxu *BookXCatalogUpdate) SetCatalogName(s string) *BookXCatalogUpdate {
	bxu.mutation.SetCatalogName(s)
	return bxu
}

// SetNillableCatalogName sets the "catalog_name" field if the given value is not nil.
func (bxu *BookXCatalogUpdate) SetNillableCatalogName(s *string) *BookXCatalogUpdate {
	if s != nil {
		bxu.SetCatalogName(*s)
	}
	return bxu
}

// ClearCatalogName clears the value of the "catalog_name" field.
func (bxu *BookXCatalogUpdate) ClearCatalogName() *BookXCatalogUpdate {
	bxu.mutation.ClearCatalogName()
	return bxu
}

// SetBookID sets the "book_id" field.
func (bxu *BookXCatalogUpdate) SetBookID(i int) *BookXCatalogUpdate {
	bxu.mutation.ResetBookID()
	bxu.mutation.SetBookID(i)
	return bxu
}

// SetNillableBookID sets the "book_id" field if the given value is not nil.
func (bxu *BookXCatalogUpdate) SetNillableBookID(i *int) *BookXCatalogUpdate {
	if i != nil {
		bxu.SetBookID(*i)
	}
	return bxu
}

// AddBookID adds i to the "book_id" field.
func (bxu *BookXCatalogUpdate) AddBookID(i int) *BookXCatalogUpdate {
	bxu.mutation.AddBookID(i)
	return bxu
}

// ClearBookID clears the value of the "book_id" field.
func (bxu *BookXCatalogUpdate) ClearBookID() *BookXCatalogUpdate {
	bxu.mutation.ClearBookID()
	return bxu
}

// Mutation returns the BookXCatalogMutation object of the builder.
func (bxu *BookXCatalogUpdate) Mutation() *BookXCatalogMutation {
	return bxu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bxu *BookXCatalogUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bxu.sqlSave, bxu.mutation, bxu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bxu *BookXCatalogUpdate) SaveX(ctx context.Context) int {
	affected, err := bxu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bxu *BookXCatalogUpdate) Exec(ctx context.Context) error {
	_, err := bxu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bxu *BookXCatalogUpdate) ExecX(ctx context.Context) {
	if err := bxu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bxu *BookXCatalogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(bookxcatalog.Table, bookxcatalog.Columns, sqlgraph.NewFieldSpec(bookxcatalog.FieldID, field.TypeInt))
	if ps := bxu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bxu.mutation.CatalogID(); ok {
		_spec.SetField(bookxcatalog.FieldCatalogID, field.TypeInt, value)
	}
	if value, ok := bxu.mutation.AddedCatalogID(); ok {
		_spec.AddField(bookxcatalog.FieldCatalogID, field.TypeInt, value)
	}
	if value, ok := bxu.mutation.Isbn(); ok {
		_spec.SetField(bookxcatalog.FieldIsbn, field.TypeString, value)
	}
	if value, ok := bxu.mutation.CatalogName(); ok {
		_spec.SetField(bookxcatalog.FieldCatalogName, field.TypeString, value)
	}
	if bxu.mutation.CatalogNameCleared() {
		_spec.ClearField(bookxcatalog.FieldCatalogName, field.TypeString)
	}
	if value, ok := bxu.mutation.BookID(); ok {
		_spec.SetField(bookxcatalog.FieldBookID, field.TypeInt, value)
	}
	if value, ok := bxu.mutation.AddedBookID(); ok {
		_spec.AddField(bookxcatalog.FieldBookID, field.TypeInt, value)
	}
	if bxu.mutation.BookIDCleared() {
		_spec.ClearField(bookxcatalog.FieldBookID, field.TypeInt)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bxu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bookxcatalog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bxu.mutation.done = true
	return n, nil
}

// BookXCatalogUpdateOne is the builder for updating a single BookXCatalog entity.
type BookXCatalogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BookXCatalogMutation
}

// SetCatalogID sets the "catalog_id" field.
func (bxuo *BookXCatalogUpdateOne) SetCatalogID(i int) *BookXCatalogUpdateOne {
	bxuo.mutation.ResetCatalogID()
	bxuo.mutation.SetCatalogID(i)
	return bxuo
}

// SetNillableCatalogID sets the "catalog_id" field if the given value is not nil.
func (bxuo *BookXCatalogUpdateOne) SetNillableCatalogID(i *int) *BookXCatalogUpdateOne {
	if i != nil {
		bxuo.SetCatalogID(*i)
	}
	return bxuo
}

// AddCatalogID adds i to the "catalog_id" field.
func (bxuo *BookXCatalogUpdateOne) AddCatalogID(i int) *BookXCatalogUpdateOne {
	bxuo.mutation.AddCatalogID(i)
	return bxuo
}

// SetIsbn sets the "isbn" field.
func (bxuo *BookXCatalogUpdateOne) SetIsbn(s string) *BookXCatalogUpdateOne {
	bxuo.mutation.SetIsbn(s)
	return bxuo
}

// SetNillableIsbn sets the "isbn" field if the given value is not nil.
func (bxuo *BookXCatalogUpdateOne) SetNillableIsbn(s *string) *BookXCatalogUpdateOne {
	if s != nil {
		bxuo.SetIsbn(*s)
	}
	return bxuo
}

// SetCatalogName sets the "catalog_name" field.
func (bxuo *BookXCatalogUpdateOne) SetCatalogName(s string) *BookXCatalogUpdateOne {
	bxuo.mutation.SetCatalogName(s)
	return bxuo
}

// SetNillableCatalogName sets the "catalog_name" field if the given value is not nil.
func (bxuo *BookXCatalogUpdateOne) SetNillableCatalogName(s *string) *BookXCatalogUpdateOne {
	if s != nil {
		bxuo.SetCatalogName(*s)
	}
	return bxuo
}

// ClearCatalogName clears the value of the "catalog_name" field.
func (bxuo *BookXCatalogUpdateOne) ClearCatalogName() *BookXCatalogUpdateOne {
	bxuo.mutation.ClearCatalogName()
	return bxuo
}

// SetBookID sets the "book_id" field.
func (bxuo *BookXCatalogUpdateOne) SetBookID(i int) *BookXCatalogUpdateOne {
	bxuo.mutation.ResetBookID()
	bxuo.mutation.SetBookID(i)
	return bxuo
}

// SetNillableBookID sets the "book_id" field if the given value is not nil.
func (bxuo *BookXCatalogUpdateOne) SetNillableBookID(i *int) *BookXCatalogUpdateOne {
	if i != nil {
		bxuo.SetBookID(*i)
	}
	return bxuo
}

// AddBookID adds i to the "book_id" field.
func (bxuo *BookXCatalogUpdateOne) AddBookID(i int) *BookXCatalogUpdateOne {
	bxuo.mutation.AddBookID(i)
	return bxuo
}

// ClearBookID clears the value of the "book_id" field.
func (bxuo *BookXCatalogUpdateOne) ClearBookID() *BookXCatalogUpdateOne {
	bxuo.mutation.ClearBookID()
	return bxuo
}

// Mutation returns the BookXCatalogMutation object of the builder.
func (bxuo *BookXCatalogUpdateOne) Mutation() *BookXCatalogMutation {
	return bxuo.mutation
}

// Where appends a list predicates to the BookXCatalogUpdate builder.
func (bxuo *BookXCatalogUpdateOne) Where(ps ...predicate.BookXCatalog) *BookXCatalogUpdateOne {
	bxuo.mutation.Where(ps...)
	return bxuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bxuo *BookXCatalogUpdateOne) Select(field string, fields ...string) *BookXCatalogUpdateOne {
	bxuo.fields = append([]string{field}, fields...)
	return bxuo
}

// Save executes the query and returns the updated BookXCatalog entity.
func (bxuo *BookXCatalogUpdateOne) Save(ctx context.Context) (*BookXCatalog, error) {
	return withHooks(ctx, bxuo.sqlSave, bxuo.mutation, bxuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bxuo *BookXCatalogUpdateOne) SaveX(ctx context.Context) *BookXCatalog {
	node, err := bxuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bxuo *BookXCatalogUpdateOne) Exec(ctx context.Context) error {
	_, err := bxuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bxuo *BookXCatalogUpdateOne) ExecX(ctx context.Context) {
	if err := bxuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bxuo *BookXCatalogUpdateOne) sqlSave(ctx context.Context) (_node *BookXCatalog, err error) {
	_spec := sqlgraph.NewUpdateSpec(bookxcatalog.Table, bookxcatalog.Columns, sqlgraph.NewFieldSpec(bookxcatalog.FieldID, field.TypeInt))
	id, ok := bxuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "BookXCatalog.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bxuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bookxcatalog.FieldID)
		for _, f := range fields {
			if !bookxcatalog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != bookxcatalog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bxuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bxuo.mutation.CatalogID(); ok {
		_spec.SetField(bookxcatalog.FieldCatalogID, field.TypeInt, value)
	}
	if value, ok := bxuo.mutation.AddedCatalogID(); ok {
		_spec.AddField(bookxcatalog.FieldCatalogID, field.TypeInt, value)
	}
	if value, ok := bxuo.mutation.Isbn(); ok {
		_spec.SetField(bookxcatalog.FieldIsbn, field.TypeString, value)
	}
	if value, ok := bxuo.mutation.CatalogName(); ok {
		_spec.SetField(bookxcatalog.FieldCatalogName, field.TypeString, value)
	}
	if bxuo.mutation.CatalogNameCleared() {
		_spec.ClearField(bookxcatalog.FieldCatalogName, field.TypeString)
	}
	if value, ok := bxuo.mutation.BookID(); ok {
		_spec.SetField(bookxcatalog.FieldBookID, field.TypeInt, value)
	}
	if value, ok := bxuo.mutation.AddedBookID(); ok {
		_spec.AddField(bookxcatalog.FieldBookID, field.TypeInt, value)
	}
	if bxuo.mutation.BookIDCleared() {
		_spec.ClearField(bookxcatalog.FieldBookID, field.TypeInt)
	}
	_node = &BookXCatalog{config: bxuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bxuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bookxcatalog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	bxuo.mutation.done = true
	return _node, nil
}
