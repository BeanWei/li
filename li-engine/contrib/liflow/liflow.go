package liflow

import "github.com/BeanWei/li/li-engine/contrib/liflow/ent/schema"

type (
	Executor interface {
		Validate(ele *schema.FlowElement) error
		Execute(ctx *FlowCtx) error
		Commit(ctx *FlowCtx) error
		Rollback(ctx *FlowCtx) error
		IsCompleted(ctx *FlowCtx) (bool, error)
		GetExecuteExecutor(ctx *FlowCtx) Executor
	}
)

type (
	RegisterElementInput struct {
		ElementName     string
		ElementExecutor Executor
	}
)

var (
	eleMap = make(map[string]Executor)
)

func RegisterElement(eles ...*RegisterElementInput) {
	for _, ele := range eles {
		eleMap[ele.ElementName] = ele.ElementExecutor
	}
}

func GetElementExecutor(name string) Executor {
	return eleMap[name]
}
