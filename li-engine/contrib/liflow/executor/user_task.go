package executor

import (
	"github.com/BeanWei/li/li-engine/contrib/liflow"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/schema"
	"github.com/gogf/gf/v2/errors/gerror"
)

type UserTaskExecutor struct {
	ElementExecutor
}

func (e *UserTaskExecutor) Validate(eleMap map[string]*schema.FlowElement, ele *schema.FlowElement) error {
	if len(ele.Incoming) > 0 {
		return gerror.NewCode(liflow.ErrCodeElementTooMuchIncoming)
	}
	return nil
}

func (e *UserTaskExecutor) Execute(ctx *liflow.FlowCtx) error {
	if ctx.CurrentNodeInstance == nil {
		return nil
	}
	if ctx.CurrentNodeInstance.Status == liflow.FlowNodeInstanceStatusCompleted {
		return nil
	}
	ctx.CurrentNodeInstance.Status = liflow.FlowNodeInstanceStatusActive
	ctx.NodeInstanceList = append(ctx.NodeInstanceList, *ctx.CurrentNodeInstance)
	return nil
}

func (e *UserTaskExecutor) Commit(ctx *liflow.FlowCtx) error {
	if ctx.SuspendNodeInstance == nil {
		return gerror.NewCode(liflow.ErrCodeCommitFailed)
	}
	if ctx.SuspendNodeInstance.NodeKey != ctx.CurrentNodeModel.Key {
		return gerror.NewCode(liflow.ErrCodeCommitFailed)
	}
	if ctx.SuspendNodeInstance.Status == liflow.FlowNodeInstanceStatusCompleted {
		return nil
	}
	if ctx.SuspendNodeInstance.Status != liflow.FlowNodeInstanceStatusActive {
		return gerror.NewCode(liflow.ErrCodeCommitFailed)
	}
	ctx.CurrentNodeInstance.Status = liflow.FlowNodeInstanceStatusCompleted
	ctx.NodeInstanceList = append(ctx.NodeInstanceList, *ctx.CurrentNodeInstance)
	return nil
}

func (e *UserTaskExecutor) Rollback(ctx *liflow.FlowCtx) error {
	if ctx.CurrentNodeInstance == nil {
		ctx.CurrentNodeInstance = ctx.SuspendNodeInstance
	} else if ctx.CurrentNodeInstance.SourceFlowNodeInstanceID != "" {
		nodeInstance, err := ent.DB().FlowNodeInstance.Get(ctx.Ctx, ctx.CurrentNodeInstance.SourceFlowNodeInstanceID)
		if err != nil {
			return gerror.WrapCode(liflow.ErrCodeGetNodeInstanceFailed, err)
		}
		ctx.CurrentNodeInstance = nodeInstance
		ctx.InstanceDataID = nodeInstance.FlowInstanceDataID
		if ctx.InstanceDataID != "" {
			currentInstanceData, err := ent.DB().FlowInstanceData.Get(ctx.Ctx, ctx.InstanceDataID)
			if err != nil {
				return gerror.WrapCode(liflow.ErrCodeGetInstanceDataFailed, err)
			}
			ctx.InstanceDataMap = currentInstanceData.Data
		}
	}

	currentStatus := ctx.CurrentNodeInstance.Status
	if currentStatus == liflow.FlowNodeInstanceStatusDisabled {
		return nil
	}
	ctx.CurrentNodeInstance.Status = liflow.FlowDefinitionStatusDisabled
	ctx.NodeInstanceList = append(ctx.NodeInstanceList, *ctx.CurrentNodeInstance)
	if currentStatus == liflow.FlowNodeInstanceStatusCompleted {
		newNodeInstance, err := ent.DB().FlowNodeInstance.
			Create().
			SetFlowInstanceID(ctx.CurrentNodeInstance.FlowInstanceID).
			SetSourceFlowNodeInstanceID(ctx.CurrentNodeInstance.SourceFlowNodeInstanceID).
			SetFlowInstanceDataID(ctx.CurrentNodeInstance.FlowInstanceDataID).
			SetNodeKey(ctx.CurrentNodeInstance.NodeKey).
			SetSourceNodeKey(ctx.CurrentNodeInstance.SourceNodeKey).
			SetStatus(liflow.FlowNodeInstanceStatusActive).
			Save(ctx.Ctx)
		if err != nil {
			return gerror.WrapCode(liflow.ErrCodeRollbackFailed, err)
		}
		ctx.CurrentNodeInstance = newNodeInstance
		ctx.NodeInstanceList = append(ctx.NodeInstanceList, *newNodeInstance)
	}
	return nil
}

func (e *UserTaskExecutor) GetExecuteExecutor(ctx *liflow.FlowCtx) liflow.Executor {
	if ctx.CurrentNodeModel == nil {
		return nil
	}
	var nextNode *schema.FlowElement
	if len(ctx.CurrentNodeModel.Outgoing) == 1 {
		nextNode = ctx.GetUniqueNextNode(ctx.CurrentNodeModel)
	} else {
		nextNode = ctx.CalculateNextNode(ctx.CurrentNodeModel, ctx.InstanceDataMap)
	}
	ctx.CurrentNodeModel = nextNode
	return liflow.GetElementExecutor(nextNode.Type)
}
