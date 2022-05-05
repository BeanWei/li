package liflow

import (
	"context"
	"fmt"

	"github.com/BeanWei/li/li-engine/contrib/liflow/ent"
	"github.com/BeanWei/li/li-engine/contrib/liflow/ent/hook"
)

func init() {
	ent.DB().FlowDeployment.Use(
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.FlowDeploymentFunc(func(ctx context.Context, m *ent.FlowDeploymentMutation) (ent.Value, error) {
					if flowModel, exists := m.Model(); exists {
						flowElementMap := flowModel.ElementMap()
						for _, ele := range flowModel {
							if ele == nil {
								return nil, fmt.Errorf("flow element cannot be empty")
							}
							if executor := executorMap[ele.Type]; executor != nil {
								if err := executor.Validate(flowElementMap, ele); err != nil {
									return nil, err
								}
							} else {
								return nil, fmt.Errorf("flow element type %s not register executor", ele.Type)
							}
						}
					}
					return next.Mutate(ctx, m)
				})
			},
			ent.OpCreate|ent.OpUpdateOne,
		),
	)
}
