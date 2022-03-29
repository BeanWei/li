module github.com/BeanWei/li/li-engine

go 1.16

require (
	entgo.io/ent v0.10.1
	github.com/BeanWei/li/li-pkg/file v0.0.0
	github.com/gogf/gf/v2 v2.0.4
	github.com/hexops/valast v1.4.1
	github.com/rs/xid v1.4.0
)

replace github.com/BeanWei/li/li-pkg/file v0.0.0 => ../li-pkg/file
