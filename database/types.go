package database

import (
	"errors"
	"strings"
)

type Type int

const (
	MYSQL Type = iota
	SQLITE
	POSTGRESQL
	DM
	GAUSS
)

func (t Type) String() string {
	switch t {
	case MYSQL:
		return "mysql"
	case SQLITE:
		return "sqlite"
	case POSTGRESQL:
		return "postgresql"
	case DM:
		return "dm"
	case GAUSS:
		return "gauss"
	default:
		return ""
	}
}

// TypeValueOf get database type by name
func TypeValueOf(name string) (Type, error) {
	var t Type
	var e error
	name = strings.ToLower(name)
	switch name {
	case "mysql":
		t = MYSQL
	case "sqlite":
		t = SQLITE
	//case "postgresql":
	//	t = POSTGRESQL
	//case "dm":
	//	t = DM
	//case "dameng":
	//	t = DM
	default:
		e = errors.New("database type " + name + " not found")
	}
	return t, e
}
