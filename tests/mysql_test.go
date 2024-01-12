package tests

import (
	"github.com/jiangliuhong/go-flyway"
	"github.com/jiangliuhong/go-flyway/database"
	"testing"
)

func TestMysqlMigrate(t *testing.T) {
	db, err := ConnectMysql()
	if err != nil {
		t.Fatal(err)
	}
	f, err := flyway.Open(database.T_MYSQL, db, &flyway.Config{Locations: []string{"db_migration/mysql"}})
	if err != nil {
		t.Fatal(err)
	}
	err = f.Migrate()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("mysql migrate success")
}
