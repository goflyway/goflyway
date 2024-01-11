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

const (
	T_MYSQL      = "mysql"      // mysql db
	T_SQLITE     = "sqlite"     // sqlite db
	T_POSTGRESQL = "postgresql" // postgresql db
	T_DM         = "dm"         // dm ,dameng db
	T_GAUSS      = "gauss"      // gauss db
)

func (t Type) String() string {
	switch t {
	case MYSQL:
		return T_MYSQL
	case SQLITE:
		return T_SQLITE
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
	case T_MYSQL:
		t = MYSQL
	case T_SQLITE:
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
