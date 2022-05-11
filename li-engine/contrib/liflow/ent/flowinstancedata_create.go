// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/flowinstancedata"
)

// FlowInstanceDataCreate is the builder for creating a FlowInstanceData entity.
type FlowInstanceDataCreate struct {
	config
	mutation *FlowInstanceDataMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (fidc *FlowInstanceDataCreate) SetCreatedAt(i int64) *FlowInstanceDataCreate {
	fidc.mutation.SetCreatedAt(i)
	return fidc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fidc *FlowInstanceDataCreate) SetNillableCreatedAt(i *int64) *FlowInstanceDataCreate {
	if i != nil {
		fidc.SetCreatedAt(*i)
	}
	return fidc
}

// SetUpdatedAt sets the "updated_at" field.
func (fidc *FlowInstanceDataCreate) SetUpdatedAt(i int64) *FlowInstanceDataCreate {
	fidc.mutation.SetUpdatedAt(i)
	return fidc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fidc *FlowInstanceDataCreate) SetNillableUpdatedAt(i *int64) *FlowInstanceDataCreate {
	if i != nil {
		fidc.SetUpdatedAt(*i)
	}
	return fidc
}

// SetDeletedAt sets the "deleted_at" field.
func (fidc *FlowInstanceDataCreate) SetDeletedAt(i int64) *FlowInstanceDataCreate {
	fidc.mutation.SetDeletedAt(i)
	return fidc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fidc *FlowInstanceDataCreate) SetNillableDeletedAt(i *int64) *FlowInstanceDataCreate {
	if i != nil {
		fidc.SetDeletedAt(*i)
	}
	return fidc
}

// SetFlowInstanceID sets the "flow_instance_id" field.
func (fidc *FlowInstanceDataCreate) SetFlowInstanceID(s string) *FlowInstanceDataCreate {
	fidc.mutation.SetFlowInstanceID(s)
	return fidc
}

// SetFlowNodeInstanceID sets the "flow_node_instance_id" field.
func (fidc *FlowInstanceDataCreate) SetFlowNodeInstanceID(s string) *FlowInstanceDataCreate {
	fidc.mutation.SetFlowNodeInstanceID(s)
	return fidc
}

// SetNillableFlowNodeInstanceID sets the "flow_node_instance_id" field if the given value is not nil.
func (fidc *FlowInstanceDataCreate) SetNillableFlowNodeInstanceID(s *string) *FlowInstanceDataCreate {
	if s != nil {
		fidc.SetFlowNodeInstanceID(*s)
	}
	return fidc
}

// SetNodeKey sets the "node_key" field.
func (fidc *FlowInstanceDataCreate) SetNodeKey(s string) *FlowInstanceDataCreate {
	fidc.mutation.SetNodeKey(s)
	return fidc
}

// SetNillableNodeKey sets the "node_key" field if the given value is not nil.
func (fidc *FlowInstanceDataCreate) SetNillableNodeKey(s *string) *FlowInstanceDataCreate {
	if s != nil {
		fidc.SetNodeKey(*s)
	}
	return fidc
}

// SetData sets the "data" field.
func (fidc *FlowInstanceDataCreate) SetData(m map[string]interface{}) *FlowInstanceDataCreate {
	fidc.mutation.SetData(m)
	return fidc
}

// SetType sets the "type" field.
func (fidc *FlowInstanceDataCreate) SetType(i int8) *FlowInstanceDataCreate {
	fidc.mutation.SetType(i)
	return fidc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (fidc *FlowInstanceDataCreate) SetNillableType(i *int8) *FlowInstanceDataCreate {
	if i != nil {
		fidc.SetType(*i)
	}
	return fidc
}

// SetID sets the "id" field.
func (fidc *FlowInstanceDataCreate) SetID(s string) *FlowInstanceDataCreate {
	fidc.mutation.SetID(s)
	return fidc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (fidc *FlowInstanceDataCreate) SetNillableID(s *string) *FlowInstanceDataCreate {
	if s != nil {
		fidc.SetID(*s)
	}
	return fidc
}

// Mutation returns the FlowInstanceDataMutation object of the builder.
func (fidc *FlowInstanceDataCreate) Mutation() *FlowInstanceDataMutation {
	return fidc.mutation
}

// Save creates the FlowInstanceData in the database.
func (fidc *FlowInstanceDataCreate) Save(ctx context.Context) (*FlowInstanceData, error) {
	var (
		err  error
		node *FlowInstanceData
	)
	fidc.defaults()
	if len(fidc.hooks) == 0 {
		if err = fidc.check(); err != nil {
			return nil, err
		}
		node, err = fidc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FlowInstanceDataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fidc.check(); err != nil {
				return nil, err
			}
			fidc.mutation = mutation
			if node, err = fidc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(fidc.hooks) - 1; i >= 0; i-- {
			if fidc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fidc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fidc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fidc *FlowInstanceDataCreate) SaveX(ctx context.Context) *FlowInstanceData {
	v, err := fidc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fidc *FlowInstanceDataCreate) Exec(ctx context.Context) error {
	_, err := fidc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fidc *FlowInstanceDataCreate) ExecX(ctx context.Context) {
	if err := fidc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fidc *FlowInstanceDataCreate) defaults() {
	if _, ok := fidc.mutation.CreatedAt(); !ok {
		v := flowinstancedata.DefaultCreatedAt()
		fidc.mutation.SetCreatedAt(v)
	}
	if _, ok := fidc.mutation.UpdatedAt(); !ok {
		v := flowinstancedata.DefaultUpdatedAt()
		fidc.mutation.SetUpdatedAt(v)
	}
	if _, ok := fidc.mutation.DeletedAt(); !ok {
		v := flowinstancedata.DefaultDeletedAt
		fidc.mutation.SetDeletedAt(v)
	}
	if _, ok := fidc.mutation.GetType(); !ok {
		v := flowinstancedata.DefaultType
		fidc.mutation.SetType(v)
	}
	if _, ok := fidc.mutation.ID(); !ok {
		v := flowinstancedata.DefaultID()
		fidc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fidc *FlowInstanceDataCreate) check() error {
	if _, ok := fidc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "FlowInstanceData.created_at"`)}
	}
	if _, ok := fidc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "FlowInstanceData.updated_at"`)}
	}
	if _, ok := fidc.mutation.FlowInstanceID(); !ok {
		return &ValidationError{Name: "flow_instance_id", err: errors.New(`ent: missing required field "FlowInstanceData.flow_instance_id"`)}
	}
	if v, ok := fidc.mutation.FlowInstanceID(); ok {
		if err := flowinstancedata.FlowInstanceIDValidator(v); err != nil {
			return &ValidationError{Name: "flow_instance_id", err: fmt.Errorf(`ent: validator failed for field "FlowInstanceData.flow_instance_id": %w`, err)}
		}
	}
	if _, ok := fidc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "FlowInstanceData.type"`)}
	}
	if v, ok := fidc.mutation.ID(); ok {
		if err := flowinstancedata.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "FlowInstanceData.id": %w`, err)}
		}
	}
	return nil
}

func (fidc *FlowInstanceDataCreate) sqlSave(ctx context.Context) (*FlowInstanceData, error) {
	_node, _spec := fidc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fidc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected FlowInstanceData.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (fidc *FlowInstanceDataCreate) createSpec() (*FlowInstanceData, *sqlgraph.CreateSpec) {
	var (
		_node = &FlowInstanceData{config: fidc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: flowinstancedata.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: flowinstancedata.FieldID,
			},
		}
	)
	_spec.OnConflict = fidc.conflict
	if id, ok := fidc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := fidc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: flowinstancedata.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := fidc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: flowinstancedata.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := fidc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: flowinstancedata.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := fidc.mutation.FlowInstanceID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowinstancedata.FieldFlowInstanceID,
		})
		_node.FlowInstanceID = value
	}
	if value, ok := fidc.mutation.FlowNodeInstanceID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowinstancedata.FieldFlowNodeInstanceID,
		})
		_node.FlowNodeInstanceID = value
	}
	if value, ok := fidc.mutation.NodeKey(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: flowinstancedata.FieldNodeKey,
		})
		_node.NodeKey = value
	}
	if value, ok := fidc.mutation.Data(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: flowinstancedata.FieldData,
		})
		_node.Data = value
	}
	if value, ok := fidc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: flowinstancedata.FieldType,
		})
		_node.Type = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.FlowInstanceData.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FlowInstanceDataUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (fidc *FlowInstanceDataCreate) OnConflict(opts ...sql.ConflictOption) *FlowInstanceDataUpsertOne {
	fidc.conflict = opts
	return &FlowInstanceDataUpsertOne{
		create: fidc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.FlowInstanceData.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (fidc *FlowInstanceDataCreate) OnConflictColumns(columns ...string) *FlowInstanceDataUpsertOne {
	fidc.conflict = append(fidc.conflict, sql.ConflictColumns(columns...))
	return &FlowInstanceDataUpsertOne{
		create: fidc,
	}
}

type (
	// FlowInstanceDataUpsertOne is the builder for "upsert"-ing
	//  one FlowInstanceData node.
	FlowInstanceDataUpsertOne struct {
		create *FlowInstanceDataCreate
	}

	// FlowInstanceDataUpsert is the "OnConflict" setter.
	FlowInstanceDataUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *FlowInstanceDataUpsert) SetCreatedAt(v int64) *FlowInstanceDataUpsert {
	u.Set(flowinstancedata.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FlowInstanceDataUpsert) UpdateCreatedAt() *FlowInstanceDataUpsert {
	u.SetExcluded(flowinstancedata.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *FlowInstanceDataUpsert) AddCreatedAt(v int64) *FlowInstanceDataUpsert {
	u.Add(flowinstancedata.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FlowInstanceDataUpsert) SetUpdatedAt(v int64) *FlowInstanceDataUpsert {
	u.Set(flowinstancedata.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FlowInstanceDataUpsert) UpdateUpdatedAt() *FlowInstanceDataUpsert {
	u.SetExcluded(flowinstancedata.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *FlowInstanceDataUpsert) AddUpdatedAt(v int64) *FlowInstanceDataUpsert {
	u.Add(flowinstancedata.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *FlowInstanceDataUpsert) SetDeletedAt(v int64) *FlowInstanceDataUpsert {
	u.Set(flowinstancedata.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *FlowInstanceDataUpsert) UpdateDeletedAt() *FlowInstanceDataUpsert {
	u.SetExcluded(flowinstancedata.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *FlowInstanceDataUpsert) AddDeletedAt(v int64) *FlowInstanceDataUpsert {
	u.Add(flowinstancedata.FieldDeletedAt, v)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *FlowInstanceDataUpsert) ClearDeletedAt() *FlowInstanceDataUpsert {
	u.SetNull(flowinstancedata.FieldDeletedAt)
	return u
}

// SetFlowInstanceID sets the "flow_instance_id" field.
func (u *FlowInstanceDataUpsert) SetFlowInstanceID(v string) *FlowInstanceDataUpsert {
	u.Set(flowinstancedata.FieldFlowInstanceID, v)
	return u
}

// UpdateFlowInstanceID sets the "flow_instance_id" field to the value that was provided on create.
func (u *FlowInstanceDataUpsert) UpdateFlowInstanceID() *FlowInstanceDataUpsert {
	u.SetExcluded(flowinstancedata.FieldFlowInstanceID)
	return u
}

// SetFlowNodeInstanceID sets the "flow_node_instance_id" field.
func (u *FlowInstanceDataUpsert) SetFlowNodeInstanceID(v string) *FlowInstanceDataUpsert {
	u.Set(flowinstancedata.FieldFlowNodeInstanceID, v)
	return u
}

// UpdateFlowNodeInstanceID sets the "flow_node_instance_id" field to the value that was provided on create.
func (u *FlowInstanceDataUpsert) UpdateFlowNodeInstanceID() *FlowInstanceDataUpsert {
	u.SetExcluded(flowinstancedata.FieldFlowNodeInstanceID)
	return u
}

// ClearFlowNodeInstanceID clears the value of the "flow_node_instance_id" field.
func (u *FlowInstanceDataUpsert) ClearFlowNodeInstanceID() *FlowInstanceDataUpsert {
	u.SetNull(flowinstancedata.FieldFlowNodeInstanceID)
	return u
}

// SetNodeKey sets the "node_key" field.
func (u *FlowInstanceDataUpsert) SetNodeKey(v string) *FlowInstanceDataUpsert {
	u.Set(flowinstancedata.FieldNodeKey, v)
	return u
}

// UpdateNodeKey sets the "node_key" field to the value that was provided on create.
func (u *FlowInstanceDataUpsert) UpdateNodeKey() *FlowInstanceDataUpsert {
	u.SetExcluded(flowinstancedata.FieldNodeKey)
	return u
}

// ClearNodeKey clears the value of the "node_key" field.
func (u *FlowInstanceDataUpsert) ClearNodeKey() *FlowInstanceDataUpsert {
	u.SetNull(flowinstancedata.FieldNodeKey)
	return u
}

// SetData sets the "data" field.
func (u *FlowInstanceDataUpsert) SetData(v map[string]interface{}) *FlowInstanceDataUpsert {
	u.Set(flowinstancedata.FieldData, v)
	return u
}

// UpdateData sets the "data" field to the value that was provided on create.
func (u *FlowInstanceDataUpsert) UpdateData() *FlowInstanceDataUpsert {
	u.SetExcluded(flowinstancedata.FieldData)
	return u
}

// ClearData clears the value of the "data" field.
func (u *FlowInstanceDataUpsert) ClearData() *FlowInstanceDataUpsert {
	u.SetNull(flowinstancedata.FieldData)
	return u
}

// SetType sets the "type" field.
func (u *FlowInstanceDataUpsert) SetType(v int8) *FlowInstanceDataUpsert {
	u.Set(flowinstancedata.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *FlowInstanceDataUpsert) UpdateType() *FlowInstanceDataUpsert {
	u.SetExcluded(flowinstancedata.FieldType)
	return u
}

// AddType adds v to the "type" field.
func (u *FlowInstanceDataUpsert) AddType(v int8) *FlowInstanceDataUpsert {
	u.Add(flowinstancedata.FieldType, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.FlowInstanceData.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(flowinstancedata.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *FlowInstanceDataUpsertOne) UpdateNewValues() *FlowInstanceDataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(flowinstancedata.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(flowinstancedata.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.FlowInstanceData.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *FlowInstanceDataUpsertOne) Ignore() *FlowInstanceDataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FlowInstanceDataUpsertOne) DoNothing() *FlowInstanceDataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FlowInstanceDataCreate.OnConflict
// documentation for more info.
func (u *FlowInstanceDataUpsertOne) Update(set func(*FlowInstanceDataUpsert)) *FlowInstanceDataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FlowInstanceDataUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *FlowInstanceDataUpsertOne) SetCreatedAt(v int64) *FlowInstanceDataUpsertOne {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *FlowInstanceDataUpsertOne) AddCreatedAt(v int64) *FlowInstanceDataUpsertOne {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FlowInstanceDataUpsertOne) UpdateCreatedAt() *FlowInstanceDataUpsertOne {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FlowInstanceDataUpsertOne) SetUpdatedAt(v int64) *FlowInstanceDataUpsertOne {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *FlowInstanceDataUpsertOne) AddUpdatedAt(v int64) *FlowInstanceDataUpsertOne {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FlowInstanceDataUpsertOne) UpdateUpdatedAt() *FlowInstanceDataUpsertOne {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *FlowInstanceDataUpsertOne) SetDeletedAt(v int64) *FlowInstanceDataUpsertOne {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *FlowInstanceDataUpsertOne) AddDeletedAt(v int64) *FlowInstanceDataUpsertOne {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *FlowInstanceDataUpsertOne) UpdateDeletedAt() *FlowInstanceDataUpsertOne {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *FlowInstanceDataUpsertOne) ClearDeletedAt() *FlowInstanceDataUpsertOne {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.ClearDeletedAt()
	})
}

// SetFlowInstanceID sets the "flow_instance_id" field.
func (u *FlowInstanceDataUpsertOne) SetFlowInstanceID(v string) *FlowInstanceDataUpsertOne {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.SetFlowInstanceID(v)
	})
}

// UpdateFlowInstanceID sets the "flow_instance_id" field to the value that was provided on create.
func (u *FlowInstanceDataUpsertOne) UpdateFlowInstanceID() *FlowInstanceDataUpsertOne {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.UpdateFlowInstanceID()
	})
}

// SetFlowNodeInstanceID sets the "flow_node_instance_id" field.
func (u *FlowInstanceDataUpsertOne) SetFlowNodeInstanceID(v string) *FlowInstanceDataUpsertOne {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.SetFlowNodeInstanceID(v)
	})
}

// UpdateFlowNodeInstanceID sets the "flow_node_instance_id" field to the value that was provided on create.
func (u *FlowInstanceDataUpsertOne) UpdateFlowNodeInstanceID() *FlowInstanceDataUpsertOne {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.UpdateFlowNodeInstanceID()
	})
}

// ClearFlowNodeInstanceID clears the value of the "flow_node_instance_id" field.
func (u *FlowInstanceDataUpsertOne) ClearFlowNodeInstanceID() *FlowInstanceDataUpsertOne {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.ClearFlowNodeInstanceID()
	})
}

// SetNodeKey sets the "node_key" field.
func (u *FlowInstanceDataUpsertOne) SetNodeKey(v string) *FlowInstanceDataUpsertOne {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.SetNodeKey(v)
	})
}

// UpdateNodeKey sets the "node_key" field to the value that was provided on create.
func (u *FlowInstanceDataUpsertOne) UpdateNodeKey() *FlowInstanceDataUpsertOne {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.UpdateNodeKey()
	})
}

// ClearNodeKey clears the value of the "node_key" field.
func (u *FlowInstanceDataUpsertOne) ClearNodeKey() *FlowInstanceDataUpsertOne {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.ClearNodeKey()
	})
}

// SetData sets the "data" field.
func (u *FlowInstanceDataUpsertOne) SetData(v map[string]interface{}) *FlowInstanceDataUpsertOne {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.SetData(v)
	})
}

// UpdateData sets the "data" field to the value that was provided on create.
func (u *FlowInstanceDataUpsertOne) UpdateData() *FlowInstanceDataUpsertOne {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.UpdateData()
	})
}

// ClearData clears the value of the "data" field.
func (u *FlowInstanceDataUpsertOne) ClearData() *FlowInstanceDataUpsertOne {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.ClearData()
	})
}

// SetType sets the "type" field.
func (u *FlowInstanceDataUpsertOne) SetType(v int8) *FlowInstanceDataUpsertOne {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.SetType(v)
	})
}

// AddType adds v to the "type" field.
func (u *FlowInstanceDataUpsertOne) AddType(v int8) *FlowInstanceDataUpsertOne {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.AddType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *FlowInstanceDataUpsertOne) UpdateType() *FlowInstanceDataUpsertOne {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.UpdateType()
	})
}

// Exec executes the query.
func (u *FlowInstanceDataUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FlowInstanceDataCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FlowInstanceDataUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *FlowInstanceDataUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: FlowInstanceDataUpsertOne.ID is not supported by MySQL driver. Use FlowInstanceDataUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *FlowInstanceDataUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// FlowInstanceDataCreateBulk is the builder for creating many FlowInstanceData entities in bulk.
type FlowInstanceDataCreateBulk struct {
	config
	builders []*FlowInstanceDataCreate
	conflict []sql.ConflictOption
}

// Save creates the FlowInstanceData entities in the database.
func (fidcb *FlowInstanceDataCreateBulk) Save(ctx context.Context) ([]*FlowInstanceData, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fidcb.builders))
	nodes := make([]*FlowInstanceData, len(fidcb.builders))
	mutators := make([]Mutator, len(fidcb.builders))
	for i := range fidcb.builders {
		func(i int, root context.Context) {
			builder := fidcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FlowInstanceDataMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, fidcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = fidcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fidcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, fidcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fidcb *FlowInstanceDataCreateBulk) SaveX(ctx context.Context) []*FlowInstanceData {
	v, err := fidcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fidcb *FlowInstanceDataCreateBulk) Exec(ctx context.Context) error {
	_, err := fidcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fidcb *FlowInstanceDataCreateBulk) ExecX(ctx context.Context) {
	if err := fidcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.FlowInstanceData.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FlowInstanceDataUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (fidcb *FlowInstanceDataCreateBulk) OnConflict(opts ...sql.ConflictOption) *FlowInstanceDataUpsertBulk {
	fidcb.conflict = opts
	return &FlowInstanceDataUpsertBulk{
		create: fidcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.FlowInstanceData.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (fidcb *FlowInstanceDataCreateBulk) OnConflictColumns(columns ...string) *FlowInstanceDataUpsertBulk {
	fidcb.conflict = append(fidcb.conflict, sql.ConflictColumns(columns...))
	return &FlowInstanceDataUpsertBulk{
		create: fidcb,
	}
}

// FlowInstanceDataUpsertBulk is the builder for "upsert"-ing
// a bulk of FlowInstanceData nodes.
type FlowInstanceDataUpsertBulk struct {
	create *FlowInstanceDataCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.FlowInstanceData.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(flowinstancedata.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *FlowInstanceDataUpsertBulk) UpdateNewValues() *FlowInstanceDataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(flowinstancedata.FieldID)
				return
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(flowinstancedata.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.FlowInstanceData.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *FlowInstanceDataUpsertBulk) Ignore() *FlowInstanceDataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FlowInstanceDataUpsertBulk) DoNothing() *FlowInstanceDataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FlowInstanceDataCreateBulk.OnConflict
// documentation for more info.
func (u *FlowInstanceDataUpsertBulk) Update(set func(*FlowInstanceDataUpsert)) *FlowInstanceDataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FlowInstanceDataUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *FlowInstanceDataUpsertBulk) SetCreatedAt(v int64) *FlowInstanceDataUpsertBulk {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *FlowInstanceDataUpsertBulk) AddCreatedAt(v int64) *FlowInstanceDataUpsertBulk {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FlowInstanceDataUpsertBulk) UpdateCreatedAt() *FlowInstanceDataUpsertBulk {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FlowInstanceDataUpsertBulk) SetUpdatedAt(v int64) *FlowInstanceDataUpsertBulk {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *FlowInstanceDataUpsertBulk) AddUpdatedAt(v int64) *FlowInstanceDataUpsertBulk {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FlowInstanceDataUpsertBulk) UpdateUpdatedAt() *FlowInstanceDataUpsertBulk {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *FlowInstanceDataUpsertBulk) SetDeletedAt(v int64) *FlowInstanceDataUpsertBulk {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *FlowInstanceDataUpsertBulk) AddDeletedAt(v int64) *FlowInstanceDataUpsertBulk {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *FlowInstanceDataUpsertBulk) UpdateDeletedAt() *FlowInstanceDataUpsertBulk {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *FlowInstanceDataUpsertBulk) ClearDeletedAt() *FlowInstanceDataUpsertBulk {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.ClearDeletedAt()
	})
}

// SetFlowInstanceID sets the "flow_instance_id" field.
func (u *FlowInstanceDataUpsertBulk) SetFlowInstanceID(v string) *FlowInstanceDataUpsertBulk {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.SetFlowInstanceID(v)
	})
}

// UpdateFlowInstanceID sets the "flow_instance_id" field to the value that was provided on create.
func (u *FlowInstanceDataUpsertBulk) UpdateFlowInstanceID() *FlowInstanceDataUpsertBulk {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.UpdateFlowInstanceID()
	})
}

// SetFlowNodeInstanceID sets the "flow_node_instance_id" field.
func (u *FlowInstanceDataUpsertBulk) SetFlowNodeInstanceID(v string) *FlowInstanceDataUpsertBulk {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.SetFlowNodeInstanceID(v)
	})
}

// UpdateFlowNodeInstanceID sets the "flow_node_instance_id" field to the value that was provided on create.
func (u *FlowInstanceDataUpsertBulk) UpdateFlowNodeInstanceID() *FlowInstanceDataUpsertBulk {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.UpdateFlowNodeInstanceID()
	})
}

// ClearFlowNodeInstanceID clears the value of the "flow_node_instance_id" field.
func (u *FlowInstanceDataUpsertBulk) ClearFlowNodeInstanceID() *FlowInstanceDataUpsertBulk {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.ClearFlowNodeInstanceID()
	})
}

// SetNodeKey sets the "node_key" field.
func (u *FlowInstanceDataUpsertBulk) SetNodeKey(v string) *FlowInstanceDataUpsertBulk {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.SetNodeKey(v)
	})
}

// UpdateNodeKey sets the "node_key" field to the value that was provided on create.
func (u *FlowInstanceDataUpsertBulk) UpdateNodeKey() *FlowInstanceDataUpsertBulk {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.UpdateNodeKey()
	})
}

// ClearNodeKey clears the value of the "node_key" field.
func (u *FlowInstanceDataUpsertBulk) ClearNodeKey() *FlowInstanceDataUpsertBulk {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.ClearNodeKey()
	})
}

// SetData sets the "data" field.
func (u *FlowInstanceDataUpsertBulk) SetData(v map[string]interface{}) *FlowInstanceDataUpsertBulk {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.SetData(v)
	})
}

// UpdateData sets the "data" field to the value that was provided on create.
func (u *FlowInstanceDataUpsertBulk) UpdateData() *FlowInstanceDataUpsertBulk {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.UpdateData()
	})
}

// ClearData clears the value of the "data" field.
func (u *FlowInstanceDataUpsertBulk) ClearData() *FlowInstanceDataUpsertBulk {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.ClearData()
	})
}

// SetType sets the "type" field.
func (u *FlowInstanceDataUpsertBulk) SetType(v int8) *FlowInstanceDataUpsertBulk {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.SetType(v)
	})
}

// AddType adds v to the "type" field.
func (u *FlowInstanceDataUpsertBulk) AddType(v int8) *FlowInstanceDataUpsertBulk {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.AddType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *FlowInstanceDataUpsertBulk) UpdateType() *FlowInstanceDataUpsertBulk {
	return u.Update(func(s *FlowInstanceDataUpsert) {
		s.UpdateType()
	})
}

// Exec executes the query.
func (u *FlowInstanceDataUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the FlowInstanceDataCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FlowInstanceDataCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FlowInstanceDataUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}