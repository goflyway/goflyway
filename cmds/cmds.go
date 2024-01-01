package cmds

import (
	"errors"
	"github.com/jiangliuhong/go-flyway/database"
	"github.com/jiangliuhong/go-flyway/history"
	"github.com/jiangliuhong/go-flyway/location"
)

type Command interface {
	// Execute 执行命令
	// schemaHistory 历史纪录
	// options 执行参数
	Execute(database database.Database, schemaHistory *history.SchemaHistory, options *Options) error
}

var commands = map[string]Command{}

func Registry(name string, cmd Command) {
	commands[name] = cmd
}

type Options struct {
	Locations []location.Location // 文件信息
}

func Execute(command string, database database.Database, schemaHistory *history.SchemaHistory, options *Options) error {
	cmd, ok := commands[command]
	if !ok {
		return errors.New("not found " + command + " command")
	}
	return cmd.Execute(database, schemaHistory, options)
}
