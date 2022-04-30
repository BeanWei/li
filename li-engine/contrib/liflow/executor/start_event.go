package executor

import (
	"github.com/BeanWei/li/li-engine/contrib/liflow"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/schema"
	"github.com/gogf/gf/v2/errors/gerror"
)

type StartEventExecutor struct {
	ElementExecutor
}

func (e *StartEventExecutor) Validate(eleMap map[string]*schema.FlowElement, ele *schema.FlowElement) error {
	if len(ele.Incoming) > 0 {
		return gerror.NewCode(liflow.ErrCodeElementTooMuchIncoming)
	}
	return nil
}

func (e *StartEventExecutor) PostExecute(ctx *liflow.FlowCtx) error {
	ctx.CurrentNodeInstance.Status = liflow.FlowNodeInstanceStatusCompleted
	ctx.NodeInstanceList = append(ctx.NodeInstanceList, *ctx.CurrentNodeInstance)
	return nil
}
