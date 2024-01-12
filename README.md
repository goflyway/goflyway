# GO flyway

![GitHub License](https://img.shields.io/github/license/jiangliuhong/go-flyway)
[![Static Badge](https://img.shields.io/badge/go.dev-reference-blue?style=flat)](https://pkg.go.dev/github.com/jiangliuhong/go-flyway)


## Install

```shell
go get -u github.com/jiangliuhong/go-flyway
```

## Quick start

```go
package main
import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/jiangliuhong/go-flyway"
	"github.com/jiangliuhong/go-flyway/database"
)
func main(){
    db, _ := sql.Open("sqlite3", "./flyway_test.db") 
	// use database.T_SQLITE 、 database.T_MYSQL or "sqlite","mysql"
    f, _ := flyway.Open(database.T_SQLITE, db, &flyway.Config{Locations:[]string{"db_migration"}})
    f.Migrate()
}
```