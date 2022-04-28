package processor

import (
	"github.com/BeanWei/li/li-engine/contrib/liflow"
	"github.com/BeanWei/li/li-engine/contrib/liflow/element"
)

func init() {
	liflow.RegisterElement(&liflow.RegisterElementInput{
		ElementName:     liflow.FlowElementTypeStartEvent,
		ElementExecutor: new(element.StartEvent),
	})
}
