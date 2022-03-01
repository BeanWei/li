package liadl

import (
	"context"
	"runtime"

	"github.com/edgedb/edgedb-go"
	"github.com/gogf/gf/v2/frame/g"
)

var db *edgedb.Client

func dbsetup(ctx context.Context) (err error) {
	var dsn = g.Cfg().MustGet(ctx, "edgedb.default.link").String()
	db, err = edgedb.CreateClientDSN(ctx, dsn, edgedb.Options{
		Concurrency: uint(runtime.NumCPU()),
		TLSOptions: edgedb.TLSOptions{
			SecurityMode: edgedb.TLSModeInsecure,
		},
	})
	return
}

func DB() *edgedb.Client {
	return db
}
