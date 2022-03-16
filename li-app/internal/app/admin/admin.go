package admin

import engine "github.com/BeanWei/li/li-engine"

func Init() {
	engine.NewApp(&engine.App{
		Title: "Li Admin",
	})
}
