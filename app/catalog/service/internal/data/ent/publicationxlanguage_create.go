// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/publication-shop/app/catalog/service/internal/data/ent/publicationxlanguage"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PublicationXLanguageCreate is the builder for creating a PublicationXLanguage entity.
type PublicationXLanguageCreate struct {
	config
	mutation *PublicationXLanguageMutation
	hooks    []Hook
}

// SetLanguageID sets the "language_id" field.
func (pxc *PublicationXLanguageCreate) SetLanguageID(i int) *PublicationXLanguageCreate {
	pxc.mutation.SetLanguageID(i)
	return pxc
}

// SetIsbn sets the "isbn" field.
func (pxc *PublicationXLanguageCreate) SetIsbn(s string) *PublicationXLanguageCreate {
	pxc.mutation.SetIsbn(s)
	return pxc
}

// SetLanguageName sets the "language_name" field.
func (pxc *PublicationXLanguageCreate) SetLanguageName(s string) *PublicationXLanguageCreate {
	pxc.mutation.SetLanguageName(s)
	return pxc
}

// SetNillableLanguageName sets the "language_name" field if the given value is not nil.
func (pxc *PublicationXLanguageCreate) SetNillableLanguageName(s *string) *PublicationXLanguageCreate {
	if s != nil {
		pxc.SetLanguageName(*s)
	}
	return pxc
}

// SetPublicationID sets the "publication_id" field.
func (pxc *PublicationXLanguageCreate) SetPublicationID(i int) *PublicationXLanguageCreate {
	pxc.mutation.SetPublicationID(i)
	return pxc
}

// SetNillablePublicationID sets the "publication_id" field if the given value is not nil.
func (pxc *PublicationXLanguageCreate) SetNillablePublicationID(i *int) *PublicationXLanguageCreate {
	if i != nil {
		pxc.SetPublicationID(*i)
	}
	return pxc
}

// SetID sets the "id" field.
func (pxc *PublicationXLanguageCreate) SetID(i int) *PublicationXLanguageCreate {
	pxc.mutation.SetID(i)
	return pxc
}

// Mutation returns the PublicationXLanguageMutation object of the builder.
func (pxc *PublicationXLanguageCreate) Mutation() *PublicationXLanguageMutation {
	return pxc.mutation
}

// Save creates the PublicationXLanguage in the database.
func (pxc *PublicationXLanguageCreate) Save(ctx context.Context) (*PublicationXLanguage, error) {
	return withHooks(ctx, pxc.sqlSave, pxc.mutation, pxc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pxc *PublicationXLanguageCreate) SaveX(ctx context.Context) *PublicationXLanguage {
	v, err := pxc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pxc *PublicationXLanguageCreate) Exec(ctx context.Context) error {
	_, err := pxc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pxc *PublicationXLanguageCreate) ExecX(ctx context.Context) {
	if err := pxc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pxc *PublicationXLanguageCreate) check() error {
	if _, ok := pxc.mutation.LanguageID(); !ok {
		return &ValidationError{Name: "language_id", err: errors.New(`ent: missing required field "PublicationXLanguage.language_id"`)}
	}
	if _, ok := pxc.mutation.Isbn(); !ok {
		return &ValidationError{Name: "isbn", err: errors.New(`ent: missing required field "PublicationXLanguage.isbn"`)}
	}
	return nil
}

func (pxc *PublicationXLanguageCreate) sqlSave(ctx context.Context) (*PublicationXLanguage, error) {
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

func (pxc *PublicationXLanguageCreate) createSpec() (*PublicationXLanguage, *sqlgraph.CreateSpec) {
	var (
		_node = &PublicationXLanguage{config: pxc.config}
		_spec = sqlgraph.NewCreateSpec(publicationxlanguage.Table, sqlgraph.NewFieldSpec(publicationxlanguage.FieldID, field.TypeInt))
	)
	if id, ok := pxc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pxc.mutation.LanguageID(); ok {
		_spec.SetField(publicationxlanguage.FieldLanguageID, field.TypeInt, value)
		_node.LanguageID = value
	}
	if value, ok := pxc.mutation.Isbn(); ok {
		_spec.SetField(publicationxlanguage.FieldIsbn, field.TypeString, value)
		_node.Isbn = value
	}
	if value, ok := pxc.mutation.LanguageName(); ok {
		_spec.SetField(publicationxlanguage.FieldLanguageName, field.TypeString, value)
		_node.LanguageName = value
	}
	if value, ok := pxc.mutation.PublicationID(); ok {
		_spec.SetField(publicationxlanguage.FieldPublicationID, field.TypeInt, value)
		_node.PublicationID = value
	}
	return _node, _spec
}

// PublicationXLanguageCreateBulk is the builder for creating many PublicationXLanguage entities in bulk.
type PublicationXLanguageCreateBulk struct {
	config
	err      error
	builders []*PublicationXLanguageCreate
}

// Save creates the PublicationXLanguage entities in the database.
func (pxcb *PublicationXLanguageCreateBulk) Save(ctx context.Context) ([]*PublicationXLanguage, error) {
	if pxcb.err != nil {
		return nil, pxcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pxcb.builders))
	nodes := make([]*PublicationXLanguage, len(pxcb.builders))
	mutators := make([]Mutator, len(pxcb.builders))
	for i := range pxcb.builders {
		func(i int, root context.Context) {
			builder := pxcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PublicationXLanguageMutation)
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
func (pxcb *PublicationXLanguageCreateBulk) SaveX(ctx context.Context) []*PublicationXLanguage {
	v, err := pxcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pxcb *PublicationXLanguageCreateBulk) Exec(ctx context.Context) error {
	_, err := pxcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pxcb *PublicationXLanguageCreateBulk) ExecX(ctx context.Context) {
	if err := pxcb.Exec(ctx); err != nil {
		panic(err)
	}
}