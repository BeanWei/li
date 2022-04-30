package processor

import (
	"context"

	"github.com/BeanWei/li/li-engine/contrib/liflow"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent"
	"github.com/gogf/gf/v2/errors/gerror"
)

type (
	terminateProcessInput struct {
		FlowInstanceID string
	}
	terminateProcessOutput struct {
		FlowInstanceID     string `json:"flow_instance_id"`
		FlowInstanceStatus int8   `json:"flow_instance_status"`
	}
)

// TerminateProcess 终止流程
// 强制终止流程实例，如果当前流程实例状态是已完成，引擎什么也不做，否则引擎会将状态置为已终止
// 注意，一旦流程已完成或者已终止，引擎将不再允许提交和回滚
func TerminateProcess(ctx context.Context, input *terminateProcessInput) (*terminateProcessOutput, error) {
	flowInstance, err := ent.DB().FlowInstance.Get(ctx, input.FlowInstanceID)
	if err != nil {
		return nil, gerror.WrapCode(liflow.ErrCodeGetFlowInstanceFailed, err)
	}
	output := new(terminateProcessOutput)
	if flowInstance.Status == liflow.FlowInstanceStatusCompleted {
		output.FlowInstanceStatus = liflow.FlowInstanceStatusCompleted
	} else {
		err = flowInstance.Update().SetStatus(liflow.FlowInstanceStatusTerminated).Exec(ctx)
		if err != nil {
			return nil, gerror.WrapCode(liflow.ErrCodeSaveFlowInstanceFailed, err)
		}
		output.FlowInstanceStatus = liflow.FlowInstanceStatusTerminated
	}
	return output, nil
}
