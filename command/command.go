package command

import (
	"context"
	"errors"
	"fmt"
	"github.com/goflyway/goflyway/database"
	"github.com/goflyway/goflyway/history"
	"github.com/goflyway/goflyway/location"
	"github.com/goflyway/goflyway/logger"
)

type Command interface {
	// Execute 执行命令
	// schemaHistory 历史纪录
	// options 执行参数
	Execute(ctx *Context) error
}

// Context 执行命令的上下文对象
type Context struct {
	context.Context
	Command       string
	Database      database.Database
	SchemaHistory *history.SchemaHistory
	Options       *Options
	Logger        logger.Interface
}

var commands = map[string]Command{}
var dispatch = &CallbackDispatch{callbacks: map[string]*callback{}}

func Registry(name string, cmd Command) {
	commands[name] = cmd
}

func Callbacks() *CallbackDispatch {
	return dispatch
}

type Options struct {
	Locations         []location.Location // 文件信息
	OutOfOrder        bool                // 是否支持乱序
	EnablePlaceholder bool
	DisableCallbacks  bool
}

// Execute 执行命令
func Execute(ctx *Context) error {
	cmd, ok := commands[ctx.Command]
	if !ok {
		return errors.New(fmt.Sprintf("not found %s command", ctx.Command))
	}
	if !ctx.Options.DisableCallbacks {
		beforeHandlers := dispatch.before(ctx.Command)
		if len(beforeHandlers) > 0 {
			for _, h := range beforeHandlers {
				h.handler(ctx)
			}
		}
	}
	err := cmd.Execute(ctx)
	if err != nil {
		return err
	}
	if !ctx.Options.DisableCallbacks {
		afterHandlers := dispatch.after(ctx.Command)
		if len(afterHandlers) > 0 {
			for _, h := range afterHandlers {
				h.handler(ctx)
			}
		}
	}
	return nil
}

func beforeExecute(ctx *Context) (string, error) {
	exists, err := ctx.SchemaHistory.Exists()
	if err != nil {
		return "", err
	}
	if !exists {
		err = ctx.SchemaHistory.Create()
		if err != nil {
			return "", err
		}
	}
	err = ctx.SchemaHistory.Schema.UseSchema()
	if err != nil {
		return "", err
	}
	err = ctx.SchemaHistory.InitBaseLineRank()
	if err != nil {
		return "", err
	}
	latestVersion := ""
	if !ctx.Options.OutOfOrder {
		_, version, err := ctx.SchemaHistory.GetLatestVersion()
		if err != nil {
			return "", err
		}
		latestVersion = version
	}
	return latestVersion, nil
}
