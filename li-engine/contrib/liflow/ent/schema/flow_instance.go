package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/BeanWei/li/li-engine/contrib/lient/mixin"
)

type FlowInstance struct {
	ent.Schema
}

func (FlowInstance) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.XID{},
		mixin.Time{},
	}
}

func (FlowInstance) Fields() []ent.Field {
	return []ent.Field{
		field.String("flow_definition_id").
			NotEmpty().
			Comment("流程定义ID"),
		field.Int8("status").
			Default(0).
			Comment("状态(1.执行完成 2.执行中 3.执行终止(强制终止))"),
	}
}

func (FlowInstance) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("flow_definition", FlowDefinition.Type).
			Ref("flow_instances").
			Field("flow_definition_id").
			Unique().
			Required(),
		edge.To("flow_node_instances", FlowNodeInstance.Type),
	}
}
