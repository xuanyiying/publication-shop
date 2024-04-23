// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/publication-shop/app/catalog/service/internal/data/ent/catalog"
	"github.com/go-kratos/publication-shop/app/catalog/service/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CatalogUpdate is the builder for updating Catalog entities.
type CatalogUpdate struct {
	config
	hooks    []Hook
	mutation *CatalogMutation
}

// Where appends a list predicates to the CatalogUpdate builder.
func (cu *CatalogUpdate) Where(ps ...predicate.Catalog) *CatalogUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetCatalogCode sets the "catalog_code" field.
func (cu *CatalogUpdate) SetCatalogCode(s string) *CatalogUpdate {
	cu.mutation.SetCatalogCode(s)
	return cu
}

// SetNillableCatalogCode sets the "catalog_code" field if the given value is not nil.
func (cu *CatalogUpdate) SetNillableCatalogCode(s *string) *CatalogUpdate {
	if s != nil {
		cu.SetCatalogCode(*s)
	}
	return cu
}

// SetCatalogName sets the "catalog_name" field.
func (cu *CatalogUpdate) SetCatalogName(s string) *CatalogUpdate {
	cu.mutation.SetCatalogName(s)
	return cu
}

// SetNillableCatalogName sets the "catalog_name" field if the given value is not nil.
func (cu *CatalogUpdate) SetNillableCatalogName(s *string) *CatalogUpdate {
	if s != nil {
		cu.SetCatalogName(*s)
	}
	return cu
}

// ClearCatalogName clears the value of the "catalog_name" field.
func (cu *CatalogUpdate) ClearCatalogName() *CatalogUpdate {
	cu.mutation.ClearCatalogName()
	return cu
}

// SetParentCatalogID sets the "parent_catalog_id" field.
func (cu *CatalogUpdate) SetParentCatalogID(s string) *CatalogUpdate {
	cu.mutation.SetParentCatalogID(s)
	return cu
}

// SetNillableParentCatalogID sets the "parent_catalog_id" field if the given value is not nil.
func (cu *CatalogUpdate) SetNillableParentCatalogID(s *string) *CatalogUpdate {
	if s != nil {
		cu.SetParentCatalogID(*s)
	}
	return cu
}

// Mutation returns the CatalogMutation object of the builder.
func (cu *CatalogUpdate) Mutation() *CatalogMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CatalogUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CatalogUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CatalogUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CatalogUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CatalogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(catalog.Table, catalog.Columns, sqlgraph.NewFieldSpec(catalog.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.CatalogCode(); ok {
		_spec.SetField(catalog.FieldCatalogCode, field.TypeString, value)
	}
	if value, ok := cu.mutation.CatalogName(); ok {
		_spec.SetField(catalog.FieldCatalogName, field.TypeString, value)
	}
	if cu.mutation.CatalogNameCleared() {
		_spec.ClearField(catalog.FieldCatalogName, field.TypeString)
	}
	if value, ok := cu.mutation.ParentCatalogID(); ok {
		_spec.SetField(catalog.FieldParentCatalogID, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{catalog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CatalogUpdateOne is the builder for updating a single Catalog entity.
type CatalogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CatalogMutation
}

// SetCatalogCode sets the "catalog_code" field.
func (cuo *CatalogUpdateOne) SetCatalogCode(s string) *CatalogUpdateOne {
	cuo.mutation.SetCatalogCode(s)
	return cuo
}

// SetNillableCatalogCode sets the "catalog_code" field if the given value is not nil.
func (cuo *CatalogUpdateOne) SetNillableCatalogCode(s *string) *CatalogUpdateOne {
	if s != nil {
		cuo.SetCatalogCode(*s)
	}
	return cuo
}

// SetCatalogName sets the "catalog_name" field.
func (cuo *CatalogUpdateOne) SetCatalogName(s string) *CatalogUpdateOne {
	cuo.mutation.SetCatalogName(s)
	return cuo
}

// SetNillableCatalogName sets the "catalog_name" field if the given value is not nil.
func (cuo *CatalogUpdateOne) SetNillableCatalogName(s *string) *CatalogUpdateOne {
	if s != nil {
		cuo.SetCatalogName(*s)
	}
	return cuo
}

// ClearCatalogName clears the value of the "catalog_name" field.
func (cuo *CatalogUpdateOne) ClearCatalogName() *CatalogUpdateOne {
	cuo.mutation.ClearCatalogName()
	return cuo
}

// SetParentCatalogID sets the "parent_catalog_id" field.
func (cuo *CatalogUpdateOne) SetParentCatalogID(s string) *CatalogUpdateOne {
	cuo.mutation.SetParentCatalogID(s)
	return cuo
}

// SetNillableParentCatalogID sets the "parent_catalog_id" field if the given value is not nil.
func (cuo *CatalogUpdateOne) SetNillableParentCatalogID(s *string) *CatalogUpdateOne {
	if s != nil {
		cuo.SetParentCatalogID(*s)
	}
	return cuo
}

// Mutation returns the CatalogMutation object of the builder.
func (cuo *CatalogUpdateOne) Mutation() *CatalogMutation {
	return cuo.mutation
}

// Where appends a list predicates to the CatalogUpdate builder.
func (cuo *CatalogUpdateOne) Where(ps ...predicate.Catalog) *CatalogUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CatalogUpdateOne) Select(field string, fields ...string) *CatalogUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Catalog entity.
func (cuo *CatalogUpdateOne) Save(ctx context.Context) (*Catalog, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CatalogUpdateOne) SaveX(ctx context.Context) *Catalog {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CatalogUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CatalogUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CatalogUpdateOne) sqlSave(ctx context.Context) (_node *Catalog, err error) {
	_spec := sqlgraph.NewUpdateSpec(catalog.Table, catalog.Columns, sqlgraph.NewFieldSpec(catalog.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Catalog.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, catalog.FieldID)
		for _, f := range fields {
			if !catalog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != catalog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.CatalogCode(); ok {
		_spec.SetField(catalog.FieldCatalogCode, field.TypeString, value)
	}
	if value, ok := cuo.mutation.CatalogName(); ok {
		_spec.SetField(catalog.FieldCatalogName, field.TypeString, value)
	}
	if cuo.mutation.CatalogNameCleared() {
		_spec.ClearField(catalog.FieldCatalogName, field.TypeString)
	}
	if value, ok := cuo.mutation.ParentCatalogID(); ok {
		_spec.SetField(catalog.FieldParentCatalogID, field.TypeString, value)
	}
	_node = &Catalog{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{catalog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}