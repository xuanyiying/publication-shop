// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/publication-shop/app/catalog/service/internal/data/ent/publicationxclassic"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PublicationXClassicCreate is the builder for creating a PublicationXClassic entity.
type PublicationXClassicCreate struct {
	config
	mutation *PublicationXClassicMutation
	hooks    []Hook
}

// SetClassicID sets the "classic_id" field.
func (pxc *PublicationXClassicCreate) SetClassicID(i int) *PublicationXClassicCreate {
	pxc.mutation.SetClassicID(i)
	return pxc
}

// SetIsbn sets the "isbn" field.
func (pxc *PublicationXClassicCreate) SetIsbn(s string) *PublicationXClassicCreate {
	pxc.mutation.SetIsbn(s)
	return pxc
}

// SetClassicName sets the "classic_name" field.
func (pxc *PublicationXClassicCreate) SetClassicName(s string) *PublicationXClassicCreate {
	pxc.mutation.SetClassicName(s)
	return pxc
}

// SetNillableClassicName sets the "classic_name" field if the given value is not nil.
func (pxc *PublicationXClassicCreate) SetNillableClassicName(s *string) *PublicationXClassicCreate {
	if s != nil {
		pxc.SetClassicName(*s)
	}
	return pxc
}

// SetPublicationID sets the "publication_id" field.
func (pxc *PublicationXClassicCreate) SetPublicationID(i int) *PublicationXClassicCreate {
	pxc.mutation.SetPublicationID(i)
	return pxc
}

// SetNillablePublicationID sets the "publication_id" field if the given value is not nil.
func (pxc *PublicationXClassicCreate) SetNillablePublicationID(i *int) *PublicationXClassicCreate {
	if i != nil {
		pxc.SetPublicationID(*i)
	}
	return pxc
}

// SetID sets the "id" field.
func (pxc *PublicationXClassicCreate) SetID(i int) *PublicationXClassicCreate {
	pxc.mutation.SetID(i)
	return pxc
}

// Mutation returns the PublicationXClassicMutation object of the builder.
func (pxc *PublicationXClassicCreate) Mutation() *PublicationXClassicMutation {
	return pxc.mutation
}

// Save creates the PublicationXClassic in the database.
func (pxc *PublicationXClassicCreate) Save(ctx context.Context) (*PublicationXClassic, error) {
	return withHooks(ctx, pxc.sqlSave, pxc.mutation, pxc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pxc *PublicationXClassicCreate) SaveX(ctx context.Context) *PublicationXClassic {
	v, err := pxc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pxc *PublicationXClassicCreate) Exec(ctx context.Context) error {
	_, err := pxc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pxc *PublicationXClassicCreate) ExecX(ctx context.Context) {
	if err := pxc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pxc *PublicationXClassicCreate) check() error {
	if _, ok := pxc.mutation.ClassicID(); !ok {
		return &ValidationError{Name: "classic_id", err: errors.New(`ent: missing required field "PublicationXClassic.classic_id"`)}
	}
	if _, ok := pxc.mutation.Isbn(); !ok {
		return &ValidationError{Name: "isbn", err: errors.New(`ent: missing required field "PublicationXClassic.isbn"`)}
	}
	return nil
}

func (pxc *PublicationXClassicCreate) sqlSave(ctx context.Context) (*PublicationXClassic, error) {
	if err := pxc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pxc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pxc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	pxc.mutation.id = &_node.ID
	pxc.mutation.done = true
	return _node, nil
}

func (pxc *PublicationXClassicCreate) createSpec() (*PublicationXClassic, *sqlgraph.CreateSpec) {
	var (
		_node = &PublicationXClassic{config: pxc.config}
		_spec = sqlgraph.NewCreateSpec(publicationxclassic.Table, sqlgraph.NewFieldSpec(publicationxclassic.FieldID, field.TypeInt))
	)
	if id, ok := pxc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pxc.mutation.ClassicID(); ok {
		_spec.SetField(publicationxclassic.FieldClassicID, field.TypeInt, value)
		_node.ClassicID = value
	}
	if value, ok := pxc.mutation.Isbn(); ok {
		_spec.SetField(publicationxclassic.FieldIsbn, field.TypeString, value)
		_node.Isbn = value
	}
	if value, ok := pxc.mutation.ClassicName(); ok {
		_spec.SetField(publicationxclassic.FieldClassicName, field.TypeString, value)
		_node.ClassicName = value
	}
	if value, ok := pxc.mutation.PublicationID(); ok {
		_spec.SetField(publicationxclassic.FieldPublicationID, field.TypeInt, value)
		_node.PublicationID = value
	}
	return _node, _spec
}

// PublicationXClassicCreateBulk is the builder for creating many PublicationXClassic entities in bulk.
type PublicationXClassicCreateBulk struct {
	config
	err      error
	builders []*PublicationXClassicCreate
}

// Save creates the PublicationXClassic entities in the database.
func (pxcb *PublicationXClassicCreateBulk) Save(ctx context.Context) ([]*PublicationXClassic, error) {
	if pxcb.err != nil {
		return nil, pxcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pxcb.builders))
	nodes := make([]*PublicationXClassic, len(pxcb.builders))
	mutators := make([]Mutator, len(pxcb.builders))
	for i := range pxcb.builders {
		func(i int, root context.Context) {
			builder := pxcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PublicationXClassicMutation)
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
					_, err = mutators[i+1].Mutate(root, pxcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pxcb.driver, spec); err != nil {
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
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, pxcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pxcb *PublicationXClassicCreateBulk) SaveX(ctx context.Context) []*PublicationXClassic {
	v, err := pxcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pxcb *PublicationXClassicCreateBulk) Exec(ctx context.Context) error {
	_, err := pxcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pxcb *PublicationXClassicCreateBulk) ExecX(ctx context.Context) {
	if err := pxcb.Exec(ctx); err != nil {
		panic(err)
	}
}
