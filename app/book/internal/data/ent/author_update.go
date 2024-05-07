// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/xuanyiying/publication-shop/app/book/internal/data/ent/author"
	"github.com/xuanyiying/publication-shop/app/book/internal/data/ent/predicate"
)

// AuthorUpdate is the builder for updating Author entities.
type AuthorUpdate struct {
	config
	hooks    []Hook
	mutation *AuthorMutation
}

// Where appends a list predicates to the AuthorUpdate builder.
func (au *AuthorUpdate) Where(ps ...predicate.Author) *AuthorUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetName sets the "name" field.
func (au *AuthorUpdate) SetName(s string) *AuthorUpdate {
	au.mutation.SetName(s)
	return au
}

// SetNillableName sets the "name" field if the given value is not nil.
func (au *AuthorUpdate) SetNillableName(s *string) *AuthorUpdate {
	if s != nil {
		au.SetName(*s)
	}
	return au
}

// SetBiography sets the "biography" field.
func (au *AuthorUpdate) SetBiography(s string) *AuthorUpdate {
	au.mutation.SetBiography(s)
	return au
}

// SetNillableBiography sets the "biography" field if the given value is not nil.
func (au *AuthorUpdate) SetNillableBiography(s *string) *AuthorUpdate {
	if s != nil {
		au.SetBiography(*s)
	}
	return au
}

// ClearBiography clears the value of the "biography" field.
func (au *AuthorUpdate) ClearBiography() *AuthorUpdate {
	au.mutation.ClearBiography()
	return au
}

// Mutation returns the AuthorMutation object of the builder.
func (au *AuthorUpdate) Mutation() *AuthorMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AuthorUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AuthorUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AuthorUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AuthorUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AuthorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(author.Table, author.Columns, sqlgraph.NewFieldSpec(author.FieldID, field.TypeInt32))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.SetField(author.FieldName, field.TypeString, value)
	}
	if value, ok := au.mutation.Biography(); ok {
		_spec.SetField(author.FieldBiography, field.TypeString, value)
	}
	if au.mutation.BiographyCleared() {
		_spec.ClearField(author.FieldBiography, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{author.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AuthorUpdateOne is the builder for updating a single Author entity.
type AuthorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AuthorMutation
}

// SetName sets the "name" field.
func (auo *AuthorUpdateOne) SetName(s string) *AuthorUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (auo *AuthorUpdateOne) SetNillableName(s *string) *AuthorUpdateOne {
	if s != nil {
		auo.SetName(*s)
	}
	return auo
}

// SetBiography sets the "biography" field.
func (auo *AuthorUpdateOne) SetBiography(s string) *AuthorUpdateOne {
	auo.mutation.SetBiography(s)
	return auo
}

// SetNillableBiography sets the "biography" field if the given value is not nil.
func (auo *AuthorUpdateOne) SetNillableBiography(s *string) *AuthorUpdateOne {
	if s != nil {
		auo.SetBiography(*s)
	}
	return auo
}

// ClearBiography clears the value of the "biography" field.
func (auo *AuthorUpdateOne) ClearBiography() *AuthorUpdateOne {
	auo.mutation.ClearBiography()
	return auo
}

// Mutation returns the AuthorMutation object of the builder.
func (auo *AuthorUpdateOne) Mutation() *AuthorMutation {
	return auo.mutation
}

// Where appends a list predicates to the AuthorUpdate builder.
func (auo *AuthorUpdateOne) Where(ps ...predicate.Author) *AuthorUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AuthorUpdateOne) Select(field string, fields ...string) *AuthorUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Author entity.
func (auo *AuthorUpdateOne) Save(ctx context.Context) (*Author, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AuthorUpdateOne) SaveX(ctx context.Context) *Author {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AuthorUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AuthorUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AuthorUpdateOne) sqlSave(ctx context.Context) (_node *Author, err error) {
	_spec := sqlgraph.NewUpdateSpec(author.Table, author.Columns, sqlgraph.NewFieldSpec(author.FieldID, field.TypeInt32))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Author.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, author.FieldID)
		for _, f := range fields {
			if !author.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != author.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.SetField(author.FieldName, field.TypeString, value)
	}
	if value, ok := auo.mutation.Biography(); ok {
		_spec.SetField(author.FieldBiography, field.TypeString, value)
	}
	if auo.mutation.BiographyCleared() {
		_spec.ClearField(author.FieldBiography, field.TypeString)
	}
	_node = &Author{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{author.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}