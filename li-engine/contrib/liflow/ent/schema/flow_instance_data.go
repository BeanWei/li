package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/BeanWei/li/li-engine/contrib/lient/mixin"
)

type FlowInstanceData struct {
	ent.Schema
}

func (FlowInstanceData) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.XID{},
		mixin.Time{},
	}
}

func (FlowInstanceData) Fields() []ent.Field {
	return []ent.Field{
		field.String("flow_instance_id").
			NotEmpty().
			Comment("流程执行实例ID"),
		field.String("flow_node_instance_id").
			NotEmpty().
			Comment("节点执行实例ID"),
		field.String("node_key").
			NotEmpty().
			Comment("节点唯一标示"),
		field.JSON("data", map[string]interface{}{}).
			Optional().
			Comment("数据列表json"),
		field.Int8("type").
			Default(0).
			Comment("操作类型(1.实例初始化 2.系统执行 3.系统主动获取 4.上游更新 5.任务提交 6.任务撤回)"),
	}
}
