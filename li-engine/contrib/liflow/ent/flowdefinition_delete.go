// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/flowdefinition"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/predicate"
)

// FlowDefinitionDelete is the builder for deleting a FlowDefinition entity.
type FlowDefinitionDelete struct {
	config
	hooks    []Hook
	mutation *FlowDefinitionMutation
}

// Where appends a list predicates to the FlowDefinitionDelete builder.
func (fdd *FlowDefinitionDelete) Where(ps ...predicate.FlowDefinition) *FlowDefinitionDelete {
	fdd.mutation.Where(ps...)
	return fdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fdd *FlowDefinitionDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fdd.hooks) == 0 {
		affected, err = fdd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FlowDefinitionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fdd.mutation = mutation
			affected, err = fdd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fdd.hooks) - 1; i >= 0; i-- {
			if fdd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fdd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fdd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (fdd *FlowDefinitionDelete) ExecX(ctx context.Context) int {
	n, err := fdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fdd *FlowDefinitionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: flowdefinition.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: flowdefinition.FieldID,
			},
		},
	}
	if ps := fdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, fdd.driver, _spec)
}

// FlowDefinitionDeleteOne is the builder for deleting a single FlowDefinition entity.
type FlowDefinitionDeleteOne struct {
	fdd *FlowDefinitionDelete
}

// Exec executes the deletion query.
func (fddo *FlowDefinitionDeleteOne) Exec(ctx context.Context) error {
	n, err := fddo.fdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{flowdefinition.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fddo *FlowDefinitionDeleteOne) ExecX(ctx context.Context) {
	fddo.fdd.ExecX(ctx)
}
