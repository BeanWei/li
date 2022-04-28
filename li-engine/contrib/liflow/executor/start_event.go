package executor

import "github.com/BeanWei/li/li-engine/contrib/liflow"

type StartEventExecutor struct {
	ElementExecutor
}

func (e *StartEventExecutor) PostExecute(ctx *liflow.FlowCtx) error {
	ctx.CurrentNodeInstance.Status = liflow.FlowNodeInstanceStatusCompleted
	ctx.NodeInstanceList = append(ctx.NodeInstanceList, *ctx.CurrentNodeInstance)
	return nil
}
