package database

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
