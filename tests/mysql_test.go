package tests

import (
	"github.com/goflyway/goflyway"
	"github.com/goflyway/goflyway/database"
	"testing"
)

func TestMysqlMigrate(t *testing.T) {
	db, err := ConnectMysql()
	if err != nil {
		t.Fatal(err)
	}
	f, err := flyway.Open(database.T_MYSQL, db, &flyway.Config{
		Locations:         []string{"db_migration/mysql"},
		BaselineOnMigrate: true,
		Schemas:           []string{"goflyway", "goflyway2"},
		CreateSchemas:     true,
		CleanDisabled:     true,
	})
	if err != nil {
		t.Fatal(err)
	}
	err = f.Migrate()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("mysql migrate success")
}
