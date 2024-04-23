// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/publication-shop/app/catalog/service/internal/data/ent/publicationauthor"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PublicationAuthorCreate is the builder for creating a PublicationAuthor entity.
type PublicationAuthorCreate struct {
	config
	mutation *PublicationAuthorMutation
	hooks    []Hook
}

// SetIsbn sets the "isbn" field.
func (pac *PublicationAuthorCreate) SetIsbn(s string) *PublicationAuthorCreate {
	pac.mutation.SetIsbn(s)
	return pac
}

// SetPublicationID sets the "publication_id" field.
func (pac *PublicationAuthorCreate) SetPublicationID(i int) *PublicationAuthorCreate {
	pac.mutation.SetPublicationID(i)
	return pac
}

// SetNillablePublicationID sets the "publication_id" field if the given value is not nil.
func (pac *PublicationAuthorCreate) SetNillablePublicationID(i *int) *PublicationAuthorCreate {
	if i != nil {
		pac.SetPublicationID(*i)
	}
	return pac
}

// SetAuthor sets the "author" field.
func (pac *PublicationAuthorCreate) SetAuthor(s string) *PublicationAuthorCreate {
	pac.mutation.SetAuthor(s)
	return pac
}

// SetNillableAuthor sets the "author" field if the given value is not nil.
func (pac *PublicationAuthorCreate) SetNillableAuthor(s *string) *PublicationAuthorCreate {
	if s != nil {
		pac.SetAuthor(*s)
	}
	return pac
}

// SetAuthorAbout sets the "author_about" field.
func (pac *PublicationAuthorCreate) SetAuthorAbout(s string) *PublicationAuthorCreate {
	pac.mutation.SetAuthorAbout(s)
	return pac
}

// SetNillableAuthorAbout sets the "author_about" field if the given value is not nil.
func (pac *PublicationAuthorCreate) SetNillableAuthorAbout(s *string) *PublicationAuthorCreate {
	if s != nil {
		pac.SetAuthorAbout(*s)
	}
	return pac
}

// SetID sets the "id" field.
func (pac *PublicationAuthorCreate) SetID(i int) *PublicationAuthorCreate {
	pac.mutation.SetID(i)
	return pac
}

// Mutation returns the PublicationAuthorMutation object of the builder.
func (pac *PublicationAuthorCreate) Mutation() *PublicationAuthorMutation {
	return pac.mutation
}

// Save creates the PublicationAuthor in the database.
func (pac *PublicationAuthorCreate) Save(ctx context.Context) (*PublicationAuthor, error) {
	return withHooks(ctx, pac.sqlSave, pac.mutation, pac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pac *PublicationAuthorCreate) SaveX(ctx context.Context) *PublicationAuthor {
	v, err := pac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pac *PublicationAuthorCreate) Exec(ctx context.Context) error {
	_, err := pac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pac *PublicationAuthorCreate) ExecX(ctx context.Context) {
	if err := pac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pac *PublicationAuthorCreate) check() error {
	if _, ok := pac.mutation.Isbn(); !ok {
		return &ValidationError{Name: "isbn", err: errors.New(`ent: missing required field "PublicationAuthor.isbn"`)}
	}
	return nil
}

func (pac *PublicationAuthorCreate) sqlSave(ctx context.Context) (*PublicationAuthor, error) {
	if err := pac.check(); err != nil {
		return nil, err
	}
	_node, _spec := pac.createSpec()
	if err := sqlgraph.CreateNode(ctx, pac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	pac.mutation.id = &_node.ID
	pac.mutation.done = true
	return _node, nil
}

func (pac *PublicationAuthorCreate) createSpec() (*PublicationAuthor, *sqlgraph.CreateSpec) {
	var (
		_node = &PublicationAuthor{config: pac.config}
		_spec = sqlgraph.NewCreateSpec(publicationauthor.Table, sqlgraph.NewFieldSpec(publicationauthor.FieldID, field.TypeInt))
	)
	if id, ok := pac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pac.mutation.Isbn(); ok {
		_spec.SetField(publicationauthor.FieldIsbn, field.TypeString, value)
		_node.Isbn = value
	}
	if value, ok := pac.mutation.PublicationID(); ok {
		_spec.SetField(publicationauthor.FieldPublicationID, field.TypeInt, value)
		_node.PublicationID = value
	}
	if value, ok := pac.mutation.Author(); ok {
		_spec.SetField(publicationauthor.FieldAuthor, field.TypeString, value)
		_node.Author = value
	}
	if value, ok := pac.mutation.AuthorAbout(); ok {
		_spec.SetField(publicationauthor.FieldAuthorAbout, field.TypeString, value)
		_node.AuthorAbout = value
	}
	return _node, _spec
}

// PublicationAuthorCreateBulk is the builder for creating many PublicationAuthor entities in bulk.
type PublicationAuthorCreateBulk struct {
	config
	err      error
	builders []*PublicationAuthorCreate
}

// Save creates the PublicationAuthor entities in the database.
func (pacb *PublicationAuthorCreateBulk) Save(ctx context.Context) ([]*PublicationAuthor, error) {
	if pacb.err != nil {
		return nil, pacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pacb.builders))
	nodes := make([]*PublicationAuthor, len(pacb.builders))
	mutators := make([]Mutator, len(pacb.builders))
	for i := range pacb.builders {
		func(i int, root context.Context) {
			builder := pacb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PublicationAuthorMutation)
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
					_, err = mutators[i+1].Mutate(root, pacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pacb *PublicationAuthorCreateBulk) SaveX(ctx context.Context) []*PublicationAuthor {
	v, err := pacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pacb *PublicationAuthorCreateBulk) Exec(ctx context.Context) error {
	_, err := pacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pacb *PublicationAuthorCreateBulk) ExecX(ctx context.Context) {
	if err := pacb.Exec(ctx); err != nil {
		panic(err)
	}
}