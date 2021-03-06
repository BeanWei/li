package processor

import (
	"context"
	"time"

	"github.com/BeanWei/li/li-engine/contrib/lient"
	"github.com/BeanWei/li/li-engine/contrib/liflow"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/schema"
	"github.com/gogf/gf/v2/errors/gerror"
)

type (
	startProcess struct {
		*liflow.FlowCtx
	}
	StartProcessInput struct {
		RefID            string                 `json:"ref_id" v:"required"`
		FlowDeploymentID string                 `json:"flow_deployment_id" v:"required"`
		Variables        map[string]interface{} `json:"variables"`
	}
	StartProcessOutput struct {
		ProcessStatus      int8                   `json:"process_status"`
		FlowDeploymentID   string                 `json:"flow_deployment_id"`
		FlowInstanceID     string                 `json:"flow_instance_id"`
		FlowInstanceStatus int8                   `json:"flow_instance_status"`
		ActiveNodeInstance *ent.FlowNodeInstance  `json:"active_node_instance"`
		Variables          map[string]interface{} `json:"variables"`
	}
)

// StartProcess 流程执行
// 创建流程实例，从开始节点开始执行，直到用户任务节点挂起或者结束节点完成。
func StartProcess(ctx context.Context, input *StartProcessInput) (*StartProcessOutput, error) {
	flow, err := ent.DB().FlowDeployment.Get(ctx, input.FlowDeploymentID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, gerror.NewCode(liflow.ErrCodeFlowNotExists)
		}
		return nil, gerror.Wrapf(err, "[processor.StartProcess] find flow deployment by id %s", input.FlowDeploymentID)
	}
	sp := &startProcess{
		FlowCtx: &liflow.FlowCtx{
			Ctx:              ctx,
			FlowElementMap:   flow.Model.ElementMap(),
			NodeInstanceList: make([]ent.FlowNodeInstance, 0),
		},
	}

	flowInstance, err := ent.DB().FlowInstance.
		Create().
		SetFlowDeploymentID(input.FlowDeploymentID).
		SetRefID(input.RefID).
		SetStatus(liflow.FlowInstanceStatusRunning).
		Save(ctx)
	if err != nil {
		return nil, gerror.WrapCode(liflow.ErrCodeSaveFlowInstanceFailed, err)
	}
	sp.FlowInstanceID = flowInstance.ID
	sp.FlowInstanceStatus = flowInstance.Status

	flowInstanceData, err := ent.DB().FlowInstanceData.
		Create().
		SetFlowInstanceID(flowInstance.ID).
		SetData(input.Variables).
		SetType(liflow.FlowInstanceDataTypeInit).
		Save(ctx)
	if err != nil {
		return nil, gerror.WrapCode(liflow.ErrCodeSaveInstanceDataFailed, err)
	}
	sp.InstanceDataID = flowInstanceData.ID
	sp.InstanceDataMap = input.Variables

	var startEvent *schema.FlowElement
	for _, ele := range flow.Model {
		if ele.FlowType == liflow.FlowElementFlowTypeStartEvent {
			startEvent = ele
			break
		}
	}
	if startEvent == nil {
		return nil, gerror.NewCode(liflow.ErrCodeStartNodeInvalid)
	}
	sp.CurrentNodeModel = startEvent
	sp.SuspendNodeInstance = &ent.FlowNodeInstance{
		NodeKey: startEvent.Key,
		Status:  liflow.FlowNodeInstanceStatusActive,
	}

	// TODO: 2022/04/30 记录错误日志
	if err := sp.doExecute(); err != nil {
		sp.ProcessStatus = liflow.ProcessStatusFailed
	} else {
		sp.ProcessStatus = liflow.ProcessStatusSuccess
	}
	if err := sp.postExecute(); err != nil {
		return nil, err
	}

	return &StartProcessOutput{
		ProcessStatus:      sp.ProcessStatus,
		FlowDeploymentID:   flowInstance.FlowDeploymentID,
		FlowInstanceID:     flowInstance.ID,
		FlowInstanceStatus: flowInstance.Status,
		ActiveNodeInstance: sp.CurrentNodeInstance,
	}, nil
}

func (sp *startProcess) doExecute() (err error) {
	executor := sp.getExecuteExecutor()
	for executor != nil {
		currentNodeInstance := &ent.FlowNodeInstance{
			ID:                 lient.NewXid(),
			CreatedAt:          time.Now().Unix(),
			UpdatedAt:          time.Now().Unix(),
			FlowInstanceID:     sp.FlowInstanceID,
			NodeKey:            sp.CurrentNodeModel.Key,
			Status:             liflow.FlowNodeInstanceStatusActive,
			FlowInstanceDataID: sp.InstanceDataID,
		}
		if sp.CurrentNodeInstance != nil {
			currentNodeInstance.SourceFlowNodeInstanceID = sp.CurrentNodeInstance.ID
			currentNodeInstance.SourceNodeKey = sp.CurrentNodeInstance.NodeKey
		}
		sp.CurrentNodeInstance = currentNodeInstance
		err := executor.Execute(sp.FlowCtx)
		if err != nil {
			return err
		}
		// 用户任务节点挂起
		if sp.CurrentNodeModel.FlowType == liflow.FlowElementFlowTypeUserTask {
			return nil
		}
		executor = executor.GetExecuteExecutor(sp.FlowCtx)
	}
	return nil
}

func (sp *startProcess) postExecute() error {
	if sp.ProcessStatus == liflow.ProcessStatusSuccess && sp.CurrentNodeInstance != nil {
		sp.SuspendNodeInstance = sp.CurrentNodeInstance
	}
	if err := sp.SaveNodeInstanceList(liflow.FlowNodeInstanceTypeExecute); err != nil {
		return err
	}
	// 更新流程实例状态
	if sp.isCompleted() {
		sp.FlowInstanceStatus = liflow.FlowInstanceStatusCompleted
		return ent.DB().FlowInstance.
			UpdateOneID(sp.FlowInstanceID).
			SetStatus(liflow.FlowInstanceStatusCompleted).
			Exec(sp.Ctx)
	}
	return nil
}

func (sp *startProcess) isCompleted() bool {
	if sp.FlowInstanceStatus == liflow.FlowInstanceStatusCompleted {
		return true
	}
	if sp.SuspendNodeInstance == nil {
		return false
	}
	if sp.SuspendNodeInstance.Status != liflow.FlowNodeInstanceStatusCompleted {
		return false
	}
	if node := sp.FlowElementMap[sp.SuspendNodeInstance.NodeKey]; node != nil && node.FlowType == liflow.FlowElementFlowTypeEndEvent {
		return true
	}
	return false
}

func (sp *startProcess) getExecuteExecutor() liflow.Executor {
	if sp.isCompleted() {
		return nil
	}
	executor := liflow.GetElementExecutor(sp.CurrentNodeModel.Type)
	return executor
}
