package command

import (
	"github.com/goflyway/goflyway/location"
	"time"
)

func defaultPlaceholderEnv(ctx *Context) (map[string]interface{}, error) {
	user, err := ctx.Database.CurrentUser()
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"defaultSchema": ctx.SchemaHistory.Schema.Name(),
		"user":          user,
		"timestamp":     time.Now().Format(time.DateTime),
		"filename":      "",
		"table":         ctx.SchemaHistory.Table.Name(),
	}, nil
}

func GenSqlPlaceholderEnv(ctx *Context, sql location.SqlFile) (map[string]interface{}, error) {
	env, err := defaultPlaceholderEnv(ctx)
	if err != nil {
		return nil, err
	}
	env["filename"] = sql.Name
	return env, nil
}
