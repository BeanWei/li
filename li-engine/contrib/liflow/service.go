package liflow

import (
	"context"

	"github.com/BeanWei/li/li-engine/contrib/liflow/ent"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/flowinstance"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/flownodeinstance"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

type (
	FlowElementInstance struct {
		NodeKey  string `json:"node_key"`
		NodeName string `json:"node_name"`
		Status   int8   `json:"status"`
	}
)

func GetHistoryUserTaskList(ctx context.Context, flowInstanceID string) ([]*ent.FlowNodeInstance, error) {
	historyNodeInstanceList, err := ent.DB().FlowNodeInstance.
		Query().
		Where(flownodeinstance.FlowInstanceIDEQ(flowInstanceID)).
		All(ctx)
	if err != nil || len(historyNodeInstanceList) == 0 {
		return historyNodeInstanceList, err
	}

	flowInstance, err := ent.DB().FlowInstance.Query().
		Where(flowinstance.IDEQ(flowInstanceID)).
		WithFlowDeployment().
		Only(ctx)
	if err != nil {
		return nil, err
	}
	if flowInstance.Edges.FlowDeployment == nil {
		return nil, gerror.NewCode(gcode.CodeInvalidRequest, "flow deployment is empty")
	}
	flowElementMap := flowInstance.Edges.FlowDeployment.Model.ElementMap()

	userTaskList := make([]*ent.FlowNodeInstance, 0)
	for _, node := range historyNodeInstanceList {
		if node.Status != FlowNodeInstanceStatusCompleted && node.Status != FlowNodeInstanceStatusActive {
			continue
		}
		if ele := flowElementMap[node.NodeKey]; ele.Type != FlowElementTypeUserTask {
			continue
		}
		userTaskList = append(userTaskList, node)
	}
	return userTaskList, nil
}

func GetHistoryElementList(ctx context.Context, flowInstanceID string) ([]*FlowElementInstance, error) {
	historyNodeInstanceList, err := ent.DB().FlowNodeInstance.
		Query().
		Where(flownodeinstance.FlowInstanceIDEQ(flowInstanceID)).
		All(ctx)
	if err != nil || len(historyNodeInstanceList) == 0 {
		return nil, err
	}

	flowInstance, err := ent.DB().FlowInstance.Query().
		Where(flowinstance.IDEQ(flowInstanceID)).
		WithFlowDeployment().
		Only(ctx)
	if err != nil {
		return nil, err
	}
	if flowInstance.Edges.FlowDeployment == nil {
		return nil, gerror.NewCode(gcode.CodeInvalidRequest, "flow deployment is empty")
	}
	flowElementMap := flowInstance.Edges.FlowDeployment.Model.ElementMap()

	elementList := make([]*FlowElementInstance, 0)
	for _, node := range historyNodeInstanceList {
		if sourceFlowElement := flowElementMap[node.SourceNodeKey]; sourceFlowElement != nil {
			sourceSequenceFlowStatus := node.Status
			if node.Status == FlowNodeInstanceStatusActive {
				sourceSequenceFlowStatus = FlowNodeInstanceStatusCompleted
			}
			elementList = append(elementList, &FlowElementInstance{
				NodeKey:  sourceFlowElement.Key,
				NodeName: gconv.String(sourceFlowElement.Properties[FlowElementPropertiesName]),
				Status:   sourceSequenceFlowStatus,
			})
		}
		if flowElement := flowElementMap[node.NodeKey]; flowElement != nil {
			elementList = append(elementList, &FlowElementInstance{
				NodeKey:  node.NodeKey,
				NodeName: gconv.String(flowElement.Properties[FlowElementPropertiesName]),
				Status:   node.Status,
			})
		}
	}
	return elementList, nil
}
