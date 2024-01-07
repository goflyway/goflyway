package history

import (
	"github.com/jiangliuhong/go-flyway/consts"
	"github.com/jiangliuhong/go-flyway/database"
)

type SchemaHistory struct {
	Database database.Database
	Table    database.Table
}

func (sh SchemaHistory) Exists() (bool, error) {
	return sh.Table.Exists()
}

func (sh SchemaHistory) Create() error {
	return sh.Table.Create()
}

func New(d database.Database, tableName string) (*SchemaHistory, error) {
	schema, err := d.CurrentSchema()
	if err != nil {
		return nil, err
	}
	if tableName == "" {
		tableName = consts.DEFAULT_HISTORY_TABLE
	}
	table, err := schema.Table(tableName)
	if err != nil {
		return nil, err
	}
	return &SchemaHistory{Database: d, Table: table}, nil
}
