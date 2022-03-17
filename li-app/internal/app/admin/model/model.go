package model

import (
	"fmt"

	"github.com/BeanWei/li/li-app/internal/db"
	"github.com/BeanWei/li/li-engine/model"
	"github.com/gogf/gf/v2/os/gctx"
)

// Migrate model
func init() {
	dbschema := model.ToDbSchema(
		new(User),
	)
	err := db.EdgeDB().Execute(gctx.New(), fmt.Sprintf(`
		START MIGRATION TO {
			%s
		};
		POPULATE MIGRATION;
		COMMIT MIGRATION;
	`, dbschema))
	if err != nil {
		panic(err)
	}
}
