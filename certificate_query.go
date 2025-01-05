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
	"github.com/open-uem/openuem_ent/certificate"
	"github.com/open-uem/openuem_ent/predicate"
)

// CertificateQuery is the builder for querying Certificate entities.
type CertificateQuery struct {
	config
	ctx        *QueryContext
	order      []certificate.OrderOption
	inters     []Interceptor
	predicates []predicate.Certificate
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CertificateQuery builder.
func (cq *CertificateQuery) Where(ps ...predicate.Certificate) *CertificateQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *CertificateQuery) Limit(limit int) *CertificateQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *CertificateQuery) Offset(offset int) *CertificateQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CertificateQuery) Unique(unique bool) *CertificateQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *CertificateQuery) Order(o ...certificate.OrderOption) *CertificateQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// First returns the first Certificate entity from the query.
// Returns a *NotFoundError when no Certificate was found.
func (cq *CertificateQuery) First(ctx context.Context) (*Certificate, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{certificate.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CertificateQuery) FirstX(ctx context.Context) *Certificate {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Certificate ID from the query.
// Returns a *NotFoundError when no Certificate ID was found.
func (cq *CertificateQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{certificate.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CertificateQuery) FirstIDX(ctx context.Context) int64 {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Certificate entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Certificate entity is found.
// Returns a *NotFoundError when no Certificate entities are found.
func (cq *CertificateQuery) Only(ctx context.Context) (*Certificate, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{certificate.Label}
	default:
		return nil, &NotSingularError{certificate.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CertificateQuery) OnlyX(ctx context.Context) *Certificate {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Certificate ID in the query.
// Returns a *NotSingularError when more than one Certificate ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CertificateQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{certificate.Label}
	default:
		err = &NotSingularError{certificate.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CertificateQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Certificates.
func (cq *CertificateQuery) All(ctx context.Context) ([]*Certificate, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryAll)
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Certificate, *CertificateQuery]()
	return withInterceptors[[]*Certificate](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *CertificateQuery) AllX(ctx context.Context) []*Certificate {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Certificate IDs.
func (cq *CertificateQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryIDs)
	if err = cq.Select(certificate.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CertificateQuery) IDsX(ctx context.Context) []int64 {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CertificateQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryCount)
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*CertificateQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CertificateQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CertificateQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryExist)
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("openuem_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CertificateQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CertificateQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CertificateQuery) Clone() *CertificateQuery {
	if cq == nil {
		return nil
	}
	return &CertificateQuery{
		config:     cq.config,
		ctx:        cq.ctx.Clone(),
		order:      append([]certificate.OrderOption{}, cq.order...),
		inters:     append([]Interceptor{}, cq.inters...),
		predicates: append([]predicate.Certificate{}, cq.predicates...),
		// clone intermediate query.
		sql:       cq.sql.Clone(),
		path:      cq.path,
		modifiers: append([]func(*sql.Selector){}, cq.modifiers...),
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Type certificate.Type `json:"type,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Certificate.Query().
//		GroupBy(certificate.FieldType).
//		Aggregate(openuem_ent.Count()).
//		Scan(ctx, &v)
func (cq *CertificateQuery) GroupBy(field string, fields ...string) *CertificateGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CertificateGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = certificate.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Type certificate.Type `json:"type,omitempty"`
//	}
//
//	client.Certificate.Query().
//		Select(certificate.FieldType).
//		Scan(ctx, &v)
func (cq *CertificateQuery) Select(fields ...string) *CertificateSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &CertificateSelect{CertificateQuery: cq}
	sbuild.label = certificate.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CertificateSelect configured with the given aggregations.
func (cq *CertificateQuery) Aggregate(fns ...AggregateFunc) *CertificateSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *CertificateQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("openuem_ent: uninitialized interceptor (forgotten import openuem_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !certificate.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("openuem_ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *CertificateQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Certificate, error) {
	var (
		nodes = []*Certificate{}
		_spec = cq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Certificate).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Certificate{config: cq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (cq *CertificateQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CertificateQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(certificate.Table, certificate.Columns, sqlgraph.NewFieldSpec(certificate.FieldID, field.TypeInt64))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, certificate.FieldID)
		for i := range fields {
			if fields[i] != certificate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *CertificateQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(certificate.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = certificate.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range cq.modifiers {
		m(selector)
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cq *CertificateQuery) Modify(modifiers ...func(s *sql.Selector)) *CertificateSelect {
	cq.modifiers = append(cq.modifiers, modifiers...)
	return cq.Select()
}

// CertificateGroupBy is the group-by builder for Certificate entities.
type CertificateGroupBy struct {
	selector
	build *CertificateQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CertificateGroupBy) Aggregate(fns ...AggregateFunc) *CertificateGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CertificateGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, ent.OpQueryGroupBy)
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CertificateQuery, *CertificateGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *CertificateGroupBy) sqlScan(ctx context.Context, root *CertificateQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CertificateSelect is the builder for selecting fields of Certificate entities.
type CertificateSelect struct {
	*CertificateQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CertificateSelect) Aggregate(fns ...AggregateFunc) *CertificateSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CertificateSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, ent.OpQuerySelect)
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CertificateQuery, *CertificateSelect](ctx, cs.CertificateQuery, cs, cs.inters, v)
}

func (cs *CertificateSelect) sqlScan(ctx context.Context, root *CertificateQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cs *CertificateSelect) Modify(modifiers ...func(s *sql.Selector)) *CertificateSelect {
	cs.modifiers = append(cs.modifiers, modifiers...)
	return cs
}
