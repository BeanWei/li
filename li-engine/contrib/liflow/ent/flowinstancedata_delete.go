// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/flowinstancedata"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/predicate"
)

// FlowInstanceDataDelete is the builder for deleting a FlowInstanceData entity.
type FlowInstanceDataDelete struct {
	config
	hooks    []Hook
	mutation *FlowInstanceDataMutation
}

// Where appends a list predicates to the FlowInstanceDataDelete builder.
func (fidd *FlowInstanceDataDelete) Where(ps ...predicate.FlowInstanceData) *FlowInstanceDataDelete {
	fidd.mutation.Where(ps...)
	return fidd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fidd *FlowInstanceDataDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fidd.hooks) == 0 {
		affected, err = fidd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FlowInstanceDataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fidd.mutation = mutation
			affected, err = fidd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fidd.hooks) - 1; i >= 0; i-- {
			if fidd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fidd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fidd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (fidd *FlowInstanceDataDelete) ExecX(ctx context.Context) int {
	n, err := fidd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fidd *FlowInstanceDataDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: flowinstancedata.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: flowinstancedata.FieldID,
			},
		},
	}
	if ps := fidd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, fidd.driver, _spec)
}

// FlowInstanceDataDeleteOne is the builder for deleting a single FlowInstanceData entity.
type FlowInstanceDataDeleteOne struct {
	fidd *FlowInstanceDataDelete
}

// Exec executes the deletion query.
func (fiddo *FlowInstanceDataDeleteOne) Exec(ctx context.Context) error {
	n, err := fiddo.fidd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{flowinstancedata.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fiddo *FlowInstanceDataDeleteOne) ExecX(ctx context.Context) {
	fiddo.fidd.ExecX(ctx)
}
