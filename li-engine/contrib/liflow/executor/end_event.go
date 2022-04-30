package executor

import (
	"github.com/BeanWei/li/li-engine/contrib/liflow"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/schema"
	"github.com/gogf/gf/v2/errors/gerror"
)

type EndEventExecutor struct {
	ElementExecutor
}

func (e *EndEventExecutor) Validate(eleMap map[string]*schema.FlowElement, ele *schema.FlowElement) error {
	if len(ele.Outgoing) > 0 {
		return gerror.NewCode(liflow.ErrCodeElementTooMuchOutgoing)
	}
	return nil
}

func (e *EndEventExecutor) Execute(ctx *liflow.FlowCtx) error {
	ctx.CurrentNodeInstance.Status = liflow.FlowNodeInstanceStatusCompleted
	ctx.NodeInstanceList = append(ctx.NodeInstanceList, *ctx.CurrentNodeInstance)
	return nil
}

func (e *EndEventExecutor) Rollback(ctx *liflow.FlowCtx) error {
	ctx.CurrentNodeInstance = ctx.SuspendNodeInstance
	ctx.NodeInstanceList = nil
	return nil
}

func (e *EndEventExecutor) GetExecuteExecutor(ctx *liflow.FlowCtx) liflow.Executor {
	return nil
}
