// Code generated by ent, DO NOT EDIT.

package openuem_ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/doncicuto/openuem_ent/predicate"
	"github.com/doncicuto/openuem_ent/revocation"
)

// RevocationQuery is the builder for querying Revocation entities.
type RevocationQuery struct {
	config
	ctx        *QueryContext
	order      []revocation.OrderOption
	inters     []Interceptor
	predicates []predicate.Revocation
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RevocationQuery builder.
func (rq *RevocationQuery) Where(ps ...predicate.Revocation) *RevocationQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit the number of records to be returned by this query.
func (rq *RevocationQuery) Limit(limit int) *RevocationQuery {
	rq.ctx.Limit = &limit
	return rq
}

// Offset to start from.
func (rq *RevocationQuery) Offset(offset int) *RevocationQuery {
	rq.ctx.Offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *RevocationQuery) Unique(unique bool) *RevocationQuery {
	rq.ctx.Unique = &unique
	return rq
}

// Order specifies how the records should be ordered.
func (rq *RevocationQuery) Order(o ...revocation.OrderOption) *RevocationQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// First returns the first Revocation entity from the query.
// Returns a *NotFoundError when no Revocation was found.
func (rq *RevocationQuery) First(ctx context.Context) (*Revocation, error) {
	nodes, err := rq.Limit(1).All(setContextOp(ctx, rq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{revocation.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *RevocationQuery) FirstX(ctx context.Context) *Revocation {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Revocation ID from the query.
// Returns a *NotFoundError when no Revocation ID was found.
func (rq *RevocationQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = rq.Limit(1).IDs(setContextOp(ctx, rq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{revocation.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rq *RevocationQuery) FirstIDX(ctx context.Context) int64 {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Revocation entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Revocation entity is found.
// Returns a *NotFoundError when no Revocation entities are found.
func (rq *RevocationQuery) Only(ctx context.Context) (*Revocation, error) {
	nodes, err := rq.Limit(2).All(setContextOp(ctx, rq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{revocation.Label}
	default:
		return nil, &NotSingularError{revocation.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *RevocationQuery) OnlyX(ctx context.Context) *Revocation {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Revocation ID in the query.
// Returns a *NotSingularError when more than one Revocation ID is found.
// Returns a *NotFoundError when no entities are found.
func (rq *RevocationQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = rq.Limit(2).IDs(setContextOp(ctx, rq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{revocation.Label}
	default:
		err = &NotSingularError{revocation.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *RevocationQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Revocations.
func (rq *RevocationQuery) All(ctx context.Context) ([]*Revocation, error) {
	ctx = setContextOp(ctx, rq.ctx, ent.OpQueryAll)
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Revocation, *RevocationQuery]()
	return withInterceptors[[]*Revocation](ctx, rq, qr, rq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rq *RevocationQuery) AllX(ctx context.Context) []*Revocation {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Revocation IDs.
func (rq *RevocationQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if rq.ctx.Unique == nil && rq.path != nil {
		rq.Unique(true)
	}
	ctx = setContextOp(ctx, rq.ctx, ent.OpQueryIDs)
	if err = rq.Select(revocation.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *RevocationQuery) IDsX(ctx context.Context) []int64 {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *RevocationQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rq.ctx, ent.OpQueryCount)
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rq, querierCount[*RevocationQuery](), rq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rq *RevocationQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *RevocationQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rq.ctx, ent.OpQueryExist)
	switch _, err := rq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("openuem_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *RevocationQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RevocationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *RevocationQuery) Clone() *RevocationQuery {
	if rq == nil {
		return nil
	}
	return &RevocationQuery{
		config:     rq.config,
		ctx:        rq.ctx.Clone(),
		order:      append([]revocation.OrderOption{}, rq.order...),
		inters:     append([]Interceptor{}, rq.inters...),
		predicates: append([]predicate.Revocation{}, rq.predicates...),
		// clone intermediate query.
		sql:       rq.sql.Clone(),
		path:      rq.path,
		modifiers: append([]func(*sql.Selector){}, rq.modifiers...),
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Reason string `json:"reason,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Revocation.Query().
//		GroupBy(revocation.FieldReason).
//		Aggregate(openuem_ent.Count()).
//		Scan(ctx, &v)
func (rq *RevocationQuery) GroupBy(field string, fields ...string) *RevocationGroupBy {
	rq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RevocationGroupBy{build: rq}
	grbuild.flds = &rq.ctx.Fields
	grbuild.label = revocation.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Reason string `json:"reason,omitempty"`
//	}
//
//	client.Revocation.Query().
//		Select(revocation.FieldReason).
//		Scan(ctx, &v)
func (rq *RevocationQuery) Select(fields ...string) *RevocationSelect {
	rq.ctx.Fields = append(rq.ctx.Fields, fields...)
	sbuild := &RevocationSelect{RevocationQuery: rq}
	sbuild.label = revocation.Label
	sbuild.flds, sbuild.scan = &rq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RevocationSelect configured with the given aggregations.
func (rq *RevocationQuery) Aggregate(fns ...AggregateFunc) *RevocationSelect {
	return rq.Select().Aggregate(fns...)
}

func (rq *RevocationQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rq.inters {
		if inter == nil {
			return fmt.Errorf("openuem_ent: uninitialized interceptor (forgotten import openuem_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rq); err != nil {
				return err
			}
		}
	}
	for _, f := range rq.ctx.Fields {
		if !revocation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("openuem_ent: invalid field %q for query", f)}
		}
	}
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *RevocationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Revocation, error) {
	var (
		nodes = []*Revocation{}
		_spec = rq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Revocation).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Revocation{config: rq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(rq.modifiers) > 0 {
		_spec.Modifiers = rq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (rq *RevocationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	if len(rq.modifiers) > 0 {
		_spec.Modifiers = rq.modifiers
	}
	_spec.Node.Columns = rq.ctx.Fields
	if len(rq.ctx.Fields) > 0 {
		_spec.Unique = rq.ctx.Unique != nil && *rq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *RevocationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(revocation.Table, revocation.Columns, sqlgraph.NewFieldSpec(revocation.FieldID, field.TypeInt64))
	_spec.From = rq.sql
	if unique := rq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rq.path != nil {
		_spec.Unique = true
	}
	if fields := rq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, revocation.FieldID)
		for i := range fields {
			if fields[i] != revocation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *RevocationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(revocation.Table)
	columns := rq.ctx.Fields
	if len(columns) == 0 {
		columns = revocation.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rq.ctx.Unique != nil && *rq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range rq.modifiers {
		m(selector)
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (rq *RevocationQuery) Modify(modifiers ...func(s *sql.Selector)) *RevocationSelect {
	rq.modifiers = append(rq.modifiers, modifiers...)
	return rq.Select()
}

// RevocationGroupBy is the group-by builder for Revocation entities.
type RevocationGroupBy struct {
	selector
	build *RevocationQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *RevocationGroupBy) Aggregate(fns ...AggregateFunc) *RevocationGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the selector query and scans the result into the given value.
func (rgb *RevocationGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rgb.build.ctx, ent.OpQueryGroupBy)
	if err := rgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RevocationQuery, *RevocationGroupBy](ctx, rgb.build, rgb, rgb.build.inters, v)
}

func (rgb *RevocationGroupBy) sqlScan(ctx context.Context, root *RevocationQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rgb.fns))
	for _, fn := range rgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rgb.flds)+len(rgb.fns))
		for _, f := range *rgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RevocationSelect is the builder for selecting fields of Revocation entities.
type RevocationSelect struct {
	*RevocationQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rs *RevocationSelect) Aggregate(fns ...AggregateFunc) *RevocationSelect {
	rs.fns = append(rs.fns, fns...)
	return rs
}

// Scan applies the selector query and scans the result into the given value.
func (rs *RevocationSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rs.ctx, ent.OpQuerySelect)
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RevocationQuery, *RevocationSelect](ctx, rs.RevocationQuery, rs, rs.inters, v)
}

func (rs *RevocationSelect) sqlScan(ctx context.Context, root *RevocationQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rs.fns))
	for _, fn := range rs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (rs *RevocationSelect) Modify(modifiers ...func(s *sql.Selector)) *RevocationSelect {
	rs.modifiers = append(rs.modifiers, modifiers...)
	return rs
}
