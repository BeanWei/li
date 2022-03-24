package driver

import (
	"context"
	"database/sql"

	"entgo.io/ent/dialect"
)

type multiDriver struct {
	master dialect.Driver
	slave  dialect.Driver
}

type DriverOption func(md *multiDriver)

func WithMasterDriver(d dialect.Driver) DriverOption {
	return func(md *multiDriver) {
		md.master = d
	}
}

func WithSlaveDriver(d dialect.Driver) DriverOption {
	return func(md *multiDriver) {
		md.slave = d
	}
}

func NewDriver(opts ...DriverOption) (*multiDriver, error) {
	master, err := Master()
	if err != nil {
		return nil, err
	}

	slave, err := Slave()
	if err != nil {
		return nil, err
	}

	md := &multiDriver{
		master: master,
		slave:  slave,
	}
	for _, opt := range opts {
		opt(md)
	}
	return md, nil
}

var _ dialect.Driver = (*multiDriver)(nil)

func (d *multiDriver) Dialect() string {
	return d.master.Dialect()
}

func (d *multiDriver) Query(ctx context.Context, query string, args, v interface{}) error {
	return d.slave.Query(ctx, query, args, v)
}

func (d *multiDriver) Exec(ctx context.Context, query string, args, v interface{}) error {
	return d.master.Exec(ctx, query, args, v)
}

func (d *multiDriver) Tx(ctx context.Context) (dialect.Tx, error) {
	return d.master.Tx(ctx)
}

func (d *multiDriver) BeginTx(ctx context.Context, opts *sql.TxOptions) (dialect.Tx, error) {
	return d.master.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
}

func (d *multiDriver) Close() error {
	merr := d.master.Close()
	serr := d.slave.Close()
	if merr != nil {
		return merr
	}
	if serr != nil {
		return serr
	}
	return nil
}
