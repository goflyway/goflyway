package sqlite

import (
	"errors"
	"fmt"
	"github.com/goflyway/goflyway/database"
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
	return &sqliteSchema{
		BaseSchema: database.BaseSchema{Schema: "main"},
		db:         d.DB,
		Database:   d,
	}, nil
}

func (d sqlite) CurrentUser() (string, error) {
	return "main", nil
}

func (d sqlite) Schema(name string) (database.Schema, error) {
	return &sqliteSchema{
		BaseSchema: database.BaseSchema{Schema: name},
		db:         d.DB,
		Database:   d,
	}, nil
}

func (d sqlite) Type() database.Type {
	return database.SQLITE
}

type sqliteSchema struct {
	database.BaseSchema
	db       *database.Session
	Database database.Database
}

func (s sqliteSchema) Exists() (bool, error) {
	return true, nil
}
func (s sqliteSchema) Create() error {
	return errors.New("sqlite does not support schema creation")
}
func (s sqliteSchema) Table(name string) (database.Table, error) {
	return &sqliteTable{db: s.db, BaseTable: database.BaseTable{Table: name, Schema: s, Database: s.Database}}, nil
}

func (s sqliteSchema) Empty() (bool, error) {
	sql := fmt.Sprintf(`select count(name) from %s.sqlite_master `, s.Name())
	count, err := s.db.Count(sql)
	if err != nil {
		return false, err
	}
	return count == 0, nil
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
