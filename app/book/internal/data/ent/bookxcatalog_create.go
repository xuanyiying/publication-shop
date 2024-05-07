// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/xuanyiying/publication-shop/app/book/internal/data/ent/bookxcatalog"
)

// BookXCatalogCreate is the builder for creating a BookXCatalog entity.
type BookXCatalogCreate struct {
	config
	mutation *BookXCatalogMutation
	hooks    []Hook
}

// SetCatalogID sets the "catalog_id" field.
func (bxc *BookXCatalogCreate) SetCatalogID(i int) *BookXCatalogCreate {
	bxc.mutation.SetCatalogID(i)
	return bxc
}

// SetIsbn sets the "isbn" field.
func (bxc *BookXCatalogCreate) SetIsbn(s string) *BookXCatalogCreate {
	bxc.mutation.SetIsbn(s)
	return bxc
}

// SetCatalogName sets the "catalog_name" field.
func (bxc *BookXCatalogCreate) SetCatalogName(s string) *BookXCatalogCreate {
	bxc.mutation.SetCatalogName(s)
	return bxc
}

// SetNillableCatalogName sets the "catalog_name" field if the given value is not nil.
func (bxc *BookXCatalogCreate) SetNillableCatalogName(s *string) *BookXCatalogCreate {
	if s != nil {
		bxc.SetCatalogName(*s)
	}
	return bxc
}

// SetBookID sets the "book_id" field.
func (bxc *BookXCatalogCreate) SetBookID(i int) *BookXCatalogCreate {
	bxc.mutation.SetBookID(i)
	return bxc
}

// SetNillableBookID sets the "book_id" field if the given value is not nil.
func (bxc *BookXCatalogCreate) SetNillableBookID(i *int) *BookXCatalogCreate {
	if i != nil {
		bxc.SetBookID(*i)
	}
	return bxc
}

// SetID sets the "id" field.
func (bxc *BookXCatalogCreate) SetID(i int) *BookXCatalogCreate {
	bxc.mutation.SetID(i)
	return bxc
}

// Mutation returns the BookXCatalogMutation object of the builder.
func (bxc *BookXCatalogCreate) Mutation() *BookXCatalogMutation {
	return bxc.mutation
}

// Save creates the BookXCatalog in the database.
func (bxc *BookXCatalogCreate) Save(ctx context.Context) (*BookXCatalog, error) {
	return withHooks(ctx, bxc.sqlSave, bxc.mutation, bxc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bxc *BookXCatalogCreate) SaveX(ctx context.Context) *BookXCatalog {
	v, err := bxc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bxc *BookXCatalogCreate) Exec(ctx context.Context) error {
	_, err := bxc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bxc *BookXCatalogCreate) ExecX(ctx context.Context) {
	if err := bxc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bxc *BookXCatalogCreate) check() error {
	if _, ok := bxc.mutation.CatalogID(); !ok {
		return &ValidationError{Name: "catalog_id", err: errors.New(`ent: missing required field "BookXCatalog.catalog_id"`)}
	}
	if _, ok := bxc.mutation.Isbn(); !ok {
		return &ValidationError{Name: "isbn", err: errors.New(`ent: missing required field "BookXCatalog.isbn"`)}
	}
	return nil
}

func (bxc *BookXCatalogCreate) sqlSave(ctx context.Context) (*BookXCatalog, error) {
	if err := bxc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bxc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bxc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	bxc.mutation.id = &_node.ID
	bxc.mutation.done = true
	return _node, nil
}

func (bxc *BookXCatalogCreate) createSpec() (*BookXCatalog, *sqlgraph.CreateSpec) {
	var (
		_node = &BookXCatalog{config: bxc.config}
		_spec = sqlgraph.NewCreateSpec(bookxcatalog.Table, sqlgraph.NewFieldSpec(bookxcatalog.FieldID, field.TypeInt))
	)
	if id, ok := bxc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := bxc.mutation.CatalogID(); ok {
		_spec.SetField(bookxcatalog.FieldCatalogID, field.TypeInt, value)
		_node.CatalogID = value
	}
	if value, ok := bxc.mutation.Isbn(); ok {
		_spec.SetField(bookxcatalog.FieldIsbn, field.TypeString, value)
		_node.Isbn = value
	}
	if value, ok := bxc.mutation.CatalogName(); ok {
		_spec.SetField(bookxcatalog.FieldCatalogName, field.TypeString, value)
		_node.CatalogName = value
	}
	if value, ok := bxc.mutation.BookID(); ok {
		_spec.SetField(bookxcatalog.FieldBookID, field.TypeInt, value)
		_node.BookID = value
	}
	return _node, _spec
}

// BookXCatalogCreateBulk is the builder for creating many BookXCatalog entities in bulk.
type BookXCatalogCreateBulk struct {
	config
	err      error
	builders []*BookXCatalogCreate
}

// Save creates the BookXCatalog entities in the database.
func (bxcb *BookXCatalogCreateBulk) Save(ctx context.Context) ([]*BookXCatalog, error) {
	if bxcb.err != nil {
		return nil, bxcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bxcb.builders))
	nodes := make([]*BookXCatalog, len(bxcb.builders))
	mutators := make([]Mutator, len(bxcb.builders))
	for i := range bxcb.builders {
		func(i int, root context.Context) {
			builder := bxcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BookXCatalogMutation)
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
					_, err = mutators[i+1].Mutate(root, bxcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bxcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bxcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bxcb *BookXCatalogCreateBulk) SaveX(ctx context.Context) []*BookXCatalog {
	v, err := bxcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bxcb *BookXCatalogCreateBulk) Exec(ctx context.Context) error {
	_, err := bxcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bxcb *BookXCatalogCreateBulk) ExecX(ctx context.Context) {
	if err := bxcb.Exec(ctx); err != nil {
		panic(err)
	}
}