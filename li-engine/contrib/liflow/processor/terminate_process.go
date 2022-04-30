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
