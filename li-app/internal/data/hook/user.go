package hook

import (
	"context"

	"github.com/BeanWei/li/li-app/internal/data"
	"github.com/BeanWei/li/li-app/internal/data/ent"
	"github.com/BeanWei/li/li-app/internal/data/ent/hook"
	"github.com/BeanWei/li/li-app/internal/data/ent/user"
	"github.com/gogf/gf/v2/util/grand"
)

func init() {
	ent.LiClient().User.Use(
		hook.If(
			func(next ent.Mutator) ent.Mutator {
				return hook.UserFunc(func(ctx context.Context, m *ent.UserMutation) (ent.Value, error) {
					password, _ := m.Password()
					salt := grand.S(10)
					m.SetSalt(salt)
					m.SetPassword(data.User.HashPassword(password, salt))
					return next.Mutate(ctx, m)
				})
			},
			hook.And(
				hook.HasOp(ent.OpCreate),
				hook.HasFields(user.FieldPassword),
			),
		),
	)
}
