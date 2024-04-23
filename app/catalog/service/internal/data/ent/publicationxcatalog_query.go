// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"github.com/go-kratos/publication-shop/app/catalog/service/internal/data/ent/predicate"
	"github.com/go-kratos/publication-shop/app/catalog/service/internal/data/ent/publicationxcatalog"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PublicationXCatalogQuery is the builder for querying PublicationXCatalog entities.
type PublicationXCatalogQuery struct {
	config
	ctx        *QueryContext
	order      []publicationxcatalog.OrderOption
	inters     []Interceptor
	predicates []predicate.PublicationXCatalog
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PublicationXCatalogQuery builder.
func (pxq *PublicationXCatalogQuery) Where(ps ...predicate.PublicationXCatalog) *PublicationXCatalogQuery {
	pxq.predicates = append(pxq.predicates, ps...)
	return pxq
}

// Limit the number of records to be returned by this query.
func (pxq *PublicationXCatalogQuery) Limit(limit int) *PublicationXCatalogQuery {
	pxq.ctx.Limit = &limit
	return pxq
}

// Offset to start from.
func (pxq *PublicationXCatalogQuery) Offset(offset int) *PublicationXCatalogQuery {
	pxq.ctx.Offset = &offset
	return pxq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pxq *PublicationXCatalogQuery) Unique(unique bool) *PublicationXCatalogQuery {
	pxq.ctx.Unique = &unique
	return pxq
}

// Order specifies how the records should be ordered.
func (pxq *PublicationXCatalogQuery) Order(o ...publicationxcatalog.OrderOption) *PublicationXCatalogQuery {
	pxq.order = append(pxq.order, o...)
	return pxq
}

// First returns the first PublicationXCatalog entity from the query.
// Returns a *NotFoundError when no PublicationXCatalog was found.
func (pxq *PublicationXCatalogQuery) First(ctx context.Context) (*PublicationXCatalog, error) {
	nodes, err := pxq.Limit(1).All(setContextOp(ctx, pxq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{publicationxcatalog.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pxq *PublicationXCatalogQuery) FirstX(ctx context.Context) *PublicationXCatalog {
	node, err := pxq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PublicationXCatalog ID from the query.
// Returns a *NotFoundError when no PublicationXCatalog ID was found.
func (pxq *PublicationXCatalogQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pxq.Limit(1).IDs(setContextOp(ctx, pxq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{publicationxcatalog.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pxq *PublicationXCatalogQuery) FirstIDX(ctx context.Context) int {
	id, err := pxq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PublicationXCatalog entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PublicationXCatalog entity is found.
// Returns a *NotFoundError when no PublicationXCatalog entities are found.
func (pxq *PublicationXCatalogQuery) Only(ctx context.Context) (*PublicationXCatalog, error) {
	nodes, err := pxq.Limit(2).All(setContextOp(ctx, pxq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{publicationxcatalog.Label}
	default:
		return nil, &NotSingularError{publicationxcatalog.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pxq *PublicationXCatalogQuery) OnlyX(ctx context.Context) *PublicationXCatalog {
	node, err := pxq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PublicationXCatalog ID in the query.
// Returns a *NotSingularError when more than one PublicationXCatalog ID is found.
// Returns a *NotFoundError when no entities are found.
func (pxq *PublicationXCatalogQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pxq.Limit(2).IDs(setContextOp(ctx, pxq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{publicationxcatalog.Label}
	default:
		err = &NotSingularError{publicationxcatalog.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pxq *PublicationXCatalogQuery) OnlyIDX(ctx context.Context) int {
	id, err := pxq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PublicationXCatalogs.
func (pxq *PublicationXCatalogQuery) All(ctx context.Context) ([]*PublicationXCatalog, error) {
	ctx = setContextOp(ctx, pxq.ctx, "All")
	if err := pxq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PublicationXCatalog, *PublicationXCatalogQuery]()
	return withInterceptors[[]*PublicationXCatalog](ctx, pxq, qr, pxq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pxq *PublicationXCatalogQuery) AllX(ctx context.Context) []*PublicationXCatalog {
	nodes, err := pxq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PublicationXCatalog IDs.
func (pxq *PublicationXCatalogQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pxq.ctx.Unique == nil && pxq.path != nil {
		pxq.Unique(true)
	}
	ctx = setContextOp(ctx, pxq.ctx, "IDs")
	if err = pxq.Select(publicationxcatalog.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pxq *PublicationXCatalogQuery) IDsX(ctx context.Context) []int {
	ids, err := pxq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pxq *PublicationXCatalogQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pxq.ctx, "Count")
	if err := pxq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pxq, querierCount[*PublicationXCatalogQuery](), pxq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pxq *PublicationXCatalogQuery) CountX(ctx context.Context) int {
	count, err := pxq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pxq *PublicationXCatalogQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pxq.ctx, "Exist")
	switch _, err := pxq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pxq *PublicationXCatalogQuery) ExistX(ctx context.Context) bool {
	exist, err := pxq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PublicationXCatalogQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pxq *PublicationXCatalogQuery) Clone() *PublicationXCatalogQuery {
	if pxq == nil {
		return nil
	}
	return &PublicationXCatalogQuery{
		config:     pxq.config,
		ctx:        pxq.ctx.Clone(),
		order:      append([]publicationxcatalog.OrderOption{}, pxq.order...),
		inters:     append([]Interceptor{}, pxq.inters...),
		predicates: append([]predicate.PublicationXCatalog{}, pxq.predicates...),
		// clone intermediate query.
		sql:  pxq.sql.Clone(),
		path: pxq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CatalogID int `json:"catalog_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PublicationXCatalog.Query().
//		GroupBy(publicationxcatalog.FieldCatalogID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pxq *PublicationXCatalogQuery) GroupBy(field string, fields ...string) *PublicationXCatalogGroupBy {
	pxq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PublicationXCatalogGroupBy{build: pxq}
	grbuild.flds = &pxq.ctx.Fields
	grbuild.label = publicationxcatalog.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CatalogID int `json:"catalog_id,omitempty"`
//	}
//
//	client.PublicationXCatalog.Query().
//		Select(publicationxcatalog.FieldCatalogID).
//		Scan(ctx, &v)
func (pxq *PublicationXCatalogQuery) Select(fields ...string) *PublicationXCatalogSelect {
	pxq.ctx.Fields = append(pxq.ctx.Fields, fields...)
	sbuild := &PublicationXCatalogSelect{PublicationXCatalogQuery: pxq}
	sbuild.label = publicationxcatalog.Label
	sbuild.flds, sbuild.scan = &pxq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PublicationXCatalogSelect configured with the given aggregations.
func (pxq *PublicationXCatalogQuery) Aggregate(fns ...AggregateFunc) *PublicationXCatalogSelect {
	return pxq.Select().Aggregate(fns...)
}

func (pxq *PublicationXCatalogQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pxq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pxq); err != nil {
				return err
			}
		}
	}
	for _, f := range pxq.ctx.Fields {
		if !publicationxcatalog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pxq.path != nil {
		prev, err := pxq.path(ctx)
		if err != nil {
			return err
		}
		pxq.sql = prev
	}
	return nil
}

func (pxq *PublicationXCatalogQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PublicationXCatalog, error) {
	var (
		nodes = []*PublicationXCatalog{}
		_spec = pxq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PublicationXCatalog).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PublicationXCatalog{config: pxq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pxq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (pxq *PublicationXCatalogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pxq.querySpec()
	_spec.Node.Columns = pxq.ctx.Fields
	if len(pxq.ctx.Fields) > 0 {
		_spec.Unique = pxq.ctx.Unique != nil && *pxq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pxq.driver, _spec)
}

func (pxq *PublicationXCatalogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(publicationxcatalog.Table, publicationxcatalog.Columns, sqlgraph.NewFieldSpec(publicationxcatalog.FieldID, field.TypeInt))
	_spec.From = pxq.sql
	if unique := pxq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pxq.path != nil {
		_spec.Unique = true
	}
	if fields := pxq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, publicationxcatalog.FieldID)
		for i := range fields {
			if fields[i] != publicationxcatalog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pxq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pxq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pxq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pxq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pxq *PublicationXCatalogQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pxq.driver.Dialect())
	t1 := builder.Table(publicationxcatalog.Table)
	columns := pxq.ctx.Fields
	if len(columns) == 0 {
		columns = publicationxcatalog.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pxq.sql != nil {
		selector = pxq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pxq.ctx.Unique != nil && *pxq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pxq.predicates {
		p(selector)
	}
	for _, p := range pxq.order {
		p(selector)
	}
	if offset := pxq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pxq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PublicationXCatalogGroupBy is the group-by builder for PublicationXCatalog entities.
type PublicationXCatalogGroupBy struct {
	selector
	build *PublicationXCatalogQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pxgb *PublicationXCatalogGroupBy) Aggregate(fns ...AggregateFunc) *PublicationXCatalogGroupBy {
	pxgb.fns = append(pxgb.fns, fns...)
	return pxgb
}

// Scan applies the selector query and scans the result into the given value.
func (pxgb *PublicationXCatalogGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pxgb.build.ctx, "GroupBy")
	if err := pxgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PublicationXCatalogQuery, *PublicationXCatalogGroupBy](ctx, pxgb.build, pxgb, pxgb.build.inters, v)
}

func (pxgb *PublicationXCatalogGroupBy) sqlScan(ctx context.Context, root *PublicationXCatalogQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pxgb.fns))
	for _, fn := range pxgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pxgb.flds)+len(pxgb.fns))
		for _, f := range *pxgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pxgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pxgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PublicationXCatalogSelect is the builder for selecting fields of PublicationXCatalog entities.
type PublicationXCatalogSelect struct {
	*PublicationXCatalogQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pxs *PublicationXCatalogSelect) Aggregate(fns ...AggregateFunc) *PublicationXCatalogSelect {
	pxs.fns = append(pxs.fns, fns...)
	return pxs
}

// Scan applies the selector query and scans the result into the given value.
func (pxs *PublicationXCatalogSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pxs.ctx, "Select")
	if err := pxs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PublicationXCatalogQuery, *PublicationXCatalogSelect](ctx, pxs.PublicationXCatalogQuery, pxs, pxs.inters, v)
}

func (pxs *PublicationXCatalogSelect) sqlScan(ctx context.Context, root *PublicationXCatalogQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pxs.fns))
	for _, fn := range pxs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pxs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pxs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}