package history

import "github.com/jiangliuhong/go-flyway/database"

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

func New(d database.Database, t database.Table) *SchemaHistory {
	return &SchemaHistory{Database: d, Table: t}
}
