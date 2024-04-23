// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/publication-shop/app/catalog/service/internal/data/ent/publicationdetail"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PublicationDetailCreate is the builder for creating a PublicationDetail entity.
type PublicationDetailCreate struct {
	config
	mutation *PublicationDetailMutation
	hooks    []Hook
}

// SetIsbn sets the "isbn" field.
func (pdc *PublicationDetailCreate) SetIsbn(s string) *PublicationDetailCreate {
	pdc.mutation.SetIsbn(s)
	return pdc
}

// SetDetailHTML sets the "detail_html" field.
func (pdc *PublicationDetailCreate) SetDetailHTML(s string) *PublicationDetailCreate {
	pdc.mutation.SetDetailHTML(s)
	return pdc
}

// SetNillableDetailHTML sets the "detail_html" field if the given value is not nil.
func (pdc *PublicationDetailCreate) SetNillableDetailHTML(s *string) *PublicationDetailCreate {
	if s != nil {
		pdc.SetDetailHTML(*s)
	}
	return pdc
}

// SetDetailImg sets the "detail_img" field.
func (pdc *PublicationDetailCreate) SetDetailImg(s string) *PublicationDetailCreate {
	pdc.mutation.SetDetailImg(s)
	return pdc
}

// SetNillableDetailImg sets the "detail_img" field if the given value is not nil.
func (pdc *PublicationDetailCreate) SetNillableDetailImg(s *string) *PublicationDetailCreate {
	if s != nil {
		pdc.SetDetailImg(*s)
	}
	return pdc
}

// SetPublicationID sets the "publication_id" field.
func (pdc *PublicationDetailCreate) SetPublicationID(i int) *PublicationDetailCreate {
	pdc.mutation.SetPublicationID(i)
	return pdc
}

// SetNillablePublicationID sets the "publication_id" field if the given value is not nil.
func (pdc *PublicationDetailCreate) SetNillablePublicationID(i *int) *PublicationDetailCreate {
	if i != nil {
		pdc.SetPublicationID(*i)
	}
	return pdc
}

// SetID sets the "id" field.
func (pdc *PublicationDetailCreate) SetID(i int) *PublicationDetailCreate {
	pdc.mutation.SetID(i)
	return pdc
}

// Mutation returns the PublicationDetailMutation object of the builder.
func (pdc *PublicationDetailCreate) Mutation() *PublicationDetailMutation {
	return pdc.mutation
}

// Save creates the PublicationDetail in the database.
func (pdc *PublicationDetailCreate) Save(ctx context.Context) (*PublicationDetail, error) {
	return withHooks(ctx, pdc.sqlSave, pdc.mutation, pdc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pdc *PublicationDetailCreate) SaveX(ctx context.Context) *PublicationDetail {
	v, err := pdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pdc *PublicationDetailCreate) Exec(ctx context.Context) error {
	_, err := pdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pdc *PublicationDetailCreate) ExecX(ctx context.Context) {
	if err := pdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pdc *PublicationDetailCreate) check() error {
	if _, ok := pdc.mutation.Isbn(); !ok {
		return &ValidationError{Name: "isbn", err: errors.New(`ent: missing required field "PublicationDetail.isbn"`)}
	}
	return nil
}

func (pdc *PublicationDetailCreate) sqlSave(ctx context.Context) (*PublicationDetail, error) {
	if err := pdc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	pdc.mutation.id = &_node.ID
	pdc.mutation.done = true
	return _node, nil
}

func (pdc *PublicationDetailCreate) createSpec() (*PublicationDetail, *sqlgraph.CreateSpec) {
	var (
		_node = &PublicationDetail{config: pdc.config}
		_spec = sqlgraph.NewCreateSpec(publicationdetail.Table, sqlgraph.NewFieldSpec(publicationdetail.FieldID, field.TypeInt))
	)
	if id, ok := pdc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pdc.mutation.Isbn(); ok {
		_spec.SetField(publicationdetail.FieldIsbn, field.TypeString, value)
		_node.Isbn = value
	}
	if value, ok := pdc.mutation.DetailHTML(); ok {
		_spec.SetField(publicationdetail.FieldDetailHTML, field.TypeString, value)
		_node.DetailHTML = value
	}
	if value, ok := pdc.mutation.DetailImg(); ok {
		_spec.SetField(publicationdetail.FieldDetailImg, field.TypeString, value)
		_node.DetailImg = value
	}
	if value, ok := pdc.mutation.PublicationID(); ok {
		_spec.SetField(publicationdetail.FieldPublicationID, field.TypeInt, value)
		_node.PublicationID = value
	}
	return _node, _spec
}

// PublicationDetailCreateBulk is the builder for creating many PublicationDetail entities in bulk.
type PublicationDetailCreateBulk struct {
	config
	err      error
	builders []*PublicationDetailCreate
}

// Save creates the PublicationDetail entities in the database.
func (pdcb *PublicationDetailCreateBulk) Save(ctx context.Context) ([]*PublicationDetail, error) {
	if pdcb.err != nil {
		return nil, pdcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pdcb.builders))
	nodes := make([]*PublicationDetail, len(pdcb.builders))
	mutators := make([]Mutator, len(pdcb.builders))
	for i := range pdcb.builders {
		func(i int, root context.Context) {
			builder := pdcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PublicationDetailMutation)
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
					_, err = mutators[i+1].Mutate(root, pdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pdcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pdcb *PublicationDetailCreateBulk) SaveX(ctx context.Context) []*PublicationDetail {
	v, err := pdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pdcb *PublicationDetailCreateBulk) Exec(ctx context.Context) error {
	_, err := pdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pdcb *PublicationDetailCreateBulk) ExecX(ctx context.Context) {
	if err := pdcb.Exec(ctx); err != nil {
		panic(err)
	}
}