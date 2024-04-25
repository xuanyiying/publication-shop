// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/xuanyiying/publication-shop/app/catalog/ent/predicate"
	"github.com/xuanyiying/publication-shop/app/catalog/ent/publicationorg"
)

// PublicationOrgDelete is the builder for deleting a PublicationOrg entity.
type PublicationOrgDelete struct {
	config
	hooks    []Hook
	mutation *PublicationOrgMutation
}

// Where appends a list predicates to the PublicationOrgDelete builder.
func (pod *PublicationOrgDelete) Where(ps ...predicate.PublicationOrg) *PublicationOrgDelete {
	pod.mutation.Where(ps...)
	return pod
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pod *PublicationOrgDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pod.sqlExec, pod.mutation, pod.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pod *PublicationOrgDelete) ExecX(ctx context.Context) int {
	n, err := pod.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pod *PublicationOrgDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(publicationorg.Table, sqlgraph.NewFieldSpec(publicationorg.FieldID, field.TypeInt))
	if ps := pod.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pod.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pod.mutation.done = true
	return affected, err
}

// PublicationOrgDeleteOne is the builder for deleting a single PublicationOrg entity.
type PublicationOrgDeleteOne struct {
	pod *PublicationOrgDelete
}

// Where appends a list predicates to the PublicationOrgDelete builder.
func (podo *PublicationOrgDeleteOne) Where(ps ...predicate.PublicationOrg) *PublicationOrgDeleteOne {
	podo.pod.mutation.Where(ps...)
	return podo
}

// Exec executes the deletion query.
func (podo *PublicationOrgDeleteOne) Exec(ctx context.Context) error {
	n, err := podo.pod.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{publicationorg.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (podo *PublicationOrgDeleteOne) ExecX(ctx context.Context) {
	if err := podo.Exec(ctx); err != nil {
		panic(err)
	}
}
