package processor

import (
	"context"

	"github.com/BeanWei/li/li-engine/contrib/liflow"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/flownodeinstance"
	"github.com/gogf/gf/v2/errors/gerror"
)

type (
	rollbackTask struct {
		*liflow.FlowCtx
	}
	RollbackTaskInput struct {
		FlowInstanceID string
		TaskInstanceID string
	}
	RollbackTaskOutput struct {
		FlowInstanceID     string                 `json:"flow_instance_id"`
		FlowInstanceStatus int8                   `json:"flow_instance_status"`
		ActiveTaskInstance *ent.FlowNodeInstance  `json:"active_task_instance"`
		Variables          map[string]interface{} `json:"variables"`
	}
)

// RollbackTask 回滚任务
// 引擎从指定的用户任务节点开始回滚，直到用户任务节点挂起或者开始节点结束
func RollbackTask(ctx context.Context, input *RollbackTaskInput) (*RollbackTaskOutput, error) {
	flowInstance, err := ent.DB().FlowInstance.Get(ctx, input.FlowInstanceID)
	if err != nil {
		return nil, gerror.WrapCode(liflow.ErrCodeGetFlowInstanceFailed, err)
	}
	if flowInstance.Status != liflow.FlowInstanceStatusRunning {
		return nil, gerror.NewCode(liflow.ErrCodeRollbackRejected)
	}
	flow, err := ent.DB().FlowDeployment.Get(ctx, flowInstance.FlowDeploymentID)
	if err != nil {
		return nil, gerror.WrapCode(liflow.ErrCodeGetFlowDeploymentFailed, err)
	}
	suspendNodeInstance, err := ent.DB().FlowNodeInstance.Get(ctx, input.TaskInstanceID)
	if err != nil {
		return nil, err
	}
	rt := &rollbackTask{
		FlowCtx: &liflow.FlowCtx{
			Ctx:                 ctx,
			FlowElementMap:      flow.Model.ElementMap(),
			FlowInstanceID:      flowInstance.ID,
			FlowInstanceStatus:  flowInstance.Status,
			SuspendNodeInstance: suspendNodeInstance,
			NodeInstanceList:    make([]ent.FlowNodeInstance, 0),
			InstanceDataID:      suspendNodeInstance.FlowInstanceDataID,
		},
	}
	rollbackNodeInstance, err := rt.getActiveUserTask(input.TaskInstanceID)
	if err != nil {
		return nil, gerror.WrapCode(liflow.ErrCodeRollbackFailed, err)
	}
	if rollbackNodeInstance == nil {
		return nil, gerror.NewCode(liflow.ErrCodeRollbackFailed)
	}
	if rt.isCompleted() {
		return nil, gerror.NewCode(liflow.ErrCodeRollbackRejected)
	}
	rt.CurrentNodeModel = rt.FlowElementMap[rt.SuspendNodeInstance.NodeKey]
	if rt.CurrentNodeModel == nil {
		return nil, gerror.NewCode(liflow.ErrCodeModelUnknownElementKey)
	}

	if rt.InstanceDataID != "" {
		instanceData, err := ent.DB().FlowInstanceData.Get(ctx, rt.InstanceDataID)
		if err != nil {
			return nil, gerror.WrapCode(liflow.ErrCodeGetInstanceDataFailed, err)
		}
		rt.InstanceDataMap = instanceData.Data
	}

	if err := rt.doRollback(); err != nil {
		return nil, gerror.WrapCode(liflow.ErrCodeRollbackFailed, err)
	} else {
		rt.ProcessStatus = liflow.ProcessStatusSuccess
		if err := rt.postRollback(); err != nil {
			return nil, err
		}
	}

	return &RollbackTaskOutput{
		FlowInstanceID:     flowInstance.ID,
		FlowInstanceStatus: rt.FlowInstanceStatus,
		ActiveTaskInstance: rt.CurrentNodeInstance,
	}, nil
}

func (rt *rollbackTask) doRollback() (err error) {
	executor := rt.getExecuteExecutor()
	for executor != nil {
		err = executor.Rollback(rt.FlowCtx)
		if err != nil {
			return err
		}
		// 用户任务节点挂起
		if rt.CurrentNodeModel.FlowType == liflow.FlowElementFlowTypeUserTask {
			return nil
		}
		executor, err = executor.GetRollbackExecutor(rt.FlowCtx)
		if err != nil {
			return
		}
	}
	return nil
}

func (rt *rollbackTask) postRollback() error {
	if rt.CurrentNodeInstance != nil {
		rt.SuspendNodeInstance = rt.CurrentNodeInstance
	}
	return rt.SaveNodeInstanceList(liflow.FlowNodeInstanceTypeRollback)
}

func (rt *rollbackTask) getActiveUserTask(suspendNodeInstanceID string) (*ent.FlowNodeInstance, error) {
	nodeInstanceList, err := ent.DB().FlowNodeInstance.
		Query().
		Where(flownodeinstance.FlowInstanceIDEQ(rt.FlowInstanceID)).
		Order(ent.Desc(flownodeinstance.FieldCreatedAt)).
		All(rt.Ctx)
	if err != nil {
		return nil, err
	}
	if len(nodeInstanceList) == 0 {
		return nil, nil
	}
	for _, node := range nodeInstanceList {
		if nodeModel := rt.FlowElementMap[node.NodeKey]; nodeModel == nil || nodeModel.FlowType != liflow.FlowElementFlowTypeUserTask {
			continue
		}
		if node.Status == liflow.FlowNodeInstanceStatusActive || node.Status == liflow.FlowNodeInstanceStatusCompleted {
			return node, nil
		}
	}
	return nil, err
}

func (rt *rollbackTask) isCompleted() bool {
	if rt.FlowInstanceStatus == liflow.FlowInstanceStatusCompleted {
		return true
	}
	if rt.SuspendNodeInstance == nil {
		return false
	}
	if rt.SuspendNodeInstance.Status != liflow.FlowNodeInstanceStatusCompleted {
		return false
	}
	if node := rt.FlowElementMap[rt.SuspendNodeInstance.NodeKey]; node != nil && node.FlowType == liflow.FlowElementFlowTypeEndEvent {
		return true
	}
	return false
}

func (rt *rollbackTask) getExecuteExecutor() liflow.Executor {
	if rt.isCompleted() {
		return nil
	}
	executor := liflow.GetElementExecutor(rt.CurrentNodeModel.Type)
	return executor
}
