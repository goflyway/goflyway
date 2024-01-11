package sqlite

import (
	"fmt"
	"github.com/jiangliuhong/go-flyway/database"
)

func init() {
	database.Registry(database.SQLITE, func(db *database.Session) (database.Database, error) {
		return &sqlite{
			database.BaseDatabase{DB: db},
		}, nil
	})
}

type sqlite struct {
	database.BaseDatabase
}

func (d sqlite) CurrentSchema() (database.Schema, error) {
	return &sqliteSchema{Schema: "main", db: d.DB, Database: d}, nil
}

func (d sqlite) CurrentUser() (string, error) {
	return "main", nil
}

func (d sqlite) Schema(name string) (database.Schema, error) {
	return nil, nil
}

func (d sqlite) Type() database.Type {
	return database.SQLITE
}

type sqliteSchema struct {
	db       *database.Session
	Schema   string
	Database database.Database
}

func (s sqliteSchema) Name() string {
	return s.Schema
}
func (s sqliteSchema) Exists() (bool, error) {
	return true, nil
}
func (s sqliteSchema) Create() error {
	return nil
}
func (s sqliteSchema) Table(name string) (database.Table, error) {
	return &sqliteTable{db: s.db, BaseTable: database.BaseTable{Table: name, Schema: s, Database: s.Database}}, nil
}

type sqliteTable struct {
	database.BaseTable
	db *database.Session
}

func (t sqliteTable) Exists() (bool, error) {
	sql := fmt.Sprintf(`select count(tbl_name) FROM %s.sqlite_master where type = 'table' and tbl_name = '%s'`, t.Schema.Name(), t.Name())
	count, err := t.db.Count(sql)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
