// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/publication-shop/app/catalog/ent/publicationimg"
)

// PublicationImgCreate is the builder for creating a PublicationImg entity.
type PublicationImgCreate struct {
	config
	mutation *PublicationImgMutation
	hooks    []Hook
}

// SetImgURL sets the "img_url" field.
func (pic *PublicationImgCreate) SetImgURL(s string) *PublicationImgCreate {
	pic.mutation.SetImgURL(s)
	return pic
}

// SetNillableImgURL sets the "img_url" field if the given value is not nil.
func (pic *PublicationImgCreate) SetNillableImgURL(s *string) *PublicationImgCreate {
	if s != nil {
		pic.SetImgURL(*s)
	}
	return pic
}

// SetIsbn sets the "isbn" field.
func (pic *PublicationImgCreate) SetIsbn(i int) *PublicationImgCreate {
	pic.mutation.SetIsbn(i)
	return pic
}

// SetImgEncode sets the "img_encode" field.
func (pic *PublicationImgCreate) SetImgEncode(s string) *PublicationImgCreate {
	pic.mutation.SetImgEncode(s)
	return pic
}

// SetNillableImgEncode sets the "img_encode" field if the given value is not nil.
func (pic *PublicationImgCreate) SetNillableImgEncode(s *string) *PublicationImgCreate {
	if s != nil {
		pic.SetImgEncode(*s)
	}
	return pic
}

// SetPublicationID sets the "publication_id" field.
func (pic *PublicationImgCreate) SetPublicationID(i int) *PublicationImgCreate {
	pic.mutation.SetPublicationID(i)
	return pic
}

// SetNillablePublicationID sets the "publication_id" field if the given value is not nil.
func (pic *PublicationImgCreate) SetNillablePublicationID(i *int) *PublicationImgCreate {
	if i != nil {
		pic.SetPublicationID(*i)
	}
	return pic
}

// SetMainFlag sets the "mainFlag" field.
func (pic *PublicationImgCreate) SetMainFlag(i int32) *PublicationImgCreate {
	pic.mutation.SetMainFlag(i)
	return pic
}

// SetNillableMainFlag sets the "mainFlag" field if the given value is not nil.
func (pic *PublicationImgCreate) SetNillableMainFlag(i *int32) *PublicationImgCreate {
	if i != nil {
		pic.SetMainFlag(*i)
	}
	return pic
}

// SetID sets the "id" field.
func (pic *PublicationImgCreate) SetID(i int) *PublicationImgCreate {
	pic.mutation.SetID(i)
	return pic
}

// Mutation returns the PublicationImgMutation object of the builder.
func (pic *PublicationImgCreate) Mutation() *PublicationImgMutation {
	return pic.mutation
}

// Save creates the PublicationImg in the database.
func (pic *PublicationImgCreate) Save(ctx context.Context) (*PublicationImg, error) {
	return withHooks(ctx, pic.sqlSave, pic.mutation, pic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pic *PublicationImgCreate) SaveX(ctx context.Context) *PublicationImg {
	v, err := pic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pic *PublicationImgCreate) Exec(ctx context.Context) error {
	_, err := pic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pic *PublicationImgCreate) ExecX(ctx context.Context) {
	if err := pic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pic *PublicationImgCreate) check() error {
	if _, ok := pic.mutation.Isbn(); !ok {
		return &ValidationError{Name: "isbn", err: errors.New(`ent: missing required field "PublicationImg.isbn"`)}
	}
	return nil
}

func (pic *PublicationImgCreate) sqlSave(ctx context.Context) (*PublicationImg, error) {
	if err := pic.check(); err != nil {
		return nil, err
	}
	_node, _spec := pic.createSpec()
	if err := sqlgraph.CreateNode(ctx, pic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	pic.mutation.id = &_node.ID
	pic.mutation.done = true
	return _node, nil
}

func (pic *PublicationImgCreate) createSpec() (*PublicationImg, *sqlgraph.CreateSpec) {
	var (
		_node = &PublicationImg{config: pic.config}
		_spec = sqlgraph.NewCreateSpec(publicationimg.Table, sqlgraph.NewFieldSpec(publicationimg.FieldID, field.TypeInt))
	)
	if id, ok := pic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pic.mutation.ImgURL(); ok {
		_spec.SetField(publicationimg.FieldImgURL, field.TypeString, value)
		_node.ImgURL = value
	}
	if value, ok := pic.mutation.Isbn(); ok {
		_spec.SetField(publicationimg.FieldIsbn, field.TypeInt, value)
		_node.Isbn = value
	}
	if value, ok := pic.mutation.ImgEncode(); ok {
		_spec.SetField(publicationimg.FieldImgEncode, field.TypeString, value)
		_node.ImgEncode = value
	}
	if value, ok := pic.mutation.PublicationID(); ok {
		_spec.SetField(publicationimg.FieldPublicationID, field.TypeInt, value)
		_node.PublicationID = value
	}
	if value, ok := pic.mutation.MainFlag(); ok {
		_spec.SetField(publicationimg.FieldMainFlag, field.TypeInt32, value)
		_node.MainFlag = value
	}
	return _node, _spec
}

// PublicationImgCreateBulk is the builder for creating many PublicationImg entities in bulk.
type PublicationImgCreateBulk struct {
	config
	err      error
	builders []*PublicationImgCreate
}

// Save creates the PublicationImg entities in the database.
func (picb *PublicationImgCreateBulk) Save(ctx context.Context) ([]*PublicationImg, error) {
	if picb.err != nil {
		return nil, picb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(picb.builders))
	nodes := make([]*PublicationImg, len(picb.builders))
	mutators := make([]Mutator, len(picb.builders))
	for i := range picb.builders {
		func(i int, root context.Context) {
			builder := picb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PublicationImgMutation)
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
					_, err = mutators[i+1].Mutate(root, picb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, picb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, picb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (picb *PublicationImgCreateBulk) SaveX(ctx context.Context) []*PublicationImg {
	v, err := picb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (picb *PublicationImgCreateBulk) Exec(ctx context.Context) error {
	_, err := picb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (picb *PublicationImgCreateBulk) ExecX(ctx context.Context) {
	if err := picb.Exec(ctx); err != nil {
		panic(err)
	}
}
