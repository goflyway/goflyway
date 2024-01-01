package database

import (
	"database/sql"
	"errors"
)

type Database interface {
	Schema(name string) (Schema, error)
	CurrentSchema() (Schema, error)
	Type() Type
	Session() *Session
}

type Schema interface {
	Name() string
	Exists() (bool, error)
	Create() error
	Table(name string) (Table, error)
}

type Table interface {
	Exists() (bool, error)
	Create() error
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
	Name     string
}

func (bt BaseTable) Create() error {
	ddl, err := loadMetadataSql(bt.Database.Type(), bt.Schema.Name(), bt.Name)
	if err != nil {
		return err
	}
	err = bt.Database.Session().ExecDDL(ddl)
	if err != nil {
		return err
	}
	return nil
}
