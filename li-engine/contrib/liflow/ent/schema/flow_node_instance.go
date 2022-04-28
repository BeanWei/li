package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/BeanWei/li/li-engine/contrib/lient/mixin"
)

type FlowNodeInstance struct {
	ent.Schema
}

func (FlowNodeInstance) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.XID{},
		mixin.Time{},
	}
}

func (FlowNodeInstance) Fields() []ent.Field {
	return []ent.Field{
		field.String("flow_instance_id").
			NotEmpty().
			Comment("流程执行实例ID"),
		field.String("source_flow_node_instance_id").
			Default("").
			Optional().
			Comment("上一个节点执行实例ID"),
		field.String("flow_instance_data_id").
			Default("").
			Optional().
			Comment("实例数据ID"),
		field.String("node_key").
			NotEmpty().
			Comment("节点唯一标示"),
		field.String("source_node_key").
			Default("").
			Optional().
			Comment("上一节点唯一标示"),
		field.Int8("status").
			Default(0).
			Comment("状态(1.处理成功 2.处理中 3.处理失败 4.处理已撤销)"),
	}
}

func (FlowNodeInstance) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("flow_instance", FlowInstance.Type).
			Ref("flow_node_instances").
			Field("flow_instance_id").
			Unique().
			Required(),
	}
}
