// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/xuanyiying/publication-shop/app/book/service/internal/data/ent/bookauthor"
	"github.com/xuanyiying/publication-shop/app/book/service/internal/data/ent/predicate"
)

// BookAuthorUpdate is the builder for updating BookAuthor entities.
type BookAuthorUpdate struct {
	config
	hooks    []Hook
	mutation *BookAuthorMutation
}

// Where appends a list predicates to the BookAuthorUpdate builder.
func (bau *BookAuthorUpdate) Where(ps ...predicate.BookAuthor) *BookAuthorUpdate {
	bau.mutation.Where(ps...)
	return bau
}

// SetAuthorID sets the "author_id" field.
func (bau *BookAuthorUpdate) SetAuthorID(i int) *BookAuthorUpdate {
	bau.mutation.ResetAuthorID()
	bau.mutation.SetAuthorID(i)
	return bau
}

// SetNillableAuthorID sets the "author_id" field if the given value is not nil.
func (bau *BookAuthorUpdate) SetNillableAuthorID(i *int) *BookAuthorUpdate {
	if i != nil {
		bau.SetAuthorID(*i)
	}
	return bau
}

// AddAuthorID adds i to the "author_id" field.
func (bau *BookAuthorUpdate) AddAuthorID(i int) *BookAuthorUpdate {
	bau.mutation.AddAuthorID(i)
	return bau
}

// SetIsbn sets the "isbn" field.
func (bau *BookAuthorUpdate) SetIsbn(s string) *BookAuthorUpdate {
	bau.mutation.SetIsbn(s)
	return bau
}

// SetNillableIsbn sets the "isbn" field if the given value is not nil.
func (bau *BookAuthorUpdate) SetNillableIsbn(s *string) *BookAuthorUpdate {
	if s != nil {
		bau.SetIsbn(*s)
	}
	return bau
}

// SetBookID sets the "book_id" field.
func (bau *BookAuthorUpdate) SetBookID(i int) *BookAuthorUpdate {
	bau.mutation.ResetBookID()
	bau.mutation.SetBookID(i)
	return bau
}

// SetNillableBookID sets the "book_id" field if the given value is not nil.
func (bau *BookAuthorUpdate) SetNillableBookID(i *int) *BookAuthorUpdate {
	if i != nil {
		bau.SetBookID(*i)
	}
	return bau
}

// AddBookID adds i to the "book_id" field.
func (bau *BookAuthorUpdate) AddBookID(i int) *BookAuthorUpdate {
	bau.mutation.AddBookID(i)
	return bau
}

// ClearBookID clears the value of the "book_id" field.
func (bau *BookAuthorUpdate) ClearBookID() *BookAuthorUpdate {
	bau.mutation.ClearBookID()
	return bau
}

// SetAuthor sets the "author" field.
func (bau *BookAuthorUpdate) SetAuthor(s string) *BookAuthorUpdate {
	bau.mutation.SetAuthor(s)
	return bau
}

// SetNillableAuthor sets the "author" field if the given value is not nil.
func (bau *BookAuthorUpdate) SetNillableAuthor(s *string) *BookAuthorUpdate {
	if s != nil {
		bau.SetAuthor(*s)
	}
	return bau
}

// ClearAuthor clears the value of the "author" field.
func (bau *BookAuthorUpdate) ClearAuthor() *BookAuthorUpdate {
	bau.mutation.ClearAuthor()
	return bau
}

// SetAuthorAbout sets the "author_about" field.
func (bau *BookAuthorUpdate) SetAuthorAbout(s string) *BookAuthorUpdate {
	bau.mutation.SetAuthorAbout(s)
	return bau
}

// SetNillableAuthorAbout sets the "author_about" field if the given value is not nil.
func (bau *BookAuthorUpdate) SetNillableAuthorAbout(s *string) *BookAuthorUpdate {
	if s != nil {
		bau.SetAuthorAbout(*s)
	}
	return bau
}

// ClearAuthorAbout clears the value of the "author_about" field.
func (bau *BookAuthorUpdate) ClearAuthorAbout() *BookAuthorUpdate {
	bau.mutation.ClearAuthorAbout()
	return bau
}

// Mutation returns the BookAuthorMutation object of the builder.
func (bau *BookAuthorUpdate) Mutation() *BookAuthorMutation {
	return bau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bau *BookAuthorUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bau.sqlSave, bau.mutation, bau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bau *BookAuthorUpdate) SaveX(ctx context.Context) int {
	affected, err := bau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bau *BookAuthorUpdate) Exec(ctx context.Context) error {
	_, err := bau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bau *BookAuthorUpdate) ExecX(ctx context.Context) {
	if err := bau.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bau *BookAuthorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(bookauthor.Table, bookauthor.Columns, sqlgraph.NewFieldSpec(bookauthor.FieldID, field.TypeInt))
	if ps := bau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bau.mutation.AuthorID(); ok {
		_spec.SetField(bookauthor.FieldAuthorID, field.TypeInt, value)
	}
	if value, ok := bau.mutation.AddedAuthorID(); ok {
		_spec.AddField(bookauthor.FieldAuthorID, field.TypeInt, value)
	}
	if value, ok := bau.mutation.Isbn(); ok {
		_spec.SetField(bookauthor.FieldIsbn, field.TypeString, value)
	}
	if value, ok := bau.mutation.BookID(); ok {
		_spec.SetField(bookauthor.FieldBookID, field.TypeInt, value)
	}
	if value, ok := bau.mutation.AddedBookID(); ok {
		_spec.AddField(bookauthor.FieldBookID, field.TypeInt, value)
	}
	if bau.mutation.BookIDCleared() {
		_spec.ClearField(bookauthor.FieldBookID, field.TypeInt)
	}
	if value, ok := bau.mutation.Author(); ok {
		_spec.SetField(bookauthor.FieldAuthor, field.TypeString, value)
	}
	if bau.mutation.AuthorCleared() {
		_spec.ClearField(bookauthor.FieldAuthor, field.TypeString)
	}
	if value, ok := bau.mutation.AuthorAbout(); ok {
		_spec.SetField(bookauthor.FieldAuthorAbout, field.TypeString, value)
	}
	if bau.mutation.AuthorAboutCleared() {
		_spec.ClearField(bookauthor.FieldAuthorAbout, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bookauthor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bau.mutation.done = true
	return n, nil
}

// BookAuthorUpdateOne is the builder for updating a single BookAuthor entity.
type BookAuthorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BookAuthorMutation
}

// SetAuthorID sets the "author_id" field.
func (bauo *BookAuthorUpdateOne) SetAuthorID(i int) *BookAuthorUpdateOne {
	bauo.mutation.ResetAuthorID()
	bauo.mutation.SetAuthorID(i)
	return bauo
}

// SetNillableAuthorID sets the "author_id" field if the given value is not nil.
func (bauo *BookAuthorUpdateOne) SetNillableAuthorID(i *int) *BookAuthorUpdateOne {
	if i != nil {
		bauo.SetAuthorID(*i)
	}
	return bauo
}

// AddAuthorID adds i to the "author_id" field.
func (bauo *BookAuthorUpdateOne) AddAuthorID(i int) *BookAuthorUpdateOne {
	bauo.mutation.AddAuthorID(i)
	return bauo
}

// SetIsbn sets the "isbn" field.
func (bauo *BookAuthorUpdateOne) SetIsbn(s string) *BookAuthorUpdateOne {
	bauo.mutation.SetIsbn(s)
	return bauo
}

// SetNillableIsbn sets the "isbn" field if the given value is not nil.
func (bauo *BookAuthorUpdateOne) SetNillableIsbn(s *string) *BookAuthorUpdateOne {
	if s != nil {
		bauo.SetIsbn(*s)
	}
	return bauo
}

// SetBookID sets the "book_id" field.
func (bauo *BookAuthorUpdateOne) SetBookID(i int) *BookAuthorUpdateOne {
	bauo.mutation.ResetBookID()
	bauo.mutation.SetBookID(i)
	return bauo
}

// SetNillableBookID sets the "book_id" field if the given value is not nil.
func (bauo *BookAuthorUpdateOne) SetNillableBookID(i *int) *BookAuthorUpdateOne {
	if i != nil {
		bauo.SetBookID(*i)
	}
	return bauo
}

// AddBookID adds i to the "book_id" field.
func (bauo *BookAuthorUpdateOne) AddBookID(i int) *BookAuthorUpdateOne {
	bauo.mutation.AddBookID(i)
	return bauo
}

// ClearBookID clears the value of the "book_id" field.
func (bauo *BookAuthorUpdateOne) ClearBookID() *BookAuthorUpdateOne {
	bauo.mutation.ClearBookID()
	return bauo
}

// SetAuthor sets the "author" field.
func (bauo *BookAuthorUpdateOne) SetAuthor(s string) *BookAuthorUpdateOne {
	bauo.mutation.SetAuthor(s)
	return bauo
}

// SetNillableAuthor sets the "author" field if the given value is not nil.
func (bauo *BookAuthorUpdateOne) SetNillableAuthor(s *string) *BookAuthorUpdateOne {
	if s != nil {
		bauo.SetAuthor(*s)
	}
	return bauo
}

// ClearAuthor clears the value of the "author" field.
func (bauo *BookAuthorUpdateOne) ClearAuthor() *BookAuthorUpdateOne {
	bauo.mutation.ClearAuthor()
	return bauo
}

// SetAuthorAbout sets the "author_about" field.
func (bauo *BookAuthorUpdateOne) SetAuthorAbout(s string) *BookAuthorUpdateOne {
	bauo.mutation.SetAuthorAbout(s)
	return bauo
}

// SetNillableAuthorAbout sets the "author_about" field if the given value is not nil.
func (bauo *BookAuthorUpdateOne) SetNillableAuthorAbout(s *string) *BookAuthorUpdateOne {
	if s != nil {
		bauo.SetAuthorAbout(*s)
	}
	return bauo
}

// ClearAuthorAbout clears the value of the "author_about" field.
func (bauo *BookAuthorUpdateOne) ClearAuthorAbout() *BookAuthorUpdateOne {
	bauo.mutation.ClearAuthorAbout()
	return bauo
}

// Mutation returns the BookAuthorMutation object of the builder.
func (bauo *BookAuthorUpdateOne) Mutation() *BookAuthorMutation {
	return bauo.mutation
}

// Where appends a list predicates to the BookAuthorUpdate builder.
func (bauo *BookAuthorUpdateOne) Where(ps ...predicate.BookAuthor) *BookAuthorUpdateOne {
	bauo.mutation.Where(ps...)
	return bauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bauo *BookAuthorUpdateOne) Select(field string, fields ...string) *BookAuthorUpdateOne {
	bauo.fields = append([]string{field}, fields...)
	return bauo
}

// Save executes the query and returns the updated BookAuthor entity.
func (bauo *BookAuthorUpdateOne) Save(ctx context.Context) (*BookAuthor, error) {
	return withHooks(ctx, bauo.sqlSave, bauo.mutation, bauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bauo *BookAuthorUpdateOne) SaveX(ctx context.Context) *BookAuthor {
	node, err := bauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bauo *BookAuthorUpdateOne) Exec(ctx context.Context) error {
	_, err := bauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bauo *BookAuthorUpdateOne) ExecX(ctx context.Context) {
	if err := bauo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bauo *BookAuthorUpdateOne) sqlSave(ctx context.Context) (_node *BookAuthor, err error) {
	_spec := sqlgraph.NewUpdateSpec(bookauthor.Table, bookauthor.Columns, sqlgraph.NewFieldSpec(bookauthor.FieldID, field.TypeInt))
	id, ok := bauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "BookAuthor.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bookauthor.FieldID)
		for _, f := range fields {
			if !bookauthor.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != bookauthor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bauo.mutation.AuthorID(); ok {
		_spec.SetField(bookauthor.FieldAuthorID, field.TypeInt, value)
	}
	if value, ok := bauo.mutation.AddedAuthorID(); ok {
		_spec.AddField(bookauthor.FieldAuthorID, field.TypeInt, value)
	}
	if value, ok := bauo.mutation.Isbn(); ok {
		_spec.SetField(bookauthor.FieldIsbn, field.TypeString, value)
	}
	if value, ok := bauo.mutation.BookID(); ok {
		_spec.SetField(bookauthor.FieldBookID, field.TypeInt, value)
	}
	if value, ok := bauo.mutation.AddedBookID(); ok {
		_spec.AddField(bookauthor.FieldBookID, field.TypeInt, value)
	}
	if bauo.mutation.BookIDCleared() {
		_spec.ClearField(bookauthor.FieldBookID, field.TypeInt)
	}
	if value, ok := bauo.mutation.Author(); ok {
		_spec.SetField(bookauthor.FieldAuthor, field.TypeString, value)
	}
	if bauo.mutation.AuthorCleared() {
		_spec.ClearField(bookauthor.FieldAuthor, field.TypeString)
	}
	if value, ok := bauo.mutation.AuthorAbout(); ok {
		_spec.SetField(bookauthor.FieldAuthorAbout, field.TypeString, value)
	}
	if bauo.mutation.AuthorAboutCleared() {
		_spec.ClearField(bookauthor.FieldAuthorAbout, field.TypeString)
	}
	_node = &BookAuthor{config: bauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bookauthor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	bauo.mutation.done = true
	return _node, nil
}
