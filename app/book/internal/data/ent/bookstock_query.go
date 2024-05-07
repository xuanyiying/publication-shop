// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/xuanyiying/publication-shop/app/book/internal/data/ent/bookstock"
	"github.com/xuanyiying/publication-shop/app/book/internal/data/ent/predicate"
)

// BookStockQuery is the builder for querying BookStock entities.
type BookStockQuery struct {
	config
	ctx        *QueryContext
	order      []bookstock.OrderOption
	inters     []Interceptor
	predicates []predicate.BookStock
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BookStockQuery builder.
func (bsq *BookStockQuery) Where(ps ...predicate.BookStock) *BookStockQuery {
	bsq.predicates = append(bsq.predicates, ps...)
	return bsq
}

// Limit the number of records to be returned by this query.
func (bsq *BookStockQuery) Limit(limit int) *BookStockQuery {
	bsq.ctx.Limit = &limit
	return bsq
}

// Offset to start from.
func (bsq *BookStockQuery) Offset(offset int) *BookStockQuery {
	bsq.ctx.Offset = &offset
	return bsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bsq *BookStockQuery) Unique(unique bool) *BookStockQuery {
	bsq.ctx.Unique = &unique
	return bsq
}

// Order specifies how the records should be ordered.
func (bsq *BookStockQuery) Order(o ...bookstock.OrderOption) *BookStockQuery {
	bsq.order = append(bsq.order, o...)
	return bsq
}

// First returns the first BookStock entity from the query.
// Returns a *NotFoundError when no BookStock was found.
func (bsq *BookStockQuery) First(ctx context.Context) (*BookStock, error) {
	nodes, err := bsq.Limit(1).All(setContextOp(ctx, bsq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{bookstock.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bsq *BookStockQuery) FirstX(ctx context.Context) *BookStock {
	node, err := bsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BookStock ID from the query.
// Returns a *NotFoundError when no BookStock ID was found.
func (bsq *BookStockQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bsq.Limit(1).IDs(setContextOp(ctx, bsq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{bookstock.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bsq *BookStockQuery) FirstIDX(ctx context.Context) int {
	id, err := bsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BookStock entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BookStock entity is found.
// Returns a *NotFoundError when no BookStock entities are found.
func (bsq *BookStockQuery) Only(ctx context.Context) (*BookStock, error) {
	nodes, err := bsq.Limit(2).All(setContextOp(ctx, bsq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{bookstock.Label}
	default:
		return nil, &NotSingularError{bookstock.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bsq *BookStockQuery) OnlyX(ctx context.Context) *BookStock {
	node, err := bsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BookStock ID in the query.
// Returns a *NotSingularError when more than one BookStock ID is found.
// Returns a *NotFoundError when no entities are found.
func (bsq *BookStockQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bsq.Limit(2).IDs(setContextOp(ctx, bsq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{bookstock.Label}
	default:
		err = &NotSingularError{bookstock.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bsq *BookStockQuery) OnlyIDX(ctx context.Context) int {
	id, err := bsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BookStocks.
func (bsq *BookStockQuery) All(ctx context.Context) ([]*BookStock, error) {
	ctx = setContextOp(ctx, bsq.ctx, "All")
	if err := bsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*BookStock, *BookStockQuery]()
	return withInterceptors[[]*BookStock](ctx, bsq, qr, bsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bsq *BookStockQuery) AllX(ctx context.Context) []*BookStock {
	nodes, err := bsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BookStock IDs.
func (bsq *BookStockQuery) IDs(ctx context.Context) (ids []int, err error) {
	if bsq.ctx.Unique == nil && bsq.path != nil {
		bsq.Unique(true)
	}
	ctx = setContextOp(ctx, bsq.ctx, "IDs")
	if err = bsq.Select(bookstock.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bsq *BookStockQuery) IDsX(ctx context.Context) []int {
	ids, err := bsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bsq *BookStockQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bsq.ctx, "Count")
	if err := bsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bsq, querierCount[*BookStockQuery](), bsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bsq *BookStockQuery) CountX(ctx context.Context) int {
	count, err := bsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bsq *BookStockQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bsq.ctx, "Exist")
	switch _, err := bsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bsq *BookStockQuery) ExistX(ctx context.Context) bool {
	exist, err := bsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BookStockQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bsq *BookStockQuery) Clone() *BookStockQuery {
	if bsq == nil {
		return nil
	}
	return &BookStockQuery{
		config:     bsq.config,
		ctx:        bsq.ctx.Clone(),
		order:      append([]bookstock.OrderOption{}, bsq.order...),
		inters:     append([]Interceptor{}, bsq.inters...),
		predicates: append([]predicate.BookStock{}, bsq.predicates...),
		// clone intermediate query.
		sql:  bsq.sql.Clone(),
		path: bsq.path,
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
//	client.BookStock.Query().
//		GroupBy(bookstock.FieldIsbn).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bsq *BookStockQuery) GroupBy(field string, fields ...string) *BookStockGroupBy {
	bsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BookStockGroupBy{build: bsq}
	grbuild.flds = &bsq.ctx.Fields
	grbuild.label = bookstock.Label
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
//	client.BookStock.Query().
//		Select(bookstock.FieldIsbn).
//		Scan(ctx, &v)
func (bsq *BookStockQuery) Select(fields ...string) *BookStockSelect {
	bsq.ctx.Fields = append(bsq.ctx.Fields, fields...)
	sbuild := &BookStockSelect{BookStockQuery: bsq}
	sbuild.label = bookstock.Label
	sbuild.flds, sbuild.scan = &bsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BookStockSelect configured with the given aggregations.
func (bsq *BookStockQuery) Aggregate(fns ...AggregateFunc) *BookStockSelect {
	return bsq.Select().Aggregate(fns...)
}

func (bsq *BookStockQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bsq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bsq); err != nil {
				return err
			}
		}
	}
	for _, f := range bsq.ctx.Fields {
		if !bookstock.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bsq.path != nil {
		prev, err := bsq.path(ctx)
		if err != nil {
			return err
		}
		bsq.sql = prev
	}
	return nil
}

func (bsq *BookStockQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BookStock, error) {
	var (
		nodes = []*BookStock{}
		_spec = bsq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*BookStock).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &BookStock{config: bsq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (bsq *BookStockQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bsq.querySpec()
	_spec.Node.Columns = bsq.ctx.Fields
	if len(bsq.ctx.Fields) > 0 {
		_spec.Unique = bsq.ctx.Unique != nil && *bsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bsq.driver, _spec)
}

func (bsq *BookStockQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(bookstock.Table, bookstock.Columns, sqlgraph.NewFieldSpec(bookstock.FieldID, field.TypeInt))
	_spec.From = bsq.sql
	if unique := bsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bsq.path != nil {
		_spec.Unique = true
	}
	if fields := bsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bookstock.FieldID)
		for i := range fields {
			if fields[i] != bookstock.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bsq *BookStockQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bsq.driver.Dialect())
	t1 := builder.Table(bookstock.Table)
	columns := bsq.ctx.Fields
	if len(columns) == 0 {
		columns = bookstock.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bsq.sql != nil {
		selector = bsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bsq.ctx.Unique != nil && *bsq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range bsq.predicates {
		p(selector)
	}
	for _, p := range bsq.order {
		p(selector)
	}
	if offset := bsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BookStockGroupBy is the group-by builder for BookStock entities.
type BookStockGroupBy struct {
	selector
	build *BookStockQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bsgb *BookStockGroupBy) Aggregate(fns ...AggregateFunc) *BookStockGroupBy {
	bsgb.fns = append(bsgb.fns, fns...)
	return bsgb
}

// Scan applies the selector query and scans the result into the given value.
func (bsgb *BookStockGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bsgb.build.ctx, "GroupBy")
	if err := bsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BookStockQuery, *BookStockGroupBy](ctx, bsgb.build, bsgb, bsgb.build.inters, v)
}

func (bsgb *BookStockGroupBy) sqlScan(ctx context.Context, root *BookStockQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bsgb.fns))
	for _, fn := range bsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bsgb.flds)+len(bsgb.fns))
		for _, f := range *bsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BookStockSelect is the builder for selecting fields of BookStock entities.
type BookStockSelect struct {
	*BookStockQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bss *BookStockSelect) Aggregate(fns ...AggregateFunc) *BookStockSelect {
	bss.fns = append(bss.fns, fns...)
	return bss
}

// Scan applies the selector query and scans the result into the given value.
func (bss *BookStockSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bss.ctx, "Select")
	if err := bss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BookStockQuery, *BookStockSelect](ctx, bss.BookStockQuery, bss, bss.inters, v)
}

func (bss *BookStockSelect) sqlScan(ctx context.Context, root *BookStockQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bss.fns))
	for _, fn := range bss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}