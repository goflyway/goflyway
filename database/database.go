package database

import (
	"database/sql"
	"errors"
)

type Database interface {
	Schema(name string) (Schema, error)
	CurrentSchema() (Schema, error)
	CurrentUser() (string, error)
	Type() Type
	Session() *Session
}

type Schema interface {
	Name() string
	Exists() (bool, error)
	Create() error
	Table(name string) (Table, error)
	// Empty 判断模式是否为空
	Empty() (bool, error)
	// Delete 删除模式
	Delete() error
	// UseSchema 设置该模式为当前使用的
	UseSchema() error
}

type Table interface {
	Exists() (bool, error)
	Create() error
	Name() string
}

type Group map[Type]func(db *Session) (Database, error)

func (g Group) Registry(t Type, f func(db *Session) (Database, error)) {
	g[t] = f
}

var groups = Group{}

func Registry(t Type, f func(db *Session) (Database, error)) {
	groups.Registry(t, f)
}

func New(t Type, db *sql.DB) (Database, error) {
	f, ok := groups[t]
	if !ok {
		return nil, errors.New("not found " + t.String() + " database")
	}
	session := newSession(db)
	return f(session)
}

type BaseDatabase struct {
	DB *Session
}

func (bd BaseDatabase) Session() *Session {
	return bd.DB
}

type BaseTable struct {
	Schema   Schema
	Database Database
	Table    string
}

func (bt BaseTable) Create() error {
	ddl, err := loadMetadataSql(bt.Database.Type(), bt.Schema.Name(), bt.Table)
	if err != nil {
		return err
	}
	err = bt.Database.Session().Exec(ddl)
	if err != nil {
		return err
	}
	return nil
}

func (bt BaseTable) Name() string {
	return bt.Table
}

type BaseSchema struct {
	Schema string
}

func (bs BaseSchema) Name() string {
	return bs.Schema

}
func (bs BaseSchema) Exists() (bool, error) {
	return false, nil
}
func (bs BaseSchema) Create() error {
	return nil
}
func (bs BaseSchema) Table(name string) (Table, error) {
	return nil, nil
}
func (bs BaseSchema) Empty() (bool, error) {
	return false, nil
}
func (bs BaseSchema) Delete() error {
	return nil
}
func (bs BaseSchema) UseSchema() error {
	return nil
}
