package command

import (
	"errors"
	"fmt"
	"github.com/goflyway/goflyway/database"
	"github.com/goflyway/goflyway/history"
	"github.com/goflyway/goflyway/location"
)

type Command interface {
	// Execute 执行命令
	// schemaHistory 历史纪录
	// options 执行参数
	Execute(ctx *Context) error
}

// Context 执行命令的上下文对象
type Context struct {
	Command       string
	Database      database.Database
	SchemaHistory *history.SchemaHistory
	Options       *Options
}

var commands = map[string]Command{}

func Registry(name string, cmd Command) {
	commands[name] = cmd
}

type Options struct {
	Locations []location.Location // 文件信息
}

// Execute 执行命令
func Execute(ctx *Context) error {
	cmd, ok := commands[ctx.Command]
	if !ok {
		return errors.New(fmt.Sprintf("not found %s command", ctx.Command))
	}
	return cmd.Execute(ctx)
}
