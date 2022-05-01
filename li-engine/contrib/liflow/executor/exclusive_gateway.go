package executor

import "github.com/BeanWei/li/li-engine/contrib/liflow"

type ExclusiveGatewayExecutor struct {
	ElementExecutor
}

func (e *ExclusiveGatewayExecutor) Execute(ctx *liflow.FlowCtx) error {
	// TODO: 2022/05/01 Support hook
	ctx.CurrentNodeInstance.FlowInstanceDataID = ctx.InstanceDataID
	ctx.CurrentNodeInstance.Status = liflow.FlowNodeInstanceStatusCompleted
	ctx.NodeInstanceList = append(ctx.NodeInstanceList, *ctx.CurrentNodeInstance)
	return nil
}

func (e *ExclusiveGatewayExecutor) GetExecuteExecutor(ctx *liflow.FlowCtx) liflow.Executor {
	nextNode := ctx.CalculateNextNode(ctx.CurrentNodeModel, ctx.InstanceDataMap)
	ctx.CurrentNodeModel = nextNode
	return liflow.GetElementExecutor(nextNode.Type)
}
