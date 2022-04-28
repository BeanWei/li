package liflow

import "github.com/BeanWei/li/li-engine/contrib/liflow/ent/schema"

type (
	Executor interface {
		Validate(ele *schema.FlowElement) error
		Execute(ctx *FlowCtx) error
		Commit(ctx *FlowCtx) error
		Rollback(ctx *FlowCtx) error
		IsCompleted(ctx *FlowCtx) bool
		GetExecuteExecutor(ctx *FlowCtx) Executor
	}
)

type (
	RegisterExecutorInput struct {
		ElementType     string
		ElementExecutor Executor
	}
)

var (
	executorMap = make(map[string]Executor)
)

func RegisterExecutor(eles ...*RegisterExecutorInput) {
	for _, ele := range eles {
		executorMap[ele.ElementType] = ele.ElementExecutor
	}
}

func GetElementExecutor(name string) Executor {
	return executorMap[name]
}
