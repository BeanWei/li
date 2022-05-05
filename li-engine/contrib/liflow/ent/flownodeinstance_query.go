// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/flowinstance"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/flownodeinstance"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/predicate"
)

// FlowNodeInstanceQuery is the builder for querying FlowNodeInstance entities.
type FlowNodeInstanceQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.FlowNodeInstance
	// eager-loading edges.
	withFlowInstance *FlowInstanceQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FlowNodeInstanceQuery builder.
func (fniq *FlowNodeInstanceQuery) Where(ps ...predicate.FlowNodeInstance) *FlowNodeInstanceQuery {
	fniq.predicates = append(fniq.predicates, ps...)
	return fniq
}

// Limit adds a limit step to the query.
func (fniq *FlowNodeInstanceQuery) Limit(limit int) *FlowNodeInstanceQuery {
	fniq.limit = &limit
	return fniq
}

// Offset adds an offset step to the query.
func (fniq *FlowNodeInstanceQuery) Offset(offset int) *FlowNodeInstanceQuery {
	fniq.offset = &offset
	return fniq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fniq *FlowNodeInstanceQuery) Unique(unique bool) *FlowNodeInstanceQuery {
	fniq.unique = &unique
	return fniq
}

// Order adds an order step to the query.
func (fniq *FlowNodeInstanceQuery) Order(o ...OrderFunc) *FlowNodeInstanceQuery {
	fniq.order = append(fniq.order, o...)
	return fniq
}

// QueryFlowInstance chains the current query on the "flow_instance" edge.
func (fniq *FlowNodeInstanceQuery) QueryFlowInstance() *FlowInstanceQuery {
	query := &FlowInstanceQuery{config: fniq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fniq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fniq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(flownodeinstance.Table, flownodeinstance.FieldID, selector),
			sqlgraph.To(flowinstance.Table, flowinstance.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, flownodeinstance.FlowInstanceTable, flownodeinstance.FlowInstanceColumn),
		)
		fromU = sqlgraph.SetNeighbors(fniq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first FlowNodeInstance entity from the query.
// Returns a *NotFoundError when no FlowNodeInstance was found.
func (fniq *FlowNodeInstanceQuery) First(ctx context.Context) (*FlowNodeInstance, error) {
	nodes, err := fniq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{flownodeinstance.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fniq *FlowNodeInstanceQuery) FirstX(ctx context.Context) *FlowNodeInstance {
	node, err := fniq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FlowNodeInstance ID from the query.
// Returns a *NotFoundError when no FlowNodeInstance ID was found.
func (fniq *FlowNodeInstanceQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = fniq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{flownodeinstance.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fniq *FlowNodeInstanceQuery) FirstIDX(ctx context.Context) string {
	id, err := fniq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FlowNodeInstance entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FlowNodeInstance entity is found.
// Returns a *NotFoundError when no FlowNodeInstance entities are found.
func (fniq *FlowNodeInstanceQuery) Only(ctx context.Context) (*FlowNodeInstance, error) {
	nodes, err := fniq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{flownodeinstance.Label}
	default:
		return nil, &NotSingularError{flownodeinstance.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fniq *FlowNodeInstanceQuery) OnlyX(ctx context.Context) *FlowNodeInstance {
	node, err := fniq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FlowNodeInstance ID in the query.
// Returns a *NotSingularError when more than one FlowNodeInstance ID is found.
// Returns a *NotFoundError when no entities are found.
func (fniq *FlowNodeInstanceQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = fniq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{flownodeinstance.Label}
	default:
		err = &NotSingularError{flownodeinstance.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fniq *FlowNodeInstanceQuery) OnlyIDX(ctx context.Context) string {
	id, err := fniq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FlowNodeInstances.
func (fniq *FlowNodeInstanceQuery) All(ctx context.Context) ([]*FlowNodeInstance, error) {
	if err := fniq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return fniq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fniq *FlowNodeInstanceQuery) AllX(ctx context.Context) []*FlowNodeInstance {
	nodes, err := fniq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FlowNodeInstance IDs.
func (fniq *FlowNodeInstanceQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := fniq.Select(flownodeinstance.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fniq *FlowNodeInstanceQuery) IDsX(ctx context.Context) []string {
	ids, err := fniq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fniq *FlowNodeInstanceQuery) Count(ctx context.Context) (int, error) {
	if err := fniq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return fniq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fniq *FlowNodeInstanceQuery) CountX(ctx context.Context) int {
	count, err := fniq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fniq *FlowNodeInstanceQuery) Exist(ctx context.Context) (bool, error) {
	if err := fniq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return fniq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fniq *FlowNodeInstanceQuery) ExistX(ctx context.Context) bool {
	exist, err := fniq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FlowNodeInstanceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fniq *FlowNodeInstanceQuery) Clone() *FlowNodeInstanceQuery {
	if fniq == nil {
		return nil
	}
	return &FlowNodeInstanceQuery{
		config:           fniq.config,
		limit:            fniq.limit,
		offset:           fniq.offset,
		order:            append([]OrderFunc{}, fniq.order...),
		predicates:       append([]predicate.FlowNodeInstance{}, fniq.predicates...),
		withFlowInstance: fniq.withFlowInstance.Clone(),
		// clone intermediate query.
		sql:    fniq.sql.Clone(),
		path:   fniq.path,
		unique: fniq.unique,
	}
}

// WithFlowInstance tells the query-builder to eager-load the nodes that are connected to
// the "flow_instance" edge. The optional arguments are used to configure the query builder of the edge.
func (fniq *FlowNodeInstanceQuery) WithFlowInstance(opts ...func(*FlowInstanceQuery)) *FlowNodeInstanceQuery {
	query := &FlowInstanceQuery{config: fniq.config}
	for _, opt := range opts {
		opt(query)
	}
	fniq.withFlowInstance = query
	return fniq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt int64 `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FlowNodeInstance.Query().
//		GroupBy(flownodeinstance.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (fniq *FlowNodeInstanceQuery) GroupBy(field string, fields ...string) *FlowNodeInstanceGroupBy {
	grbuild := &FlowNodeInstanceGroupBy{config: fniq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fniq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fniq.sqlQuery(ctx), nil
	}
	grbuild.label = flownodeinstance.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt int64 `json:"created_at,omitempty"`
//	}
//
//	client.FlowNodeInstance.Query().
//		Select(flownodeinstance.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (fniq *FlowNodeInstanceQuery) Select(fields ...string) *FlowNodeInstanceSelect {
	fniq.fields = append(fniq.fields, fields...)
	selbuild := &FlowNodeInstanceSelect{FlowNodeInstanceQuery: fniq}
	selbuild.label = flownodeinstance.Label
	selbuild.flds, selbuild.scan = &fniq.fields, selbuild.Scan
	return selbuild
}

func (fniq *FlowNodeInstanceQuery) prepareQuery(ctx context.Context) error {
	for _, f := range fniq.fields {
		if !flownodeinstance.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fniq.path != nil {
		prev, err := fniq.path(ctx)
		if err != nil {
			return err
		}
		fniq.sql = prev
	}
	return nil
}

func (fniq *FlowNodeInstanceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FlowNodeInstance, error) {
	var (
		nodes       = []*FlowNodeInstance{}
		_spec       = fniq.querySpec()
		loadedTypes = [1]bool{
			fniq.withFlowInstance != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*FlowNodeInstance).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &FlowNodeInstance{config: fniq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fniq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := fniq.withFlowInstance; query != nil {
		ids := make([]string, 0, len(nodes))
		nodeids := make(map[string][]*FlowNodeInstance)
		for i := range nodes {
			fk := nodes[i].FlowInstanceID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(flowinstance.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "flow_instance_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.FlowInstance = n
			}
		}
	}

	return nodes, nil
}

func (fniq *FlowNodeInstanceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fniq.querySpec()
	_spec.Node.Columns = fniq.fields
	if len(fniq.fields) > 0 {
		_spec.Unique = fniq.unique != nil && *fniq.unique
	}
	return sqlgraph.CountNodes(ctx, fniq.driver, _spec)
}

func (fniq *FlowNodeInstanceQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := fniq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (fniq *FlowNodeInstanceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   flownodeinstance.Table,
			Columns: flownodeinstance.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: flownodeinstance.FieldID,
			},
		},
		From:   fniq.sql,
		Unique: true,
	}
	if unique := fniq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := fniq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, flownodeinstance.FieldID)
		for i := range fields {
			if fields[i] != flownodeinstance.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fniq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fniq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fniq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fniq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fniq *FlowNodeInstanceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fniq.driver.Dialect())
	t1 := builder.Table(flownodeinstance.Table)
	columns := fniq.fields
	if len(columns) == 0 {
		columns = flownodeinstance.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fniq.sql != nil {
		selector = fniq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fniq.unique != nil && *fniq.unique {
		selector.Distinct()
	}
	for _, p := range fniq.predicates {
		p(selector)
	}
	for _, p := range fniq.order {
		p(selector)
	}
	if offset := fniq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fniq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FlowNodeInstanceGroupBy is the group-by builder for FlowNodeInstance entities.
type FlowNodeInstanceGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fnigb *FlowNodeInstanceGroupBy) Aggregate(fns ...AggregateFunc) *FlowNodeInstanceGroupBy {
	fnigb.fns = append(fnigb.fns, fns...)
	return fnigb
}

// Scan applies the group-by query and scans the result into the given value.
func (fnigb *FlowNodeInstanceGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := fnigb.path(ctx)
	if err != nil {
		return err
	}
	fnigb.sql = query
	return fnigb.sqlScan(ctx, v)
}

func (fnigb *FlowNodeInstanceGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range fnigb.fields {
		if !flownodeinstance.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := fnigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fnigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fnigb *FlowNodeInstanceGroupBy) sqlQuery() *sql.Selector {
	selector := fnigb.sql.Select()
	aggregation := make([]string, 0, len(fnigb.fns))
	for _, fn := range fnigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(fnigb.fields)+len(fnigb.fns))
		for _, f := range fnigb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(fnigb.fields...)...)
}

// FlowNodeInstanceSelect is the builder for selecting fields of FlowNodeInstance entities.
type FlowNodeInstanceSelect struct {
	*FlowNodeInstanceQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (fnis *FlowNodeInstanceSelect) Scan(ctx context.Context, v interface{}) error {
	if err := fnis.prepareQuery(ctx); err != nil {
		return err
	}
	fnis.sql = fnis.FlowNodeInstanceQuery.sqlQuery(ctx)
	return fnis.sqlScan(ctx, v)
}

func (fnis *FlowNodeInstanceSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fnis.sql.Query()
	if err := fnis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
