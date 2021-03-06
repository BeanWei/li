// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/flownodeinstancelog"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/predicate"
)

// FlowNodeInstanceLogQuery is the builder for querying FlowNodeInstanceLog entities.
type FlowNodeInstanceLogQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.FlowNodeInstanceLog
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FlowNodeInstanceLogQuery builder.
func (fnilq *FlowNodeInstanceLogQuery) Where(ps ...predicate.FlowNodeInstanceLog) *FlowNodeInstanceLogQuery {
	fnilq.predicates = append(fnilq.predicates, ps...)
	return fnilq
}

// Limit adds a limit step to the query.
func (fnilq *FlowNodeInstanceLogQuery) Limit(limit int) *FlowNodeInstanceLogQuery {
	fnilq.limit = &limit
	return fnilq
}

// Offset adds an offset step to the query.
func (fnilq *FlowNodeInstanceLogQuery) Offset(offset int) *FlowNodeInstanceLogQuery {
	fnilq.offset = &offset
	return fnilq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fnilq *FlowNodeInstanceLogQuery) Unique(unique bool) *FlowNodeInstanceLogQuery {
	fnilq.unique = &unique
	return fnilq
}

// Order adds an order step to the query.
func (fnilq *FlowNodeInstanceLogQuery) Order(o ...OrderFunc) *FlowNodeInstanceLogQuery {
	fnilq.order = append(fnilq.order, o...)
	return fnilq
}

// First returns the first FlowNodeInstanceLog entity from the query.
// Returns a *NotFoundError when no FlowNodeInstanceLog was found.
func (fnilq *FlowNodeInstanceLogQuery) First(ctx context.Context) (*FlowNodeInstanceLog, error) {
	nodes, err := fnilq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{flownodeinstancelog.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fnilq *FlowNodeInstanceLogQuery) FirstX(ctx context.Context) *FlowNodeInstanceLog {
	node, err := fnilq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FlowNodeInstanceLog ID from the query.
// Returns a *NotFoundError when no FlowNodeInstanceLog ID was found.
func (fnilq *FlowNodeInstanceLogQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = fnilq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{flownodeinstancelog.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fnilq *FlowNodeInstanceLogQuery) FirstIDX(ctx context.Context) string {
	id, err := fnilq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FlowNodeInstanceLog entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FlowNodeInstanceLog entity is found.
// Returns a *NotFoundError when no FlowNodeInstanceLog entities are found.
func (fnilq *FlowNodeInstanceLogQuery) Only(ctx context.Context) (*FlowNodeInstanceLog, error) {
	nodes, err := fnilq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{flownodeinstancelog.Label}
	default:
		return nil, &NotSingularError{flownodeinstancelog.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fnilq *FlowNodeInstanceLogQuery) OnlyX(ctx context.Context) *FlowNodeInstanceLog {
	node, err := fnilq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FlowNodeInstanceLog ID in the query.
// Returns a *NotSingularError when more than one FlowNodeInstanceLog ID is found.
// Returns a *NotFoundError when no entities are found.
func (fnilq *FlowNodeInstanceLogQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = fnilq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{flownodeinstancelog.Label}
	default:
		err = &NotSingularError{flownodeinstancelog.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fnilq *FlowNodeInstanceLogQuery) OnlyIDX(ctx context.Context) string {
	id, err := fnilq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FlowNodeInstanceLogs.
func (fnilq *FlowNodeInstanceLogQuery) All(ctx context.Context) ([]*FlowNodeInstanceLog, error) {
	if err := fnilq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return fnilq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fnilq *FlowNodeInstanceLogQuery) AllX(ctx context.Context) []*FlowNodeInstanceLog {
	nodes, err := fnilq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FlowNodeInstanceLog IDs.
func (fnilq *FlowNodeInstanceLogQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := fnilq.Select(flownodeinstancelog.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fnilq *FlowNodeInstanceLogQuery) IDsX(ctx context.Context) []string {
	ids, err := fnilq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fnilq *FlowNodeInstanceLogQuery) Count(ctx context.Context) (int, error) {
	if err := fnilq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return fnilq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fnilq *FlowNodeInstanceLogQuery) CountX(ctx context.Context) int {
	count, err := fnilq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fnilq *FlowNodeInstanceLogQuery) Exist(ctx context.Context) (bool, error) {
	if err := fnilq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return fnilq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fnilq *FlowNodeInstanceLogQuery) ExistX(ctx context.Context) bool {
	exist, err := fnilq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FlowNodeInstanceLogQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fnilq *FlowNodeInstanceLogQuery) Clone() *FlowNodeInstanceLogQuery {
	if fnilq == nil {
		return nil
	}
	return &FlowNodeInstanceLogQuery{
		config:     fnilq.config,
		limit:      fnilq.limit,
		offset:     fnilq.offset,
		order:      append([]OrderFunc{}, fnilq.order...),
		predicates: append([]predicate.FlowNodeInstanceLog{}, fnilq.predicates...),
		// clone intermediate query.
		sql:    fnilq.sql.Clone(),
		path:   fnilq.path,
		unique: fnilq.unique,
	}
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
//	client.FlowNodeInstanceLog.Query().
//		GroupBy(flownodeinstancelog.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (fnilq *FlowNodeInstanceLogQuery) GroupBy(field string, fields ...string) *FlowNodeInstanceLogGroupBy {
	grbuild := &FlowNodeInstanceLogGroupBy{config: fnilq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fnilq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fnilq.sqlQuery(ctx), nil
	}
	grbuild.label = flownodeinstancelog.Label
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
//	client.FlowNodeInstanceLog.Query().
//		Select(flownodeinstancelog.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (fnilq *FlowNodeInstanceLogQuery) Select(fields ...string) *FlowNodeInstanceLogSelect {
	fnilq.fields = append(fnilq.fields, fields...)
	selbuild := &FlowNodeInstanceLogSelect{FlowNodeInstanceLogQuery: fnilq}
	selbuild.label = flownodeinstancelog.Label
	selbuild.flds, selbuild.scan = &fnilq.fields, selbuild.Scan
	return selbuild
}

func (fnilq *FlowNodeInstanceLogQuery) prepareQuery(ctx context.Context) error {
	for _, f := range fnilq.fields {
		if !flownodeinstancelog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fnilq.path != nil {
		prev, err := fnilq.path(ctx)
		if err != nil {
			return err
		}
		fnilq.sql = prev
	}
	return nil
}

func (fnilq *FlowNodeInstanceLogQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FlowNodeInstanceLog, error) {
	var (
		nodes = []*FlowNodeInstanceLog{}
		_spec = fnilq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*FlowNodeInstanceLog).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &FlowNodeInstanceLog{config: fnilq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fnilq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	return nodes, nil
}

func (fnilq *FlowNodeInstanceLogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fnilq.querySpec()
	_spec.Node.Columns = fnilq.fields
	if len(fnilq.fields) > 0 {
		_spec.Unique = fnilq.unique != nil && *fnilq.unique
	}
	return sqlgraph.CountNodes(ctx, fnilq.driver, _spec)
}

func (fnilq *FlowNodeInstanceLogQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := fnilq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (fnilq *FlowNodeInstanceLogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   flownodeinstancelog.Table,
			Columns: flownodeinstancelog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: flownodeinstancelog.FieldID,
			},
		},
		From:   fnilq.sql,
		Unique: true,
	}
	if unique := fnilq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := fnilq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, flownodeinstancelog.FieldID)
		for i := range fields {
			if fields[i] != flownodeinstancelog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fnilq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fnilq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fnilq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fnilq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fnilq *FlowNodeInstanceLogQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fnilq.driver.Dialect())
	t1 := builder.Table(flownodeinstancelog.Table)
	columns := fnilq.fields
	if len(columns) == 0 {
		columns = flownodeinstancelog.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fnilq.sql != nil {
		selector = fnilq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fnilq.unique != nil && *fnilq.unique {
		selector.Distinct()
	}
	for _, p := range fnilq.predicates {
		p(selector)
	}
	for _, p := range fnilq.order {
		p(selector)
	}
	if offset := fnilq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fnilq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FlowNodeInstanceLogGroupBy is the group-by builder for FlowNodeInstanceLog entities.
type FlowNodeInstanceLogGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fnilgb *FlowNodeInstanceLogGroupBy) Aggregate(fns ...AggregateFunc) *FlowNodeInstanceLogGroupBy {
	fnilgb.fns = append(fnilgb.fns, fns...)
	return fnilgb
}

// Scan applies the group-by query and scans the result into the given value.
func (fnilgb *FlowNodeInstanceLogGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := fnilgb.path(ctx)
	if err != nil {
		return err
	}
	fnilgb.sql = query
	return fnilgb.sqlScan(ctx, v)
}

func (fnilgb *FlowNodeInstanceLogGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range fnilgb.fields {
		if !flownodeinstancelog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := fnilgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fnilgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fnilgb *FlowNodeInstanceLogGroupBy) sqlQuery() *sql.Selector {
	selector := fnilgb.sql.Select()
	aggregation := make([]string, 0, len(fnilgb.fns))
	for _, fn := range fnilgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(fnilgb.fields)+len(fnilgb.fns))
		for _, f := range fnilgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(fnilgb.fields...)...)
}

// FlowNodeInstanceLogSelect is the builder for selecting fields of FlowNodeInstanceLog entities.
type FlowNodeInstanceLogSelect struct {
	*FlowNodeInstanceLogQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (fnils *FlowNodeInstanceLogSelect) Scan(ctx context.Context, v interface{}) error {
	if err := fnils.prepareQuery(ctx); err != nil {
		return err
	}
	fnils.sql = fnils.FlowNodeInstanceLogQuery.sqlQuery(ctx)
	return fnils.sqlScan(ctx, v)
}

func (fnils *FlowNodeInstanceLogSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fnils.sql.Query()
	if err := fnils.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
