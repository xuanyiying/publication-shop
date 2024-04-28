// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/xuanyiying/publication-shop/app/book/service/internal/data/ent/bookxlanguage"
	"github.com/xuanyiying/publication-shop/app/book/service/internal/data/ent/predicate"
)

// BookXLanguageDelete is the builder for deleting a BookXLanguage entity.
type BookXLanguageDelete struct {
	config
	hooks    []Hook
	mutation *BookXLanguageMutation
}

// Where appends a list predicates to the BookXLanguageDelete builder.
func (bxd *BookXLanguageDelete) Where(ps ...predicate.BookXLanguage) *BookXLanguageDelete {
	bxd.mutation.Where(ps...)
	return bxd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bxd *BookXLanguageDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, bxd.sqlExec, bxd.mutation, bxd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bxd *BookXLanguageDelete) ExecX(ctx context.Context) int {
	n, err := bxd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bxd *BookXLanguageDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(bookxlanguage.Table, sqlgraph.NewFieldSpec(bookxlanguage.FieldID, field.TypeInt))
	if ps := bxd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bxd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bxd.mutation.done = true
	return affected, err
}

// BookXLanguageDeleteOne is the builder for deleting a single BookXLanguage entity.
type BookXLanguageDeleteOne struct {
	bxd *BookXLanguageDelete
}

// Where appends a list predicates to the BookXLanguageDelete builder.
func (bxdo *BookXLanguageDeleteOne) Where(ps ...predicate.BookXLanguage) *BookXLanguageDeleteOne {
	bxdo.bxd.mutation.Where(ps...)
	return bxdo
}

// Exec executes the deletion query.
func (bxdo *BookXLanguageDeleteOne) Exec(ctx context.Context) error {
	n, err := bxdo.bxd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{bookxlanguage.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bxdo *BookXLanguageDeleteOne) ExecX(ctx context.Context) {
	if err := bxdo.Exec(ctx); err != nil {
		panic(err)
	}
}
