// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/xuanyiying/publication-shop/app/catalog/ent/language"
	"github.com/xuanyiying/publication-shop/app/catalog/ent/predicate"
)

// LanguageUpdate is the builder for updating Language entities.
type LanguageUpdate struct {
	config
	hooks    []Hook
	mutation *LanguageMutation
}

// Where appends a list predicates to the LanguageUpdate builder.
func (lu *LanguageUpdate) Where(ps ...predicate.Language) *LanguageUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetLanguageName sets the "language_name" field.
func (lu *LanguageUpdate) SetLanguageName(s string) *LanguageUpdate {
	lu.mutation.SetLanguageName(s)
	return lu
}

// SetNillableLanguageName sets the "language_name" field if the given value is not nil.
func (lu *LanguageUpdate) SetNillableLanguageName(s *string) *LanguageUpdate {
	if s != nil {
		lu.SetLanguageName(*s)
	}
	return lu
}

// ClearLanguageName clears the value of the "language_name" field.
func (lu *LanguageUpdate) ClearLanguageName() *LanguageUpdate {
	lu.mutation.ClearLanguageName()
	return lu
}

// SetLanguageCode sets the "language_code" field.
func (lu *LanguageUpdate) SetLanguageCode(s string) *LanguageUpdate {
	lu.mutation.SetLanguageCode(s)
	return lu
}

// SetNillableLanguageCode sets the "language_code" field if the given value is not nil.
func (lu *LanguageUpdate) SetNillableLanguageCode(s *string) *LanguageUpdate {
	if s != nil {
		lu.SetLanguageCode(*s)
	}
	return lu
}

// ClearLanguageCode clears the value of the "language_code" field.
func (lu *LanguageUpdate) ClearLanguageCode() *LanguageUpdate {
	lu.mutation.ClearLanguageCode()
	return lu
}

// Mutation returns the LanguageMutation object of the builder.
func (lu *LanguageUpdate) Mutation() *LanguageMutation {
	return lu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LanguageUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LanguageUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LanguageUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LanguageUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lu *LanguageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(language.Table, language.Columns, sqlgraph.NewFieldSpec(language.FieldID, field.TypeInt))
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.LanguageName(); ok {
		_spec.SetField(language.FieldLanguageName, field.TypeString, value)
	}
	if lu.mutation.LanguageNameCleared() {
		_spec.ClearField(language.FieldLanguageName, field.TypeString)
	}
	if value, ok := lu.mutation.LanguageCode(); ok {
		_spec.SetField(language.FieldLanguageCode, field.TypeString, value)
	}
	if lu.mutation.LanguageCodeCleared() {
		_spec.ClearField(language.FieldLanguageCode, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{language.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}

// LanguageUpdateOne is the builder for updating a single Language entity.
type LanguageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LanguageMutation
}

// SetLanguageName sets the "language_name" field.
func (luo *LanguageUpdateOne) SetLanguageName(s string) *LanguageUpdateOne {
	luo.mutation.SetLanguageName(s)
	return luo
}

// SetNillableLanguageName sets the "language_name" field if the given value is not nil.
func (luo *LanguageUpdateOne) SetNillableLanguageName(s *string) *LanguageUpdateOne {
	if s != nil {
		luo.SetLanguageName(*s)
	}
	return luo
}

// ClearLanguageName clears the value of the "language_name" field.
func (luo *LanguageUpdateOne) ClearLanguageName() *LanguageUpdateOne {
	luo.mutation.ClearLanguageName()
	return luo
}

// SetLanguageCode sets the "language_code" field.
func (luo *LanguageUpdateOne) SetLanguageCode(s string) *LanguageUpdateOne {
	luo.mutation.SetLanguageCode(s)
	return luo
}

// SetNillableLanguageCode sets the "language_code" field if the given value is not nil.
func (luo *LanguageUpdateOne) SetNillableLanguageCode(s *string) *LanguageUpdateOne {
	if s != nil {
		luo.SetLanguageCode(*s)
	}
	return luo
}

// ClearLanguageCode clears the value of the "language_code" field.
func (luo *LanguageUpdateOne) ClearLanguageCode() *LanguageUpdateOne {
	luo.mutation.ClearLanguageCode()
	return luo
}

// Mutation returns the LanguageMutation object of the builder.
func (luo *LanguageUpdateOne) Mutation() *LanguageMutation {
	return luo.mutation
}

// Where appends a list predicates to the LanguageUpdate builder.
func (luo *LanguageUpdateOne) Where(ps ...predicate.Language) *LanguageUpdateOne {
	luo.mutation.Where(ps...)
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LanguageUpdateOne) Select(field string, fields ...string) *LanguageUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Language entity.
func (luo *LanguageUpdateOne) Save(ctx context.Context) (*Language, error) {
	return withHooks(ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LanguageUpdateOne) SaveX(ctx context.Context) *Language {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LanguageUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LanguageUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (luo *LanguageUpdateOne) sqlSave(ctx context.Context) (_node *Language, err error) {
	_spec := sqlgraph.NewUpdateSpec(language.Table, language.Columns, sqlgraph.NewFieldSpec(language.FieldID, field.TypeInt))
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Language.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, language.FieldID)
		for _, f := range fields {
			if !language.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != language.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.LanguageName(); ok {
		_spec.SetField(language.FieldLanguageName, field.TypeString, value)
	}
	if luo.mutation.LanguageNameCleared() {
		_spec.ClearField(language.FieldLanguageName, field.TypeString)
	}
	if value, ok := luo.mutation.LanguageCode(); ok {
		_spec.SetField(language.FieldLanguageCode, field.TypeString, value)
	}
	if luo.mutation.LanguageCodeCleared() {
		_spec.ClearField(language.FieldLanguageCode, field.TypeString)
	}
	_node = &Language{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{language.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}
