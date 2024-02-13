package tests

import (
	"fmt"
	flyway "github.com/goflyway/goflyway"
	"github.com/goflyway/goflyway/command"
	"github.com/goflyway/goflyway/database"
	"testing"
)

func TestSqliteMigrate(t *testing.T) {
	db, err := ConnSqlite()
	if err != nil {
		t.Fatal(err)
	}
	f, err := flyway.Open(database.T_SQLITE, db, &flyway.Config{Locations: []string{"db_migration/sqlite"}})
	if err != nil {
		t.Fatal(err)
	}
	f.Callbacks().Registry("before:migrate", "test", func(context *command.Context) {
		fmt.Println("test before migrate")
	})
	err = f.Migrate()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("sqlite migrate success")
}
