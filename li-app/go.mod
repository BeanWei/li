module github.com/BeanWei/li/li-app

go 1.16

require (
	entgo.io/ent v0.10.2-0.20220408082730-6e4e4da89674
	github.com/BeanWei/li/li-engine v0.0.0
	github.com/gogf/gf/contrib/drivers/pgsql/v2 v2.0.0-20220215154347-6ffdff70950e
	github.com/gogf/gf/v2 v2.0.4
	golang.org/x/crypto v0.0.0-20220321153916-2c7772ba3064
)

replace github.com/BeanWei/li/li-engine v0.0.0 => ../li-engine
