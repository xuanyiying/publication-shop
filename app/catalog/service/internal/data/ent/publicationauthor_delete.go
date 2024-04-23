// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"github.com/go-kratos/publication-shop/app/catalog/service/internal/data/ent/predicate"
	"github.com/go-kratos/publication-shop/app/catalog/service/internal/data/ent/publicationauthor"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PublicationAuthorDelete is the builder for deleting a PublicationAuthor entity.
type PublicationAuthorDelete struct {
	config
	hooks    []Hook
	mutation *PublicationAuthorMutation
}

// Where appends a list predicates to the PublicationAuthorDelete builder.
func (pad *PublicationAuthorDelete) Where(ps ...predicate.PublicationAuthor) *PublicationAuthorDelete {
	pad.mutation.Where(ps...)
	return pad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pad *PublicationAuthorDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pad.sqlExec, pad.mutation, pad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pad *PublicationAuthorDelete) ExecX(ctx context.Context) int {
	n, err := pad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pad *PublicationAuthorDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(publicationauthor.Table, sqlgraph.NewFieldSpec(publicationauthor.FieldID, field.TypeInt))
	if ps := pad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pad.mutation.done = true
	return affected, err
}

// PublicationAuthorDeleteOne is the builder for deleting a single PublicationAuthor entity.
type PublicationAuthorDeleteOne struct {
	pad *PublicationAuthorDelete
}

// Where appends a list predicates to the PublicationAuthorDelete builder.
func (pado *PublicationAuthorDeleteOne) Where(ps ...predicate.PublicationAuthor) *PublicationAuthorDeleteOne {
	pado.pad.mutation.Where(ps...)
	return pado
}

// Exec executes the deletion query.
func (pado *PublicationAuthorDeleteOne) Exec(ctx context.Context) error {
	n, err := pado.pad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{publicationauthor.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pado *PublicationAuthorDeleteOne) ExecX(ctx context.Context) {
	if err := pado.Exec(ctx); err != nil {
		panic(err)
	}
}
