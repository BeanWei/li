// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/flowdeployment"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/flowinstance"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/flownodeinstance"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/predicate"
)

// FlowInstanceQuery is the builder for querying FlowInstance entities.
type FlowInstanceQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.FlowInstance
	// eager-loading edges.
	withFlowDeployment    *FlowDeploymentQuery
	withFlowNodeInstances *FlowNodeInstanceQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FlowInstanceQuery builder.
func (fiq *FlowInstanceQuery) Where(ps ...predicate.FlowInstance) *FlowInstanceQuery {
	fiq.predicates = append(fiq.predicates, ps...)
	return fiq
}

// Limit adds a limit step to the query.
func (fiq *FlowInstanceQuery) Limit(limit int) *FlowInstanceQuery {
	fiq.limit = &limit
	return fiq
}

// Offset adds an offset step to the query.
func (fiq *FlowInstanceQuery) Offset(offset int) *FlowInstanceQuery {
	fiq.offset = &offset
	return fiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fiq *FlowInstanceQuery) Unique(unique bool) *FlowInstanceQuery {
	fiq.unique = &unique
	return fiq
}

// Order adds an order step to the query.
func (fiq *FlowInstanceQuery) Order(o ...OrderFunc) *FlowInstanceQuery {
	fiq.order = append(fiq.order, o...)
	return fiq
}

// QueryFlowDeployment chains the current query on the "flow_deployment" edge.
func (fiq *FlowInstanceQuery) QueryFlowDeployment() *FlowDeploymentQuery {
	query := &FlowDeploymentQuery{config: fiq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(flowinstance.Table, flowinstance.FieldID, selector),
			sqlgraph.To(flowdeployment.Table, flowdeployment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, flowinstance.FlowDeploymentTable, flowinstance.FlowDeploymentColumn),
		)
		fromU = sqlgraph.SetNeighbors(fiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFlowNodeInstances chains the current query on the "flow_node_instances" edge.
func (fiq *FlowInstanceQuery) QueryFlowNodeInstances() *FlowNodeInstanceQuery {
	query := &FlowNodeInstanceQuery{config: fiq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(flowinstance.Table, flowinstance.FieldID, selector),
			sqlgraph.To(flownodeinstance.Table, flownodeinstance.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, flowinstance.FlowNodeInstancesTable, flowinstance.FlowNodeInstancesColumn),
		)
		fromU = sqlgraph.SetNeighbors(fiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first FlowInstance entity from the query.
// Returns a *NotFoundError when no FlowInstance was found.
func (fiq *FlowInstanceQuery) First(ctx context.Context) (*FlowInstance, error) {
	nodes, err := fiq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{flowinstance.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fiq *FlowInstanceQuery) FirstX(ctx context.Context) *FlowInstance {
	node, err := fiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FlowInstance ID from the query.
// Returns a *NotFoundError when no FlowInstance ID was found.
func (fiq *FlowInstanceQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = fiq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{flowinstance.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fiq *FlowInstanceQuery) FirstIDX(ctx context.Context) string {
	id, err := fiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FlowInstance entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FlowInstance entity is found.
// Returns a *NotFoundError when no FlowInstance entities are found.
func (fiq *FlowInstanceQuery) Only(ctx context.Context) (*FlowInstance, error) {
	nodes, err := fiq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{flowinstance.Label}
	default:
		return nil, &NotSingularError{flowinstance.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fiq *FlowInstanceQuery) OnlyX(ctx context.Context) *FlowInstance {
	node, err := fiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FlowInstance ID in the query.
// Returns a *NotSingularError when more than one FlowInstance ID is found.
// Returns a *NotFoundError when no entities are found.
func (fiq *FlowInstanceQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = fiq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{flowinstance.Label}
	default:
		err = &NotSingularError{flowinstance.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fiq *FlowInstanceQuery) OnlyIDX(ctx context.Context) string {
	id, err := fiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FlowInstances.
func (fiq *FlowInstanceQuery) All(ctx context.Context) ([]*FlowInstance, error) {
	if err := fiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return fiq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fiq *FlowInstanceQuery) AllX(ctx context.Context) []*FlowInstance {
	nodes, err := fiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FlowInstance IDs.
func (fiq *FlowInstanceQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := fiq.Select(flowinstance.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fiq *FlowInstanceQuery) IDsX(ctx context.Context) []string {
	ids, err := fiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fiq *FlowInstanceQuery) Count(ctx context.Context) (int, error) {
	if err := fiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return fiq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fiq *FlowInstanceQuery) CountX(ctx context.Context) int {
	count, err := fiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fiq *FlowInstanceQuery) Exist(ctx context.Context) (bool, error) {
	if err := fiq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return fiq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fiq *FlowInstanceQuery) ExistX(ctx context.Context) bool {
	exist, err := fiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FlowInstanceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fiq *FlowInstanceQuery) Clone() *FlowInstanceQuery {
	if fiq == nil {
		return nil
	}
	return &FlowInstanceQuery{
		config:                fiq.config,
		limit:                 fiq.limit,
		offset:                fiq.offset,
		order:                 append([]OrderFunc{}, fiq.order...),
		predicates:            append([]predicate.FlowInstance{}, fiq.predicates...),
		withFlowDeployment:    fiq.withFlowDeployment.Clone(),
		withFlowNodeInstances: fiq.withFlowNodeInstances.Clone(),
		// clone intermediate query.
		sql:    fiq.sql.Clone(),
		path:   fiq.path,
		unique: fiq.unique,
	}
}

// WithFlowDeployment tells the query-builder to eager-load the nodes that are connected to
// the "flow_deployment" edge. The optional arguments are used to configure the query builder of the edge.
func (fiq *FlowInstanceQuery) WithFlowDeployment(opts ...func(*FlowDeploymentQuery)) *FlowInstanceQuery {
	query := &FlowDeploymentQuery{config: fiq.config}
	for _, opt := range opts {
		opt(query)
	}
	fiq.withFlowDeployment = query
	return fiq
}

// WithFlowNodeInstances tells the query-builder to eager-load the nodes that are connected to
// the "flow_node_instances" edge. The optional arguments are used to configure the query builder of the edge.
func (fiq *FlowInstanceQuery) WithFlowNodeInstances(opts ...func(*FlowNodeInstanceQuery)) *FlowInstanceQuery {
	query := &FlowNodeInstanceQuery{config: fiq.config}
	for _, opt := range opts {
		opt(query)
	}
	fiq.withFlowNodeInstances = query
	return fiq
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
//	client.FlowInstance.Query().
//		GroupBy(flowinstance.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (fiq *FlowInstanceQuery) GroupBy(field string, fields ...string) *FlowInstanceGroupBy {
	grbuild := &FlowInstanceGroupBy{config: fiq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fiq.sqlQuery(ctx), nil
	}
	grbuild.label = flowinstance.Label
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
//	client.FlowInstance.Query().
//		Select(flowinstance.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (fiq *FlowInstanceQuery) Select(fields ...string) *FlowInstanceSelect {
	fiq.fields = append(fiq.fields, fields...)
	selbuild := &FlowInstanceSelect{FlowInstanceQuery: fiq}
	selbuild.label = flowinstance.Label
	selbuild.flds, selbuild.scan = &fiq.fields, selbuild.Scan
	return selbuild
}

func (fiq *FlowInstanceQuery) prepareQuery(ctx context.Context) error {
	for _, f := range fiq.fields {
		if !flowinstance.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fiq.path != nil {
		prev, err := fiq.path(ctx)
		if err != nil {
			return err
		}
		fiq.sql = prev
	}
	return nil
}

func (fiq *FlowInstanceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FlowInstance, error) {
	var (
		nodes       = []*FlowInstance{}
		_spec       = fiq.querySpec()
		loadedTypes = [2]bool{
			fiq.withFlowDeployment != nil,
			fiq.withFlowNodeInstances != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*FlowInstance).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &FlowInstance{config: fiq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := fiq.withFlowDeployment; query != nil {
		ids := make([]string, 0, len(nodes))
		nodeids := make(map[string][]*FlowInstance)
		for i := range nodes {
			fk := nodes[i].FlowDeploymentID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(flowdeployment.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "flow_deployment_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.FlowDeployment = n
			}
		}
	}

	if query := fiq.withFlowNodeInstances; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[string]*FlowInstance)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.FlowNodeInstances = []*FlowNodeInstance{}
		}
		query.Where(predicate.FlowNodeInstance(func(s *sql.Selector) {
			s.Where(sql.InValues(flowinstance.FlowNodeInstancesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.FlowInstanceID
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "flow_instance_id" returned %v for node %v`, fk, n.ID)
			}
			node.Edges.FlowNodeInstances = append(node.Edges.FlowNodeInstances, n)
		}
	}

	return nodes, nil
}

func (fiq *FlowInstanceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fiq.querySpec()
	_spec.Node.Columns = fiq.fields
	if len(fiq.fields) > 0 {
		_spec.Unique = fiq.unique != nil && *fiq.unique
	}
	return sqlgraph.CountNodes(ctx, fiq.driver, _spec)
}

func (fiq *FlowInstanceQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := fiq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (fiq *FlowInstanceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   flowinstance.Table,
			Columns: flowinstance.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: flowinstance.FieldID,
			},
		},
		From:   fiq.sql,
		Unique: true,
	}
	if unique := fiq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := fiq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, flowinstance.FieldID)
		for i := range fields {
			if fields[i] != flowinstance.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fiq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fiq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fiq *FlowInstanceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fiq.driver.Dialect())
	t1 := builder.Table(flowinstance.Table)
	columns := fiq.fields
	if len(columns) == 0 {
		columns = flowinstance.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fiq.sql != nil {
		selector = fiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fiq.unique != nil && *fiq.unique {
		selector.Distinct()
	}
	for _, p := range fiq.predicates {
		p(selector)
	}
	for _, p := range fiq.order {
		p(selector)
	}
	if offset := fiq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fiq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FlowInstanceGroupBy is the group-by builder for FlowInstance entities.
type FlowInstanceGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (figb *FlowInstanceGroupBy) Aggregate(fns ...AggregateFunc) *FlowInstanceGroupBy {
	figb.fns = append(figb.fns, fns...)
	return figb
}

// Scan applies the group-by query and scans the result into the given value.
func (figb *FlowInstanceGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := figb.path(ctx)
	if err != nil {
		return err
	}
	figb.sql = query
	return figb.sqlScan(ctx, v)
}

func (figb *FlowInstanceGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range figb.fields {
		if !flowinstance.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := figb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := figb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (figb *FlowInstanceGroupBy) sqlQuery() *sql.Selector {
	selector := figb.sql.Select()
	aggregation := make([]string, 0, len(figb.fns))
	for _, fn := range figb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(figb.fields)+len(figb.fns))
		for _, f := range figb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(figb.fields...)...)
}

// FlowInstanceSelect is the builder for selecting fields of FlowInstance entities.
type FlowInstanceSelect struct {
	*FlowInstanceQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (fis *FlowInstanceSelect) Scan(ctx context.Context, v interface{}) error {
	if err := fis.prepareQuery(ctx); err != nil {
		return err
	}
	fis.sql = fis.FlowInstanceQuery.sqlQuery(ctx)
	return fis.sqlScan(ctx, v)
}

func (fis *FlowInstanceSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fis.sql.Query()
	if err := fis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
