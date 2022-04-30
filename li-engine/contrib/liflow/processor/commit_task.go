package processor

import (
	"context"

	"github.com/BeanWei/li/li-engine/contrib/liflow"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type (
	commitTask struct {
		*liflow.FlowCtx
	}
	CommitTaskInput struct {
		FlowInstanceID string
		TaskInstanceID string
		Variables      map[string]interface{}
	}
	CommitTaskOutput struct {
		FlowInstanceID     string                 `json:"flow_instance_id"`
		FlowInstanceStatus int8                   `json:"flow_instance_status"`
		ActiveTaskInstance *ent.FlowNodeInstance  `json:"active_task_instance"`
		Variables          map[string]interface{} `json:"variables"`
	}
)

func CommitTask(ctx context.Context, input *CommitTaskInput) (*CommitTaskOutput, error) {
	flowInstance, err := ent.DB().FlowInstance.Get(ctx, input.FlowInstanceID)
	if err != nil {
		return nil, err
	}
	if flowInstance.Status == liflow.FlowInstanceStatusTerminated {
		return nil, gerror.New("flowInstance is terminated")
	}
	if flowInstance.Status == liflow.FlowInstanceStatusCompleted {
		return nil, gerror.New("flowInstance is completed")
	}
	flow, err := ent.DB().FlowDeployment.Get(ctx, flowInstance.FlowDeploymentID)
	if err != nil {
		return nil, err
	}
	suspendNodeInstance, err := ent.DB().FlowNodeInstance.Get(ctx, input.TaskInstanceID)
	if err != nil {
		return nil, err
	}
	ct := &commitTask{
		FlowCtx: &liflow.FlowCtx{
			Ctx:                 ctx,
			FlowElementMap:      flow.Model.ElementMap(),
			FlowInstanceID:      flowInstance.ID,
			FlowInstanceStatus:  flowInstance.Status,
			SuspendNodeInstance: suspendNodeInstance,
			NodeInstanceList:    make([]ent.FlowNodeInstance, 0),
		},
	}
	if ct.isCompleted() {
		return nil, gerror.NewCode(gcode.CodeInvalidOperation, "flow has been processed completely")
	}
	ct.CurrentNodeModel = ct.FlowElementMap[ct.SuspendNodeInstance.NodeKey]
	if ct.CurrentNodeModel == nil {
		return nil, gerror.NewCode(gcode.CodeInvalidOperation, "cannot get current node model")
	}

	if err := ct.doCommit(); err != nil {
		return nil, err
	}
	if err := ct.postCommit(); err != nil {
		return nil, err
	}

	return &CommitTaskOutput{
		FlowInstanceID:     flowInstance.ID,
		FlowInstanceStatus: ct.FlowInstanceStatus,
		ActiveTaskInstance: ct.CurrentNodeInstance,
	}, nil
}

func (ct *commitTask) doCommit() error {
	executor := ct.getExecuteExecutor()
	if executor == nil {
		return nil
	}
	if err := executor.Commit(ct.FlowCtx); err != nil {
		return err
	}
	for executor != nil {
		err := executor.Execute(ct.FlowCtx)
		if err != nil {
			return err
		}
		if ct.ProcessStatus == liflow.ProcessStatusSuccess || ct.ProcessStatus == liflow.ProcessStatusCommitSuspend {
			return nil
		}
		executor = executor.GetExecuteExecutor(ct.FlowCtx)
	}
	return nil
}

func (ct *commitTask) postCommit() error {
	if ct.CurrentNodeInstance != nil {
		ct.SuspendNodeInstance = ct.CurrentNodeInstance
	}
	if err := ct.SaveNodeInstanceList(liflow.FlowNodeInstanceTypeCommit); err != nil {
		return err
	}
	// 更新流程实例状态
	if ct.isCompleted() {
		ct.FlowInstanceStatus = liflow.FlowInstanceStatusCompleted
		return ent.DB().FlowInstance.
			UpdateOneID(ct.FlowInstanceID).
			SetStatus(liflow.FlowInstanceStatusCompleted).
			Exec(ct.Ctx)
	}
	return nil
}

func (ct *commitTask) isCompleted() bool {
	if ct.FlowInstanceStatus == liflow.FlowInstanceStatusCompleted {
		return true
	}
	if ct.SuspendNodeInstance == nil {
		return false
	}
	if ct.SuspendNodeInstance.Status != liflow.FlowNodeInstanceStatusCompleted {
		return false
	}
	if node := ct.FlowElementMap[ct.SuspendNodeInstance.NodeKey]; node != nil && node.FlowType == liflow.FlowElementFlowTypeEndEvent {
		return true
	}
	return false
}

func (ct *commitTask) getExecuteExecutor() liflow.Executor {
	if ct.isCompleted() {
		return nil
	}
	executor := liflow.GetElementExecutor(ct.CurrentNodeModel.Type)
	return executor
}
