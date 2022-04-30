package executor

import (
	"time"

	"github.com/BeanWei/li/li-engine/contrib/lient"
	"github.com/BeanWei/li/li-engine/contrib/liflow"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/schema"
	"github.com/gogf/gf/v2/errors/gerror"
)

type ElementExecutor struct{}

func (e *ElementExecutor) Validate(eleMap map[string]*schema.FlowElement, ele *schema.FlowElement) error {
	return nil
}

func (e *ElementExecutor) Execute(ctx *liflow.FlowCtx) error {
	if err := e.PreExecute(ctx); err != nil {
		return err
	}
	if err := e.DoExecute(ctx); err != nil {
		return err
	}
	return e.PostExecute(ctx)
}

func (e *ElementExecutor) Commit(ctx *liflow.FlowCtx) error {
	if err := e.PreCommit(ctx); err != nil {
		return err
	}
	if err := e.DoCommit(ctx); err != nil {
		return err
	}
	return e.PostCommit(ctx)
}

func (e *ElementExecutor) Rollback(ctx *liflow.FlowCtx) error {
	if err := e.PreRollback(ctx); err != nil {
		return err
	}
	if err := e.DoRollback(ctx); err != nil {
		return err
	}
	return e.PostRollback(ctx)
}

func (e *ElementExecutor) IsCompleted(ctx *liflow.FlowCtx) bool {
	// case1. startEvent
	if ctx.CurrentNodeInstance == nil {
		return false
	}
	// case2. begin to process the node
	if ctx.CurrentNodeModel.Key != ctx.CurrentNodeInstance.NodeKey {
		return false
	}
	// case3. process completed
	if ctx.CurrentNodeInstance.Status == liflow.FlowNodeInstanceStatusCompleted {
		return true
	}
	// case4. to preocess
	return false
}

func (e *ElementExecutor) PreExecute(ctx *liflow.FlowCtx) (err error) {
	currentNodeInstance := &ent.FlowNodeInstance{
		ID:                 lient.NewXid(),
		CreatedAt:          time.Now().Unix(),
		UpdatedAt:          time.Now().Unix(),
		FlowInstanceID:     ctx.FlowInstanceID,
		NodeKey:            ctx.CurrentNodeModel.Key,
		Status:             liflow.FlowNodeInstanceStatusActive,
		FlowInstanceDataID: ctx.InstanceDataID,
	}
	if ctx.CurrentNodeInstance != nil {
		currentNodeInstance.SourceFlowNodeInstanceID = ctx.CurrentNodeInstance.ID
		currentNodeInstance.SourceNodeKey = ctx.CurrentNodeInstance.NodeKey
		ctx.CurrentNodeInstance, err = ent.DB().FlowNodeInstance.Get(ctx.Ctx, ctx.CurrentNodeInstance.ID)
		if err != nil {
			return gerror.WrapCode(liflow.ErrCodeGetNodeInstanceFailed, err)
		}
	}
	ctx.CurrentNodeInstance = currentNodeInstance
	return nil
}

func (e *ElementExecutor) DoExecute(ctx *liflow.FlowCtx) error {
	return nil
}

func (e *ElementExecutor) PostExecute(ctx *liflow.FlowCtx) error {
	return nil
}

// 仅用户节点能够提交, 用户节点必须实现该方法
func (e *ElementExecutor) PreCommit(ctx *liflow.FlowCtx) error {
	return gerror.NewCode(liflow.ErrCodeUnsupportElementType)
}

func (e *ElementExecutor) DoCommit(ctx *liflow.FlowCtx) error {
	return nil
}

func (e *ElementExecutor) PostCommit(ctx *liflow.FlowCtx) error {
	return nil
}

func (e *ElementExecutor) PreRollback(ctx *liflow.FlowCtx) error {
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
	return nil
}

func (e *ElementExecutor) DoRollback(ctx *liflow.FlowCtx) error {
	return nil
}

func (e *ElementExecutor) PostRollback(ctx *liflow.FlowCtx) error {
	if ctx.CurrentNodeInstance != nil {
		ctx.CurrentNodeInstance.Status = liflow.FlowNodeInstanceStatusDisabled
		ctx.NodeInstanceList = append(ctx.NodeInstanceList, *ctx.CurrentNodeInstance)
	}
	return nil
}

func (e *ElementExecutor) GetExecuteExecutor(ctx *liflow.FlowCtx) liflow.Executor {
	flowElement := ctx.GetUniqueNextNode(ctx.CurrentNodeModel)
	ctx.CurrentNodeModel = flowElement
	return liflow.GetElementExecutor(flowElement.Type)
}

func (e *ElementExecutor) GetRollbackExecutor(ctx *liflow.FlowCtx) (liflow.Executor, error) {
	if ctx.CurrentNodeInstance != nil && ctx.CurrentNodeInstance.SourceFlowNodeInstanceID == "" {
		return nil, nil
	}
	sourceNodeInstance, err := ent.DB().FlowNodeInstance.Get(ctx.Ctx, ctx.CurrentNodeInstance.SourceFlowNodeInstanceID)
	if err != nil {
		return nil, gerror.WrapCode(liflow.ErrCodeGetNodeFailed, err)
	}
	ctx.CurrentNodeModel = ctx.FlowElementMap[sourceNodeInstance.NodeKey]
	return liflow.GetElementExecutor(ctx.CurrentNodeModel.Type), nil
}

var _ liflow.Executor = (*ElementExecutor)(nil)
