package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/BeanWei/li/li-engine/contrib/lient/mixin"
)

type (
	FlowDefinition struct {
		ent.Schema
	}
	FlowModel   []*FlowElement
	FlowElement struct {
		Key        string
		Type       string
		Outgoing   []string
		Incoming   []string
		Properties map[string]interface{}
	}
)

func (fm *FlowModel) ElementMap() map[string]*FlowElement {
	femap := make(map[string]*FlowElement)
	for _, fe := range *fm {
		femap[fe.Key] = fe
	}
	return femap
}

func (FlowDefinition) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.XID{},
		mixin.Time{},
	}
}

func (FlowDefinition) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Comment("流程名称"),
		field.Int8("status").
			Default(0).
			Comment("状态(0.编辑中 1.已上线 2.已下线)"),
		field.JSON("model", FlowModel{}).
			Comment("流程模型"),
		field.String("remark").
			Default("").
			Optional().
			Comment("备注"),
	}
}

func (FlowDefinition) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("flow_instances", FlowInstance.Type),
	}
}
