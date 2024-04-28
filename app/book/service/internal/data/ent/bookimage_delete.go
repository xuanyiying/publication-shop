// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/xuanyiying/publication-shop/app/book/service/internal/data/ent/bookimage"
	"github.com/xuanyiying/publication-shop/app/book/service/internal/data/ent/predicate"
)

// BookImageDelete is the builder for deleting a BookImage entity.
type BookImageDelete struct {
	config
	hooks    []Hook
	mutation *BookImageMutation
}

// Where appends a list predicates to the BookImageDelete builder.
func (bid *BookImageDelete) Where(ps ...predicate.BookImage) *BookImageDelete {
	bid.mutation.Where(ps...)
	return bid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bid *BookImageDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, bid.sqlExec, bid.mutation, bid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bid *BookImageDelete) ExecX(ctx context.Context) int {
	n, err := bid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bid *BookImageDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(bookimage.Table, sqlgraph.NewFieldSpec(bookimage.FieldID, field.TypeInt))
	if ps := bid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bid.mutation.done = true
	return affected, err
}

// BookImageDeleteOne is the builder for deleting a single BookImage entity.
type BookImageDeleteOne struct {
	bid *BookImageDelete
}

// Where appends a list predicates to the BookImageDelete builder.
func (bido *BookImageDeleteOne) Where(ps ...predicate.BookImage) *BookImageDeleteOne {
	bido.bid.mutation.Where(ps...)
	return bido
}

// Exec executes the deletion query.
func (bido *BookImageDeleteOne) Exec(ctx context.Context) error {
	n, err := bido.bid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{bookimage.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bido *BookImageDeleteOne) ExecX(ctx context.Context) {
	if err := bido.Exec(ctx); err != nil {
		panic(err)
	}
}
