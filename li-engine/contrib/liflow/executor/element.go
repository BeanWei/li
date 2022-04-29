package executor

import (
	"time"

	"github.com/BeanWei/li/li-engine/contrib/lient"
	"github.com/BeanWei/li/li-engine/contrib/liflow"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/schema"
)

type ElementExecutor struct{}

func (e *ElementExecutor) Validate(ele *schema.FlowElement) error {
	return nil
}

func (e *ElementExecutor) Execute(ctx *liflow.FlowCtx) error {
	if err := e.PreExecute(ctx); err != nil {
		return err
	}
	if err := e.DoExecute(ctx); err != nil {
		return err
	}
	if err := e.PostExecute(ctx); err != nil {
		return err
	}
	return nil
}

func (e *ElementExecutor) Commit(ctx *liflow.FlowCtx) error {
	return nil
}

func (e *ElementExecutor) Rollback(ctx *liflow.FlowCtx) error {
	return nil
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
		return nil, err
	}
	ctx.CurrentNodeModel = ctx.FlowElementMap[sourceNodeInstance.NodeKey]
	return liflow.GetElementExecutor(ctx.CurrentNodeModel.Type), nil
}

func (e *ElementExecutor) PreExecute(ctx *liflow.FlowCtx) error {
	currentNodeInstance := &ent.FlowNodeInstance{
		ID:             lient.NewXid(),
		CreatedAt:      time.Now().Unix(),
		UpdatedAt:      time.Now().Unix(),
		FlowInstanceID: ctx.FlowInstanceID,
		NodeKey:        ctx.CurrentNodeModel.Key,
		Status:         liflow.FlowNodeInstanceStatusActive,
	}
	if ctx.CurrentNodeInstance != nil {
		currentNodeInstance.SourceFlowNodeInstanceID = ctx.CurrentNodeInstance.ID
		currentNodeInstance.NodeKey = ctx.CurrentNodeInstance.NodeKey
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

func (e *ElementExecutor) PreCommit(ctx *liflow.FlowCtx) error {
	return nil
}

func (e *ElementExecutor) DoCommit(ctx *liflow.FlowCtx) error {
	return nil
}

func (e *ElementExecutor) PostCommit(ctx *liflow.FlowCtx) error {
	return nil
}

func (e *ElementExecutor) PreRollback(ctx *liflow.FlowCtx) error {
	return nil
}

func (e *ElementExecutor) DoRollback(ctx *liflow.FlowCtx) error {
	return nil
}

func (e *ElementExecutor) PostRollback(ctx *liflow.FlowCtx) error {
	return nil
}

var _ liflow.Executor = (*ElementExecutor)(nil)
