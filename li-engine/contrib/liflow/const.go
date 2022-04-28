package liflow

const (
	ProcessStatusSuccess          int8 = iota // 处理成功
	ProcessStatusReentrantWarning             // 重复处理
	ProcessStatusCommitSuspend                // 任务待提交
	ProcessStatusRollbackSuspend              // 任务待撤销
)

const (
	FlowDefinitionStatusEditing  int8 = iota // 编辑中
	FlowDefinitionStatusDeployed             // 已部署
	FlowDefinitionStatusDisabled             // 已下线, 暂未使用
)

const (
	FlowInstanceStatusDefault    int8 = iota // 数据库默认值
	FlowInstanceStatusCompleted              // 执行完成
	FlowInstanceStatusRunning                // 执行中
	FlowInstanceStatusTerminated             // 已终止
)

const (
	FlowInstanceDataTypeDefault  int8 = iota // 数据库默认值
	FlowInstanceDataTypeInit                 // 实例初始化
	FlowInstanceDataTypeExecute              // 系统执行
	FlowInstanceDataTypeHook                 // 系统主动获取
	FlowInstanceDataTypeUpdate               // 上游更新
	FlowInstanceDataTypeCommit               // 任务提交
	FlowInstanceDataTypeRollback             //任务回滚(暂时无用, 回滚时不产生新数据, 只修改数据版本号(dbId))
)

const (
	FlowNodeInstanceStatusDefault   int8 = iota // 数据库默认值
	FlowNodeInstanceStatusCompleted             // 处理成功
	FlowNodeInstanceStatusActive                // 处理中
	FlowNodeInstanceStatusFailed                // 处理失败
	FlowNodeInstanceStatusDisabled              // 处理已撤销
)

const (
	FlowNodeInstanceTypeDetaul   int8 = iota // 数据库默认值
	FlowNodeInstanceTypeExecute              // 系统执行
	FlowNodeInstanceTypeCommit               // 任务提交
	FlowNodeInstanceTypeRollback             // 任务撤销
)

const (
	FlowElementTypeSequenceFlow     = "SequenceFlow"
	FlowElementTypeStartEvent       = "StartEvent"
	FlowElementTypeEndEvent         = "EndEvent"
	FlowElementTypeUserTask         = "UserTask"
	FlowElementTypeExclusiveGateway = "ExclusiveGateway"
)
