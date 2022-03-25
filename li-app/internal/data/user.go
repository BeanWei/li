package data

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/BeanWei/li/li-app/internal/data/ent"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gcache"
	"golang.org/x/crypto/pbkdf2"
)

var User = new(usersvc)

type usersvc struct{}

func (usersvc) HashPassword(pwd, salt string) string {
	passwd := pbkdf2.Key([]byte(pwd), []byte(salt), 10000, 50, sha256.New)
	return fmt.Sprintf("%x", passwd)
}

func (usersvc) GetUser(ctx context.Context, uid string) (usr *ent.User, err error) {
	val, err := gcache.GetOrSetFunc(ctx, uid, func(ctx context.Context) (interface{}, error) {
		return ent.LiClient().User.Get(ctx, uid)
	}, 0)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err)
	}
	if err = val.Struct(&usr); err != nil {
		return nil, err
	}
	return
}

func (usersvc) RemoveUserCache(ctx context.Context, uids ...string) (err error) {
	_, err = gcache.Remove(ctx, uids)
	return
}
