// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-uem/ent/agent"
	"github.com/open-uem/ent/operatingsystem"
	"github.com/open-uem/ent/predicate"
)

// OperatingSystemQuery is the builder for querying OperatingSystem entities.
type OperatingSystemQuery struct {
	config
	ctx        *QueryContext
	order      []operatingsystem.OrderOption
	inters     []Interceptor
	predicates []predicate.OperatingSystem
	withOwner  *AgentQuery
	withFKs    bool
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OperatingSystemQuery builder.
func (osq *OperatingSystemQuery) Where(ps ...predicate.OperatingSystem) *OperatingSystemQuery {
	osq.predicates = append(osq.predicates, ps...)
	return osq
}

// Limit the number of records to be returned by this query.
func (osq *OperatingSystemQuery) Limit(limit int) *OperatingSystemQuery {
	osq.ctx.Limit = &limit
	return osq
}

// Offset to start from.
func (osq *OperatingSystemQuery) Offset(offset int) *OperatingSystemQuery {
	osq.ctx.Offset = &offset
	return osq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (osq *OperatingSystemQuery) Unique(unique bool) *OperatingSystemQuery {
	osq.ctx.Unique = &unique
	return osq
}

// Order specifies how the records should be ordered.
func (osq *OperatingSystemQuery) Order(o ...operatingsystem.OrderOption) *OperatingSystemQuery {
	osq.order = append(osq.order, o...)
	return osq
}

// QueryOwner chains the current query on the "owner" edge.
func (osq *OperatingSystemQuery) QueryOwner() *AgentQuery {
	query := (&AgentClient{config: osq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := osq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := osq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(operatingsystem.Table, operatingsystem.FieldID, selector),
			sqlgraph.To(agent.Table, agent.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, operatingsystem.OwnerTable, operatingsystem.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(osq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OperatingSystem entity from the query.
// Returns a *NotFoundError when no OperatingSystem was found.
func (osq *OperatingSystemQuery) First(ctx context.Context) (*OperatingSystem, error) {
	nodes, err := osq.Limit(1).All(setContextOp(ctx, osq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{operatingsystem.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (osq *OperatingSystemQuery) FirstX(ctx context.Context) *OperatingSystem {
	node, err := osq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OperatingSystem ID from the query.
// Returns a *NotFoundError when no OperatingSystem ID was found.
func (osq *OperatingSystemQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = osq.Limit(1).IDs(setContextOp(ctx, osq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{operatingsystem.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (osq *OperatingSystemQuery) FirstIDX(ctx context.Context) int {
	id, err := osq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OperatingSystem entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OperatingSystem entity is found.
// Returns a *NotFoundError when no OperatingSystem entities are found.
func (osq *OperatingSystemQuery) Only(ctx context.Context) (*OperatingSystem, error) {
	nodes, err := osq.Limit(2).All(setContextOp(ctx, osq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{operatingsystem.Label}
	default:
		return nil, &NotSingularError{operatingsystem.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (osq *OperatingSystemQuery) OnlyX(ctx context.Context) *OperatingSystem {
	node, err := osq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OperatingSystem ID in the query.
// Returns a *NotSingularError when more than one OperatingSystem ID is found.
// Returns a *NotFoundError when no entities are found.
func (osq *OperatingSystemQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = osq.Limit(2).IDs(setContextOp(ctx, osq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{operatingsystem.Label}
	default:
		err = &NotSingularError{operatingsystem.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (osq *OperatingSystemQuery) OnlyIDX(ctx context.Context) int {
	id, err := osq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OperatingSystems.
func (osq *OperatingSystemQuery) All(ctx context.Context) ([]*OperatingSystem, error) {
	ctx = setContextOp(ctx, osq.ctx, ent.OpQueryAll)
	if err := osq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OperatingSystem, *OperatingSystemQuery]()
	return withInterceptors[[]*OperatingSystem](ctx, osq, qr, osq.inters)
}

// AllX is like All, but panics if an error occurs.
func (osq *OperatingSystemQuery) AllX(ctx context.Context) []*OperatingSystem {
	nodes, err := osq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OperatingSystem IDs.
func (osq *OperatingSystemQuery) IDs(ctx context.Context) (ids []int, err error) {
	if osq.ctx.Unique == nil && osq.path != nil {
		osq.Unique(true)
	}
	ctx = setContextOp(ctx, osq.ctx, ent.OpQueryIDs)
	if err = osq.Select(operatingsystem.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (osq *OperatingSystemQuery) IDsX(ctx context.Context) []int {
	ids, err := osq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (osq *OperatingSystemQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, osq.ctx, ent.OpQueryCount)
	if err := osq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, osq, querierCount[*OperatingSystemQuery](), osq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (osq *OperatingSystemQuery) CountX(ctx context.Context) int {
	count, err := osq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (osq *OperatingSystemQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, osq.ctx, ent.OpQueryExist)
	switch _, err := osq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (osq *OperatingSystemQuery) ExistX(ctx context.Context) bool {
	exist, err := osq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OperatingSystemQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (osq *OperatingSystemQuery) Clone() *OperatingSystemQuery {
	if osq == nil {
		return nil
	}
	return &OperatingSystemQuery{
		config:     osq.config,
		ctx:        osq.ctx.Clone(),
		order:      append([]operatingsystem.OrderOption{}, osq.order...),
		inters:     append([]Interceptor{}, osq.inters...),
		predicates: append([]predicate.OperatingSystem{}, osq.predicates...),
		withOwner:  osq.withOwner.Clone(),
		// clone intermediate query.
		sql:       osq.sql.Clone(),
		path:      osq.path,
		modifiers: append([]func(*sql.Selector){}, osq.modifiers...),
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (osq *OperatingSystemQuery) WithOwner(opts ...func(*AgentQuery)) *OperatingSystemQuery {
	query := (&AgentClient{config: osq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	osq.withOwner = query
	return osq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Type string `json:"type,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OperatingSystem.Query().
//		GroupBy(operatingsystem.FieldType).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (osq *OperatingSystemQuery) GroupBy(field string, fields ...string) *OperatingSystemGroupBy {
	osq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OperatingSystemGroupBy{build: osq}
	grbuild.flds = &osq.ctx.Fields
	grbuild.label = operatingsystem.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Type string `json:"type,omitempty"`
//	}
//
//	client.OperatingSystem.Query().
//		Select(operatingsystem.FieldType).
//		Scan(ctx, &v)
func (osq *OperatingSystemQuery) Select(fields ...string) *OperatingSystemSelect {
	osq.ctx.Fields = append(osq.ctx.Fields, fields...)
	sbuild := &OperatingSystemSelect{OperatingSystemQuery: osq}
	sbuild.label = operatingsystem.Label
	sbuild.flds, sbuild.scan = &osq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OperatingSystemSelect configured with the given aggregations.
func (osq *OperatingSystemQuery) Aggregate(fns ...AggregateFunc) *OperatingSystemSelect {
	return osq.Select().Aggregate(fns...)
}

func (osq *OperatingSystemQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range osq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, osq); err != nil {
				return err
			}
		}
	}
	for _, f := range osq.ctx.Fields {
		if !operatingsystem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if osq.path != nil {
		prev, err := osq.path(ctx)
		if err != nil {
			return err
		}
		osq.sql = prev
	}
	return nil
}

func (osq *OperatingSystemQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OperatingSystem, error) {
	var (
		nodes       = []*OperatingSystem{}
		withFKs     = osq.withFKs
		_spec       = osq.querySpec()
		loadedTypes = [1]bool{
			osq.withOwner != nil,
		}
	)
	if osq.withOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, operatingsystem.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OperatingSystem).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OperatingSystem{config: osq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(osq.modifiers) > 0 {
		_spec.Modifiers = osq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, osq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := osq.withOwner; query != nil {
		if err := osq.loadOwner(ctx, query, nodes, nil,
			func(n *OperatingSystem, e *Agent) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (osq *OperatingSystemQuery) loadOwner(ctx context.Context, query *AgentQuery, nodes []*OperatingSystem, init func(*OperatingSystem), assign func(*OperatingSystem, *Agent)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*OperatingSystem)
	for i := range nodes {
		if nodes[i].agent_operatingsystem == nil {
			continue
		}
		fk := *nodes[i].agent_operatingsystem
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(agent.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "agent_operatingsystem" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (osq *OperatingSystemQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := osq.querySpec()
	if len(osq.modifiers) > 0 {
		_spec.Modifiers = osq.modifiers
	}
	_spec.Node.Columns = osq.ctx.Fields
	if len(osq.ctx.Fields) > 0 {
		_spec.Unique = osq.ctx.Unique != nil && *osq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, osq.driver, _spec)
}

func (osq *OperatingSystemQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(operatingsystem.Table, operatingsystem.Columns, sqlgraph.NewFieldSpec(operatingsystem.FieldID, field.TypeInt))
	_spec.From = osq.sql
	if unique := osq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if osq.path != nil {
		_spec.Unique = true
	}
	if fields := osq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, operatingsystem.FieldID)
		for i := range fields {
			if fields[i] != operatingsystem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := osq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := osq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := osq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := osq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (osq *OperatingSystemQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(osq.driver.Dialect())
	t1 := builder.Table(operatingsystem.Table)
	columns := osq.ctx.Fields
	if len(columns) == 0 {
		columns = operatingsystem.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if osq.sql != nil {
		selector = osq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if osq.ctx.Unique != nil && *osq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range osq.modifiers {
		m(selector)
	}
	for _, p := range osq.predicates {
		p(selector)
	}
	for _, p := range osq.order {
		p(selector)
	}
	if offset := osq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := osq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (osq *OperatingSystemQuery) Modify(modifiers ...func(s *sql.Selector)) *OperatingSystemSelect {
	osq.modifiers = append(osq.modifiers, modifiers...)
	return osq.Select()
}

// OperatingSystemGroupBy is the group-by builder for OperatingSystem entities.
type OperatingSystemGroupBy struct {
	selector
	build *OperatingSystemQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (osgb *OperatingSystemGroupBy) Aggregate(fns ...AggregateFunc) *OperatingSystemGroupBy {
	osgb.fns = append(osgb.fns, fns...)
	return osgb
}

// Scan applies the selector query and scans the result into the given value.
func (osgb *OperatingSystemGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, osgb.build.ctx, ent.OpQueryGroupBy)
	if err := osgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OperatingSystemQuery, *OperatingSystemGroupBy](ctx, osgb.build, osgb, osgb.build.inters, v)
}

func (osgb *OperatingSystemGroupBy) sqlScan(ctx context.Context, root *OperatingSystemQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(osgb.fns))
	for _, fn := range osgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*osgb.flds)+len(osgb.fns))
		for _, f := range *osgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*osgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := osgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OperatingSystemSelect is the builder for selecting fields of OperatingSystem entities.
type OperatingSystemSelect struct {
	*OperatingSystemQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (oss *OperatingSystemSelect) Aggregate(fns ...AggregateFunc) *OperatingSystemSelect {
	oss.fns = append(oss.fns, fns...)
	return oss
}

// Scan applies the selector query and scans the result into the given value.
func (oss *OperatingSystemSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, oss.ctx, ent.OpQuerySelect)
	if err := oss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OperatingSystemQuery, *OperatingSystemSelect](ctx, oss.OperatingSystemQuery, oss, oss.inters, v)
}

func (oss *OperatingSystemSelect) sqlScan(ctx context.Context, root *OperatingSystemQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(oss.fns))
	for _, fn := range oss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*oss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (oss *OperatingSystemSelect) Modify(modifiers ...func(s *sql.Selector)) *OperatingSystemSelect {
	oss.modifiers = append(oss.modifiers, modifiers...)
	return oss
}
