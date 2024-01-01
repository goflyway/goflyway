package tests

import (
	"github.com/jiangliuhong/go-flyway"
	"github.com/jiangliuhong/go-flyway/database"
	"testing"
)

func TestSqliteMigrate(t *testing.T) {
	db, err := ConnSqlite()
	if err != nil {
		t.Fatal(err)
	}
	f, err := flyway.Open(database.SQLITE, db, &flyway.Config{})
	if err != nil {
		t.Fatal(err)
	}
	err = f.Migrate()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("sqlite migrate success")
}
