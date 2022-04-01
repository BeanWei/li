package driver

import (
	gosql "database/sql"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func getDriver(master bool, name ...string) (drv *sql.Driver, err error) {
	var sqlDB *gosql.DB
	if master {
		sqlDB, err = g.DB(name...).Master()
	} else {
		sqlDB, err = g.DB(name...).Slave()
	}
	if err != nil {
		return nil, err
	}
	if sqlDB == nil {
		return nil, gerror.New("failed to connect master/slave db")
	}

	driver := g.DB(name...).GetConfig().Type
	if driver == "pgsql" {
		driver = dialect.Postgres
	}

	drv = sql.OpenDB(driver, sqlDB)
	return
}

// Master 主数据库驱动
func Master(name ...string) (drv *sql.Driver, err error) {
	return getDriver(true, name...)
}

// Slave 从数据库驱动
func Slave(name ...string) (drv *sql.Driver, err error) {
	return getDriver(false, name...)
}
