package schema

// import (
// 	"entgo.io/ent"
// 	"entgo.io/ent/schema/field"
// 	"entgo.io/ent/schema/index"
// 	"github.com/BeanWei/li/li-engine/contrib/lient/mixin"
// )

// type FlowInstanceComment struct {
// 	ent.Schema
// }

// func (FlowInstanceComment) Mixin() []ent.Mixin {
// 	return []ent.Mixin{
// 		mixin.XID{},
// 	}
// }

// func (FlowInstanceComment) Fields() []ent.Field {
// 	return []ent.Field{
// 		field.String("flow_instance_id").
// 			NotEmpty().
// 			Comment("流程实例ID"),
// 		field.String("flow_node_instance_id").
// 			NotEmpty().
// 			Comment("节点实例ID"),
// 		field.Int8("type").
// 			Default(0).
// 			Comment("类型(1:事件 2:意见)"),
// 		field.Int8("action").
// 			Default(0).
// 			Comment("操作(1:通过 2:拒绝 3:驳回)"),
// 		field.String("user_id").
// 			Default("").
// 			Optional().
// 			Comment("审批人"),
// 		field.Int64("start_at").
// 			Optional().
// 			Comment("开始时间"),
// 		field.Int64("end_at").
// 			Optional().
// 			Comment("结束时间"),
// 		field.String("content").
// 			Default("").
// 			Optional().
// 			Comment("审批信息"),
// 		field.Strings("attachments").
// 			Default([]string{}).
// 			Optional().
// 			Comment("审批附件"),
// 	}
// }

// func (FlowInstanceComment) Indexes() []ent.Index {
// 	return []ent.Index{
// 		index.Fields("flow_instance_id"),
// 		index.Fields("flow_node_instance_id"),
// 	}
// }
