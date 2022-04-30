package processor

import (
	"context"

	"github.com/BeanWei/li/li-engine/contrib/liflow"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/schema"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type (
	startProcess struct {
		*liflow.FlowCtx
	}
	StartProcessInput struct {
		RefID            string
		FlowDeploymentID string
		Variables        map[string]interface{}
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
	// 1: 获取流程模型
	flow, err := ent.DB().FlowDeployment.Get(ctx, input.FlowDeploymentID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, gerror.WrapCodef(gcode.CodeInvalidParameter, err, "not exists flow deployment: %s", input.FlowDeploymentID)
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
	// 2: 初始化全局数据

	// 3: 执行流程
	flowInstance, err := ent.DB().FlowInstance.Create().
		SetFlowDeploymentID(input.FlowDeploymentID).
		SetRefID(input.RefID).
		SetStatus(liflow.FlowInstanceStatusRunning).
		Save(ctx)
	if err != nil {
		return nil, gerror.Wrap(err, "[processor.StartProcess] save flowInstance")
	}
	sp.FlowInstanceID = flowInstance.ID
	sp.FlowInstanceStatus = flowInstance.Status

	var startEvent *schema.FlowElement
	for _, ele := range flow.Model {
		if ele.FlowType == liflow.FlowElementFlowTypeStartEvent {
			startEvent = ele
			break
		}
	}
	if startEvent == nil {
		return nil, gerror.Newf("cannot get startEvent node from %s flowDefinition", flow.ID)
	}
	sp.CurrentNodeModel = startEvent
	sp.SuspendNodeInstance = &ent.FlowNodeInstance{
		NodeKey: startEvent.Key,
		Status:  liflow.FlowNodeInstanceStatusActive,
	}

	if err := sp.doExecute(); err != nil {
		return nil, err
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

func (sp *startProcess) doExecute() error {
	executor := sp.getExecuteExecutor()
	for executor != nil {
		err := executor.Execute(sp.FlowCtx)
		if err != nil {
			return err
		}
		if sp.ProcessStatus == liflow.ProcessStatusSuccess || sp.ProcessStatus == liflow.ProcessStatusCommitSuspend {
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