package processor

import (
	"github.com/BeanWei/li/li-engine/contrib/liflow"
	"github.com/BeanWei/li/li-engine/contrib/liflow/executor"
)

func init() {
	liflow.RegisterExecutor(
		&liflow.RegisterExecutorInput{
			ElementType:     "StartEvent",
			ElementExecutor: new(executor.StartEventExecutor),
		},
		&liflow.RegisterExecutorInput{
			ElementType:     "EndEvent",
			ElementExecutor: new(executor.EndEventExecutor),
		},
	)
}
