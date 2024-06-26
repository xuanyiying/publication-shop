// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/xuanyiying/publication-shop/app/user/service/internal/data/ent/address"
	"github.com/xuanyiying/publication-shop/app/user/service/internal/data/ent/predicate"
	"github.com/xuanyiying/publication-shop/app/user/service/internal/data/ent/user"
)

// AddressUpdate is the builder for updating Address entities.
type AddressUpdate struct {
	config
	hooks    []Hook
	mutation *AddressMutation
}

// Where adds a new predicate for the AddressUpdate builder.
func (au *AddressUpdate) Where(ps ...predicate.Address) *AddressUpdate {
	au.mutation.predicates = append(au.mutation.predicates, ps...)
	return au
}

// SetName sets the "name" field.
func (au *AddressUpdate) SetName(s string) *AddressUpdate {
	au.mutation.SetName(s)
	return au
}

// SetMobile sets the "mobile" field.
func (au *AddressUpdate) SetMobile(s string) *AddressUpdate {
	au.mutation.SetMobile(s)
	return au
}

// SetAddress sets the "address" field.
func (au *AddressUpdate) SetAddress(s string) *AddressUpdate {
	au.mutation.SetAddress(s)
	return au
}

// SetPostCode sets the "post_code" field.
func (au *AddressUpdate) SetPostCode(s string) *AddressUpdate {
	au.mutation.SetPostCode(s)
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *AddressUpdate) SetCreatedAt(t time.Time) *AddressUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *AddressUpdate) SetNillableCreatedAt(t *time.Time) *AddressUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AddressUpdate) SetUpdatedAt(t time.Time) *AddressUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (au *AddressUpdate) SetNillableUpdatedAt(t *time.Time) *AddressUpdate {
	if t != nil {
		au.SetUpdatedAt(*t)
	}
	return au
}

// SetUserID sets the "user" edge to the User entity by ID.
func (au *AddressUpdate) SetUserID(id int64) *AddressUpdate {
	au.mutation.SetUserID(id)
	return au
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (au *AddressUpdate) SetNillableUserID(id *int64) *AddressUpdate {
	if id != nil {
		au = au.SetUserID(*id)
	}
	return au
}

// SetUser sets the "user" edge to the User entity.
func (au *AddressUpdate) SetUser(u *User) *AddressUpdate {
	return au.SetUserID(u.ID)
}

// Mutation returns the AddressMutation object of the builder.
func (au *AddressUpdate) Mutation() *AddressMutation {
	return au.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (au *AddressUpdate) ClearUser() *AddressUpdate {
	au.mutation.ClearUser()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AddressUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AddressMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AddressUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AddressUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AddressUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AddressUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   address.Table,
			Columns: address.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: address.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: address.FieldName,
		})
	}
	if value, ok := au.mutation.Mobile(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: address.FieldMobile,
		})
	}
	if value, ok := au.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: address.FieldAddress,
		})
	}
	if value, ok := au.mutation.PostCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: address.FieldPostCode,
		})
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: address.FieldCreatedAt,
		})
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: address.FieldUpdatedAt,
		})
	}
	if au.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   address.UserTable,
			Columns: []string{address.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   address.UserTable,
			Columns: []string{address.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{address.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AddressUpdateOne is the builder for updating a single Address entity.
type AddressUpdateOne struct {
	config
	hooks    []Hook
	mutation *AddressMutation
}

// SetName sets the "name" field.
func (auo *AddressUpdateOne) SetName(s string) *AddressUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetMobile sets the "mobile" field.
func (auo *AddressUpdateOne) SetMobile(s string) *AddressUpdateOne {
	auo.mutation.SetMobile(s)
	return auo
}

// SetAddress sets the "address" field.
func (auo *AddressUpdateOne) SetAddress(s string) *AddressUpdateOne {
	auo.mutation.SetAddress(s)
	return auo
}

// SetPostCode sets the "post_code" field.
func (auo *AddressUpdateOne) SetPostCode(s string) *AddressUpdateOne {
	auo.mutation.SetPostCode(s)
	return auo
}

// SetCreatedAt sets the "created_at" field.
func (auo *AddressUpdateOne) SetCreatedAt(t time.Time) *AddressUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *AddressUpdateOne) SetNillableCreatedAt(t *time.Time) *AddressUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AddressUpdateOne) SetUpdatedAt(t time.Time) *AddressUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (auo *AddressUpdateOne) SetNillableUpdatedAt(t *time.Time) *AddressUpdateOne {
	if t != nil {
		auo.SetUpdatedAt(*t)
	}
	return auo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (auo *AddressUpdateOne) SetUserID(id int64) *AddressUpdateOne {
	auo.mutation.SetUserID(id)
	return auo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (auo *AddressUpdateOne) SetNillableUserID(id *int64) *AddressUpdateOne {
	if id != nil {
		auo = auo.SetUserID(*id)
	}
	return auo
}

// SetUser sets the "user" edge to the User entity.
func (auo *AddressUpdateOne) SetUser(u *User) *AddressUpdateOne {
	return auo.SetUserID(u.ID)
}

// Mutation returns the AddressMutation object of the builder.
func (auo *AddressUpdateOne) Mutation() *AddressMutation {
	return auo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (auo *AddressUpdateOne) ClearUser() *AddressUpdateOne {
	auo.mutation.ClearUser()
	return auo
}

// Save executes the query and returns the updated Address entity.
func (auo *AddressUpdateOne) Save(ctx context.Context) (*Address, error) {
	var (
		err  error
		node *Address
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AddressMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AddressUpdateOne) SaveX(ctx context.Context) *Address {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AddressUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AddressUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AddressUpdateOne) sqlSave(ctx context.Context) (_node *Address, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   address.Table,
			Columns: address.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: address.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Address.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: address.FieldName,
		})
	}
	if value, ok := auo.mutation.Mobile(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: address.FieldMobile,
		})
	}
	if value, ok := auo.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: address.FieldAddress,
		})
	}
	if value, ok := auo.mutation.PostCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: address.FieldPostCode,
		})
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: address.FieldCreatedAt,
		})
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: address.FieldUpdatedAt,
		})
	}
	if auo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   address.UserTable,
			Columns: []string{address.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   address.UserTable,
			Columns: []string{address.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Address{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{address.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
