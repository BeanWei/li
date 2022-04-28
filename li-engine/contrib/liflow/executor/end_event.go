package executor

import "github.com/BeanWei/li/li-engine/contrib/liflow"

type EndEventExecutor struct {
	ElementExecutor
}

func (e *EndEventExecutor) PostExecute(ctx *liflow.FlowCtx) error {
	ctx.CurrentNodeInstance.Status = liflow.FlowNodeInstanceStatusCompleted
	ctx.NodeInstanceList = append(ctx.NodeInstanceList, *ctx.CurrentNodeInstance)
	return nil
}

func (e *EndEventExecutor) GetExecuteExecutor(ctx *liflow.FlowCtx) liflow.Executor {
	return nil
}
