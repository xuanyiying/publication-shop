// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/xuanyiying/publication-shop/app/catalog/ent/predicate"
	"github.com/xuanyiying/publication-shop/app/catalog/ent/publicationstock"
)

// PublicationStockQuery is the builder for querying PublicationStock entities.
type PublicationStockQuery struct {
	config
	ctx        *QueryContext
	order      []publicationstock.OrderOption
	inters     []Interceptor
	predicates []predicate.PublicationStock
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PublicationStockQuery builder.
func (psq *PublicationStockQuery) Where(ps ...predicate.PublicationStock) *PublicationStockQuery {
	psq.predicates = append(psq.predicates, ps...)
	return psq
}

// Limit the number of records to be returned by this query.
func (psq *PublicationStockQuery) Limit(limit int) *PublicationStockQuery {
	psq.ctx.Limit = &limit
	return psq
}

// Offset to start from.
func (psq *PublicationStockQuery) Offset(offset int) *PublicationStockQuery {
	psq.ctx.Offset = &offset
	return psq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (psq *PublicationStockQuery) Unique(unique bool) *PublicationStockQuery {
	psq.ctx.Unique = &unique
	return psq
}

// Order specifies how the records should be ordered.
func (psq *PublicationStockQuery) Order(o ...publicationstock.OrderOption) *PublicationStockQuery {
	psq.order = append(psq.order, o...)
	return psq
}

// First returns the first PublicationStock entity from the query.
// Returns a *NotFoundError when no PublicationStock was found.
func (psq *PublicationStockQuery) First(ctx context.Context) (*PublicationStock, error) {
	nodes, err := psq.Limit(1).All(setContextOp(ctx, psq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{publicationstock.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (psq *PublicationStockQuery) FirstX(ctx context.Context) *PublicationStock {
	node, err := psq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PublicationStock ID from the query.
// Returns a *NotFoundError when no PublicationStock ID was found.
func (psq *PublicationStockQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = psq.Limit(1).IDs(setContextOp(ctx, psq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{publicationstock.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (psq *PublicationStockQuery) FirstIDX(ctx context.Context) int {
	id, err := psq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PublicationStock entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PublicationStock entity is found.
// Returns a *NotFoundError when no PublicationStock entities are found.
func (psq *PublicationStockQuery) Only(ctx context.Context) (*PublicationStock, error) {
	nodes, err := psq.Limit(2).All(setContextOp(ctx, psq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{publicationstock.Label}
	default:
		return nil, &NotSingularError{publicationstock.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (psq *PublicationStockQuery) OnlyX(ctx context.Context) *PublicationStock {
	node, err := psq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PublicationStock ID in the query.
// Returns a *NotSingularError when more than one PublicationStock ID is found.
// Returns a *NotFoundError when no entities are found.
func (psq *PublicationStockQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = psq.Limit(2).IDs(setContextOp(ctx, psq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{publicationstock.Label}
	default:
		err = &NotSingularError{publicationstock.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (psq *PublicationStockQuery) OnlyIDX(ctx context.Context) int {
	id, err := psq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PublicationStocks.
func (psq *PublicationStockQuery) All(ctx context.Context) ([]*PublicationStock, error) {
	ctx = setContextOp(ctx, psq.ctx, "All")
	if err := psq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PublicationStock, *PublicationStockQuery]()
	return withInterceptors[[]*PublicationStock](ctx, psq, qr, psq.inters)
}

// AllX is like All, but panics if an error occurs.
func (psq *PublicationStockQuery) AllX(ctx context.Context) []*PublicationStock {
	nodes, err := psq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PublicationStock IDs.
func (psq *PublicationStockQuery) IDs(ctx context.Context) (ids []int, err error) {
	if psq.ctx.Unique == nil && psq.path != nil {
		psq.Unique(true)
	}
	ctx = setContextOp(ctx, psq.ctx, "IDs")
	if err = psq.Select(publicationstock.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (psq *PublicationStockQuery) IDsX(ctx context.Context) []int {
	ids, err := psq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (psq *PublicationStockQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, psq.ctx, "Count")
	if err := psq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, psq, querierCount[*PublicationStockQuery](), psq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (psq *PublicationStockQuery) CountX(ctx context.Context) int {
	count, err := psq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (psq *PublicationStockQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, psq.ctx, "Exist")
	switch _, err := psq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (psq *PublicationStockQuery) ExistX(ctx context.Context) bool {
	exist, err := psq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PublicationStockQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (psq *PublicationStockQuery) Clone() *PublicationStockQuery {
	if psq == nil {
		return nil
	}
	return &PublicationStockQuery{
		config:     psq.config,
		ctx:        psq.ctx.Clone(),
		order:      append([]publicationstock.OrderOption{}, psq.order...),
		inters:     append([]Interceptor{}, psq.inters...),
		predicates: append([]predicate.PublicationStock{}, psq.predicates...),
		// clone intermediate query.
		sql:  psq.sql.Clone(),
		path: psq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Isbn string `json:"isbn,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PublicationStock.Query().
//		GroupBy(publicationstock.FieldIsbn).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (psq *PublicationStockQuery) GroupBy(field string, fields ...string) *PublicationStockGroupBy {
	psq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PublicationStockGroupBy{build: psq}
	grbuild.flds = &psq.ctx.Fields
	grbuild.label = publicationstock.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Isbn string `json:"isbn,omitempty"`
//	}
//
//	client.PublicationStock.Query().
//		Select(publicationstock.FieldIsbn).
//		Scan(ctx, &v)
func (psq *PublicationStockQuery) Select(fields ...string) *PublicationStockSelect {
	psq.ctx.Fields = append(psq.ctx.Fields, fields...)
	sbuild := &PublicationStockSelect{PublicationStockQuery: psq}
	sbuild.label = publicationstock.Label
	sbuild.flds, sbuild.scan = &psq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PublicationStockSelect configured with the given aggregations.
func (psq *PublicationStockQuery) Aggregate(fns ...AggregateFunc) *PublicationStockSelect {
	return psq.Select().Aggregate(fns...)
}

func (psq *PublicationStockQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range psq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, psq); err != nil {
				return err
			}
		}
	}
	for _, f := range psq.ctx.Fields {
		if !publicationstock.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if psq.path != nil {
		prev, err := psq.path(ctx)
		if err != nil {
			return err
		}
		psq.sql = prev
	}
	return nil
}

func (psq *PublicationStockQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PublicationStock, error) {
	var (
		nodes = []*PublicationStock{}
		_spec = psq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PublicationStock).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PublicationStock{config: psq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, psq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (psq *PublicationStockQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := psq.querySpec()
	_spec.Node.Columns = psq.ctx.Fields
	if len(psq.ctx.Fields) > 0 {
		_spec.Unique = psq.ctx.Unique != nil && *psq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, psq.driver, _spec)
}

func (psq *PublicationStockQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(publicationstock.Table, publicationstock.Columns, sqlgraph.NewFieldSpec(publicationstock.FieldID, field.TypeInt))
	_spec.From = psq.sql
	if unique := psq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if psq.path != nil {
		_spec.Unique = true
	}
	if fields := psq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, publicationstock.FieldID)
		for i := range fields {
			if fields[i] != publicationstock.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := psq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := psq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := psq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := psq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (psq *PublicationStockQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(psq.driver.Dialect())
	t1 := builder.Table(publicationstock.Table)
	columns := psq.ctx.Fields
	if len(columns) == 0 {
		columns = publicationstock.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if psq.sql != nil {
		selector = psq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if psq.ctx.Unique != nil && *psq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range psq.predicates {
		p(selector)
	}
	for _, p := range psq.order {
		p(selector)
	}
	if offset := psq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := psq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PublicationStockGroupBy is the group-by builder for PublicationStock entities.
type PublicationStockGroupBy struct {
	selector
	build *PublicationStockQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (psgb *PublicationStockGroupBy) Aggregate(fns ...AggregateFunc) *PublicationStockGroupBy {
	psgb.fns = append(psgb.fns, fns...)
	return psgb
}

// Scan applies the selector query and scans the result into the given value.
func (psgb *PublicationStockGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, psgb.build.ctx, "GroupBy")
	if err := psgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PublicationStockQuery, *PublicationStockGroupBy](ctx, psgb.build, psgb, psgb.build.inters, v)
}

func (psgb *PublicationStockGroupBy) sqlScan(ctx context.Context, root *PublicationStockQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(psgb.fns))
	for _, fn := range psgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*psgb.flds)+len(psgb.fns))
		for _, f := range *psgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*psgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := psgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PublicationStockSelect is the builder for selecting fields of PublicationStock entities.
type PublicationStockSelect struct {
	*PublicationStockQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pss *PublicationStockSelect) Aggregate(fns ...AggregateFunc) *PublicationStockSelect {
	pss.fns = append(pss.fns, fns...)
	return pss
}

// Scan applies the selector query and scans the result into the given value.
func (pss *PublicationStockSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pss.ctx, "Select")
	if err := pss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PublicationStockQuery, *PublicationStockSelect](ctx, pss.PublicationStockQuery, pss, pss.inters, v)
}

func (pss *PublicationStockSelect) sqlScan(ctx context.Context, root *PublicationStockQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pss.fns))
	for _, fn := range pss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
