package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/BeanWei/li/li-engine/contrib/lient/mixin"
)

type FlowDeployment struct {
	ent.Schema
}

func (FlowDeployment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.XID{},
		mixin.Time{},
	}
}

func (FlowDeployment) Fields() []ent.Field {
	return []ent.Field{
		field.String("flow_definition_id").
			NotEmpty().
			Comment("流程定义ID"),
		field.String("name").
			NotEmpty().
			Comment("流程名称"),
		field.Int8("status").
			Default(0).
			Comment("状态(1.已部署 2.已下线)"),
		field.JSON("model", FlowModel{}).
			Comment("流程模型"),
		field.String("remark").
			Default("").
			Optional().
			Comment("备注"),
	}
}

func (FlowDeployment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("flow_definition", FlowDefinition.Type).
			Ref("flow_deployments").
			Field("flow_definition_id").
			Unique().
			Required(),
		edge.To("flow_instances", FlowInstance.Type),
	}
}

func (FlowDeployment) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("flow_definition_id"),
	}
}
