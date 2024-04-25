// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/xuanyiying/publication-shop/app/catalog/ent/predicate"
	"github.com/xuanyiying/publication-shop/app/catalog/ent/publicationxclassic"
)

// PublicationXClassicUpdate is the builder for updating PublicationXClassic entities.
type PublicationXClassicUpdate struct {
	config
	hooks    []Hook
	mutation *PublicationXClassicMutation
}

// Where appends a list predicates to the PublicationXClassicUpdate builder.
func (pxu *PublicationXClassicUpdate) Where(ps ...predicate.PublicationXClassic) *PublicationXClassicUpdate {
	pxu.mutation.Where(ps...)
	return pxu
}

// SetClassicID sets the "classic_id" field.
func (pxu *PublicationXClassicUpdate) SetClassicID(i int) *PublicationXClassicUpdate {
	pxu.mutation.ResetClassicID()
	pxu.mutation.SetClassicID(i)
	return pxu
}

// SetNillableClassicID sets the "classic_id" field if the given value is not nil.
func (pxu *PublicationXClassicUpdate) SetNillableClassicID(i *int) *PublicationXClassicUpdate {
	if i != nil {
		pxu.SetClassicID(*i)
	}
	return pxu
}

// AddClassicID adds i to the "classic_id" field.
func (pxu *PublicationXClassicUpdate) AddClassicID(i int) *PublicationXClassicUpdate {
	pxu.mutation.AddClassicID(i)
	return pxu
}

// SetIsbn sets the "isbn" field.
func (pxu *PublicationXClassicUpdate) SetIsbn(s string) *PublicationXClassicUpdate {
	pxu.mutation.SetIsbn(s)
	return pxu
}

// SetNillableIsbn sets the "isbn" field if the given value is not nil.
func (pxu *PublicationXClassicUpdate) SetNillableIsbn(s *string) *PublicationXClassicUpdate {
	if s != nil {
		pxu.SetIsbn(*s)
	}
	return pxu
}

// SetClassicName sets the "classic_name" field.
func (pxu *PublicationXClassicUpdate) SetClassicName(s string) *PublicationXClassicUpdate {
	pxu.mutation.SetClassicName(s)
	return pxu
}

// SetNillableClassicName sets the "classic_name" field if the given value is not nil.
func (pxu *PublicationXClassicUpdate) SetNillableClassicName(s *string) *PublicationXClassicUpdate {
	if s != nil {
		pxu.SetClassicName(*s)
	}
	return pxu
}

// ClearClassicName clears the value of the "classic_name" field.
func (pxu *PublicationXClassicUpdate) ClearClassicName() *PublicationXClassicUpdate {
	pxu.mutation.ClearClassicName()
	return pxu
}

// SetPublicationID sets the "publication_id" field.
func (pxu *PublicationXClassicUpdate) SetPublicationID(i int) *PublicationXClassicUpdate {
	pxu.mutation.ResetPublicationID()
	pxu.mutation.SetPublicationID(i)
	return pxu
}

// SetNillablePublicationID sets the "publication_id" field if the given value is not nil.
func (pxu *PublicationXClassicUpdate) SetNillablePublicationID(i *int) *PublicationXClassicUpdate {
	if i != nil {
		pxu.SetPublicationID(*i)
	}
	return pxu
}

// AddPublicationID adds i to the "publication_id" field.
func (pxu *PublicationXClassicUpdate) AddPublicationID(i int) *PublicationXClassicUpdate {
	pxu.mutation.AddPublicationID(i)
	return pxu
}

// ClearPublicationID clears the value of the "publication_id" field.
func (pxu *PublicationXClassicUpdate) ClearPublicationID() *PublicationXClassicUpdate {
	pxu.mutation.ClearPublicationID()
	return pxu
}

// Mutation returns the PublicationXClassicMutation object of the builder.
func (pxu *PublicationXClassicUpdate) Mutation() *PublicationXClassicMutation {
	return pxu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pxu *PublicationXClassicUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pxu.sqlSave, pxu.mutation, pxu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pxu *PublicationXClassicUpdate) SaveX(ctx context.Context) int {
	affected, err := pxu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pxu *PublicationXClassicUpdate) Exec(ctx context.Context) error {
	_, err := pxu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pxu *PublicationXClassicUpdate) ExecX(ctx context.Context) {
	if err := pxu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pxu *PublicationXClassicUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(publicationxclassic.Table, publicationxclassic.Columns, sqlgraph.NewFieldSpec(publicationxclassic.FieldID, field.TypeInt))
	if ps := pxu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pxu.mutation.ClassicID(); ok {
		_spec.SetField(publicationxclassic.FieldClassicID, field.TypeInt, value)
	}
	if value, ok := pxu.mutation.AddedClassicID(); ok {
		_spec.AddField(publicationxclassic.FieldClassicID, field.TypeInt, value)
	}
	if value, ok := pxu.mutation.Isbn(); ok {
		_spec.SetField(publicationxclassic.FieldIsbn, field.TypeString, value)
	}
	if value, ok := pxu.mutation.ClassicName(); ok {
		_spec.SetField(publicationxclassic.FieldClassicName, field.TypeString, value)
	}
	if pxu.mutation.ClassicNameCleared() {
		_spec.ClearField(publicationxclassic.FieldClassicName, field.TypeString)
	}
	if value, ok := pxu.mutation.PublicationID(); ok {
		_spec.SetField(publicationxclassic.FieldPublicationID, field.TypeInt, value)
	}
	if value, ok := pxu.mutation.AddedPublicationID(); ok {
		_spec.AddField(publicationxclassic.FieldPublicationID, field.TypeInt, value)
	}
	if pxu.mutation.PublicationIDCleared() {
		_spec.ClearField(publicationxclassic.FieldPublicationID, field.TypeInt)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pxu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{publicationxclassic.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pxu.mutation.done = true
	return n, nil
}

// PublicationXClassicUpdateOne is the builder for updating a single PublicationXClassic entity.
type PublicationXClassicUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PublicationXClassicMutation
}

// SetClassicID sets the "classic_id" field.
func (pxuo *PublicationXClassicUpdateOne) SetClassicID(i int) *PublicationXClassicUpdateOne {
	pxuo.mutation.ResetClassicID()
	pxuo.mutation.SetClassicID(i)
	return pxuo
}

// SetNillableClassicID sets the "classic_id" field if the given value is not nil.
func (pxuo *PublicationXClassicUpdateOne) SetNillableClassicID(i *int) *PublicationXClassicUpdateOne {
	if i != nil {
		pxuo.SetClassicID(*i)
	}
	return pxuo
}

// AddClassicID adds i to the "classic_id" field.
func (pxuo *PublicationXClassicUpdateOne) AddClassicID(i int) *PublicationXClassicUpdateOne {
	pxuo.mutation.AddClassicID(i)
	return pxuo
}

// SetIsbn sets the "isbn" field.
func (pxuo *PublicationXClassicUpdateOne) SetIsbn(s string) *PublicationXClassicUpdateOne {
	pxuo.mutation.SetIsbn(s)
	return pxuo
}

// SetNillableIsbn sets the "isbn" field if the given value is not nil.
func (pxuo *PublicationXClassicUpdateOne) SetNillableIsbn(s *string) *PublicationXClassicUpdateOne {
	if s != nil {
		pxuo.SetIsbn(*s)
	}
	return pxuo
}

// SetClassicName sets the "classic_name" field.
func (pxuo *PublicationXClassicUpdateOne) SetClassicName(s string) *PublicationXClassicUpdateOne {
	pxuo.mutation.SetClassicName(s)
	return pxuo
}

// SetNillableClassicName sets the "classic_name" field if the given value is not nil.
func (pxuo *PublicationXClassicUpdateOne) SetNillableClassicName(s *string) *PublicationXClassicUpdateOne {
	if s != nil {
		pxuo.SetClassicName(*s)
	}
	return pxuo
}

// ClearClassicName clears the value of the "classic_name" field.
func (pxuo *PublicationXClassicUpdateOne) ClearClassicName() *PublicationXClassicUpdateOne {
	pxuo.mutation.ClearClassicName()
	return pxuo
}

// SetPublicationID sets the "publication_id" field.
func (pxuo *PublicationXClassicUpdateOne) SetPublicationID(i int) *PublicationXClassicUpdateOne {
	pxuo.mutation.ResetPublicationID()
	pxuo.mutation.SetPublicationID(i)
	return pxuo
}

// SetNillablePublicationID sets the "publication_id" field if the given value is not nil.
func (pxuo *PublicationXClassicUpdateOne) SetNillablePublicationID(i *int) *PublicationXClassicUpdateOne {
	if i != nil {
		pxuo.SetPublicationID(*i)
	}
	return pxuo
}

// AddPublicationID adds i to the "publication_id" field.
func (pxuo *PublicationXClassicUpdateOne) AddPublicationID(i int) *PublicationXClassicUpdateOne {
	pxuo.mutation.AddPublicationID(i)
	return pxuo
}

// ClearPublicationID clears the value of the "publication_id" field.
func (pxuo *PublicationXClassicUpdateOne) ClearPublicationID() *PublicationXClassicUpdateOne {
	pxuo.mutation.ClearPublicationID()
	return pxuo
}

// Mutation returns the PublicationXClassicMutation object of the builder.
func (pxuo *PublicationXClassicUpdateOne) Mutation() *PublicationXClassicMutation {
	return pxuo.mutation
}

// Where appends a list predicates to the PublicationXClassicUpdate builder.
func (pxuo *PublicationXClassicUpdateOne) Where(ps ...predicate.PublicationXClassic) *PublicationXClassicUpdateOne {
	pxuo.mutation.Where(ps...)
	return pxuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pxuo *PublicationXClassicUpdateOne) Select(field string, fields ...string) *PublicationXClassicUpdateOne {
	pxuo.fields = append([]string{field}, fields...)
	return pxuo
}

// Save executes the query and returns the updated PublicationXClassic entity.
func (pxuo *PublicationXClassicUpdateOne) Save(ctx context.Context) (*PublicationXClassic, error) {
	return withHooks(ctx, pxuo.sqlSave, pxuo.mutation, pxuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pxuo *PublicationXClassicUpdateOne) SaveX(ctx context.Context) *PublicationXClassic {
	node, err := pxuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pxuo *PublicationXClassicUpdateOne) Exec(ctx context.Context) error {
	_, err := pxuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pxuo *PublicationXClassicUpdateOne) ExecX(ctx context.Context) {
	if err := pxuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pxuo *PublicationXClassicUpdateOne) sqlSave(ctx context.Context) (_node *PublicationXClassic, err error) {
	_spec := sqlgraph.NewUpdateSpec(publicationxclassic.Table, publicationxclassic.Columns, sqlgraph.NewFieldSpec(publicationxclassic.FieldID, field.TypeInt))
	id, ok := pxuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PublicationXClassic.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pxuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, publicationxclassic.FieldID)
		for _, f := range fields {
			if !publicationxclassic.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != publicationxclassic.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pxuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pxuo.mutation.ClassicID(); ok {
		_spec.SetField(publicationxclassic.FieldClassicID, field.TypeInt, value)
	}
	if value, ok := pxuo.mutation.AddedClassicID(); ok {
		_spec.AddField(publicationxclassic.FieldClassicID, field.TypeInt, value)
	}
	if value, ok := pxuo.mutation.Isbn(); ok {
		_spec.SetField(publicationxclassic.FieldIsbn, field.TypeString, value)
	}
	if value, ok := pxuo.mutation.ClassicName(); ok {
		_spec.SetField(publicationxclassic.FieldClassicName, field.TypeString, value)
	}
	if pxuo.mutation.ClassicNameCleared() {
		_spec.ClearField(publicationxclassic.FieldClassicName, field.TypeString)
	}
	if value, ok := pxuo.mutation.PublicationID(); ok {
		_spec.SetField(publicationxclassic.FieldPublicationID, field.TypeInt, value)
	}
	if value, ok := pxuo.mutation.AddedPublicationID(); ok {
		_spec.AddField(publicationxclassic.FieldPublicationID, field.TypeInt, value)
	}
	if pxuo.mutation.PublicationIDCleared() {
		_spec.ClearField(publicationxclassic.FieldPublicationID, field.TypeInt)
	}
	_node = &PublicationXClassic{config: pxuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pxuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{publicationxclassic.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pxuo.mutation.done = true
	return _node, nil
}