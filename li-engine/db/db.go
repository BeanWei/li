package db

import (
	"context"
	"runtime"
	"sync"

	"github.com/edgedb/edgedb-go"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	edgedbclient     *edgedb.Client
	edgedbclientOnce sync.Once
)

func clien() *edgedb.Client {
	edgedbclientOnce.Do(func() {
		var (
			err error
			ctx = gctx.New()
			dsn = g.Cfg().MustGet(ctx, "edgedb.default.link").String()
		)
		edgedbclient, err = edgedb.CreateClientDSN(ctx, dsn, edgedb.Options{
			Concurrency: uint(runtime.NumCPU()),
			TLSOptions: edgedb.TLSOptions{
				SecurityMode: edgedb.TLSModeInsecure,
			},
		})
		if err != nil {
			panic(err)
		}
	})
	return edgedbclient
}

func Exec(ctx context.Context, cmd string, args ...interface{}) (res []byte, err error) {
	err = clien().QuerySingleJSON(ctx, cmd, &res, args...)
	return
}
