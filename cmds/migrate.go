package cmds

import (
	"fmt"
	"github.com/jiangliuhong/go-flyway/consts"
	"github.com/jiangliuhong/go-flyway/database"
	"github.com/jiangliuhong/go-flyway/history"
)

func init() {
	Registry(consts.CMD_NAME_MIGRATE, &Migrate{})
}

type Migrate struct {
}

func (m Migrate) Execute(database database.Database, schemaHistory *history.SchemaHistory, options *Options) error {
	exists, err := schemaHistory.Exists()
	if err != nil {
		return err
	}
	if !exists {
		err = schemaHistory.Create()
		if err != nil {
			return err
		}
	}
	for _, location := range options.Locations {
		for _, sql := range location.Sqls {
			fmt.Print(sql.Name)
		}
	}
	return nil
}
