// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/xuanyiying/publication-shop/app/catalog/ent/predicate"
	"github.com/xuanyiying/publication-shop/app/catalog/ent/publicationxclassic"
)

// PublicationXClassicDelete is the builder for deleting a PublicationXClassic entity.
type PublicationXClassicDelete struct {
	config
	hooks    []Hook
	mutation *PublicationXClassicMutation
}

// Where appends a list predicates to the PublicationXClassicDelete builder.
func (pxd *PublicationXClassicDelete) Where(ps ...predicate.PublicationXClassic) *PublicationXClassicDelete {
	pxd.mutation.Where(ps...)
	return pxd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pxd *PublicationXClassicDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pxd.sqlExec, pxd.mutation, pxd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pxd *PublicationXClassicDelete) ExecX(ctx context.Context) int {
	n, err := pxd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pxd *PublicationXClassicDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(publicationxclassic.Table, sqlgraph.NewFieldSpec(publicationxclassic.FieldID, field.TypeInt))
	if ps := pxd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pxd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pxd.mutation.done = true
	return affected, err
}

// PublicationXClassicDeleteOne is the builder for deleting a single PublicationXClassic entity.
type PublicationXClassicDeleteOne struct {
	pxd *PublicationXClassicDelete
}

// Where appends a list predicates to the PublicationXClassicDelete builder.
func (pxdo *PublicationXClassicDeleteOne) Where(ps ...predicate.PublicationXClassic) *PublicationXClassicDeleteOne {
	pxdo.pxd.mutation.Where(ps...)
	return pxdo
}

// Exec executes the deletion query.
func (pxdo *PublicationXClassicDeleteOne) Exec(ctx context.Context) error {
	n, err := pxdo.pxd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{publicationxclassic.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pxdo *PublicationXClassicDeleteOne) ExecX(ctx context.Context) {
	if err := pxdo.Exec(ctx); err != nil {
		panic(err)
	}
}