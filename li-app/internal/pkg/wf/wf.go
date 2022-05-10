package wf

import (
	"github.com/BeanWei/li/li-app/internal/pkg/wf/executor"
	"github.com/BeanWei/li/li-engine/contrib/liflow"
)

func init() {
	liflow.RegisterExecutor(
		&liflow.RegisterExecutorInput{
			ElementType:     "ApprovalTask",
			ElementExecutor: new(executor.ApprovalTask),
		},
	)
}
