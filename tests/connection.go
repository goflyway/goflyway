package tests

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func ConnSqlite() (*sql.DB, error) {
	return sql.Open("sqlite3", "./flyway_test.db")
}

func ConnectMysql() (*sql.DB, error) {
	return sql.Open("mysql", "goflyway:goflyway@tcp(localhost:9910)/goflyway?charset=utf8")
}
