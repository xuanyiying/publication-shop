// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/xuanyiying/publication-shop/app/user/service/internal/data/ent/address"
	"github.com/xuanyiying/publication-shop/app/user/service/internal/data/ent/user"
)

// AddressCreate is the builder for creating a Address entity.
type AddressCreate struct {
	config
	mutation *AddressMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ac *AddressCreate) SetName(s string) *AddressCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetMobile sets the "mobile" field.
func (ac *AddressCreate) SetMobile(s string) *AddressCreate {
	ac.mutation.SetMobile(s)
	return ac
}

// SetAddress sets the "address" field.
func (ac *AddressCreate) SetAddress(s string) *AddressCreate {
	ac.mutation.SetAddress(s)
	return ac
}

// SetPostCode sets the "post_code" field.
func (ac *AddressCreate) SetPostCode(s string) *AddressCreate {
	ac.mutation.SetPostCode(s)
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *AddressCreate) SetCreatedAt(t time.Time) *AddressCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AddressCreate) SetNillableCreatedAt(t *time.Time) *AddressCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AddressCreate) SetUpdatedAt(t time.Time) *AddressCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AddressCreate) SetNillableUpdatedAt(t *time.Time) *AddressCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AddressCreate) SetID(i int64) *AddressCreate {
	ac.mutation.SetID(i)
	return ac
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ac *AddressCreate) SetUserID(id int64) *AddressCreate {
	ac.mutation.SetUserID(id)
	return ac
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (ac *AddressCreate) SetNillableUserID(id *int64) *AddressCreate {
	if id != nil {
		ac = ac.SetUserID(*id)
	}
	return ac
}

// SetUser sets the "user" edge to the User entity.
func (ac *AddressCreate) SetUser(u *User) *AddressCreate {
	return ac.SetUserID(u.ID)
}

// Mutation returns the AddressMutation object of the builder.
func (ac *AddressCreate) Mutation() *AddressMutation {
	return ac.mutation
}

// Save creates the Address in the database.
func (ac *AddressCreate) Save(ctx context.Context) (*Address, error) {
	var (
		err  error
		node *Address
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AddressMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			node, err = ac.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AddressCreate) SaveX(ctx context.Context) *Address {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (ac *AddressCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := address.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := address.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AddressCreate) check() error {
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := ac.mutation.Mobile(); !ok {
		return &ValidationError{Name: "mobile", err: errors.New("ent: missing required field \"mobile\"")}
	}
	if _, ok := ac.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New("ent: missing required field \"address\"")}
	}
	if _, ok := ac.mutation.PostCode(); !ok {
		return &ValidationError{Name: "post_code", err: errors.New("ent: missing required field \"post_code\"")}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	return nil
}

func (ac *AddressCreate) sqlSave(ctx context.Context) (*Address, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	if _node.ID == 0 {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (ac *AddressCreate) createSpec() (*Address, *sqlgraph.CreateSpec) {
	var (
		_node = &Address{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: address.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: address.FieldID,
			},
		}
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: address.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ac.mutation.Mobile(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: address.FieldMobile,
		})
		_node.Mobile = value
	}
	if value, ok := ac.mutation.Address(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: address.FieldAddress,
		})
		_node.Address = value
	}
	if value, ok := ac.mutation.PostCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: address.FieldPostCode,
		})
		_node.PostCode = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: address.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: address.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := ac.mutation.UserIDs(); len(nodes) > 0 {
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
		_node.user_addresses = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AddressCreateBulk is the builder for creating many Address entities in bulk.
type AddressCreateBulk struct {
	config
	builders []*AddressCreate
}

// Save creates the Address entities in the database.
func (acb *AddressCreateBulk) Save(ctx context.Context) ([]*Address, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Address, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AddressMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				if nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AddressCreateBulk) SaveX(ctx context.Context) []*Address {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
