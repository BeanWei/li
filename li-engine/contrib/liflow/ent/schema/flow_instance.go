package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
		field.String("flow_deployment_id").
			NotEmpty().
			Comment("部署的流程ID"),
		field.String("ref_id").
			NotEmpty().
			Comment("引用/调用方ID"),
		field.Int8("status").
			Default(0).
			Comment("状态(1.执行完成 2.执行中 3.执行终止(强制终止))"),
	}
}

func (FlowInstance) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("flow_deployment", FlowDeployment.Type).
			Ref("flow_instances").
			Field("flow_deployment_id").
			Unique().
			Required(),
		edge.To("flow_node_instances", FlowNodeInstance.Type),
	}
}

func (FlowInstance) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("flow_deployment_id"),
		index.Fields("ref_id"),
		index.Fields("status"),
	}
}
