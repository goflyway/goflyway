package tests

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func ConnSqlite() (*sql.DB, error) {
	return sql.Open("sqlite3", "./flyway_test.db")
}
