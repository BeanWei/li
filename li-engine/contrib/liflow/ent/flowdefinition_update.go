// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/flowdefinition"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/flowdeployment"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/predicate"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/schema"
)

// FlowDefinitionUpdate is the builder for updating FlowDefinition entities.
type FlowDefinitionUpdate struct {
	config
	hooks    []Hook
	mutation *FlowDefinitionMutation
}

// Where appends a list predicates to the FlowDefinitionUpdate builder.
func (fdu *FlowDefinitionUpdate) Where(ps ...predicate.FlowDefinition) *FlowDefinitionUpdate {
	fdu.mutation.Where(ps...)
	return fdu
}

// SetUpdatedAt sets the "updated_at" field.
func (fdu *FlowDefinitionUpdate) SetUpdatedAt(i int64) *FlowDefinitionUpdate {
	fdu.mutation.ResetUpdatedAt()
	fdu.mutation.SetUpdatedAt(i)
	return fdu
}

// AddUpdatedAt adds i to the "updated_at" field.
func (fdu *FlowDefinitionUpdate) AddUpdatedAt(i int64) *FlowDefinitionUpdate {
	fdu.mutation.AddUpdatedAt(i)
	return fdu
}

// SetDeletedAt sets the "deleted_at" field.
func (fdu *FlowDefinitionUpdate) SetDeletedAt(i int64) *FlowDefinitionUpdate {
	fdu.mutation.ResetDeletedAt()
	fdu.mutation.SetDeletedAt(i)
	return fdu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fdu *FlowDefinitionUpdate) SetNillableDeletedAt(i *int64) *FlowDefinitionUpdate {
	if i != nil {
		fdu.SetDeletedAt(*i)
	}
	return fdu
}

// AddDeletedAt adds i to the "deleted_at" field.
func (fdu *FlowDefinitionUpdate) AddDeletedAt(i int64) *FlowDefinitionUpdate {
	fdu.mutation.AddDeletedAt(i)
	return fdu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (fdu *FlowDefinitionUpdate) ClearDeletedAt() *FlowDefinitionUpdate {
	fdu.mutation.ClearDeletedAt()
	return fdu
}

// SetName sets the "name" field.
func (fdu *FlowDefinitionUpdate) SetName(s string) *FlowDefinitionUpdate {
	fdu.mutation.SetName(s)
	return fdu
}

// SetStatus sets the "status" field.
func (fdu *FlowDefinitionUpdate) SetStatus(i int8) *FlowDefinitionUpdate {
	fdu.mutation.ResetStatus()
	fdu.mutation.SetStatus(i)
	return fdu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (fdu *FlowDefinitionUpdate) SetNillableStatus(i *int8) *FlowDefinitionUpdate {
	if i != nil {
		fdu.SetStatus(*i)
	}
	return fdu
}

// AddStatus adds i to the "status" field.
func (fdu *FlowDefinitionUpdate) AddStatus(i int8) *FlowDefinitionUpdate {
	fdu.mutation.AddStatus(i)
	return fdu
}

// SetModel sets the "model" field.
func (fdu *FlowDefinitionUpdate) SetModel(sm schema.FlowModel) *FlowDefinitionUpdate {
	fdu.mutation.SetModel(sm)
	return fdu
}

// SetRemark sets the "remark" field.
func (fdu *FlowDefinitionUpdate) SetRemark(s string) *FlowDefinitionUpdate {
	fdu.mutation.SetRemark(s)
	return fdu
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (fdu *FlowDefinitionUpdate) SetNillableRemark(s *string) *FlowDefinitionUpdate {
	if s != nil {
		fdu.SetRemark(*s)
	}
	return fdu
}

// ClearRemark clears the value of the "remark" field.
func (fdu *FlowDefinitionUpdate) ClearRemark() *FlowDefinitionUpdate {
	fdu.mutation.ClearRemark()
	return fdu
}

// AddFlowDeploymentIDs adds the "flow_deployments" edge to the FlowDeployment entity by IDs.
func (fdu *FlowDefinitionUpdate) AddFlowDeploymentIDs(ids ...string) *FlowDefinitionUpdate {
	fdu.mutation.AddFlowDeploymentIDs(ids...)
	return fdu
}

// AddFlowDeployments adds the "flow_deployments" edges to the FlowDeployment entity.
func (fdu *FlowDefinitionUpdate) AddFlowDeployments(f ...*FlowDeployment) *FlowDefinitionUpdate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fdu.AddFlowDeploymentIDs(ids...)
}

// Mutation returns the FlowDefinitionMutation object of the builder.
func (fdu *FlowDefinitionUpdate) Mutation() *FlowDefinitionMutation {
	return fdu.mutation
}

// ClearFlowDeployments clears all "flow_deployments" edges to the FlowDeployment entity.
func (fdu *FlowDefinitionUpdate) ClearFlowDeployments() *FlowDefinitionUpdate {
	fdu.mutation.ClearFlowDeployments()
	return fdu
}

// RemoveFlowDeploymentIDs removes the "flow_deployments" edge to FlowDeployment entities by IDs.
func (fdu *FlowDefinitionUpdate) RemoveFlowDeploymentIDs(ids ...string) *FlowDefinitionUpdate {
	fdu.mutation.RemoveFlowDeploymentIDs(ids...)
	return fdu
}

// RemoveFlowDeployments removes "flow_deployments" edges to FlowDeployment entities.
func (fdu *FlowDefinitionUpdate) RemoveFlowDeployments(f ...*FlowDeployment) *FlowDefinitionUpdate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fdu.RemoveFlowDeploymentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fdu *FlowDefinitionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	fdu.defaults()
	if len(fdu.hooks) == 0 {
		if err = fdu.check(); err != nil {
			return 0, err
		}
		affected, err = fdu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FlowDefinitionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fdu.check(); err != nil {
				return 0, err
			}
			fdu.mutation = mutation
			affected, err = fdu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fdu.hooks) - 1; i >= 0; i-- {
			if fdu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fdu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fdu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fdu *FlowDefinitionUpdate) SaveX(ctx context.Context) int {
	affected, err := fdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fdu *FlowDefinitionUpdate) Exec(ctx context.Context) error {
	_, err := fdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fdu *FlowDefinitionUpdate) ExecX(ctx context.Context) {
	if err := fdu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fdu *FlowDefinitionUpdate) defaults() {
	if _, ok := fdu.mutation.UpdatedAt(); !ok {
		v := flowdefinition.UpdateDefaultUpdatedAt()
		fdu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fdu *FlowDefinitionUpdate) check() error {
	if v, ok := fdu.mutation.Name(); ok {
		if err := flowdefinition.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "FlowDefinition.name": %w`, err)}
		}
	}
	return nil
}

func (fdu *FlowDefinitionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   flowdefinition.Table,
			Columns: flowdefinition.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: flowdefinition.FieldID,
			},
		},
	}
	if ps := fdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fdu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: flowdefinition.FieldUpdatedAt,
		})
	}
	if value, ok := fdu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: flowdefinition.FieldUpdatedAt,
		})
	}
	if value, ok := fdu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: flowdefinition.FieldDeletedAt,
		})
	}
	if value, ok := fdu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: flowdefinition.FieldDeletedAt,
		})
	}
	if fdu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: flowdefinition.FieldDeletedAt,
		})
	}
	if value, ok := fdu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowdefinition.FieldName,
		})
	}
	if value, ok := fdu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: flowdefinition.FieldStatus,
		})
	}
	if value, ok := fdu.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: flowdefinition.FieldStatus,
		})
	}
	if value, ok := fdu.mutation.Model(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: flowdefinition.FieldModel,
		})
	}
	if value, ok := fdu.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowdefinition.FieldRemark,
		})
	}
	if fdu.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: flowdefinition.FieldRemark,
		})
	}
	if fdu.mutation.FlowDeploymentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flowdefinition.FlowDeploymentsTable,
			Columns: []string{flowdefinition.FlowDeploymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: flowdeployment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fdu.mutation.RemovedFlowDeploymentsIDs(); len(nodes) > 0 && !fdu.mutation.FlowDeploymentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flowdefinition.FlowDeploymentsTable,
			Columns: []string{flowdefinition.FlowDeploymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: flowdeployment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fdu.mutation.FlowDeploymentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flowdefinition.FlowDeploymentsTable,
			Columns: []string{flowdefinition.FlowDeploymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: flowdeployment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{flowdefinition.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// FlowDefinitionUpdateOne is the builder for updating a single FlowDefinition entity.
type FlowDefinitionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FlowDefinitionMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (fduo *FlowDefinitionUpdateOne) SetUpdatedAt(i int64) *FlowDefinitionUpdateOne {
	fduo.mutation.ResetUpdatedAt()
	fduo.mutation.SetUpdatedAt(i)
	return fduo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (fduo *FlowDefinitionUpdateOne) AddUpdatedAt(i int64) *FlowDefinitionUpdateOne {
	fduo.mutation.AddUpdatedAt(i)
	return fduo
}

// SetDeletedAt sets the "deleted_at" field.
func (fduo *FlowDefinitionUpdateOne) SetDeletedAt(i int64) *FlowDefinitionUpdateOne {
	fduo.mutation.ResetDeletedAt()
	fduo.mutation.SetDeletedAt(i)
	return fduo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fduo *FlowDefinitionUpdateOne) SetNillableDeletedAt(i *int64) *FlowDefinitionUpdateOne {
	if i != nil {
		fduo.SetDeletedAt(*i)
	}
	return fduo
}

// AddDeletedAt adds i to the "deleted_at" field.
func (fduo *FlowDefinitionUpdateOne) AddDeletedAt(i int64) *FlowDefinitionUpdateOne {
	fduo.mutation.AddDeletedAt(i)
	return fduo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (fduo *FlowDefinitionUpdateOne) ClearDeletedAt() *FlowDefinitionUpdateOne {
	fduo.mutation.ClearDeletedAt()
	return fduo
}

// SetName sets the "name" field.
func (fduo *FlowDefinitionUpdateOne) SetName(s string) *FlowDefinitionUpdateOne {
	fduo.mutation.SetName(s)
	return fduo
}

// SetStatus sets the "status" field.
func (fduo *FlowDefinitionUpdateOne) SetStatus(i int8) *FlowDefinitionUpdateOne {
	fduo.mutation.ResetStatus()
	fduo.mutation.SetStatus(i)
	return fduo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (fduo *FlowDefinitionUpdateOne) SetNillableStatus(i *int8) *FlowDefinitionUpdateOne {
	if i != nil {
		fduo.SetStatus(*i)
	}
	return fduo
}

// AddStatus adds i to the "status" field.
func (fduo *FlowDefinitionUpdateOne) AddStatus(i int8) *FlowDefinitionUpdateOne {
	fduo.mutation.AddStatus(i)
	return fduo
}

// SetModel sets the "model" field.
func (fduo *FlowDefinitionUpdateOne) SetModel(sm schema.FlowModel) *FlowDefinitionUpdateOne {
	fduo.mutation.SetModel(sm)
	return fduo
}

// SetRemark sets the "remark" field.
func (fduo *FlowDefinitionUpdateOne) SetRemark(s string) *FlowDefinitionUpdateOne {
	fduo.mutation.SetRemark(s)
	return fduo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (fduo *FlowDefinitionUpdateOne) SetNillableRemark(s *string) *FlowDefinitionUpdateOne {
	if s != nil {
		fduo.SetRemark(*s)
	}
	return fduo
}

// ClearRemark clears the value of the "remark" field.
func (fduo *FlowDefinitionUpdateOne) ClearRemark() *FlowDefinitionUpdateOne {
	fduo.mutation.ClearRemark()
	return fduo
}

// AddFlowDeploymentIDs adds the "flow_deployments" edge to the FlowDeployment entity by IDs.
func (fduo *FlowDefinitionUpdateOne) AddFlowDeploymentIDs(ids ...string) *FlowDefinitionUpdateOne {
	fduo.mutation.AddFlowDeploymentIDs(ids...)
	return fduo
}

// AddFlowDeployments adds the "flow_deployments" edges to the FlowDeployment entity.
func (fduo *FlowDefinitionUpdateOne) AddFlowDeployments(f ...*FlowDeployment) *FlowDefinitionUpdateOne {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fduo.AddFlowDeploymentIDs(ids...)
}

// Mutation returns the FlowDefinitionMutation object of the builder.
func (fduo *FlowDefinitionUpdateOne) Mutation() *FlowDefinitionMutation {
	return fduo.mutation
}

// ClearFlowDeployments clears all "flow_deployments" edges to the FlowDeployment entity.
func (fduo *FlowDefinitionUpdateOne) ClearFlowDeployments() *FlowDefinitionUpdateOne {
	fduo.mutation.ClearFlowDeployments()
	return fduo
}

// RemoveFlowDeploymentIDs removes the "flow_deployments" edge to FlowDeployment entities by IDs.
func (fduo *FlowDefinitionUpdateOne) RemoveFlowDeploymentIDs(ids ...string) *FlowDefinitionUpdateOne {
	fduo.mutation.RemoveFlowDeploymentIDs(ids...)
	return fduo
}

// RemoveFlowDeployments removes "flow_deployments" edges to FlowDeployment entities.
func (fduo *FlowDefinitionUpdateOne) RemoveFlowDeployments(f ...*FlowDeployment) *FlowDefinitionUpdateOne {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fduo.RemoveFlowDeploymentIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fduo *FlowDefinitionUpdateOne) Select(field string, fields ...string) *FlowDefinitionUpdateOne {
	fduo.fields = append([]string{field}, fields...)
	return fduo
}

// Save executes the query and returns the updated FlowDefinition entity.
func (fduo *FlowDefinitionUpdateOne) Save(ctx context.Context) (*FlowDefinition, error) {
	var (
		err  error
		node *FlowDefinition
	)
	fduo.defaults()
	if len(fduo.hooks) == 0 {
		if err = fduo.check(); err != nil {
			return nil, err
		}
		node, err = fduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FlowDefinitionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fduo.check(); err != nil {
				return nil, err
			}
			fduo.mutation = mutation
			node, err = fduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fduo.hooks) - 1; i >= 0; i-- {
			if fduo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fduo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fduo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fduo *FlowDefinitionUpdateOne) SaveX(ctx context.Context) *FlowDefinition {
	node, err := fduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fduo *FlowDefinitionUpdateOne) Exec(ctx context.Context) error {
	_, err := fduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fduo *FlowDefinitionUpdateOne) ExecX(ctx context.Context) {
	if err := fduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fduo *FlowDefinitionUpdateOne) defaults() {
	if _, ok := fduo.mutation.UpdatedAt(); !ok {
		v := flowdefinition.UpdateDefaultUpdatedAt()
		fduo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fduo *FlowDefinitionUpdateOne) check() error {
	if v, ok := fduo.mutation.Name(); ok {
		if err := flowdefinition.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "FlowDefinition.name": %w`, err)}
		}
	}
	return nil
}

func (fduo *FlowDefinitionUpdateOne) sqlSave(ctx context.Context) (_node *FlowDefinition, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   flowdefinition.Table,
			Columns: flowdefinition.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: flowdefinition.FieldID,
			},
		},
	}
	id, ok := fduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "FlowDefinition.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, flowdefinition.FieldID)
		for _, f := range fields {
			if !flowdefinition.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != flowdefinition.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fduo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: flowdefinition.FieldUpdatedAt,
		})
	}
	if value, ok := fduo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: flowdefinition.FieldUpdatedAt,
		})
	}
	if value, ok := fduo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: flowdefinition.FieldDeletedAt,
		})
	}
	if value, ok := fduo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: flowdefinition.FieldDeletedAt,
		})
	}
	if fduo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: flowdefinition.FieldDeletedAt,
		})
	}
	if value, ok := fduo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowdefinition.FieldName,
		})
	}
	if value, ok := fduo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: flowdefinition.FieldStatus,
		})
	}
	if value, ok := fduo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: flowdefinition.FieldStatus,
		})
	}
	if value, ok := fduo.mutation.Model(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: flowdefinition.FieldModel,
		})
	}
	if value, ok := fduo.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowdefinition.FieldRemark,
		})
	}
	if fduo.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: flowdefinition.FieldRemark,
		})
	}
	if fduo.mutation.FlowDeploymentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flowdefinition.FlowDeploymentsTable,
			Columns: []string{flowdefinition.FlowDeploymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: flowdeployment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fduo.mutation.RemovedFlowDeploymentsIDs(); len(nodes) > 0 && !fduo.mutation.FlowDeploymentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flowdefinition.FlowDeploymentsTable,
			Columns: []string{flowdefinition.FlowDeploymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: flowdeployment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fduo.mutation.FlowDeploymentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flowdefinition.FlowDeploymentsTable,
			Columns: []string{flowdefinition.FlowDeploymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: flowdeployment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &FlowDefinition{config: fduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{flowdefinition.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
