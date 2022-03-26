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
		// 生成盐值和加密密码
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
		// 移除缓存中的用户信息
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.UserFunc(func(ctx context.Context, m *ent.UserMutation) (ent.Value, error) {
					val, err := next.Mutate(ctx, m)
					if err != nil {
						return val, err
					}
					var ids []string
					if id, ok := m.ID(); ok {
						ids = append(ids, id)
					}
					data.User.RemoveUserCache(ctx, ids...)
					return val, err
				})
			},
			ent.OpUpdate|ent.OpUpdateOne|ent.OpDelete|ent.OpDeleteOne,
		),
	)
}
