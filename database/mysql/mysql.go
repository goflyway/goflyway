package mysql

import (
	"errors"
	"fmt"
	"github.com/goflyway/goflyway/database"
	"reflect"
)

func init() {
	database.Registry(database.MYSQL, func(db *database.Session) (database.Database, error) {
		return &mysql{
			database.BaseDatabase{DB: db},
		}, nil
	})
}

type mysql struct {
	database.BaseDatabase
}

func (m mysql) CurrentSchema() (database.Schema, error) {
	sql := `SELECT DATABASE() as currdb`
	res, err := m.DB.SelectOneForMap(sql)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, errors.New("not found current schema")
	}
	currdb := res["currdb"]
	if reflect.TypeOf(currdb).Kind() != reflect.String {
		return nil, errors.New("current database is " + reflect.TypeOf(currdb).Kind().String() + " , not String")
	}
	currdbStr := currdb.(string)
	return &mysqlSchema{BaseSchema: database.BaseSchema{Schema: currdbStr}, db: m.DB, Database: m}, nil
}

func (m mysql) CurrentUser() (string, error) {
	sql := `SELECT SUBSTRING_INDEX(USER(),'@',1) as curruser`
	res, err := m.DB.SelectOneForMap(sql)
	if err != nil {
		return "", err
	}
	if res == nil {
		return "", errors.New("not found current user")
	}
	curruser := res["curruser"]
	if reflect.TypeOf(curruser).Kind() != reflect.String {
		return "", errors.New("current user is " + reflect.TypeOf(curruser).Kind().String() + " , not String")
	}
	return curruser.(string), nil
}

func (m mysql) Schema(name string) (database.Schema, error) {
	return &mysqlSchema{BaseSchema: database.BaseSchema{
		Schema: name,
	},
		db:       m.DB,
		Database: m,
	}, nil
}

func (m mysql) Type() database.Type {
	return database.MYSQL
}

type mysqlSchema struct {
	database.BaseSchema
	db       *database.Session
	Database database.Database
}

func (s mysqlSchema) Exists() (bool, error) {
	sql := `select count(SCHEMA_NAME) from information_schema.SCHEMATA where SCHEMA_NAME = ?`
	count, err := s.db.Count(sql, s.Name())
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
func (s mysqlSchema) Create() error {
	sql := fmt.Sprintf(`create database %s `, s.Name())
	return s.db.Exec(sql)
}
func (s mysqlSchema) Table(name string) (database.Table, error) {
	return &mysqlTable{db: s.db, BaseTable: database.BaseTable{Table: name, Schema: s, Database: s.Database}}, nil
}

func (s mysqlSchema) Empty() (bool, error) {
	sql := `select count(TABLE_NAME) from information_schema.tables where TABLE_SCHEMA = ?`
	count, err := s.db.Count(sql, s.Name())
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

type mysqlTable struct {
	database.BaseTable
	db *database.Session
}

func (t mysqlTable) Exists() (bool, error) {
	sql := `select count(TABLE_NAME) FROM information_schema.tables where TABLE_SCHEMA = ? and TABLE_NAME = ?`
	count, err := t.db.Count(sql, t.Schema.Name(), t.Name())
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
