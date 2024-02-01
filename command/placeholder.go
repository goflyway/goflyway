package command

import (
	"github.com/goflyway/goflyway/location"
	"time"
)

func defaultPlaceholderEnv(ctx *Context) map[string]interface{} {
	return map[string]interface{}{
		"defaultSchema": ctx.SchemaHistory.Schema.Name(),
		"user":          ctx.Database.CurrentUser(),
		"timestamp":     time.Now().Format(time.DateTime),
		"filename":      "",
		"table":         ctx.SchemaHistory.Table.Name(),
	}
}

func GenSqlPlaceholderEnv(ctx *Context, sql location.SqlFile) map[string]interface{} {
	env := defaultPlaceholderEnv(ctx)
	env["filename"] = sql.Name
	return env
}
