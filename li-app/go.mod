module github.com/BeanWei/li/li-app

go 1.16

require (
	entgo.io/ent v0.10.2-0.20220510065209-fc03c8d283d3
	github.com/BeanWei/li/li-engine v0.0.0
	github.com/gogf/gf/contrib/drivers/pgsql/v2 v2.0.0-20220427091526-ae5891068e0e
	github.com/gogf/gf/v2 v2.0.6
	golang.org/x/crypto v0.0.0-20220321153916-2c7772ba3064
)

replace github.com/BeanWei/li/li-engine v0.0.0 => ../li-engine
