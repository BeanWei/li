package liflow

import (
	"context"

	"github.com/BeanWei/li/li-engine/contrib/liflow/ent"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/flownodeinstance"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/flownodeinstancelog"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/schema"
	"github.com/antonmedv/expr"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

type FlowCtx struct {
	Ctx                 context.Context
	FlowElementMap      map[string]*schema.FlowElement
	FlowInstanceID      string
	FlowInstanceStatus  int8
	SuspendNodeInstance *ent.FlowNodeInstance
	NodeInstanceList    []ent.FlowNodeInstance
	CurrentNodeModel    *schema.FlowElement
	CurrentNodeInstance *ent.FlowNodeInstance
	InstanceDataID      string
	InstanceDataMap     map[string]interface{}
	ProcessStatus       int8
}

func (ctx *FlowCtx) SaveNodeInstanceList(nodeInstanceType int8) error {
	if len(ctx.NodeInstanceList) == 0 {
		return nil
	}
	bulkNodeInstance := make([]*ent.FlowNodeInstanceCreate, len(ctx.NodeInstanceList))
	bulkNodeInstanceLog := make([]*ent.FlowNodeInstanceLogCreate, len(ctx.NodeInstanceList))
	for i, node := range ctx.NodeInstanceList {
		bulkNodeInstance[i] = ent.DB().FlowNodeInstance.Create().
			SetID(node.ID).
			SetCreatedAt(node.CreatedAt).
			SetUpdatedAt(node.UpdatedAt).
			SetFlowInstanceID(node.FlowInstanceID).
			SetSourceFlowNodeInstanceID(node.SourceFlowNodeInstanceID).
			SetFlowInstanceDataID(node.FlowInstanceDataID).
			SetNodeKey(node.NodeKey).
			SetSourceNodeKey(node.SourceNodeKey).
			SetStatus(node.Status)
		bulkNodeInstanceLog[i] = ent.DB().FlowNodeInstanceLog.Create().
			SetFlowInstanceID(node.FlowInstanceID).
			SetFlowNodeInstanceID(node.ID).
			SetFlowInstanceDataID(node.FlowInstanceDataID).
			SetNodeKey(node.NodeKey).
			SetType(nodeInstanceType).
			SetStatus(node.Status)
	}
	err := ent.DB().FlowNodeInstance.
		CreateBulk(bulkNodeInstance...).
		OnConflictColumns(flownodeinstance.FieldID).
		UpdateNewValues().
		Exec(ctx.Ctx)
	if err != nil {
		return gerror.WrapCode(ErrCodeSaveNodeInstanceListFailed, err)
	}
	err = ent.DB().FlowNodeInstanceLog.
		CreateBulk(bulkNodeInstanceLog...).
		OnConflictColumns(flownodeinstancelog.FieldID).
		UpdateNewValues().
		Exec(ctx.Ctx)
	if err != nil {
		return gerror.WrapCode(ErrCodeSaveNodeInstanceLogListFailed, err)
	}
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
	for nextFlowElement != nil && nextFlowElement.FlowType == FlowElementFlowTypeSequenceFlow {
		nextFlowElement = ctx.GetUniqueNextNode(nextFlowElement)
	}
	return nextFlowElement
}

func (ctx *FlowCtx) CalculateNextNode(currentFlowElement *schema.FlowElement, instanceDataMap map[string]interface{}) *schema.FlowElement {
	var nextFlowElement *schema.FlowElement
	for _, og := range currentFlowElement.Outgoing {
		outgoingSequenceFlow := ctx.FlowElementMap[og]
		if outgoingSequenceFlow == nil || outgoingSequenceFlow.Properties == nil {
			continue
		}
		condition := gconv.String(outgoingSequenceFlow.Properties[FlowElementPropertiesCondition])
		expr.Compile(condition, expr.Env(instanceDataMap))
		if condition != "" {
			if output, err := expr.Eval(condition, instanceDataMap); err != nil && gconv.Bool(output) {
				return outgoingSequenceFlow
			}
		}
		if gconv.Bool(outgoingSequenceFlow.Properties[FlowElementPropertiesDefaultCondition]) {
			return outgoingSequenceFlow
		}
	}
	for nextFlowElement.FlowType == FlowElementFlowTypeSequenceFlow {
		nextFlowElement = ctx.GetUniqueNextNode(nextFlowElement)
	}
	return nextFlowElement
}
