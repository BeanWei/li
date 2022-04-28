package liflow

import (
	"context"

	"github.com/BeanWei/li/li-engine/contrib/liflow/ent"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/schema"
)

type FlowCtx struct {
	Ctx                 context.Context
	ProcessStatus       int8
	FlowElementMap      map[string]*schema.FlowElement
	CurrentNodeModel    *schema.FlowElement
	FlowInstanceID      string
	FlowInstanceStatus  int8
	NodeInstanceList    []*ent.FlowNodeInstance
	CurrentNodeInstance *ent.FlowNodeInstance
	SuspendNodeInstance *ent.FlowNodeInstance
}

func (ctx *FlowCtx) SaveNodeInstanceList() error {
	return nil
}
