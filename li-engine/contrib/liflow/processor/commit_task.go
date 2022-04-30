package processor

import (
	"context"

	"github.com/BeanWei/li/li-engine/contrib/liflow"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gutil"
)

type (
	commitTask struct {
		*liflow.FlowCtx
	}
	CommitTaskInput struct {
		FlowInstanceID string                 `json:"flow_instance_id" v:"required"`
		TaskInstanceID string                 `json:"task_instance_id" v:"required"`
		Variables      map[string]interface{} `json:"variables"`
	}
	CommitTaskOutput struct {
		FlowInstanceID     string                 `json:"flow_instance_id"`
		FlowInstanceStatus int8                   `json:"flow_instance_status"`
		ActiveTaskInstance *ent.FlowNodeInstance  `json:"active_task_instance"`
		Variables          map[string]interface{} `json:"variables"`
	}
)

// CommitTask 提交任务
// 引擎从指定的用户任务节点开始执行，直到用户任务节点挂起或者结束节点完成
func CommitTask(ctx context.Context, input *CommitTaskInput) (*CommitTaskOutput, error) {
	flowInstance, err := ent.DB().FlowInstance.Get(ctx, input.FlowInstanceID)
	if err != nil {
		return nil, gerror.WrapCode(liflow.ErrCodeGetFlowInstanceFailed, err)
	}
	if flowInstance.Status == liflow.FlowInstanceStatusTerminated {
		return nil, gerror.NewCode(liflow.ErrCodeCommitRejected)
	}
	if flowInstance.Status == liflow.FlowInstanceStatusCompleted {
		return nil, nil
	}
	flow, err := ent.DB().FlowDeployment.Get(ctx, flowInstance.FlowDeploymentID)
	if err != nil {
		return nil, gerror.WrapCode(liflow.ErrCodeGetFlowDeploymentFailed, err)
	}
	suspendNodeInstance, err := ent.DB().FlowNodeInstance.Get(ctx, input.TaskInstanceID)
	if err != nil {
		return nil, gerror.WrapCode(liflow.ErrCodeGetNodeInstanceFailed, err)
	}
	ct := &commitTask{
		FlowCtx: &liflow.FlowCtx{
			Ctx:                 ctx,
			FlowElementMap:      flow.Model.ElementMap(),
			FlowInstanceID:      flowInstance.ID,
			FlowInstanceStatus:  flowInstance.Status,
			SuspendNodeInstance: suspendNodeInstance,
			CurrentNodeInstance: suspendNodeInstance,
			NodeInstanceList:    make([]ent.FlowNodeInstance, 0),
			InstanceDataID:      suspendNodeInstance.FlowInstanceDataID,
		},
	}
	if ct.isCompleted() {
		return nil, nil
	}
	ct.CurrentNodeModel = ct.FlowElementMap[ct.SuspendNodeInstance.NodeKey]
	if ct.CurrentNodeModel == nil {
		return nil, gerror.NewCode(liflow.ErrCodeModelUnknownElementKey)
	}

	instanceDataMap := make(map[string]interface{})
	if ct.InstanceDataID != "" {
		instanceData, err := ent.DB().FlowInstanceData.Get(ctx, ct.InstanceDataID)
		if err != nil {
			return nil, gerror.WrapCode(liflow.ErrCodeGetInstanceDataFailed, err)
		}
		instanceDataMap = instanceData.Data
	}
	if input.Variables != nil {
		gutil.MapMerge(instanceDataMap, input.Variables)
		flowInstanceData, err := ent.DB().FlowInstanceData.Create().
			SetFlowInstanceID(flowInstance.ID).
			SetData(instanceDataMap).
			SetType(liflow.FlowInstanceDataTypeCommit).
			SetNodeKey(ct.CurrentNodeModel.Key).
			Save(ctx)
		if err != nil {
			return nil, gerror.WrapCode(liflow.ErrCodeSaveInstanceDataFailed, err)
		}
		ct.InstanceDataID = flowInstanceData.ID
	}
	ct.InstanceDataMap = instanceDataMap

	if err := ct.doCommit(); err != nil {
		ct.ProcessStatus = liflow.ProcessStatusFailed
	} else {
		ct.ProcessStatus = liflow.ProcessStatusSuccess
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
		// 用户节点执行完成之后退出
		if ct.CurrentNodeModel.FlowType == liflow.FlowElementFlowTypeUserTask {
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
