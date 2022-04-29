package liflow

import (
	"context"

	"github.com/BeanWei/li/li-engine/contrib/liflow/ent"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/schema"
)

type FlowCtx struct {
	Ctx                 context.Context
	ProcessStatus       int8
	FlowElementMap      map[string]*schema.FlowElement
	CurrentNodeModel    *schema.FlowElement
	FlowInstanceID      string
	FlowInstanceStatus  int8
	NodeInstanceList    []ent.FlowNodeInstance
	CurrentNodeInstance *ent.FlowNodeInstance
	SuspendNodeInstance *ent.FlowNodeInstance
}

func (ctx *FlowCtx) SaveNodeInstanceList(nodeInstanceType int8) error {
	return nil
}

func (ctx *FlowCtx) GetUniqueNextNode(currentFlowElement *schema.FlowElement) *schema.FlowElement {
	if currentFlowElement == nil {
		return nil
	}
	if len(currentFlowElement.Outgoing) == 0 {
		return nil
	}
	nextFlowElement := ctx.FlowElementMap[currentFlowElement.Outgoing[0]]
	if nextFlowElement == nil {
		return nil
	}
	for nextFlowElement != nil && nextFlowElement.Type == FlowElementTypeSequenceFlow {
		nextFlowElement = ctx.GetUniqueNextNode(nextFlowElement)
	}
	return nextFlowElement
}
