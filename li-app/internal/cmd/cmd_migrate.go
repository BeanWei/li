package cmd

import (
	"context"

	"github.com/BeanWei/li/li-app/internal/data/ent"
	entmigrate "github.com/BeanWei/li/li-app/internal/data/ent/migrate"
	"github.com/gogf/gf/v2/os/gcmd"
)

func migrate(ctx context.Context, parser *gcmd.Parser) (err error) {
	return ent.LiClient().Schema.Create(ctx,
		entmigrate.WithForeignKeys(false),
		entmigrate.WithDropIndex(true),
		entmigrate.WithDropColumn(true),
	)
}
