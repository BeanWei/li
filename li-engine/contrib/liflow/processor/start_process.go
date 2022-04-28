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
	startProcessor struct {
		*liflow.FlowCtx
	}
	StartProcessInput struct {
		FlowDefinitionID string
		Variables        map[string]interface{}
	}
	StartProcessOutput struct {
		ProcessStatus      int8                   `json:"process_status"`
		FlowDefinitionID   string                 `json:"flow_definition_id"`
		FlowInstanceID     string                 `json:"flow_instance_id"`
		FlowInstanceStatus int8                   `json:"flow_instance_status"`
		ActiveNodeInstance *ent.FlowNodeInstance  `json:"active_node_instance"`
		Variables          map[string]interface{} `json:"variables"`
	}
)

// StartProcess 流程执行
// 创建流程实例，从开始节点开始执行，直到用户任务节点挂起或者结束节点完成。
func StartProcess(ctx context.Context, input *StartProcessInput) (*StartProcessOutput, error) {
	// 1: 获取流程信息
	flow, err := ent.DB().FlowDefinition.Get(ctx, input.FlowDefinitionID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, gerror.WrapCodef(gcode.CodeInvalidParameter, err, "not exists flow definition: %s", input.FlowDefinitionID)
		}
		return nil, gerror.Wrapf(err, "[processor.StartProcess] find flow definition by id %s", input.FlowDefinitionID)
	}
	spc := &startProcessor{
		FlowCtx: &liflow.FlowCtx{
			Ctx:            ctx,
			FlowElementMap: flow.Model.ElementMap(),
		},
	}
	// 2: 初始化全局数据

	// 3: 执行流程
	flowInstance, err := ent.DB().FlowInstance.Create().
		SetFlowDefinitionID(flow.ID).
		SetStatus(liflow.FlowInstanceStatusRunning).
		Save(ctx)
	if err != nil {
		return nil, gerror.Wrapf(err, "[processor.StartProcess] save flowInstance", input.FlowDefinitionID)
	}
	spc.FlowInstanceID = flowInstance.ID
	spc.FlowInstanceStatus = flowInstance.Status

	var startEvent *schema.FlowElement
	for _, ele := range flow.Model {
		if ele.Type == liflow.FlowElementTypeStartEvent {
			startEvent = ele
			break
		}
	}
	if startEvent == nil {
		return nil, gerror.Newf("cannot get startEvent node from %s flowDefinition", flow.ID)
	}
	spc.CurrentNodeModel = startEvent
	spc.SuspendNodeInstance = &ent.FlowNodeInstance{
		NodeKey: startEvent.Key,
		Status:  liflow.FlowNodeInstanceStatusActive,
	}

	spc.doExecute()
	spc.postExecute()

	return &StartProcessOutput{
		ProcessStatus:      spc.ProcessStatus,
		FlowDefinitionID:   flow.ID,
		FlowInstanceID:     flowInstance.ID,
		FlowInstanceStatus: flowInstance.Status,
		ActiveNodeInstance: spc.CurrentNodeInstance,
	}, nil
}

func (spc *startProcessor) doExecute() {
	executor := spc.getExecuteExecutor()
	for executor != nil {
		executor.Execute(spc.FlowCtx)
		if spc.ProcessStatus == liflow.ProcessStatusSuccess || spc.ProcessStatus == liflow.ProcessStatusCommitSuspend {
			return
		}
		executor = executor.GetExecuteExecutor(spc.FlowCtx)
	}
}

func (spc *startProcessor) postExecute() {
	if spc.ProcessStatus == liflow.ProcessStatusSuccess {
		if spc.CurrentNodeInstance != nil {
			spc.SuspendNodeInstance = spc.CurrentNodeInstance
		}
	}
	// TODO: 保存全部执行节点实例到数据库

	// 更新流程实例状态
	if spc.isCompleted() {
		ent.DB().FlowInstance.
			UpdateOneID(spc.FlowInstanceID).
			SetStatus(liflow.FlowInstanceStatusCompleted).
			ExecX(spc.Ctx)
		spc.FlowInstanceStatus = liflow.FlowInstanceStatusCompleted
	}
}

func (spc *startProcessor) isCompleted() bool {
	if spc.FlowInstanceStatus == liflow.FlowInstanceStatusCompleted {
		return true
	}
	if spc.SuspendNodeInstance == nil {
		return false
	}
	if spc.SuspendNodeInstance.Status != liflow.FlowNodeInstanceStatusCompleted {
		return false
	}
	if node := spc.FlowElementMap[spc.SuspendNodeInstance.NodeKey]; node != nil && node.Type == liflow.FlowElementTypeEndEvent {
		return true
	}
	return false
}

func (spc *startProcessor) getExecuteExecutor() liflow.Executor {
	if spc.isCompleted() {
		return nil
	}
	executor := liflow.GetElementExecutor(spc.CurrentNodeModel.Type)
	return executor
}
