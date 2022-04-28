package processor

import (
	"github.com/BeanWei/li/li-engine/contrib/liflow"
	"github.com/BeanWei/li/li-engine/contrib/liflow/executor"
)

func init() {
	liflow.RegisterExecutor(&liflow.RegisterExecutorInput{
		ElementType:     liflow.FlowElementTypeStartEvent,
		ElementExecutor: new(executor.StartEventExecutor),
	})
}
