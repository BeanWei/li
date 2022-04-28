package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/BeanWei/li/li-engine/contrib/lient/mixin"
)

type FlowNodeInstanceLog struct {
	ent.Schema
}

func (FlowNodeInstanceLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.XID{},
		mixin.Time{},
	}
}

func (FlowNodeInstanceLog) Fields() []ent.Field {
	return []ent.Field{
		field.String("flow_instance_id").
			NotEmpty().
			Comment("流程执行实例ID"),
		field.String("flow_node_instance_id").
			NotEmpty().
			Comment("节点执行实例ID"),
		field.String("flow_instance_data_id").
			Default("").
			Optional().
			Comment("实例数据ID"),
		field.String("node_key").
			NotEmpty().
			Comment("节点唯一标示"),
		field.Int8("type").
			Default(0).
			Comment("操作类型(1.系统执行 2.任务提交 3.任务撤销)"),
		field.Int8("status").
			Default(0).
			Comment("'状态(1.处理成功 2.处理中 3.处理失败 4.处理已撤销)"),
	}
}
