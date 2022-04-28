package liflow

const (
	ProcessStatusSuccess          int8 = iota // 处理成功
	ProcessStatusReentrantWarning             // 重复处理
	ProcessStatusCommitSuspend                // 任务待提交
	ProcessStatusRollbackSuspend              // 任务待撤销
)

const (
	FlowInstanceStatusDefault    int8 = iota // 数据库默认值
	FlowInstanceStatusCompleted              // 执行完成
	FlowInstanceStatusRunning                // 执行中
	FlowInstanceStatusTerminated             // 已终止
)

const (
	FlowNodeInstanceStatusDefault   int8 = iota // 数据库默认值
	FlowNodeInstanceStatusCompleted             // 处理成功
	FlowNodeInstanceStatusActive                // 处理中
	FlowNodeInstanceStatusFailed                // 处理失败
	FlowNodeInstanceStatusDisabled              // 处理已撤销
)

const (
	FlowElementTypeStartEvent = "StartEvent"
	FlowElementTypeEndEvent   = "EndEvent"
)
