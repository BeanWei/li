package element

import (
	"github.com/BeanWei/li/li-engine/contrib/liflow"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/schema"
)

type ElementExecutor struct{}

func (e *ElementExecutor) Validate(ele *schema.FlowElement) error {
	return nil
}

func (e *ElementExecutor) Execute(ctx *liflow.FlowCtx) error {
	return nil
}

func (e *ElementExecutor) Commit(ctx *liflow.FlowCtx) error {
	return nil
}

func (e *ElementExecutor) Rollback(ctx *liflow.FlowCtx) error {
	return nil
}

func (e *ElementExecutor) IsCompleted(ctx *liflow.FlowCtx) (bool, error) {
	return false, nil
}

func (e *ElementExecutor) GetExecuteExecutor(ctx *liflow.FlowCtx) liflow.Executor {
	return nil
}

func (e *ElementExecutor) PreExecute() error {
	return nil
}

func (e *ElementExecutor) DoExecute() error {
	return nil
}

func (e *ElementExecutor) PostExecute() error {
	return nil
}

func (e *ElementExecutor) PreCommit() error {
	return nil
}

func (e *ElementExecutor) DoCommit() error {
	return nil
}

func (e *ElementExecutor) PostCommit() error {
	return nil
}

func (e *ElementExecutor) PreRollback() error {
	return nil
}

func (e *ElementExecutor) DoRollback() error {
	return nil
}

func (e *ElementExecutor) PostRollback() error {
	return nil
}

var _ liflow.Executor = (*ElementExecutor)(nil)
